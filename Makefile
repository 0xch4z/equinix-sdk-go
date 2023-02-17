.PHONY: all gen patch fetch

CURRENT_UID := $(shell id -u)
CURRENT_GID := $(shell id -g)

# https://github.com/OpenAPITools/openapi-generator-cli

SPEC_URL:=https://api.equinix.com/metal/v1/api-docs
SPEC_ROOT_FILE:=openapi3.yaml
SPEC_FETCHED_DIR=spec/oas3.fetched/
SPEC_PATCH_DIR=patches/spec.fetched.json
SPEC_PATCHED_DIR=spec/oas3.patched/
SPEC_COMBINED_DIR=spec/oas3.combined/

SPEC_FETCHED_FILE:=spec.fetched.json
SPEC_PATCHED_FILE:=spec.patched.json
OPENAPI_IMAGE_TAG=v6.3.0
OPENAPI_IMAGE=openapitools/openapi-generator-cli:${OPENAPI_IMAGE_TAG}
GIT_ORG=equinix-labs
GIT_REPO=metal-go
PACKAGE_PREFIX=metal
PACKAGE_MAJOR=v1
CRI=docker # nerdctl

# Pull in custom-built generator jar so we can use unmerged bugfixes
# Custom generator is built from https://github.com/ctreatma/openapi-generator/tree/local-generator-testing
SWAGGER=${CRI} run --rm -u ${CURRENT_UID}:${CURRENT_GID} -v $(CURDIR):/local ${OPENAPI_IMAGE}
GOLANGCI_LINT=golangci-lint

all: pull patch combine_spec clean gen mod docs move-other patch-post fmt test stage

pull:
	${CRI} pull ${OPENAPI_IMAGE}

patch:
	rm -rf ${SPEC_PATCHED_DIR}
	cp -r ${SPEC_FETCHED_DIR} ${SPEC_PATCHED_DIR}

	for diff in $(shell set -x; find ${SPEC_PATCH_DIR} -name '*.patch' | sort -n); do \
		patch --no-backup-if-mismatch -N -t -p1 -i $$diff; \
	done

combine_spec:
	${SWAGGER} generate \
		-i /local/${SPEC_PATCHED_DIR}/${SPEC_ROOT_FILE} \
		-g openapi-yaml \
		-p skipOperationExample=true \
		-o /local/${SPEC_COMBINED_DIR} -p outputFile=${SPEC_ROOT_FILE}

patch-post:
	# patch is idempotent, always starting with the generated files
	for diff in $(shell find patches/post -name \*.patch | sort -n); do \
		patch --no-backup-if-mismatch -N -t -p1 -i $$diff; \
	done

clean:
	rm -rf $(PACKAGE_PREFIX)

gen:
	${SWAGGER} generate -g go \
		--package-name ${PACKAGE_MAJOR} \
		-p isGoSubmodule=true \
		--model-package types \
		--api-package models \
		--git-user-id ${GIT_ORG} \
		--git-repo-id ${GIT_REPO}/${PACKAGE_PREFIX} \
		-o /local/${PACKAGE_PREFIX}/${PACKAGE_MAJOR} \
		-i /local/${SPEC_COMBINED_DIR}/${SPEC_ROOT_FILE}

validate:
	${SWAGGER} validate \
		--recommend \
		-i /local/${SPEC_PATCHED_DIR}

mod:
	rm -f go.mod go.sum ${PACKAGE_PREFIX}/${PACKAGE_MAJOR}/go.mod ${PACKAGE_PREFIX}/${PACKAGE_MAJOR}/go.sum
	go mod init github.com/${GIT_ORG}/${GIT_REPO}
	go mod tidy

test:
	go test -v ./...

clean-docs:
	rm -rf docs API.md

move-docs:
	mv ${PACKAGE_PREFIX}/${PACKAGE_MAJOR}/README.md API.md
	mv ${PACKAGE_PREFIX}/${PACKAGE_MAJOR}/docs .

docs: clean-docs move-docs

move-other:
	rm -rf api .travis.yml git_push.sh
	rm -f ${PACKAGE_PREFIX}/${PACKAGE_MAJOR}/.travis.yml
	mv ${PACKAGE_PREFIX}/${PACKAGE_MAJOR}/api .
	rm ${PACKAGE_PREFIX}/${PACKAGE_MAJOR}/git_push.sh

# https://github.com/OpenAPITools/openapi-generator/issues/741#issuecomment-569791780
remove-dupe-requests: ## Removes duplicate Request structs from the generated code
	@for struct in $$(grep -h 'type .\{1,\} struct' $(PACKAGE_PREFIX)/$(PACKAGE_MAJOR)/*.go | grep Request  | sort | uniq -c | grep -v '^      1' | awk '{print $$3}'); do \
	  for f in $$(/bin/ls $(PACKAGE_PREFIX)/$(PACKAGE_MAJOR)/*.go); do \
	    if grep -qF "type $${struct} struct" "$${f}"; then \
	      if eval "test -z \$${$${struct}}"; then \
	        echo "skipping first appearance of $${struct} in file $${f}"; \
	        eval "export $${struct}=1"; \
	      else \
	        echo "removing dupe $${struct} from file $${f}"; \
	        tr '\n' '\r' <"$${f}" | sed 's~// '"$${struct}"'.\{1,\}type '"$${struct}"' struct {[^}]\{1,\}}~~' | tr '\r' '\n' >"$${f}.tmp"; \
	        mv -f "$${f}.tmp" "$${f}"; \
	      fi; \
	    fi \
	  done \
	done

lint:
	@$(GOLANGCI_LINT) run -v --no-config --fast=false --fix --disable-all --enable goimports $(PACKAGE_PREFIX)

fmt:
	go run mvdan.cc/gofumpt@v0.3.1 -l -w $(PACKAGE_PREFIX)

stage:
	test -d .git && git add --intent-to-add API.md docs ${PACKAGE_PREFIX} go.mod go.sum
