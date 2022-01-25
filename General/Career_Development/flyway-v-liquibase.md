# flyway V liquibase

## liquibase

Integrates with Spring-Boot.
Same as flyway, just add the dependency and put the migration files in the expected location.

By default migrations are written in XML
But it looks like a migration can be written in Java: https://docs.liquibase.com/change-types/community/custom-change.html

### Rollback

What to execute on rollback can be auto generated per change set.
However, you can add your own into each change set

I'm not yet sure when the rollback commands get executed. I think it's a manual step using either the liquibase cli or maven plugin
