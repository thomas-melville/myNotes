# Commands

git commands are translated by git into git-<command> and it then looks for this script
so... you can create a script name it git-<command> and place it on the path
when you type "git <command>" git will look for this script on the path

## rebase
in interactive mode there is another option: pick, squash **and fixup**
fixup squashes the commit into the commit above it.

before your branch is merged with master the history is malleable.
Make the commit history a story about the work.

## pull
can take an argument, -rebase, which tells it to rebase instead of merge
rebase on pull can also be set in your git config so you don't have to specify it every time.

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

## log

grep
	you can grep the contents of the commit messages

	git log --grep='what am I looking for'

	-G 'regex'

pickaxe feature
	search for text in the changes in each commit, the text can be added or removed

	git log -S <text to search for>

--stat
	what changed in each commit

-c
	the actual changes

## reflog

history of local HEAD

## blame

find out in what commit each line of code was changed.
(must be what Intellij uses for annotate)

## stash

temporarily store uncommitted changes
stash actually creates a commit.

### save

git stash save "description"

stash the current uncommitted changes
**Deprecated in favour of git stash push**

-u / --include-untracked: stash un-tracked files too

### push

git stash push "description"

stash the current uncommitted changes

-u include un-tracked
-p select what to store in the stash

### list

git stash list

List the current stashed changes

### apply

git stash apply <stash>

apply the stashed change to the working tree
keep the stash in the list

### pop

git stash pop <stash>

apply the stashed change to the working tree
delete it from the list

### show

git stash show <stash>

show the files changed in the stash

-p show all the changes within the files

### branch

git stash branch <name> <stash>

creates a new branch and pops the stash onto it

### clear

git stash clear

delete all stashes

### drop

git stash drop <stash>

delete the specified stash