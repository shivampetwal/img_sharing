cd .. #move to root dir

# Clean up any old container with the same name
docker rm -f go-app 2>/dev/null || true

# Build the Docker image
docker build -f deploy/Dockerfile -t go .

# Bring down any existing docker-compose services to free the port
docker compose -f ./deploy/docker-compose.yml down

docker compose -f ./deploy/docker-compose.yml up -d

echo "Application started on http://localhost:6009"


#TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO
# []ADD DOCKER COMPOSE UP + DOWN