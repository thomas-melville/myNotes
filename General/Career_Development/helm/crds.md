# helm & CRDs

https://github.com/helm/helm/blob/master/docs/chart_best_practices/custom_resource_definitions.md

The CRD must be registered on the cluster before any resources of that CRDs kind can be used.

Solutions:

1. separate charts

2. crd-install hook
  CRD will not be deleted when the chart is deleted
  this causes issues if the same chart is installed again
  in v2 only runs on install
    2.14.3 has fixes for upgrading charts with crds
  dropped in helm3, crd directory instead


CRDs are global on the cluster, what happens if another chart tries to registry the CRD?
charts installing different versions of CRDs
deleting a CRD deletes all the instances of it, CR
