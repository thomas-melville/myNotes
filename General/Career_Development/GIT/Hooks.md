Content-Type: text/x-zim-wiki
Wiki-Format: zim 0.4
Creation-Date: 2017-08-10T13:19:07+01:00

###### Hooks ######
Created Thursday 10 August 2017

https://git-scm.com/docs/githooks

scripts with specific names that execute at certains points in the git process.
I know they can be in bash or python, **are there other languages you can use?**
	Any properly named executable script will work fine


2 types: client-side & server-side

we have no control over server-side so I'll stick with client-side

Client-side hooks go into the hooks subdirectory of the .git directory of your repository.
	**Is there any way to have the hooks centralized on your machine?**

Hooks:
There's a lot so I'm just going to list the ones I think would be useful.

pre-commit
	runs when you execute commit before the dialog opens to enter the commit message
	could be used for some static analysis checks, linting
prepare-commit-msg
	runs before the commit message editor opens.
	It gives the type of commit so you could have different messages depending on the type, merge, rebase, ...
	could be used to populate default commit message
commit-msg
	runs after you save the commit
	could be used to enforce the format of the commit message or project state
post-commit
	runs after you save the commit message
	could be used for notifications are commiting
pre-rebase
	runs when you execute the rebase commands, but before the rebase actually happens
	could be used to make sure the project is in a good state to rebase
post-checkout
	runs after you do a checkout
	could be used to set up the working directory properly
pre-push
	runs during a push, after remote refs have been updated but before any objects are transferred



