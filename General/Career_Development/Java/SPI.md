Content-Type: text/x-zim-wiki
Wiki-Format: zim 0.4
Creation-Date: 2017-08-16T14:51:06+01:00

###### SPI ######
Created Wednesday 16 August 2017

Service Provider Interface
Not really specific to java but don't know where else to put it

This allows an interface to be implemented / implementation used overridden by a third party

src/main/resources/META-INF/services/<fqn_of_interface>

contents of file should be fqn of implementation

fqn # fully qwualified name # package + class_name
