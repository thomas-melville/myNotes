# user account mgmt

## user accounts

linux sys provides a multi-user env which permits people and processes to have separate simultaneous working envs
purpose of individual accounts:

* each user has their own individual private space
* creating particular user accounts for particular purposes
* distinguishing privileges among users.

**root** is a special user account, who is able to do anything. should only be used when absolutely necessary.
normal user accounts are for people who will work on the system
there are other special user accounts which allow processes to run as a user other than root.

* daemon is an example

### attributes

each user has a corresponding line in **/etc/passwd** that describes their basic account attributes

<username>:<user_password>:<user_id_number>:<group_id_number>:<comment>:<home_directory>:<login_shell>

if /etc/shadow is used then the user_password filed holds an x

the convention is any user id below 1000 is special and belongs to the system

each user also has a corresponding line in **/etc/shadow** that contains information
/etc/shadow enables password aging on a per user basis and maintains greater security of hashed passwords

/etc/passwd has permissions of 644, everyone can read but only the owner, root, can edit
/etc/shadow has permissions of 400, which means only root can access the file

<username>:<hashed_password>:<last_change>:<mindays>:<maxdays>:<warn>:<grace>:<expire>:<reserved>

* last change: days since Jan 1 1970 that the password was last changed
* mindays: min days before password can be changed
* maxdays: max days after which the password must be changed
* warn: days before expiry that the user is warned
* grace: days after expiry that account is disabled

order of /etc/passwd and /etc/shadow must match exactly

### creating a user account

sudo useradd <username>

* next available UID is assigned to the user
* a group with the username is created and GID is assigned the value of the UID
* * these are called User Private Groups
* home directory is created. **On Ubuntu and OpenSUSE this doesn't happen by default, you should specufy the -m flag**
* contents of /etc/skel are copied to home directory
* an entry !! / ! is placed in the password field of the /etc/shadow file 

defaults are contained in /etc/default/useradd

### removing a user account

sudo userdel <username>

* all references to username are deleted from /etc/ config files
* home directory is not removed
* * to delete add the -r flag

### modify a user account

sudo usermod <username>

### locked accounts

linux ships with some sys accounts which are locked
this means they can run programs but never login
/sbin/login for shell in /etc/passwd file achieves this

you can lock an account using the -L flag in usermod

another way is to use chage to change the expiration date to a date in the past

yet another way is to edit the /etc/shadow file and replace hashed password with !!

### password management

**passwd** command is used

normal user can change only their own password, root can change any password
password choice is examined by pam_cracklib.so
    this checks the password for strength

**chage** command is used to set password aging

### restrictions

#### restricted shell

bash -r restricts what a user can do in a shell

* can't cd
* can't redefine certain env vars
* can't specify absolute path of executables
* prevents the user from redirecting inp/output

#### restricted accounts

a user account can be given a restricted shell
limit system programs and user apps
limit system resources
limit access time
limit access location

### root account

should only be user for administrative purposes
never use as a regular account
by default, root logins are prohibited

ssh as root is configured in /etc/ssh/sshd_config and PAM (Pluggable Authentication Module)

generally recommended that all root access is through su or sudo, so that there is an audit trail of who did it

## ssh

Secure SHell
one often needs to login through the network to a remote system
uses encryptions based on strong algorithms

ssh can be configured on your machine to login machines without a password.
user specific config files are stored in $HOME/.ssh

* id_rsa            users private encryption key
* id_rsa.pub        users public encryption key
* authorized_keys   a list of public keys that are permitted to login
* known_hosts       a list of hosts from which logins have been allowed in the past
* config            file for specifying various config options

private and public keys generated with ssh-keygen

the public key can be given to any machine with which you want to permit password-less access

## scp

securely/ssh copy files

## remote graphical login

VNC (Virtual Network Computing) used
a common impl is tigervnc
I use a Google VNC Viewer