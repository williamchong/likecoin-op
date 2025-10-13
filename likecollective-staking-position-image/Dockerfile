FROM python:3.13-slim as deps

WORKDIR /app

RUN pip install poetry && poetry self add poetry-plugin-export

COPY pyproject.toml poetry.lock ./

RUN poetry export -f requirements.txt --output requirements.txt

FROM python:3.13-slim

WORKDIR /app

COPY --from=deps /app/requirements.txt ./

RUN pip install -r requirements.txt

COPY . ./

CMD ["uvicorn", "staking_image.server:app", "--host", "0.0.0.0", "--port", "8000"]
