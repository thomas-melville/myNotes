Content-Type: text/x-zim-wiki
Wiki-Format: zim 0.4
Creation-Date: 2016-08-30T11:11:37+01:00

###### Inheritance in Test Classes ######
Created Tuesday 30 August 2016

The use of inheritance to reuse common code goes against the idea behind inheritance in OOP, polymorphism blah blah blah
I came across strange behaviour by testng when there was a crazy class heirarchy and the testware owners weren't sure of their own class hierarchy.
We now have a rule which states that classes should only extend TafTestBase and all configuration should be delegated.

This adds some difficulties/complexities
1. Executing a single test locally on it's own
	a. You need to comment lines out of the suite.xml file.
2. Reusing common configuration code across classes.
	a. Your common configuration needs checks to see if it's already been executed.

@Before/After Suite
Separate class of configuration methods which is first in the suite.xml
What do I do when I want to execute one class locally? Commenting out stuff in the suite.xml file

@Before/After Class
What about when I want to reuse configuration across classes.

This is going to be a trade off

