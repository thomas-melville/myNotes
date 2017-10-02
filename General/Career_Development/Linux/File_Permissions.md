# File Permissions


## ls command output

first character reveals the type:

- regular file
d directory
l symbolic link

the next 9 characters are the permissions for user, group and other

r 	read the file 		allow file names in the directory to be read
w 	write to the file 	allow the entries to be modified within the directory
x 	execute the file 	allow access to the contents and metadata for entries

user	u
group	g
other	o
all 	a

These don't show up in ls output but can be used to change permissions

**There may be an extra character at the end of the permissions string**
This means an alternative access control method has been applied

. SELinux (Security Enhanced Linux) security context
+ Access Control Lists

**suid**
s is a special type of file permission, suid (Set owner User ID up on execution)
	when a user executes the file with s it uses the access permissions of the owner rather than the executor

**sgid**
s is a special type of file permission, suid (Set Group ID up on execution)
	when a user executes the file with it uses the access permissions of the group rather than the executor

**sticky bit**
prevent a user from deleting another users files even if they would normally have permissions to do so.
Set on a directory, t character at the end of 

## groups

every user is a member of at least one group, called their primary group

'groups' will show the groups you belong to, add in a userid to see the groups they belong to

## changing permissions

chmod mode file

2 ways to specify mode:

### symbolic mode

chmod user_category operator permission

user category 	: ugoa
operator		: +-=  add, substract, set
permission 		: rwx

### octal mode
3 sets of octal numbers

rwx
000	0
001	1
010	2
011	3
100	4
101	5
110	6
111	7

chgrp

## file creation mask

umask

determines access file will be given on it's creation
when setting mask, numbers apply the opposite affect to chgrp/chmod