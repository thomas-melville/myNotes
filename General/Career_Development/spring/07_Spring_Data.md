# spring data with JPA

https://app.pluralsight.com/library/courses/spring-data-jpa-getting-started/table-of-contents

builds on and enhances JPA
simplifies your data access Layer by reducing the code
Repository generator
There's a query DSL
  for inspecting your repository interface methods
auditing & paging
you can go around it when needs be

Repository + Entity = DAO

Entities go in model package
Repositories go in repository package

Add @EnableJpaRepositories to the main spring class to auto configure it.

@Repository is a course stereotype
The repositry idea allows you to switch out the actual data implementation without your service layer knowing
The repository sits between the database and the service layer.

A JpaRepository instance should map to a single Entity

When you implement a repository your public methods become part of the DAO contract.
You can hide impl details in private methods

JpaRepositories should always be interfaces:

```java
extend the JpaRepository<U, L> interface
  // you get a lot of operations by default

  // Query DSL
  // CRUD
  // paging & sorting
  // helpers

  // you can add your own methods to the interface and spring will try to create the sql from it
```
It uses generics to know what entity type to map to, the U above

**Recommendations**

* He recommends creating a super type with the audit information for each entity
* And also creating an interface for the audit operations which my other interfaces extend

**Best Practices**

packages: model & repository
  can limit jpa scanning to repository package
have repository in class name
leave jpa out, in case we want to change
no need for repository annotation when extending the jparepository interface
