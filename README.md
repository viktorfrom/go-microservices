# Kubernetes/Docker - go-microservices
Simple proof of concept Docker container for playing around with local clusters in Kubernetes.

## Requirements
* Golang 1.13.8+
* Docker 20.10.7+
* Minikube 1.25.2+

### Setup
Install dependencies
```
go get
```

### Usage
Minikube init.
```
minikube start --driver=docker
```

Run the following command before building Docker containers to import said container into Minikube. NOTE: Needs to be re-run for every new terminal. 
```
eval $(minikube docker-env)
```

Run server locally.
```
go run main.go
```

Build Docker container where "my-app" is the name of the project.
```
docker build -t my-app:latest . 
docker run -dp 3000:3000 --name web my-app:latest
```

Or use Docker-Compose instead.
```
docker-compose build
docker-compose up -d
docker-compose ps 
docker-compose down
```

Deploy a local cluster with x number of replicas defined in kubernetes/deployment.yml and connect to an arbitrary pod. 

```
kubectl apply -f kubernetes/
kubectl get po,svc
minikube ip
add my-app-service port to minikube ip 
```

### Tests
Run unit tests
```
npm test
```

## Authors
* Viktor From [viktorfrom](https://github.com/viktorfrom)

## License
Licensed under the MIT license. See [LICENSE](LICENSE) for details.