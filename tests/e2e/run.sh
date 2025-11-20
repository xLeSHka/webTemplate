#!/bin/bash

# install uv and deps
curl -LsSf https://astral.sh/uv/install.sh | sh
uv venv && uv sync

# start backend
docker compose up -d --build

# run tests
uv run pytest -n auto

# stop backend
docker compose down
