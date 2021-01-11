# jpa

Annotate a Pojo to decide how it's persisted.
A specific persistance library does the work, for example hibernate
jpa is a specification with a list of annotations.

## annotations

@Entity
@Id
@Column

### Relationships

@OneToOne

Two entities have a one to one Relationship.
There is a parent and child, not sure of the correct terms here.

Parent: @OneToOne(mappedBy = "field name in child class")

@OneToMany
@ManyToOne
@JoinColumn

These annotations are used together
One on the parent class and one on the child class

The parent class field is a list of child classes, annotated with @OneToMany
It uses the mappedBy attribute to say what field on the child class to map on

The child class has a field of type parent class annotated with @ManyToOne
And the @JoinColumn which states what the column name is which will have the id of the parent

@ManyToMany
@JoinTable

One side uses the mappedBy attribute
The other side is more involved:

```java

@ManyToMany
   @JoinTable(name="BOOK_AUTHORS",
          joinColumns=@JoinColumn(name="BOOK_ID"),
          inverseJoinColumns=@JoinColumn(name="AUTHOR_ID"))

```

The table to hold the relations has to be specified
And the colums in the table which will hold the primary keys of the table.s

### lifecycle events

There are a number of annotations that enable logic to be executed at a certain point in an operation.
Pre/PostPersist,Update,Remove, ...
https://www.baeldung.com/jpa-entity-lifecycle-events#lifecycle

You can create methods on the class or you can create a class which contains the methods and annotations and link it to the entity class using ther @EntityListener annotation.

```java

public class UserEntityListener {

  @PrePersist
  public void PrePersist(User user){
    ...
  }
}

@EntityListener(UserEntityListener.class)
public class User {

  ...

}

```

## Entity lifecycle model

### persistence context

a set of all entity objects that you used in your current use case.
each of them represents a record in the database.

4 States in the model

1. transient
2. managed
3. removed
4. detached

### Transient

newly instantiated entity object.
has not been persisted yet, so does not represent any db record.

the persistence context does not know about this object, so it doesn't perform any SQL insert statement on the object.

### Managed

all entity objects connected to the persistence context are managed.
persistence provider will detect any changes and generate the required SQL statements when it flushes the context.

There are different ways to get entity into managed:

1. call EntityManager.persist
2. load an entity from the DB
3. You can merge a detached entity.

### Detached

an entity that was previously managed but is no longer attached to the persistence context.
it gets detached when you close the persistence context.
typically happens after a request gets processed.
database transcation gets committed, persistence context gets closed and entity object gets returned to caller.
you can programmatically detach an object by calling detach on the EntityManager

### Removed

when you call remove method on EntityManager the mapped db record does not get removed immediately.
state changes to removed.
During next flush, SQL delete statement is generated.
