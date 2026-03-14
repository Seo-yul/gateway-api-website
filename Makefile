HUGO_VERSION     ?= 0.133.0

# The CONTAINER_ENGINE variable is used for specifying the container engine. By default 'docker' is used
# but this can be overridden when calling make, e.g.
# CONTAINER_ENGINE=podman make container-image
CONTAINER_ENGINE ?= docker
CONTAINER_IMAGE  ?= gateway-api-website-hugo:v$(HUGO_VERSION)

CONTAINER_HUGO_MOUNTS = \
	--read-only \
	--mount type=bind,source=$(CURDIR)/.git,target=/src/.git,readonly \
	--mount type=bind,source=$(CURDIR)/content,target=/src/content,readonly \
	--mount type=bind,source=$(CURDIR)/layouts,target=/src/layouts,readonly \
	--mount type=bind,source=$(CURDIR)/assets,target=/src/assets,readonly \
	--mount type=bind,source=$(CURDIR)/static,target=/src/static,readonly \
	--mount type=bind,source=$(CURDIR)/i18n,target=/src/i18n,readonly \
	--mount type=bind,source=$(CURDIR)/examples,target=/src/examples,readonly \
	--mount type=bind,source=$(CURDIR)/hugo.toml,target=/src/hugo.toml,readonly \
	--mount type=tmpfs,destination=/tmp,tmpfs-mode=01777

.PHONY: help serve build container-image container-serve update-geps verify-geps

help: ## Show this help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

serve: ## Boot the development server locally.
	npm ci
	hugo server --buildDrafts --buildFuture --environment development

build: ## Build the static site into ./public.
	npm ci
	hugo --cleanDestinationDir --minify --environment development

container-image: ## Build a container image for the website.
	$(CONTAINER_ENGINE) build . \
		--network=host \
		--tag $(CONTAINER_IMAGE) \
		--build-arg HUGO_VERSION=$(HUGO_VERSION)

container-serve: ## Boot the development server using a container.
	"$(CONTAINER_ENGINE)" run --rm --interactive --tty \
		--cap-drop=ALL --cap-add=AUDIT_WRITE \
		$(CONTAINER_HUGO_MOUNTS) \
		-p 1313:1313 $(CONTAINER_IMAGE) \
		hugo server --buildDrafts --buildFuture --environment development --bind 0.0.0.0 --destination /tmp/public --cleanDestinationDir --noBuildLock

update-geps: ## Regenerate the GEP listing page from metadata.yaml files.
	hack/update-geps.sh

verify-geps: ## Verify the GEP listing page is up to date.
	hack/verify-geps.sh
