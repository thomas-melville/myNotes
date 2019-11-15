# junit 5

## order test method execution

https://www.mkyong.com/junit5/junit-5-test-execution-order/

annotation on class
value is a class of ordering from MethodOrderer

```java

@TestMethodOrder(MethodOrderer.Alphanumeric.class)
public class MyTest{

}

```

Orderings:

* Alphanumeric
* Define the order on the tests
  * @Order(int) on each test method
  * no annotation is presumed last
* Random
  * It can be random for each iteration
  * or you can define a seed to have the same random execution.
* Custom
  * You can implement *MethodOrderer* interface
  * override orderMethods to define your own ordering
    * Takes a MethodOrdererContext
    * from which you can get a lot of information about each method.
