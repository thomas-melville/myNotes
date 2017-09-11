# UI acceptance testing using Graphene

Drone integrates Selenium with Arquillian
	Dependency injection using CDI
	manages selenium lifecycle

Graphene
	adds layer on selenium to simplify and add extra functionality for ui testing
	WebDriver in graphene wraps the Selenium WebDriver
	@Location annotation
		location of page relative to website root
	nice fluent api for waiting
	guardAjax
		because arquillian is managing the lifecycle it has access to the ajax requests by the browser?

Arquillian
	arquillian.xml

Chris works backwards when writing given, when, then

serverside
	is mocked
		using node.js & express
	or can be deployed in a docker container

docker ui environment is very useful for exploratory ui testing