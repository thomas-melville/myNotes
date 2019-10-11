# java in a container

## PID 1 problem

Signals are used to get running processes to behave in certain ways.
Some can be caught and give possibilities to implement graceful shutdown logic.
A common case is to send SIGTERM to terminate a process nicely

kill <pid> sends a SIGTERM
when you have a process in the foreground and you hit ctrl + c you send SIGINT

jvm gives exit code 130 when we ctrl + c a running jvm
143 for kill <pid>

jvm can't handle KILL signal, you'll get an exception if you try to handle it
jvm gives 137 when you use the KILL signal

Use the *Signal* API to set up handlers for different signals

docker kill by default sends the KILL signal
  use the -s argument to set the signal to give

docker stop will send TERM signal
  and give your container 10 seconds by default to shutdown
  otherwise it gets the KILL signal

when using CMD in a Dockerfile, that command gets pid 1
If some logic is required before starting the process it gets put inside an entrypoint script
  The script then gets pid 1 and the long running process gets some other pid
  The long running process will no longer get the kill signals
  prepend the long running process command in the entrpoint script with *exec*
    exec will replace the process in the pid with this process

What if you don't want to be pid 1?
  you can run a tiny init process which deals with these things for you
  add --init in your docker run command
  does k8s do this?
