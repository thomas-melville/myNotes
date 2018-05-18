# pam

Pluggable Authentication Modules

allows Authentication to be done in one uniform way, libpam*

It's a library of Modules which provides enormous flexibility and consistency with respect to authentication, session, password and account services

PAM incorporates the following components

* PAM-aware apps (login, ssh, su)
* Configuration files in /etc/pam.d/
* PAM modules in the libpam* libraries

Several steps involved:

* user invokes pam aware app
* app calls libpam*
* library checks for files in /etc/pam.d/
* * these delineate which PAM modules to invoked
* each referenced module is executed in accordance with the rules

each file in /etc/pam.d/ corresponds to a service
each line in each file corresponds to a rule, excluding commented out lines

<type> <control> <module-path> <module-arguments>

type specifies the mgmt group the module is to be associated with

* auth      prompt the for id
* account   checks on aspects of the users account
* password  resp for updating the user auth token
* session   provide functions before/after the session is established

control controls how the success/failure of a module affects the overall auth process.

* required      must return success for access to be granted
* requisite     same as required, just where it fails is different
* optional      module is not required
* sufficient    if this module succeeds no other modules are executed.

module-path gives the file name of the library

module-arguments can be given to modify the behaviour of an individual module.