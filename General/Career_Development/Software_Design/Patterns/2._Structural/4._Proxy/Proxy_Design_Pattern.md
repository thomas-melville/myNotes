Content-Type: text/x-zim-wiki
Wiki-Format: zim 0.4
Creation-Date: 2015-07-14T08:13:40+01:00

###### Proxy Design Pattern ######
Created Tuesday 14 July 2015

UI static block
	creates an instance of UiWindowProviderRegistry
		UiWindowProviderRegistry constructor uses ServiceLoader to load implementations of UiWindowProviderFactory, SeleniumUiWindowProviderFactory
		Here the UiWindowProviderProxy is setup, UiWindowProviderInvocationHandler is added to it and returned.
Browser.open 
	WindowProvider.get
	Proxy is called which invokes the wanted method on UiWindowProvider
	Goes through all the code of opening the window on the selenium side
	Creates an instance of BrowserTabImpl
	Creates a proxy for BrowserTab
		All calls to the browserTab instance go through this proxy!
		uses the BrowserTabImpl instance.
		Switch the windowProvider to this browserTabImpl instance
		invoke whatever method was called on the proxy


SeleniumUiWindowProvider.createNewTab
	Anytime a method is called on an instance of BrowserTab it goes through this proxy.
	(http://docs.oracle.com/javase/7/docs/api/java/lang/reflect/Proxy.html)
