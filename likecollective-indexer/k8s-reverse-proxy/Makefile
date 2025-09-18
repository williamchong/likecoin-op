-include .env
export

RPROXY_PORT_LOCAL?=50022
APP_PORT_LOCAL?=8091

.PHONY: ensure-env
ensure-env:
ifndef NAMESPACE
	$(error missing NAMESPACE env. Consider adding it in .env)
endif
ifndef APP
	$(error missing APP env. Consider adding it in .env)
endif

.PRECIOUS: rproxy-$(APP)-$(NAMESPACE).pub
rproxy-$(APP)-$(NAMESPACE).pub:
	ssh-keygen -t rsa-sha2-256 -b 4096 -q -N "" -C "user" -f rproxy-$(APP)-$(NAMESPACE)

.PRECIOUS: k8s-secret-$(APP)-$(NAMESPACE)
ssh-rproxy-$(APP)-$(NAMESPACE)-public-key: rproxy-$(APP)-$(NAMESPACE).pub
	$(MAKE) ensure-env
	kubectl -n $$NAMESPACE create secret generic $@ --from-file=$^ 2>&1 | tee $@

.PRECIOUS: rproxy-$(APP)-$(NAMESPACE).yaml
rproxy-$(APP)-$(NAMESPACE).yaml: rproxy.yaml.tmpl
	cat rproxy.yaml.tmpl | NAME="$(APP)-$(NAMESPACE)" NAMESPACE="$(NAMESPACE)" go run github.com/hairyhenderson/gomplate/v4/cmd/gomplate@latest > $@

rproxy-$(APP)-$(NAMESPACE)-container: rproxy-$(APP)-$(NAMESPACE).yaml
	$(MAKE) ensure-env
	$(MAKE) ssh-rproxy-$(APP)-$(NAMESPACE)-public-key
	kubectl -n $$NAMESPACE apply -f $^
	kubectl -n $$NAMESPACE wait --for='jsonpath={.status.readyReplicas}=1' deployment/rproxy-$(APP)-$(NAMESPACE)
	kubectl -n $$NAMESPACE port-forward deployment/rproxy-$(APP)-$(NAMESPACE) $(RPROXY_PORT_LOCAL):2222

rproxy-container: rproxy-$(APP)-$(NAMESPACE)-container

rproxy-$(APP)-$(NAMESPACE)-ssh:
	ssh -i rproxy-$(APP)-$(NAMESPACE) user@localhost \
		-p $(RPROXY_PORT_LOCAL) \
		-R 0.0.0.0:$(APP_PORT_LOCAL):127.0.0.1:$(APP_PORT_LOCAL) \
		-F .ssh.config

rproxy-ssh: rproxy-$(APP)-$(NAMESPACE)-ssh

rproxy-$(APP)-$(NAMESPACE)-container/clean: ssh-rproxy-$(APP)-$(NAMESPACE)-public-key rproxy-$(APP)-$(NAMESPACE).yaml
	$(MAKE) ensure-env
	kubectl -n $$NAMESPACE delete secret ssh-rproxy-$(APP)-$(NAMESPACE)-public-key
	kubectl -n $$NAMESPACE delete -f rproxy-$(APP)-$(NAMESPACE).yaml
	rm -f \
		ssh-rproxy-$(APP)-$(NAMESPACE)-public-key \
		rproxy-$(APP)-$(NAMESPACE).yaml \
		rproxy-$(APP)-$(NAMESPACE) \
		rproxy-$(APP)-$(NAMESPACE).pub \
		.known_hosts

rproxy-container/clean: rproxy-$(APP)-$(NAMESPACE)-container/clean

.PHONY: start-test-server
start-test-server:
	npx http-server test -p $(APP_PORT_LOCAL)

.PHONY: test
test:
	curl https://rproxy-$(APP)-$(NAMESPACE).pandawork.com

