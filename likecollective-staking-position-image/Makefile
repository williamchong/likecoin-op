PYTHON_VERSION = python3.13

PY_MODULES = staking_image

VENV = venv

.PRECIOUS: $(VENV)
$(VENV):
	$(PYTHON_VERSION) -m venv $(VENV)

.PHONY: setup-dev
setup-dev: $(VENV)
	source $(VENV)/bin/activate; \
		pip install poetry; \
		poetry install

.PHONY: format
format: $(VENV)
	source $(VENV)/bin/activate; \
		poetry run black $(PY_MODULES); \
		poetry run isort $(PY_MODULES)

.PHONY: dev
dev: $(VENV)
	source $(VENV)/bin/activate; \
		fastapi dev staking_image/server.py
