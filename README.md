[![Build Status](https://travis-ci.com/sriram-yeluri/sonar-cli.svg?branch=master)](https://travis-ci.com/sriram-yeluri/sonar-cli)
[![Go Report Card](https://goreportcard.com/badge/github.com/sriram-yeluri/sonar-cli)](https://goreportcard.com/report/github.com/sriram-yeluri/sonar-cli)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

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
docker run -d --name sonarqube -p 9000:9000 -p 9092:9092 sonarqube:6.7.4 
```
#### Local instance of sonarqube can be reached on http://localhost:9000 , with default credentials (admin/admin)


#### using pagination for Projects: [stackoverflow](https://stackoverflow.com/questions/47889780/how-to-get-more-than-500-issues-from-sonarqube-api)
api/issues/search?componentKeys=PROJECT_KEY&ps=500&p=1,
then api/issues/search?componentKeys=PROJECT_KEY&ps=500&p=2, etc.

The total number of page can be retrieved from the response:  "paging" -> "total".
Example:
http://localhost:9000/api/components/search?qualifiers=TRK&p=2&ps=500
```
	"paging": {
		"pageIndex": 2,
		"pageSize": 500,
		"total": 2001
	}
```
	
