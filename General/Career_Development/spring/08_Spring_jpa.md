# spring jpa

started off in hibernate and was extracted
It's an ORM
POJO based
built around best practices & patterns
pluggable persistance providers

## transaction management

https://www.baeldung.com/transaction-configuration-with-jpa-and-spring

two ways to configure transactions:

1. Annotations
2. AOP

### Annotations

You can enable transaction management with @EnableTransactionManagement on a @Configuration class.
In which you create beans for a LocalContainerEntityManagerFactoryBean and a PlatformTransactionManager.

In spring boot if you bring in a spring data dependency then this is done by default.

Once configured you can annotate classes/methods to make them transactional with @Transactional
This should be done on a Service, not on a Controller. Minimize the scope.

The annotation has some fields:

* Propagation Type
  * REQUIRED: create one if none exists
  * SUPPORTS: use one if it exists, otherwise don't
  * MANDATORY: use one if it exists, otherwise throw exception
  * REQUIRES_NEW: create a new one, suspend the current one
  * NOT_SUPPORTED: suspend the current transactrion
  * NEVER: execute without a transaction, throw an exception if one existsd.
* Isolation Level
  * DEFAULT: delegate to default in underlying data store.
  * READ_UNCOMMITTED: does not take changes made outside the transaction into account.
  * READ_COMMITED: prohibits reading a row with uncomiited changes.,
  * REPEATABLE_READ:
  * SERIALIZABLE
* timeout
  * the timeout value for the transaction
  * default to underlying system
* readOnly
  * state that the transaction is read only
  * allows for optimizations at runtime
* Rollback rules
  * what exceptions will/not trigger a rollback

### Potential Pitfalls

Spring creates proxies for all classes annotated with @Transactional. This allows the framework to inject the transaction logic before and after running the method.
If the transactional bean is implementing an interface, by default the proxy will be a Java Dynamic proxy.

Only public methods should be annotated, if any other methods are annotated the annotation will be ignored.

Read-only transactions are dependent on the underlying system supporting it, it will not fail if it doesn't.
I think it will just down grade to the next level of functionality (OCP prep)
