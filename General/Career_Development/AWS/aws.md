# hackathon

app: nodeStoreRestService
db: mongo
deploy to docker
docker-compose
aws codepipeline????

So much functionality that it gets confusing to know where to start
quickly get tied into all things Amazon
roles in AWS are very important

## AWS Code pipeline

Amazons Jenkins

nicely integrates with github

### buildspec.yml

They have their own build file for stages of the pipeline
you commit it to your repo and specify it in the stage

phases:
	install
	build
	pre_build
	post_build
		commands


## AWS Elastic Container Service

Amazons docker solution 

## Elastic Container Registry

Amazons docker hub registry