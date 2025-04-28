FROM golang:1.22.2-alpine
RUN echo "...............LAYER 1 completed ✅..............."

#WORKDIR inside container, if not there, it will create
WORKDIR /code/idk
RUN echo "...............LAYER 2 completed ✅..............."

COPY go.mod  ./
#add  GO.SUM later
RUN echo "...............LAYER 3 completed ✅..............."

COPY . .
RUN echo "...............LAYER 4 completed ✅..............."

# Build the application
RUN go build -o app ./cmd/app
RUN echo "...............LAYER 5 completed ✅..............."
# -o app specifies the output filename ("app")
# . tells it to build the code in the current directory


EXPOSE 6009
#exposed port

# Double quotes required, not single quotes
CMD ["./app"]
