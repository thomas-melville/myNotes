Content-Type: text/x-zim-wiki
Wiki-Format: zim 0.4
Creation-Date: 2017-08-10T17:18:58+01:00

###### Checked Builder Pattern ######
Created Thursday 10 August 2017

https://dev.to/schreiber_chris/creating-complex-objects-using-checked-builder-pattern

for every step of the building process a simple interface is defined
which contains the setters methods for this stage and returns an instance of the next stage.
The builder then implements all of these interfaces

[[+ives]]
	explicitly tells the user how the object can be built so they never end up with a bad object

-ives
	a lot of extra interfaces, which means more maintenance
	
