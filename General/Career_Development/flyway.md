# Flyway

migration tool for database
snippets that change the database schema
often written in SQL
can be written in code also
each sql script is a migration which changes the schema somehow
  initially to create tables
  after a while to alter tables and migrate data

Flyway adds its own schema_version table to the schema

## options

sqlMigrationPrefix    default is V, you can set it to whatever you want
baselineOnMigrate     I think this means if a database already exists it won't start from the beginning. It will find out what is the baseline and start from there.
outOfOrder            will I also snippets to be applied out of order

## snippet file naming conventions

V{version}\__{name}.sql

version can be semantic, but underscores instead of dots
