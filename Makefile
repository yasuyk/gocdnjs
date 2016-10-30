OSARCH="linux/386 linux/amd64 darwin/amd64 windows/386 windows/amd64"
RELEASE_WORKING_DIR=release/gocdnjs
PROJECT_ROOT=$(shell pwd)

.PHONY: gox-exists
gox-exists: ; @which gox > /dev/null 2>&1

.PHONY: ghr-exists
ghr-exists: ; @which ghr > /dev/null 2>&1

.PHONY : check-release-dependencies
check-release-dependencies: gox-exists ghr-exists

.PHONY : clean-release
clean-release:
	rm -rf $(RELEASE_WORKING_DIR)

.PHONY : check-release-tag-variable
check-release-tag-variable:
	@git tag -l | grep -P "^$$RELEASE_TAG$$" > /dev/null || (echo $(RELEASE_TAG) is not found in git tags && exit 1)
	@echo RELEASE_TAG=$(RELEASE_TAG)

.PHONY : release
release: check-release-dependencies clean-release check-release-tag-variable
	@mkdir -p $(RELEASE_WORKING_DIR)
	git clone -b $(RELEASE_TAG) . $(RELEASE_WORKING_DIR)
	cd $(RELEASE_WORKING_DIR)
	cd $(RELEASE_WORKING_DIR) && gox --osarch $(OSARCH) -output "dist/{{.Dir}}_$(RELEASE_TAG)_{{.OS}}_{{.Arch}}/{{.Dir}}"
	cd $(RELEASE_WORKING_DIR) && bash $(PROJECT_ROOT)/script/zip-releases.sh
	ghr -t $(GITHUB_TOKEN) $(RELEASE_TAG) $(RELEASE_WORKING_DIR)/pkg/

