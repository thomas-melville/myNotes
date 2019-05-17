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
