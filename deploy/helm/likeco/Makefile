.PHONY: secrets
secrets:
	kubectl create secret generic basic-auth --from-file=auth
	kubectl create secret generic \
		signer-private-key --from-env-file=signer-private-key.env

