.PHONY: operator-key-link
operator-key-link:
	ln -s ../deploy/env.operator likecoin/.env
	ln -s ../deploy/env.operator likenft/.env

.PHONY: remove-operator-key-link
remove-operator-key-link:
	rm -f likecoin/.env
	rm -f likenft/.env

.PHONY: local-node
local-node:
	(sleep 1 && make -C operation init-local-state) &
	(sleep 2 && make -C likecoin deploy-local) &
	(sleep 3 && make -C likenft deploy-local) &
	make -C operation local-node