# tar

bundle a group of files/directors into an arhive

tar args tafFilename pattern

## args

c 	create
x 	extract
t 	list contents in a table
v 	verbose output
f 	the tarfile to perform the actions against

in modern versions of tar gzip is built in!

z argument compresses archived file


use - instead of filename to send output to pipe
	so you could then pipe it into gzip to compress the archived contents