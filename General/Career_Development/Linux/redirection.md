# redirection & pipes

## pipes

done using |
pipe redirects the output of one command into the input of the next

## redirection

done using >

3 default types of in/output
	standard in		stdin	0
	standard out 	stdout	1
	standard error	stderr	2

each one has a FileDescriptor
	a FileDescirptor is a number that represents an open file

\> redirect output to file (overwriting contents)
\>> redirect output to file (appending to file)
\< redirect input from a particular location, rather than stdin

put the FileDescriptor number before the number to specify

2>&1 redirect stderr to stdout so that all information is in one FileDescriptor
	if you just redirect 1 to a file you may miss some information on 2

redirect output nowhere!!!
	\>/dev/null
	i.e. ignore the output
