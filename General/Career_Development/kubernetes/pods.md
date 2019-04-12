# pods

## phases of lifecycle

* Pending
  * Accepted but no container created.
  * This includes waiting for system to pick it up and downloading images over the network
* Init: N/M
* Init: Error
* Init: CrashLoopBackOff
* Running
* Succeeded
* Failed
* Unknown

https://confluence.lmera.ericsson.se/display/AA/Approach+for+ADP+microservice+LCM?focusedCommentId=145454279#comment-145454279
