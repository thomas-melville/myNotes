Content-Type: text/x-zim-wiki
Wiki-Format: zim 0.4
Creation-Date: 2017-02-03T08:44:45+00:00

###### try stack ######
Created Friday 03 February 2017

starting an instance

1. network
	a. dns: 8.8.8.8
2. launch instance
	a. Key Pairs
		a. paste id_rsa.pub
3. router
	a. set gateway to external
	b. add interface to internal network
4. Setup access rules
	a. security groups
		a. 2 new rules
			a. custom icmp rule
			b. type: -1
			c. code:  -1
			d. cutom tcp rule
			e. port, 22
5. Floating ip address

##### SSH to instance #####

There look to be 2 ways to enable ssh'ing to an instance:

1. upload your public key
2. create a key pem file from openstack, download it and generate the private key, and the use the private key to ssh.
	a. It's asking me for a passphrase for the key????

##### Volumes #####

you can create a volume from an image, add some extra stuff to it and launch from that instance

you can:
Boot from an image
Create a volume from an image and boot from that
attach a volume to a running instance

##### Boot from ... when ... #####

image		don't want to store any state
volume	store state of vm between terminations

##### Security Groups #####

only expose a small number of your vms to ther outside world.
Create specific security groups for your vms
