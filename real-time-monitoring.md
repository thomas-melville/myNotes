# real time monitoring course

The usual speil about things change but they stay the same, new terms old practices.
  We won't get all the advantages of moving to microservices, ...

collect request counts bucketed by latency, instead of an average.
  Taking an average can hide spikes.

is this functionality better suited in supporting containers?
  cAdvisor, node-exporter, logging, ...

## graphite

graphite, real time graphing system using a time series database
  you need to push, it's not pull
  metrics pushed in plain text
  python pickle for large data
  metric names are heirarchial, like package names
  you can add tags to give some context.
  statsD, proxy for graphite to aggregate data before sending

## grafana

No datastore, only a visualization tool

it's pull

It has it's own concepts

datasource    where the data is coming from. (graphite, db, prometheus)
dashboard     main building block
widgets       a graph from a datasource
panel         or widget
              query editor to extract & display data from a single datasource
organizations add users to view/edit certain dashboards
query ed      how to retrieve the data from your data source
              specific to datasource
row           of widgets/panels
              repeated rows func is useful

you can put menus on dashbooards which can filter down into widgets
  they are called template variables

annotates point of interest in your data
playlists: rotate through dashboards

### dynamic dashbooards

drop down menu on dashboard

$ / [[]] in query to get variable, use name from dashboard configuration
nice little feature to get multiple selections of the variable to appear on separate panels.

can do a lot with custom time ranges

## prometheus

Data is stored as a time series stream of time stamped data
each is identified by it's metic name
labels are added to bits of data
  prometheus attaches two labels automatically
  job & instance

yaml file to configure where to scrape metrics from
it talks about service discovery so is that for where to pull metrics from?
there are client libraries to produce metrics
exporter applications available
  which sit beside your application. sidecar container
Java libraries for creating metrics to be exported
  spring-boot integration??? yup, spring-metrics
  you have to add logic to the code though.
  not actually with the spring-boot dependency!
    It looks to expose a lot, you just have to find them

don't overload with visualizing data from the outset.
Start with, did it work? when there's a failure what data do I need to troubleshoot This
put that in for next time
