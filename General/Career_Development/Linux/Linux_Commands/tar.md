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

the order in which the files are added/extracted from the archive is file system dependent.
tar makes a kernel call to readdir() which will return the files to be added to the archive.
In some file systems this returns the file in the order they were last accessed/modified.
