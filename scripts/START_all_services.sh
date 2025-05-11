cd .. #move to root dir



# ──────────────────────────────────────────────────────────────────────────────
# ensure go.mod exists, otherwise initialize module and tidy
if [ ! -f go.mod ]; then
  echo "⚙️  go.mod not found; initializing module code/idk"
  go mod init code/idk
  go mod tidy
fi
# ──────────────────────────────────────────────────────────────────────────────


# Clean up any old container with the same name
docker rm -f go-app 2>/dev/null || true

# Build the Docker image idk-backend
docker build -f deploy/Dockerfile -t idk-backend .


echo "🔴 Shutting down existing services..."
docker compose -f deploy/docker-compose.yml down


echo " Running migrations…"
./scripts/migrate.sh

docker compose -f ./deploy/docker-compose.yml up -d

echo "Application started on http://localhost:6009"
