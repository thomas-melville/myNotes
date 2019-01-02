# Kata

https://www.katacoda.com/courses/docker/1

docker search -s 1 <name>
	-s minimum number of stars

docker run <options> <name>
	-d background (daemon)
	--name <name> give the running container a meaningful name

docker ps
	list running containers

docker inspect <name/id>
	View the details about a running container
	big json object

docker logs <name/id>
	print logs for running container

**Binding Ports**
	Without binding ports a container is only accessible from the machine it is running on.
	Binding Ports allows you to access a container from outside the host machine
	(Kind of like port forwarding in a vApp)
	-p :Open Source Projects:Docker:port:port option when starting the container
		**what is source and destination port?**
	Specifying source and dest limits number of instances of container which can be running
	-p port will expose the port and assign a random port
		**This binds the port on the container to a port on the host machine?**

**Binding Host Directories**
	Containers are designed to be stateless
	any data we want to persist needs to be saved to the host machine
	Achieved by mounting/binding host directories in the container
	-v dir:dir

**Images are built from base images**
	you get a lot of the low level configuration done for you
	not advised to use latest.
		new version could be incompatible with your image

**Dockerfile Syntax**
FROM
	the base image to use
RUN
	execute any command in the container just like a shell command
COPY
	copy files from the directory which contains the Dockerfile to the container
EXPOSE
	tell docker which ports should be open on the container and bound to
	can specify single, multiple or range
CMD
	Defines default command to run when the container is launched
	If the command takes arguments an array needs to be passed to the CMD keyword
	["command", "argument"...]
ENTRYPOINT
	similar to CMD but allows commands to be passed in when the container is launched
WORKDIR
	Define a directory from which all future commands are executed from relative to the application
