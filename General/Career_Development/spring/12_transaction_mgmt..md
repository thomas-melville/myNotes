# transaction mgmt

https://docs.spring.io/spring/docs/5.0.x/spring-framework-reference/data-access.html#spring-data-tier

Provides consistent abstraction for transaction mgmt across different transaction APIs: JTA, JDBC, Hibernate, JPA, ...

Spring provides both declarative and programmatic transaction models.
*declarative is recommended in most cases.*

The key to Springs transaction abstraction is the notion of a transaction strategy.
It is defined by the *PlatformTransactionManager* interface.
This is primarily an SPI, although it can be used programmatically from your application code.
Implementations are defined like any other object/bean in the Spring IoC Container.

*TransactionException* is unchecked and can be thrown by any of the *PlatformTransactionManager* interface methods.
Transaction infrastructurte failures are invariably fatal.

The *getTranscation* method returns a *TransactionStatus* object, depending on the *TransactionDefinition* parameter.
The returned *TransactionStatus* might represent a new transaction, or an existing transaction if one is present on the call stack.

The *TransactionDefinition* interface specifies:

* Isolation
  * The degree to which this transaction is isolated from the work of other transactions
* Propagation
  * inherit the existing transaction, create a new one or execute outside a transaction
* Timeout
  * length of time until the transaction times out and rollback
* Read-only status
  * a transaction which only reads data
  * no writes

The *TransactionStatus* interface provides a simple way for transactional code to control transaction execution and query status.

Defining the correct *PlatformTransactionManager* implementation is essential.
It is defined through dependency injection.
Implementations require knowledge of the env in which they work, JDBC, Hibernate, JTA, ...

https://docs.spring.io/spring/docs/5.0.x/spring-framework-reference/data-access.html#spring-data-tier
