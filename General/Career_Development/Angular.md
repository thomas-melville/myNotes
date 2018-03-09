# Angular

started it's life in google, angularjs.org.

plnkr.co
    in browser angularjs site editor

Is an MVW framework in javascript which aims to simplify UI creation

MVW = Model - View - Whatever works for you!

So it could be MVC, MVVM, MVP

## Concepts

Template:       HTML with additional markup
Directives:     extend HTML with custom attributes and elements
Model:          data shown to the user in the view and with which the user interacts
Scope:          context where the model is stored so that controllers, expressions and directives can access it.
Expressions:    access variables and functions from the scope
Compiler:       parses the template and instantiates directives and expressions
Filter:         formats the value of an expression for display to the user
View:           What the user sees, DOM
Data Binding:   sync data between model and the view
Controller:     the business logic behind the views

## implementation

script tag in the index.html which points to location of angular.js file
ng-app attribute in element to tell Angular to boostrap/initialize itself and start managing everything below this element in the dom

Controller is the controller javascript function
Model is the $scope object passed to the controller function
View is the html with angular expressions which are linked to the model through data binding

### Directives

allow for indirect model view interaction.
    directives can move data from the model to view and vice versa, some do both.
a lot of HTML elements have custom implementations in angular with the below attributes

{{}}            is a data binding directive
                refers to a variable in the model ($scope) or a method exposed by the controller

ng-app          application directive. Enables all other ng functionality for this html

ng-init         give variables in the scope a default value

ng-model        link the view (html element) and the model
                ng-model="field"
                field is pushed into the model as a property, so you can do $scope.field in the controller
                can be applied to many elements and automatically keeps the view and model in sync
                {{field}} will be auto updated anywhere else in the DOM when the element that ng-model="field" is on

ng-controller   behaviour of content under this element will be managed by the specified controller
                ng-controller="JavascriptFilename as myController"
                **Lots of attributes which invoke actions on the controller**

ng-click        what action to call in the controller when this element is clicked
                normally a method call, and you can pass data into the method call from the view
                ng-click="myController.doSomeAction

ng-repeat       iterate over a collection creating an element for each entry
                <tr ng-repeat="repo in repos">
                    <td>{{repo.name}}</td>
                </tr>

ng-submit       what action to call in the controller when the form is submitted
                if form entries are required use the html validation attribute required

ng-show         display an element based on a condition
                if the condition is truthy in javascript then the element is shown.
                This means we can give the attribute a reference to an object, null => false, an object => true
                0 = false, ...
                if the condition is not met the element is hidden

ng-hide         hide an element based on a condition
                same syntax as ng-show

ng-if           display/remove element based on condition
                if the condition is not met the element is removed from the DOM

ng-src          use it instead of src in elements like imgs
                this is because they browser tries to download the image before angular has replaced the expression

ng-include      bring in html from another source
                attribute value is exact location of markup to include

ng-class        dynamically set the css classes on HTML elements
                the value for the directive/attribute must be in the css format { k : v; k : v; ...}

ng-disabled     based on the condition disable a button

.... so many mores....

there are also custom directives which have been open sourced for your use
xeditable and ui-grid are 2 of them

### Binding expression

Double curly braces: {{expression | filter}}
when the compiler encounters this it evaluates the expression, formats the output as per the filter and display the result in the view.
Expressions can be text within an element or values for attributes of an element

If the expression references something that doesn't exist no error will be thrown in the console and nothing will be displayed in the UI.
Angular's philosophy is not to throw errors because there may be a valid reason that this data does not exist yet in the model.s

### filters

There are a number of built in filters:
    currency
    date
    number
    limitTo
    upper/lowercase
    orderBy ( on ng-repeat) +/- for ascending/descending order
    filter
    ...
syntax of filters is "filterName:parameters"
    parameters can be anything for example in number :1 means include one place after the decimal

### Controllers

**Directive**: ng-controller

```html
<div ng-controller="MainController">
```

MainController is a function. Controller is just a function
$scope is passed to the function.
    **properties and functions attached to the $scope will be the model.**

One of the central pieces of the angular fwk
in charge of an area of application
responsible for building a model
    (and then data-binding is used to link the model to the view and display the data)
In a complex app there can be multiple controllers responsible for different areas of the DOM
These controllers can be nested within each other.

#### modules

Controllers usually live in modules so they don't pollute the global namespace.
You first have to create an angular app and then register functions as controller.

```javascript
var app = angular.module('appName',['dependencies'])

var MainController = function($scope, ...){
    //...
}

app.controller("MainController", MainController);
```

ng-app then takes the name of the app in the module call, in the above example appName

### Promises

a promise object is an object which promises to give you a result sometime in the future.
    call then() on the promise object, passing it a function which will be invoked when the http method returns successfully.
    pass a second function which will be called if the http method call fails.

when then is invoked in a method and it's value is returned from the method. It returns another promise, not the value is computes.

### Services

Components performs a particular job.
Services are used by Controllers, Models, directives and other services

Services can be methods which are invoked or objects which have methods which can be invoked.

Services provide a separation of concerns and de-coupling for Controllers.

Some examples: http, interval, timeout, animate, window, browser, log, location

#### building your own service

many ways to expose service within an app
pluralsight course recommends:

```javascript
var module = angular("app-name");
module.factory(identifier, object);
```
Service should be a function which returns an object.
The object is the service functions/properties tied to the object properties.

#### $http

make http calls to a backend

second parameter to Controller function, $http

```javascript
$http.get(...)
```

all these methods return a promise object

#### $timeout

wraps the javascript timeout functionality

#### $interval

wraps the javascript interval functionality

```javascript
var functionVariable = $interval(function, intervalInMilliseconds, numberOfIntervals)

$interval.cancel(functionVariable);
```

#### $log

log information to the console
all the normal logging levels
you can configure it log back to the web server

#### $location

if you need to read or write the location in the address bar from the controller this is your man.
This comes in useful when moving between views in a big app where you use routing

#### $anchorScroll

force the browser to scroll to a specific element with an id

#### $scope

this is a service which gives the controller access to the model and vice versa.
It does more than that.
you can bind functions to events in scope too.

```javascript
$scope.get(this).$on(event, this.function.bind(this));

$scope.get(this).$broadcast(event);
```


#### $rootScope

This service allows a controller access to the root scope of the web application.
The normal scope is limited to the current model/location in the DOM

### Routing

Allows you to manage multiple views in an application
essential feature for building larger apps where having everything in one view and controller is unmanageable
    allow a user to navigate between different screens (views)
    pass parameters between the controllers that manage these screens
    tap into the browser back and forward buttons

routing is based on the url
    used to locate resource on the client, view / controller
routing engine captures requests and figures out what to do
routing engine configured with rules which tell it what to do when a certain url is encountered.

1. another file needed from angular, angular-route.js
2. which gives use the module ng-route
3. configure rules in $routeProvider
    ```javascript
    .when("url", {
        templateUrl: "html file",
        controller: "ControllerName"
    }
    ```
    which is called in a function which is passed to app.config
4. main html becomes layout view
    header/footer
    routing takes care of what's displayed in the middle
    ng-view directive

The url will now have a # in it. Angular routing uses the fragment part of the URI for routing

when you want to extract parameters from the url configure it:
    .when("/user/:username)
    get this param in the controller using $routeParams
        this makes them available as properties of the $routeParams object