Content-Type: text/x-zim-wiki
Wiki-Format: zim 0.4
Creation-Date: 2014-12-18T08:41:16+00:00

###### Fix broken Release ######
Created Thursday 18 December 2014

ISSUE:
Jenkins job all broken becuase they can't push to other branches

CAUSE:
Putting in a full version in the G sprint update job.

SOLUTION:
Run the sprint update job again with the next snapshot version
Jobs will still fail.
Merge Release into master, fix the conflicts and force push master


(08:15:09 AM) ethomev@wetalk.lmera.ericsson.se/5711663361418890362753467: Mornin' Could I get your expertise and user rights to help me fix the Jenkins jobs?
(08:28:47 AM) ethomev@wetalk.lmera.ericsson.se/5711663361418890362753467: I think I have to merge Release into master, fix all the conflicts and force push master?
(08:29:09 AM) taf-team@conference.wetalk.lmera.ericsson.se/JJ: :-)
(08:29:18 AM) taf-team@conference.wetalk.lmera.ericsson.se/JJ: I took a quick look last night before I went home
(08:29:25 AM) taf-team@conference.wetalk.lmera.ericsson.se/JJ: it's more than that
(08:29:45 AM) taf-team@conference.wetalk.lmera.ericsson.se/JJ: the version on the master branch is 2.5.1 
(08:29:47 AM) ethomev@wetalk.lmera.ericsson.se/5711663361418890362753467: master has to go in Release then
(08:29:53 AM) ethomev@wetalk.lmera.ericsson.se/5711663361418890362753467: Acceptance
(08:29:54 AM) taf-team@conference.wetalk.lmera.ericsson.se/JJ: it should be 2.5.1-SNAOSHOT
(08:29:55 AM) ethomev@wetalk.lmera.ericsson.se/5711663361418890362753467: not release
(08:30:03 AM) taf-team@conference.wetalk.lmera.ericsson.se/JJ: so things are a bit messed up
(08:30:07 AM) ethomev@wetalk.lmera.ericsson.se/5711663361418890362753467: they are indeed
(08:30:18 AM) taf-team@conference.wetalk.lmera.ericsson.se/JJ: merging release and master is one of the steps to sort it all out
(08:30:24 AM) ethomev@wetalk.lmera.ericsson.se/5711663361418890362753467: 2.5.1 is already in nexus so I should probably go to 2.5.2
(08:30:28 AM) taf-team@conference.wetalk.lmera.ericsson.se/JJ: I got side tracked this morning
(08:30:29 AM) ethomev@wetalk.lmera.ericsson.se/5711663361418890362753467: -SNAPSHOT
(08:30:41 AM) taf-team@conference.wetalk.lmera.ericsson.se/JJ: need to go to 2.5.2-SNAPSHOT
(08:30:48 AM) taf-team@conference.wetalk.lmera.ericsson.se/JJ: but all the poms have to be updated
(08:30:52 AM) ethomev@wetalk.lmera.ericsson.se/5711663361418890362753467: yup
(08:31:02 AM) ethomev@wetalk.lmera.ericsson.se/5711663361418890362753467: sure I can do the donkey work
(08:31:11 AM) ethomev@wetalk.lmera.ericsson.se/5711663361418890362753467: it was me who screwed it up :-)
(08:31:15 AM) taf-team@conference.wetalk.lmera.ericsson.se/JJ: ah no 
(08:31:32 AM) taf-team@conference.wetalk.lmera.ericsson.se/JJ: I left the opening in the jenkins job 
(08:31:41 AM) taf-team@conference.wetalk.lmera.ericsson.se/JJ: for the potiential to screw it up
(08:31:47 AM) ethomev@wetalk.lmera.ericsson.se/5711663361418890362753467: oh????
(08:31:58 AM) ethomev@wetalk.lmera.ericsson.se/5711663361418890362753467: Run sprint release with 2.5.2-SNAPSHOT?
(08:32:06 AM) taf-team@conference.wetalk.lmera.ericsson.se/JJ: do you want me to sort it seeing as I have push rights to the repo
(08:32:22 AM) taf-team@conference.wetalk.lmera.ericsson.se/JJ: yes to Run sprint release with 2.5.2-SNAPSHOT?
(08:32:30 AM) taf-team@conference.wetalk.lmera.ericsson.se/JJ: that would be the first step
(08:32:55 AM) ethomev@wetalk.lmera.ericsson.se/5711663361418890362753467: Sound, I'll kick that off now
(08:33:12 AM) taf-team@conference.wetalk.lmera.ericsson.se/JJ: then we need to sort out all the branches most likely with force pushes
(08:33:20 AM) taf-team@conference.wetalk.lmera.ericsson.se/JJ: so I guess I will have to do thatn
(08:33:21 AM) taf-team@conference.wetalk.lmera.ericsson.se/JJ: that
(08:34:01 AM) ethomev@wetalk.lmera.ericsson.se/5711663361418890362753467: Ok.
(08:34:25 AM) ethomev@wetalk.lmera.ericsson.se/5711663361418890362753467: If you're too busy and are happy to give me force push rights I could do it
(08:34:41 AM) ethomev@wetalk.lmera.ericsson.se/5711663361418890362753467: trying to multi task was the cause of the error yesterday ;-P
(08:34:49 AM) taf-team@conference.wetalk.lmera.ericsson.se/JJ: I would if I could but that's a gerrit thing
(08:34:55 AM) ethomev@wetalk.lmera.ericsson.se/5711663361418890362753467: ah right
(08:39:57 AM) ethomev@wetalk.lmera.ericsson.se/5711663361418890362753467: sort out branches # merge Release into master, fix conflicts, force push, then merge master into Acceptance, fix conflicts again and force push again?
(08:41:07 AM) taf-team@conference.wetalk.lmera.ericsson.se/JJ: I would start by updating master and then see if the regular jobs pass and can update Acceptance
(08:41:25 AM) taf-team@conference.wetalk.lmera.ericsson.se/JJ: I would leave Release branch and job till last
(08:41:37 AM) ethomev@wetalk.lmera.ericsson.se/5711663361418890362753467: right.
(08:41:42 AM) taf-team@conference.wetalk.lmera.ericsson.se/JJ: rest should be doable without force puses
(08:41:46 AM) taf-team@conference.wetalk.lmera.ericsson.se/JJ: pushes
(08:41:51 AM) taf-team@conference.wetalk.lmera.ericsson.se/JJ: I think 
