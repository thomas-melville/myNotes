# Find

find [path] [expression]

Find files of a certain name/extension
	find . -name groovy-all*

Find files that contain a certain content:
	 find -type f | xargs grep -l "cron"

	**Much simpler one!!!!**
	grep -rl MeContext

## expression

is made up of options, tests and actions which can be combined using logical boolean operators

-name [pattern]
-iname [pattern] 		ignore case!
-mtime number_of_days 	find file x days old, can be a positive or negative number for older/newer than
-size +/-[num][unit]
-newer [file]			find file newer than this file
-exec command {} ;			execute a command against all found files. {} placeholder for current file, terminate command with ;
-type					directory / regular file / link
-path [pattern]		the path to match for

-print0		if there is a chance that the paths you are finding could contain spaces or newline characters then use this argument to enable the results to be correctly processed by the command that you are piping the output of this command into. Corresponds to -0 argument in xargs

### options

-d 					perform a depth first traversal
-daystart		measure time from the start of the day instead of 24 hours ago
-min/maxdepth controls how many directories to go down, default: min 0 and max unlimited.

### tests

-name
-regex
-type
-a/min/newer/time access time
-c/min/newer/time created time
-m/min/newer/time modified time
-newer if the file is newer than another file.
-perm match a given permission

### actions

actions executed on the files that match the test.
default is to print the fule path.

-ls
-print/0/f print extra info to stdout
-fprint/0/f print extra info to a file
-delete
-exec

### exclusions

-path <path> -prune -o
	don't descend into the directory specified by <path>
-o is for logical OR
-a is for logical AND

**ACK-GREP**
Use this anymore for searching file content.
Ignores scm configuration folders and the output is nice and colourful.

ack-grep "cron"

**if you want to find a file and you know parts of it's name but not where it is then locate is a good command**
