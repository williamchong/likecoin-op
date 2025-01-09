GIT_SHORT_HASH ?= $(shell git rev-parse --short=8 HEAD)
APP_VERSION ?= $(GIT_SHORT_HASH)

ENV ?= dev

K8S_NAMESPACE := likecoin-op

VALUE_FILES := -f helm/values.yaml \
	-f helm/$(ENV)/values.yaml \
	-f helm/$(ENV)/values.secret.yaml
ADDITIONAL_VALUES := \
	--set-string appVersion="$(APP_VERSION)"

.PHONY: debug-template
debug-template:
	helm template --debug helm ./helm/ \
		$(VALUE_FILES) \
		$(ADDITIONAL_VALUES)

.PHONY: lint
lint:
	helm lint helm \
		$(VALUE_FILES) \
		$(ADDITIONAL_VALUES)

.PHONY: deploy
deploy:
	helm upgrade likecoin-op helm --install \
		--wait --wait-for-jobs --cleanup-on-fail \
		-n $(K8S_NAMESPACE) \
		$(VALUE_FILES) \
		$(ADDITIONAL_VALUES)
