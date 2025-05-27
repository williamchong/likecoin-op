.PHONY: htpasswd
htpasswd:
	htpasswd -b -c auth \
		`yq .migrationAdmin.env.ADMIN_USERNAME values.secret.yaml` \
		`yq .migrationAdmin.env.ADMIN_PASSWORD values.secret.yaml`

.PHONY: secrets
secrets:
	kubectl -n likecoin-op \
		create secret generic basic-auth --from-file=auth
