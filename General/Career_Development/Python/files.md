# working with files

built in method open("filename", "access modes")
returns a file object

access modes

* a	append
* x exclusive access, fails if it already exists
* r	read
* * if the file doesn't exist it throws an exception
* w	write, will overwrite existing data. It's called clobbering
* * if the file doesn't exist it will create it
* b	open the file in binary mode, Windows only! but doesn't hurt when your developing on linux, makes the module platform independent

.write(...)
.close()

always remember to close files.
files have locks and if it's already open it won't be possible to open it again

open returns an object which is a Context Manager which means you can use the keyword 'with' when opening it to auto close it when the variable goes out of scope

when reading and printing lines put a , at the end to stop blanks lines between each line
