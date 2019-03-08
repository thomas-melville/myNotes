# Objects

## Pod

ReadinessProbe (ensure the app is ready to recieve traffic)
  exec
  httpGet
  tcpSocket

LivenessProbe (check the health of the container, terminate it if it's not. A controller should bring up another one)

Resources (define what the max the pod require are, what happens if they exceed it is not specific)
  limits.cpu/memory
  requests.cpu/memory

  limits/requests.ephemeral-storage (temporary storage the pod needs, for the likes of log file before they are shipped)
