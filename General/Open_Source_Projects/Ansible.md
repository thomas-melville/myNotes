# Ansible 

Automate application deployment
Provision onto bare metal or into cloud
Configuration manage deployments (update yaml and re-roll) - immutable infrastructure
among other things

**In yaml formatting, spacing and indentation are very important**

Integrations:
	AWS
	Docker
		Playbook ontop of docker container
	OpenStack
	RedHat
	Windows
	Vagrant

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
	Inventory
	Facts
	Plays
	Tasks
	Playbooks

## Inventory

what hosts to execute against.
Selected from hosts file in /etc/ansible/hosts
hosts can be grouped.
And individual hosts can be configured, ssh port, user, ...

## Modules

Ansible has a module for almost any task
	git, apt-get, shell, ....
	all listed here: http://docs.ansible.com/ansible/latest/modules_by_category.html

## Playbooks

put lots of module calls together into a playbook
format is YAML, who's history is python
A playbook is composed of plays
	a play is a series of tasks executed against a group of hosts

**http://docs.ansible.com/ansible/latest/playbooks_intro.html Basics**

## tasks
all arguments can be on one line or split on several lines in key: value pairs

## vars

you can set variables at the top of the playbook and reference them through out the play book

  vars:
    var: value

{{ var }}
if it makes up part of entry then it all needs to be enclosed in inverted commans

## sudo commands

  become: true
  become_user: root

## looping in a task

  with_items:
    x
    y
    z
  {{ item }}