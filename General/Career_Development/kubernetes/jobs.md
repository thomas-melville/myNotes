# jobs

for something that runs once, or on a recurring basis

A job creates one or pods and ensures that a specified number of them successfully terminate

A job can run multiple pods in parallel
  split the work into multiple pods to increase throughput

names of cronjobs must be less than 52 characters

```yaml

apiVersion: batch/v1
kind: Job
...
spec:
  completions: 4 # 4 pods must complete successfully
  parallelism: 2 # 2 pods can run in parallel
  template:
    ...
    spec:
      restartPolicy: Never/OnFailure
      containers:
        ...

```



## cronjob

use cron format or @hourly, @daily, @weekly, @monthly, @yearly

\*/5 every 5 units (minutes, hours, ...)

https://crontab.guru/

```yaml


apiVersion: batch/v1beta1
kind: CronJob
...
jobTemplate:
  spec:
    concurrencyPolicy: Allow # allow job execution to overlap.
    schedule: "..."
    ...

```
