cd .. #move to root dir

# Clean up any old container with the same name
docker rm -f go-app 2>/dev/null || true

# Build the Docker image
docker build -f deploy/Dockerfile -t go .

# Run the container
docker run -d -p 6009:6009 --name go-app go
# deatached ,
#-p portmapping
#--name of container
#name of image.. go




echo "Application started on http://localhost:6009"


#TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO
# []ADD DOCKER COMPOSE UP + DOWN