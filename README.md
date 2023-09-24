## Prerequisites

* Docker CLI

## Basic Usage

```bash
docker build -t yatzy-api .
```

```bash
docker run -p 8080:8080 yatzy-api
```

## Running Tests
To run tests, go to the directory you want to run the tests in.

Then run:  
```bash
go test -v
```

## Docker hub
For this project we use docker hub for our container registry, since they offer 1 free private repo.

docker push leebadal1/yatzy-api:tagname


## Pushing to docker hub

PS. REMEMBER TO: docker login 

Locate root directory

Build docker image for kubernetes: 
1. docker build -t yatzy-api:v1(+1) . 
Tag local image
2. docker tag yatzy-api:v1 leebadal1/yatzy-api:v1(+1)
Push image to intenal registry
3. docker push leebadal1/yatzy-api:tagname
Create kubernetes deployment (see below)
4. 
spec:
  containers:
    - name: api-container
      image: leebadal1/yatzy-api:v1.10(+1)
  imagePullSecrets:
        - name: regcred


Since the repo is private you must create a secret:

kubectl create secret docker-registry regcred --docker-server=https://index.docker.io/v1/ --docker-username=leebadal1 --docker-password=<your-pword> --docker-email=<your-email>

Once you have pushed to docker hub and updated deployment, apply all api-*.yaml's with kubectl (see below)

You will have to get a new minikube service url by running:

minikube service api-service --url

### Minikube

Enable kubernetes in docker

download minikube & run: minikube start --driver=docker

run: minikube dashboard or minikube ip to verify

Build docker image for kubernetes: 
1. docker build -t yatzy-api:v1 .
Tag local image
2. docker tag yatzy-api:v1 $(minikube ip):5000/yatzy-api:v1
Push image to intenal registry
3. docker push $(minikube ip):5000/yatzy-api:v1
Create kubernetes deployment (see below)
4. 
spec:
  containers:
    - name: api-container
      image: $(minikube ip):5000/yatzy-api:v1

apply deployment

### First time running database
For running the database for the first time, you will need to generate a user/password for psql:
In this case we choose to store the secrets in the k8 cluster using the database-creds.yaml
As this contains sensitive data the repo contains a database-creds-template.yaml in which you can add you username/password for the service.

Enter a username/password and deploy the secret to using kubectl:

1. kubectl apply -f database-credentials-secret.yaml



### Applying deployment

1. kubectl apply -f api-deployment.yaml 
2. kubectl apply -f api-ingress.yaml
3. kubectl apply -f api-service.yaml
4. kubectl scale deployment api-deployment --replicas=3  (optional)
5. kubectl port-forward svc/api-service 8080:80 ##
6. CLEANUP 
kubectl delete deployment api-deployment
kubectl delete ingress api-ingress
kubectl delete service api-service

## When running
### Get status
kubectl get deployment

### Get pods
kubectl get pods

### View logs of pods
kubectl logs <pod-name>

### Get node-ip & node-port
to reach the service, we have exposed a node-port. which is 30000 defined in service
to find this use: 
kubectl describe svc api-service 
look for NodePort


### Making changes and redeploying to minikube
1. Making changes to code
2. Clean up deployment


### Troubleshooting

If you intalled minikube as an elevated user, run elevated command prompts when runnign commands.


If you have issues with $(minikube ip) resolving, try these steps:

Check Docker Contexts:

Run the following command to list your Docker contexts:

sh
Copy code
docker context ls
If you see any issues with the contexts, such as missing or misconfigured context, try to fix them.

Reset the Docker CLI Context:

If the context associated with Minikube is causing problems, you can reset the Docker CLI context by running:

sh
Copy code
docker context use default
Replace "default" with the appropriate context name if it's different.

Check Minikube Status:

Ensure that Minikube is running and its status is normal:

sh
Copy code
minikube status
If Minikube is not running, start it using:

sh
Copy code
minikube start
Restart Docker and Minikube:

Sometimes, restarting Docker and Minikube can resolve various context-related issues:

Restart Docker Desktop.
Restart Minikube using minikube stop followed by minikube start.


### finding URL for your service:
minikube service <service-name> --url

(on windows terminal has to remain open, url will change on restart)

### Quick launch
1. Start Docker Desktop, use minikube context
2. start minikube (minikube start)
3. Apply everything (see above)
4. Get the url for the service (see finding url for your service)



### Api2

protoc --go_out=plugins=grpc:. dbservice.proto

example:
protoc -I api/ api/user.proto --go_out=plugins=grpc:service1/
protoc -I api/ api/user.proto --go_out=plugins=grpc:service2/

