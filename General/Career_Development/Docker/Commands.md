# Commands

##### Docs #####
https://docs.docker.com/engine/reference/commandline/docker/
https://docs.docker.com/compose/compose-file/
https://docs.docker.com/compose/compose-file/compose-file-v2/

##### Docker commands #####
docker ps							// list containers - running
docker ps -a						// list containers - all
docker search -s 1 <name>			// search docker hub for <name>     -s 1 # only show results that have rating 1 star or above
docker rm <container>				// remove a container
docker inspect myredis				// print config for myredis
docker logs myredis					// show what the container wrote to std out and std err
docker build -t webserver-image:v1		// build a docker image	
docker build .								// build a docker image in the current dir (need to contain a docker file)
doc build -t armdocker.rnd.ericsson.se/proj_taf_te/taf_te_base_slave:latest --build-arg TE_SLAVE_ARTIFACT_VERSION#1.0.71-SNAPSHOT --build-arg ARM_REPOSITORY#snapshots .
docker push armdocker.rnd.ericsson.se/proj_oss/<image-name>:<tag>		// push an image
docker pull armdocker.rnd.ericsson.se/proj_oss/<image-name>:<tag>		// pull an image
docker images									// list docker images
docker rmi  <image-id>							// delete a docker image
docker ps										// View what docker containers are running 
docker exec -ti <container-id> bash					// Access a running container
docker stop `docker ps -q`						// Stops all running containers
docker info									// View number of docker containers and image1
docker rm $(docker ps -a -q)						//remove all docker containers
docker rm -v $(docker ps -a -q -f status#exited)		//Delete all exited containers
docker rmi $(docker images -f "dangling#true" -q)		//Remove unwanted ‘dangling’ images
docker run -ti <image-id/image-name> bash			// Create a container from an image and run inside it
docker run -it --entrypoint /bin/bash <image> 			// Run but use a differnt entry point
docker tag 4bf94ff45502 armdocker.rnd.ericsson.se/proj_taf_te/taf_te_slave:latest		//tag1 an image
docker login armdocker.rnd.ericsson.se			// login

### docker run options ###
-d <image>				// -d run in the background
-name <image>		// give the image a friendly name
		docker run --name myredis -d redis:latest
-p 6379:6379			// set the port
-p 6379				//let docker assign the port on the host
		docker run -d --name redisFixed -p 6379:6379 redis:latest		// hard wire the port 
		docker run -d --name redisDynamic -p 6379 redis:latest		// let docker assing the port 
-v "$PWD/data":/data redis			//map a folder on the host to a folder in the container
		docker run -d --name redisMapped -v "$PWD/data":/data redis		
-e NODE_ENV#production			// set env variable
-it		//interactive shell

##### Docker-compopse commands #####
docker-compose up -d
docker-compose ps
docker-compose -f docker-compose.yml -f docker-compose.admin.yml build
docker-compose stop
docker-compose pull
docker-compose rm
docker-compose kill
docker-compose logs
docker-compose build --pull					//Always attempt to pull a newer version of the image
docker-compose up --force-recreate -d 		//Recreate containers even if their configuration image haven't changed.

#### Delete docker images in jenkins ####
{{{code: linenumbers#"True"
function delete_images {
    echo "Deleting image $1"
    all_images#$(docker images -q $1)
    if [ -n "$all_images" ]; then
         # Remove duplicates
         sorted_unique_ids#$(echo "${all_images[@]}" | tr ' ' '\n' | sort -u | tr '\n' ' ')
         docker rmi -f $sorted_unique_ids
    fi
}

#delete_images armdocker.rnd.ericsson.se/proj_taf_te/taf_te_master
#delete_images armdocker.rnd.ericsson.se/proj_taf_te/taf_te_slave
#delete_images armdocker.rnd.ericsson.se/proj_taf_te/te_message_bus
}}}
