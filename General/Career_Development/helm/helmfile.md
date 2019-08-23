# helmfile

https://github.com/roboll/helmfile

manage many releases in one cluster with a single file and command.
specify the charts in a yaml file, with the path to the associated values file, with a specific cluster, and then sync them in a single command

Not at version 1 yet!

Integrates with kustomize!

```yaml

context: eks_staging

releases:
- chart: charts/apps/event-processor
  name: staging-event-processor
  values:
  - chart/sapps/event-processor/values/staging/values.yaml

- chart: charts/apps/delivery-manager
  name: staging-delivery-manager
  values:
  - chart/sapps/delivery-manager/values/staging/values.yaml

```

the chart can reference a url too
