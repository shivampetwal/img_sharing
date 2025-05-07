#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")/.."
export $(grep -v '^#' deploy/.env | xargs)

docker build -f deploy/Dockerfile-migration -t migrate_db .
docker compose -f deploy/docker-compose.yml up -d db

until docker compose -f deploy/docker-compose.yml exec db pg_isready \
    -U "$POSTGRES_USER" >/dev/null 2>&1; do sleep 1; done

MIGRATE="docker compose -f deploy/docker-compose.yml run --rm migrate"

# ────────────────────────────────────────────────────────────
# 1) Check for a dirty database version and fix it
VERSION_OUTPUT=$($MIGRATE version 2>&1)
if echo "$VERSION_OUTPUT" | grep -q "Dirty"; then
  DIRTY=$(echo "$VERSION_OUTPUT" | awk '/Dirty/ {print $1}')
  CLEAN=$((DIRTY - 1))
  echo "⚠️  Dirty migration detected (v${DIRTY}). Forcing to v${CLEAN}…"
  $MIGRATE force "${CLEAN}"
fi

# 2) Apply any pending up-migrations
echo "🚀 Applying migrations…"
$MIGRATE up

echo "✅ Migrations completed!"