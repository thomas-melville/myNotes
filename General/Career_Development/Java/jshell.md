# jshell

interactive shell which allows you to experiment with snippets of java code

REPL for the java language

Read    type code
Eval    executes code
Print   print to console
Loop    repeat

all just plain java

## why?

quickly test out ideas
  no need for Main
  direct feedback with no compilation
explore new APIs
  code completion
great teaching tool
  no ide, just java

## usage

```bash

jshell
|  Welcome to JShell -- Version 11.0.2
|  For an introduction type: /help intro

jshell> 1+1
$1 ==> 2


jshell> /exit
| Goodbye

```

You can print any valid java expression in jshell

1+1
evaluates it to 2

Result of expression is given an identifier! So you can use it in later expression
Unless, it's a statement! Then there is no identifier

no need for semicolon, unless you want multiple line statements!

No need to catch checked exceptions, takes care of them automatically
  You can catch them if you want

ctrl-c interrupts a running thread in jshell

tab does a lot in jshell
  complete methods
  see overloaded methods
  see docs of each method

## regex in JsShell

can be usefule for testing regular expressions
  gives Error messages when there's an issue in the regex
  for example not double escaping \

## declarations

/help   show commands available
/set    change configuration
        for example increase "feedback"
/vars   list of specified variables in the jshell
        they are global and available in any created method

jshell will interpret certain characters to mean the use wants to add more code
for example { will be interpreted as the start of a method

/imports  list imported classes

you can use the normal import keyword to import a class

/methods  list availble methods

you can create a variable in a method which hasn't been created yet!
  verbose feedback will tell you that the variable hans't been declated yet

/types shows types created in this jshell

## jshell and the environment

/save <name>.jsh save the session to a file
  writes every statement as a separate snippet in the file
/open <name>.jsh open previously saved session
  evaluates all of the snippets in file
  could end up printing to console, interacting with files, ...
/open <class>.java
  you can open any source file
  no need to compile it

--class-path flag to jshell
  get access to files in jars outside base modules
  if jars files have dependencies you need to add those dependencies to the classpath!

## jshell API

a jshell in your own application
source code analysis
IDEs can make good use of this
  create jshell for currently open project including complete classpath

```java

JShell shell = JShell.create();

List<SnippetEvent> events = shell.eval("int i = 0;");

Stream<VarSnippet> vars = jshell.variables();

```
