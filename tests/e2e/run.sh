#!/bin/bash
set -euo pipefail

cleanup() {
    docker compose logs backend
    echo "Cleaning up Docker containers..."
    docker compose down --remove-orphans || true
    rm -rf ./.venv || true
}
trap cleanup EXIT

# remove old venv
rm -rf ./.venv

# install uv and deps
if [ -x "$HOME/.local/bin/uv" ]; then
    uv_path="$HOME/.local/bin/uv"
    echo "Found uv at $uv_path"
else
  echo "uv not found, installing..."
  # Install uv
  curl -LsSf https://astral.sh/uv/install.sh | sh
fi
export PATH="$HOME/.local/bin:$PATH"
uv venv && uv sync

# start backend
docker compose up -d --build

# wait for services to be ready
echo "Waiting for services to be ready..."
sleep 5

# run tests with explicit failure handling
echo "Running Tavern tests..."
TEST_EXIT_CODE=0

# run tests and capture exit code
uv run pytest -n auto --maxfail=1 --exitfirst || TEST_EXIT_CODE=$?

# check if all tests failed
if [ $TEST_EXIT_CODE -ne 0 ]; then
    echo "ERROR: Tests failed with exit code $TEST_EXIT_CODE"
    exit $TEST_EXIT_CODE
fi
