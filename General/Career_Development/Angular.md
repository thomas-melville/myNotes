# Angular

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

src/index.html in the html tag: ng-app="app" tells Angular to be active

### Directives

a lot of HTML elements have custom implementations in angular with the below attributes

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
...

Double curly braces: {{expression | filter}}
when the compile encounters this it evaluates the expression, formats the output as per the filter and display the result in the view.

### Controllers

angular.module('controllerName',['dependencies'])