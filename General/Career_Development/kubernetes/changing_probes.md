# changing probes

K8s uses probes to periodically perform diagnostic on a pod, if it's not healthy then take action
If a container in the pod fails a health check it can be re-started without re-creating the whole Pod

## types

### liveness

is the pod healthy and running as expected

### readiness

is the pod ready to receive requests

## actions

ExecAction - executes an action inside the container
TCPCockedAction - TCP check against the containers ip address & port
HTTPGetAction - HTTP Get request against the container

Successful
Fail
Unknown

changing the readiness probe to fail would be good for taking it out of traffic so an issue can be troubleshooted.

not finding much on line for the best practices for changing the response of the probe.

## options

1. create a rest endpoint which you use to set a flag which is used in the probe
2. update a file on the fs which is cat'd in the probe
3.
