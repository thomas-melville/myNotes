# Shit! I forgot to put a change in a previous commit ######

**N.B. This is only if the commits aren't on origin master**

git add <file>
	add the file/change to be tracked
git commit --fixup HEAD~1 
	This gives the commit the same subject line as the commit at HEAD~1 **prepended with "fixup!"**
	This commit is temporary as it will be squashed up into the commit you want to add to
git rebase -i --autosquash HEAD~<index_of_commit+1>
	this command orders the rebase of commits
	the commit at HEAD~<index_of_commit> is first
	followed by any commits which start with squash! or fixup!, ...
save and exit the editor

--autosquash can be set to true in the gitconfig file!!!!

be very careful in the git rebase!!!
	make sure the commit you want to fixup into is first in the list of commits!!!!!