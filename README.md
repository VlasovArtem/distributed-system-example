# Distributed Communication

This repository is a skeleton for further work on the course "Distributed Communication".

All microservices are ready to be run as Docker images, you can find **Dockerfile** for each microservice and **docker-compose.yml** to start all microservices with ports binding.

## Build and Run

- ```docker-compose 	up --build```  - builds, (re)creates, starts, and attaches to containers for a service.

- ```docker-compose 	down``` - stops containers and removes containers, networks, volumes, and images created by **up**. 	 	

Existing **docker-compose.yml** file contains required microservices port bindings and dependencies:
- Authors MS is available by port 8094.
- Books MS is available by port 8095
- Frontend MS is available by port 8096
- RabbitMQ is available by ports 5672 and management port 15672; you can check management dashboard console with **guest/guest** as user/password in your browser

## Useful information
- Protobuf files (.proto) and the generated ones should be placed in a standalone Go module. 
- You can use any popular framework for HTTP endpoints implementation, but the most lightweight simple, and popular at the same time is gorilla mux.
- if you use private git repository, most probably you should customize GOPRIVATE env and project's git config: [Configuration for downloading non-public code ](https://tip.golang.org/cmd/go/#hdr-Configuration_for_downloading_non_public_code), [Why does "go get" use HTTPS when cloning a repository?](https://golang.org/doc/faq#git_https)

### Additional Links

[Example Go REST API ](https://thenewstack.io/make-a-restful-json-api-go/)

[Language Guide (proto3)](https://developers.google.com/protocol-buffers/docs/proto3)

[GRPC Quickstart](https://grpc.io/docs/languages/go/quickstart/)

[GRPC Basics](https://grpc.io/docs/languages/go/basics/)

[Series of rabbitMQ tutorials](https://www.rabbitmq.com/tutorials/tutorial-one-go.html)