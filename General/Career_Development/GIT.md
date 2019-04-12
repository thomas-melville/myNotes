# GIT

https://codewords.recurse.com/issues/two/git-from-the-inside-out

.git folder
	objects
		folder which contains blobs of compressed contents of files at particular points in time.
	index
		a list which contains every file git has been told to keep track of.
		maps a tracked file to the hash of its contents, commit message & parent hash at the moment it was committed.
	refs
		a folder which contains refs & tags
		refs are files which contain hashes of commits
		tags are human readable refs


## commit command has 3 steps.
1. Creates a tree graph
2. Creates a commit object
3. Points the current branch at the new commit object

## Tree graph
	each commit has it's own tree graph
	This is the state of the project for this commit
	records location and content of every file in the project
	2 types of objects:
	1. blobs. Already covered.
	2. trees. represents a directory in the working tree

## Commit object
	points at the tree graph
		the hash is for the tree object that represents the root of the repo.
		This is linked to all other blobs and trees at that point in time.
		If the blob tree didn't change then the new tree points to the obj from the previous commit
	If it's not the first commit it points to it's parent tree object, i.e. the tree graph of the previous commit
	the last line is the commit message


## Point current branch at the new commit
	move HEAD to point at the new commit
	HEAD is a ref, a ref is a label used by git/user to identify specific commit.
		Think release labels
	HEAD points to refs/heads/master which points to the commit that was just made.


### working copy V index#
	index, & HEAD use hashes to refer to blobs
	working copy is text stored in a different place
	once you make a change to a file the working copy & index diverge until you make a commit
	index refers to the staged area.


### Graph properties#
	content is stored as a tree of objects -> only diffs are stored in objects database.
		if blob/tree hasn't change in commit the new tree references the existing blob/tree
	each commit has a parent
		-> stores history
	refs are entries points to particular points in the history
		-> can give meaningful names to points in time.
		GIT uses symbolic references for commands that manipulate the commit history
	nodes in objects/ are immutable.
		-> content is edited, not deleted. entire history of repo is in objects folder
	refs are mutable
		The meaning of a ref can change
	recent history is easier to recall but also changes more often.
		Working copy constantly changes.
		HEAD changes most often next.
		Refs are easily accessible
		The further you go from refs the harder it is to recall the history.


### checkout a commit has 4 steps#
1. get the commit and the tree graph it points at.
2. write the file entries in the tree graph to the working copy
3. write the file entries in the tree graph to the index
4. the content of HEAD is set to the hash of a2 commit

When checking out a commit HEAD goes into detached state.
Any commits made can easily be lost.

### Create a branch#
a new file in refs/heads/ is created that contains the hash that HEAD is pointing at.

when checking out a branch git follows the same steps as checking out a commit.
A branch is just a file in refs/heads/ which contains the hash of a commit.
