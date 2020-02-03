# Flyway

migration tool for database
snippets that change the database schema
often written in SQL
can be written in code also
Each migration can be a different language! As long as the versioning is done correctly.
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

## migrations

### sql

for any database

### java

Must implement the JavaMigration interface. Or for convenience extend the BaseJavaMigration

By default don't have a checksum so don't take part in the change detection of Flyways validation.
  To include it implement the getChecksum() method.

You can use the java.sql APIs, PreparedStatement, ResultSet, etc.
Don't close the DB connection

If you use Spring you can use the JDBC template instead of JDBC directly.
