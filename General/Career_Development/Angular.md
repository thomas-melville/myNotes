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

a lot of HTML elements have custom implementations in angular with the below attributes

ng-app - application directive. Enables all other ng functionality for this html
ng-init - give variables in the scope a default value
ng-model - link the view (html element) and the model
    ng-model="field"
    {{field}} will be auto updated anywhere else in the DOM when the element that ng-model="field" is on
ng-controller - behaviour of content under this element will be managed by the specified controller
    ng-controller="JavascriptFilename as myController"
**Lots of attributes which invoke actions on the controller**
ng-click - what action to call in the controller when this element is clicked
    ng-click="myController.doSomeAction
ng-repeat
    unroll a collection
ng-submit
    submit a form
ng-show
    display an element based on a condition
    if the condition is not met the element is hidden
ng-if
    display/remove element based on condition
    if the condition is not met the element is removed from the DOM
ng-src
    use it instead of src in elements like imgs
    this is because they browser tries to download the image before angular has replaced the expression

### Binding expression

Double curly braces: {{expression | filter}}
when the compiler encounters this it evaluates the expression, formats the output as per the filter and display the result in the view.
Expressions can be text within an element or values for attributes of an element

If the expression references something that doesn't exist no error will be thrown in the console and nothing will be displayed in the UI.
Angular's philosophy is not to throw errors because there may be a valid reason that this data does not exist yet in the model.s

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

var MainController = function($scope){
    //...
}

app.controller("MainController", MainController);
```

ng-app then takes the name of the app in the module call, in the above example appName

### Services

#### $http

make http calls to a backend

second parameter to Controller function, $http

```javascript
$http.get(...)
```

all these methods return a promise object

a promise object is an object which promises to give you a result sometime in the future.
    call then() on the promise object, passing it a function which will be invoked when the http method returns successfully.
    pass a second function which will be called if the http method call fails.
