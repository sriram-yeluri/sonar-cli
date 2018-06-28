# sonar-cli
sonarqube cli build with golang

## Project Setup

```
mkdir -p sonar-cli/src/github.com/sriram-yeluri/sonar-cli  
mkdir -p sonar-cli/src/github.com/sriram-yeluri/sonar-cli/utils  
mkdir -p sonar-cli/src/github.com/sriram-yeluri/sonar-cli/sonarlib  
mkdir -p sonar-cli/pkg  
mkdir -p sonar-cli/bin  
touch README.md  

export GOPATH=/home/sriram/goProjects/sonar-cli/  
```
### For testing purpose, docker container is used on local instance.  
Pull docker image of sonarqube and spin a container
[Docker-hub-sonarqube](https://hub.docker.com/_/sonarqube/)  
```
docker pull sonarqube  
docker run -d --name sonarqube -p 9000:9000 -p 9092:9092 sonarqube  
```
#### Local instance of sonarqube can be reached on http://localhost:9000 , with default credentials (admin/admin)
