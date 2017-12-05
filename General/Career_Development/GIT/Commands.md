# Commands

git commands are translated by git in git-<command> and it then looks for this script
so... you can create a script name it git-<command> and place it on the path
when you type "git <command>" git will look for this script on the path

## rebase
in interactive mode there is another option: pick, squash **and fixup**
fixup squashes the commit into the commit above it.

## pull
can take an argument, -rebase, which tells it to rebase instead of merge

## add
you can add chunks of files to the staging area interactively, https://git-scm.com/book/en/v2/Git-Tools-Interactive-Staging
git add -i
	this allows you to stage whole or parts of files to add to the commit
	making your commits smaller and more focused.
	
**--patch is an argument to lots of GIT commands to allow you to decide what changes to work with**

## rev-list

list commits in reverse chronological order
has useful arguments for skipping and max count

## describe

describe a commit using the most recent tag reachable from it