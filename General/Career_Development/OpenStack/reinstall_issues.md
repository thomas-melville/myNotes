Content-Type: text/x-zim-wiki
Wiki-Format: zim 0.4
Creation-Date: 2016-12-22T15:32:36+00:00

###### reinstall issues ######
Created Thursday 22 December 2016

##### PhysicalResourceNameAbiguity #####

https://fem114-eiffel004.lmera.ericsson.se:8443/jenkins/view/08.ThunderBee_Test_Area/job/MT_Micro_ENM/210/console

u'stack_status_reason': u'Resource CREATE failed: PhysicalResourceNameAmbiguity: resources.mssnmpfm_interface1_port: Multiple physical resources were found with name (enm_security_group_enmcloud11).',

Cloud UI -> Project -> Compute -> Access & Security
You'll see 2 groups with the same name.
Delete both of them.

##### Gets stuck waiting for UI to come up #####
https://fem114-eiffel004.lmera.ericsson.se:8443/jenkins/view/08.ThunderBee_Test_Area/job/MT_Micro_ENM/212/console

5420 seconds until timeout

deploying...
5410 seconds until timeout

deploying...
5400 seconds until timeout

deploying...
5390 seconds until timeout

deploying...
5380 seconds until timeout
