-include .env
export

PROJECT_NAME := likecollective-indexer

.PHONY: secret
secret:
	@curl -s https://raw.githubusercontent.com/oursky/devsecops-secret/v1.0/generate-secret.sh \
		| bash -s -- \
		.env.example .env
