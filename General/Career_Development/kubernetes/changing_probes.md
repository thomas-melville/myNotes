# changing probes

changing the readiness probe to fail would be good for taking it out of traffic so an issue can be troubleshooted.

not finding much on line for the best practices for changing the response of the probe.

## options

1. create a rest endpoint which you use to set a flag which is used in the probe
2. update a file on the fs which is cat'd in the probe
3.
