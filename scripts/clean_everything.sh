#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")/.."
docker compose -f deploy/docker-compose.yml down

echo "✅ All Compose services stopped."