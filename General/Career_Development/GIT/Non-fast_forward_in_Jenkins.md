# Non-fast forward in Jenkins

If a job that pushes code back to master from Release fails then you'll end up with a non-fast forward change

You need +2 rights to do this.

Make sure you have the latest code on the master and Release branch
checkout master
git pull

checkout Release
git pull

checkout master
git merge Release
git push
