# File Permissions

in ls:

first bit is file type:
	d directory
	- regular file
	l link


In directory permissions there is r,w,x and s

**suid**

s is a special type of file permission, suid (Set owner User ID up on execution)
	when a user executes the file with s it uses the access permissions of the owner rather than the executor

**sgid**
s is a special type of file permission, suid (Set Group ID up on execution)
	when a user executes the file with it uses the access permissions of the group rather than the executor


**sticky bit**
prevent a user from deleting another users files even if they would normally have permissions to do so.
Set on a directory, t character at the end of 