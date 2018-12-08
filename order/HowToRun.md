# GOLANG REST API for order service

## Running GOLANG API locally

#### 0. Open a terminal

#### 1. Set your GOPATH to the project directoy

``` 
export GOPATH="Your Project directory"
```

- Note you might need to setup your environment before running the API

#### 2. Build your app
```
go build burger-order
```

#### 3. Run the app from terminal
```
./burger-order
```

#### 4. See where the app runs
```
[negroni] listening on :8000
```
#### 5. Test APIs using postman
```
Try ping the API

curl -X GET \
  http://localhost:8000/order/ping \
  -H 'Postman-Token: ceca2e35-2963-4bd8-9a41-6504f9186f78' \
  -H 'cache-control: no-cache'

{
    "Test": "Burger order API Server Working on machine: 10.250.236.69"
}

```

## Running the GO API in EC2 using docker

#### 1. Install Docker 

#### 2. Start Docker
```
sudo systemctl start docker
sudo systemctl is-active docker
```

#### 3. Login to your docker hub account
```
sudo docker login
```

#### 4. Create Docker file 
```
sudo vi Dockerfile

FROM golang:latest 
EXPOSE 8000
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
ENV GOPATH /app
RUN cd /app ; go install burger-order
CMD ["/app/bin/burger-order"]
```

#### 5. Build the docker image locally
```
sudo docker build -t burger-order .
sudo docker images
```

#### 6. Push docker image to dockerhub
```
docker push burger-order:latest
```

#### 7. Create Public EC2 Instance

Configuration:
1. AMI:             CentOS 7 (x86_64) - with Updates HVM
2. Instance Type:   t2.micro
3. VPC:             cmpe281
4. Network:         Public subnet (us-west-1c)
5. Auto Public IP:  Yes
6. SG Open Ports:   22, 80, 8080, 3000, 8000
7. Key Pair:        cmpe281-us-west-1

#### 8. ssh to your ec2 instance, user name is centos

#### 9. Create docker-compose yml file (with the environment variables set up)

#### 10. Deploy go API for order sevice
```
docker-compose up
```

#### 11. Clean Up docker environment when finished
```
docker stop burger-order
docker rm burger-order
docker rmi {imageid}
``` 

## Configuring Mongo Shard
#### 1. ssh to the querry-router instance for shard cluster

#### 2. Login to mongo shell
```
mongo mongo-router:27017 -u mongo_admin -p --authenticationDatabase admin
pass: yourpassword
use burger
db.order.getShardDistribution()
```

#### 3. Use burger database
```
use burger
```

#### 4. Enable Sharding
```
sh.enableSharding("burger")
```

#### 5. Select shard key for your collection
```
db.order.ensureIndex( { orderId : "hashed" } )
sh.shardCollection( "burger.order", { "orderId" : "hashed" } )
```

#### 6. Get the shard distribution
```
db.order.getShardDistribution() 
```

## Microsoft AKS
Official tutorial: https://docs.microsoft.com/en-us/azure/aks/tutorial-kubernetes-prepare-app

#### 0. Create your resource group 
```
az group create --name myResourceGroup --location westus
```

#### 1. Create your container registry
```
az acr create --resource-group myResourceGroup --name <acrName> --sku Basic
```

#### 2. Login to your container registry
```
az acr login --name <acrName>
```

#### 3. Get the login server address
```
az acr list --resource-group myResourceGroup --query "[].{acrLoginServer:loginServer}" --output table
```

#### 4. (Locally) tag your local docker image
```
docker tag burger-order <acrLoginServer>/burger-order:v1
```

#### 5. Push the docker image and check the image and tag in the registry
```
docker push <acrLoginServer>/burger-order:v1
az acr repository list --name <acrName> --output table 
4az acr repository show-tags --name <acrName> --repository burger-order --output table
```

#### 6. Create a service principal
```
az ad sp create-for-rbac --skip-assignment

Sample Output:
{
  "appId": "e7596ae3-6864-4cb8-94fc-20164b1588a9",
  "displayName": "azure-cli-2018-06-29-19-14-37",
  "name": "http://azure-cli-2018-06-29-19-14-37",
  "password": "52c95f25-bd1e-4314-bd31-d8112b293521",
  "tenant": "72f988bf-86f1-41af-91ab-2d7cd011db48"
}
```

#### 7. Get the resource ID
```
az acr show --resource-group myResourceGroup --name <acrName> --query "id" --output tsv
```

#### 8. Create a role assignment
```
az role assignment create --assignee <appId> --scope <acrId> --role Reader
```

#### 9. Create the AKS cluster
```
az aks create \
    --resource-group myResourceGroup \
    --name myAKSCluster \
    --node-count 1 \
    --service-principal <appId> \
    --client-secret <password> \
    --generate-ssh-keys
```

#### 10. Connect to the kubernetes server
```
az aks get-credentials --resource-group myResourceGroup --name myAKSCluster
```

#### 11. Create a yaml file with all the deployment, service, and environment variables setup

#### 12. Apply the yml file 
```
kubectl apply -f burger-order.yaml
```

#### 13. Check the pods, deployments, and service
```
kubectl get pods -o wide
kubectl get deployments -o wide
kubectl get services -o wide
```

#### 14. Test the API using postman via the external IP of the exposed service
```
curl -X GET \
  http://168.61.40.253:8000/order/ping \
  -H 'Postman-Token: 0978d781-b3b4-4075-b13b-4eae2d7f4a0b' \
  -H 'cache-control: no-cache'
```

























