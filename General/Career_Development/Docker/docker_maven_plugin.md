# DockerMavenPlugin

## Documentation
https://github.com/spotify/docker-maven-plugin

## Command line
mvn clean package docker:build			//build an image with the above configurations by running this command.

## options
'''
	-DpushImage 		//push to registry
	-DpushImageTag	//push only specific tags of the image
	-DpushImageTags -DdockerImageTags#latest -DdockerImageTags#another-tag		//overwrite tags and push to registry
'''

mvn docker:removeImage -DimageName#foobar	//remove an image
