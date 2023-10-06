# postgres

Open source, general purpose and object-relational database system
The most advanced open source database system
originated in Berkley
designed to run on unix like platforms, but is portable
SQL based

## feature highlights

* user-defined types
* Table Inheritance
* Sophisticated locking mechanism
* Foreign Key referential integrity
* Views, rules, subquery
* Nested transactions
* multi-version concurrency control
* asynchronous replication

you can add custom functions developer in a number of languages

## using it

psql command line tool to connect to a database instance

\?              list the possible commands
\c [DB_NAME]    connect to database
\q              quit the connection

It's SQL based, the syntax is very similar to what I'm used to in MYSQL

I'm sure there's some fancy stuff it can do but I'll leave that until I need it.

## Data types

String
Integer
Date Time
and many more ...

## security

authentication, authorization and encryption.

username and password unencrypted
username and password encrypted
client certificate encrypted.

everything is passed in the jdbc url: https://jdbc.postgresql.org/documentation/use/

username, password.

For encryption, using SSL use the *sslMode* parameter
It can be:
  disable, allow, prefer, require, verify-ca, verify-full

For one way TLS, i.e. client verifies identity of the server, the CA of the server is required by the client.
Pass the path to the CA cert using the *sslrootcert* parameter.

The server also has to be configured with certificates and keys.

/var/lib/postgresql/data/pg
  postgresl.conf has a number of ssl parameters
  pg_hba.conf configures the security method on the server side.
    https://www.postgresql.org/docs/current/auth-pg-hba-conf.html

For two way TLS, i.e. client verifies identity of the server, then the server verifies the identity of the client, the client needs to provide it's cert and key.
*sslcert* and *sslkey* are the parameters to use to pass the paths to these files.
Server will need access to the CA that signed the client certificate,
