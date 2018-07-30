# gitflow

branching model presented by Vincent Driessen

There is a repo which is a set of git subcommands built for git-flow: https://github.com/nvie/gitflow
these just make it easier to follow this model by giving one word commands instead of the individual git commands

**What about the review process??????**
ECM are using GitFlow & Gerrit, working quite well according to Santhosh

## branches

### master

The HEAD of this branch always reflects a production ready state.
Once all the commits from the develop branch which are on the release branch are merged back into master this is considered a release and tagged as such
this enables the possibility of a git hook script on the server to deploy the new release to a production server.

### develop

The HEAD of this branch always reflects the latest delivered changes for the next release
Nightly automatic builds are run from this, it can be thought of as a staging area for the latest code
eventually all commits on this branch make there way back onto master.

### feature/topic branches

may branch off the develop branch
must merge back into develop

* when merging back in use the --no-ff flag (no fast forward)
* always create a new commit object on the develop branch, even if the commit could be done with a fast forward
* this keeps history of the feature and groups together all the commits that together make the feature
* this can be seen using my alias **git lds**

naming convention can be anything but master, develop, release-*, hotfix-*

are used to develop new features for the upcoming or distant release

typically exist in developer repos, and not in the origin
**I would take issue with this, I would push all branches to origin**

### release branches

may branch off the develop branch
must merge back into develop and master
naming convention: release-*

used to support preparation of a new production release
allow for last minute tidy up, polishing, last minute bugfixes and preparing meta-data for release.
It enables a deploy -> test -> fix -> redeploy cycle without affecting the develop branch
By doing all this on the release branch the develop branch can still continue with the latest features

When to branch release from develop: When the develop branch almost reflects the desired state of the new release.
At the start of the release branch the version number gets updated

only bug fixes can be applied to this branch

when the release branch is ready it is merged into master, using the --no-ff option again
the merge commit is then tagged
it then needs to be merged back into develop, so that any commit aren't lost
this will probably result in a merge conflict, fix it and merge.
Then delete the release branch as we don't need it anymore

### hotfix branches

may branch off from master
must merge back into master and develop
naming convention: hotfix-*

very much like release branches in that they are in prep for the next production release.
However, they are unplanned and a reaction to a critical issue that must be fixed immediately.
They allow the developers to continue unhindered on the next releases features.

don't forget to bump the version after you've checked out the branch

when the bug is fixed merge it into master and develop with the --no-ff tag
If there is a release branch, it should be merged into that instead of the develop branch. It'll eventually make it back into the develop branch

## versioning

https://www.fredonism.com/a-practical-take-on-gitflow-and-semantic-versioning

bump the version when branching from:

* develop to release
* master to hotfix

Follow the semantic versioning rules, https://semver.org/