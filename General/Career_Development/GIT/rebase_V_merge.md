Content-Type: text/x-zim-wiki
Wiki-Format: zim 0.4
Creation-Date: 2017-04-19T16:34:22+01:00

###### rebase V merge ######
Created Wednesday 19 April 2017

They solve the same problem, but in very different ways
	integrate changes from one branch into another.
	
merge
	a mege commit is created which ties the tip of both branches together
	In the history the commits from each branch are mixed up
	
	non-destructive operation
	adds extra commits though
	
	use merge in public branches!
	
rebase
	the commits in one branch are taken and put on top of the commits in the other branch
	In the history, the commits from each branch are separated.
	
	re-write history
	cleaner history
	
	use rebase in private branches
