# submodules

https://git-scm.com/book/en/v2/Git-Tools-Submodules

Keep a Git repository as a subdirectory of another Git repository

Useful when ...

    you want to use another library but ...

        the deployment/dependency mechanism isn't straight forward
        or you want to modify it and the contribution process is slow

## adding a submodule

git submodule add <absolute / relative url>

.gitmodule file is created
this file needs to be version controlled in the top level git repo
git status will show the directory of the new submodule, but won't show the commits
as it knows it's a submodule
add the submodule to the commit also
it goes in with a special mode

## cloning a project with submodules

when you clone a project with submodules you'll get the submodule directories but no files in the directories.
you must run 2 commands to get the files

```bash

git submodule init
git submodule update

```

alternatively you can pass --recurse-submodules to clone command

## keeping your version of the submodule up to date with it's master

cd into submodule directory and do a git pull

or from top level directory do

git submodule update --remote

## contributing to submodule from submodule