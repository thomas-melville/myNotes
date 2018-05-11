# group management

all files fall into 3 categories

* owner
* group
* world

users belong to one or more groups

Purposes include:

* share work area
* file permissions to allow access but not to the world
* permit certain users to access resources they wouldn't otherwise be able to access

defined in **/etc/group**

each line in the file is as follows:

<group_name>:<password>:<group_id>:user1,user2,...

for group id values between 0 and 99 are system groups
100 and GID_MIN are considered special
anything over GID_MIN are for User Private Groups

a user has one primary group and can belong to 0 to 15 secondary groups
the primary group is the group used when a user creates a file

there can be a corresponding **/etc/gshadow** file for groups

## commands

* groupadd  add a new group
* groupdel  delete a group
* groupmod  modify a groups properties
* usermod   modify a users group membership (add or remove)
* groups    see group membership details

## User Private Groups

each user will have their own group
not guaranteed to be private, users may be added to the group in /etc/group

By default, useradd gives GID = UID and group name equals username
