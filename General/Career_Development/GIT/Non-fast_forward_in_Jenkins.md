Content-Type: text/x-zim-wiki
Wiki-Format: zim 0.4
Creation-Date: 2015-10-27T15:59:55+00:00

###### Non-fast forward in Jenkins ######
Created Tuesday 27 October 2015

If a job that pushes code back to master from Release fails then you'll end up with a non-fast forward change

You need +2 rights to do this.

Make sure you have the latest code on the master and Release branch
checkout master
git pull
checkout Release git pull

checkout master
git merge Release
git push
