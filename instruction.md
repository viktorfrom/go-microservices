#Setup
#Minikube
1. minikube start --driver=docker
2. eval $(minikube docker-env)

#Run
#Docker 
1. docker build -t my-app:latest . 
2. docker run -dp 3000:3000 --name web my-app:latest

or docker-compose
1. docker-compose build
2. docker-compose up -d
3. docker-compose ps 
4. docker-compose down

#Kubernetes
1. kubectl apply -f kubernetes/
2. kubectl get po,svc
3. minikube ip
4. add my-app-service port to minikube ip 