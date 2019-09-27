# commands

## upgrade

upgrade a release to a specified version of a chart and/or updates chart values.

### mandatory arguments

release
chart
  url
  directory
  package

### optional arguments

values/f
set/set-string/set-file

reuse-values, to edit or append to the existing customized values
  this causes any non-customized values from the original values.yaml to override any in the new chart.
  for example, we have image tag in the values.yaml so the old one is taken.
  to make this work we would need to extract the values.yaml file from the chart and pass it on the command line.
  Can also cause issues when a value is reused which no longer exists in the new chart

If no chart value arguments are provided on the cli, any existing customized values are carried forward.

reset-values, to revert to the values provided in the chart

If a value is specified multiple times, the last one of the cli takes precedence

Special characters in the value side of can cause issues on the cli, use single characters and \ escape any special characters

## get

get details about a named release
by default it gives the latest revision, you can get a previous revision by adding the --revision argument and using the number from the helm history command

hooks, manifest, notes, values

the values subcommand by default just prints the values passed in
If you want to see all values add the --all flag
