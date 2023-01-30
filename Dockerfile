FROM golang:latest

WORKDIR /app

COPY ./src ./src
COPY ./config/config.yaml ./config/config.yaml

# Build the project
RUN cd src && go build -o my-project

# Expose the port
#EXPOSE 8080

# Run the project
CMD ["./src/my-project", "serve"]
