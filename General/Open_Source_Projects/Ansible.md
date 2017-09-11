Content-Type: text/x-zim-wiki
Wiki-Format: zim 0.4
Creation-Date: 2015-12-15T12:38:34+00:00

###### Ansible ######
Created Tuesday 15 December 2015

Automate application deployment
Provision onto bare metal or into cloud
Configuration manage deployments (update yaml and re-roll) - immutable infrastructure
among other things

Integrations:
	AWS
	Docker
		Playbook ontop of docker container
	OpenStack
	RedHat
	Windows

Playbooks are a level above integrations so they can be run on any of the integrations
Ansible can model containers and non-containers

No agent running on deployment
	ssh to
	load modules
	execute
	delete modules
	
	Or as another guy describes it
		ssh to server
		run configured task

Provision multiple hosts at once.
	On central machine update [[/etc/ansible/hosts]] file with ip's of deployed applications
	update script
	run against all ip's in hosts file!

Terminologies
	Facts
	Plays
	Tasks
	Playbooks
