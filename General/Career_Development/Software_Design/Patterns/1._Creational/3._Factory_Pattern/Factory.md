Content-Type: text/x-zim-wiki
Wiki-Format: zim 0.4
Creation-Date: 2015-08-23T19:18:52+01:00

###### Factory ######
Created Sunday 23 August 2015

Factory Method pattern:
	Define an interface for creating an object.
	Let subclasses decide which class to instaniate
	Defer instantiation to subclass
	
Creator -> Product

+Ives
	Easily extendable to new subclasses
	Eliminates binding app specific classes in your code.
		code only deals with interfaces, any user specific subclasses can be used.
	More flexible than creating an object directly.
	Gives subclasses a hook for providing an extended version of an object.


-Ives
	Must subclass 2 classes to create their own specific impl
	Object interface extensions could be problematic depending on locations of impl's
	
2 flavours of Factories
	Interface which leaves all implemenation details to subclasses.

	Abstract which returns a default implementation of the object.

Interfaces:
	UiWindowProviderFactory
		newInstance
		getType
	UiWindowProvider
		all window interactions
		get, switch, close, ...
		
Implementations:
	SeleniumUiWindowProviderFactory
	SikuliUiWindowProviderFactory
	
	Selenium...
	Sikuli...
