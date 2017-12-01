# VI

### Normal Mode ###

normal mode is for navigation and text manipulation, I used to call it command mode
ESC	go from insert to normal mode
I've changed mine to "rty"

. repeat last command

### Insert Mode ###

i		go from command mode to insert mode
		on the cursor
a		go from command mode to insert mode
		after the cursor
A		go from command mode to insert mode
		at the end of the line
o		go from command mode to insert mode
		on a new line inserted below the line the cursor is on
O		go from command mode to insert mode
		on a new line inserted above the line the cursor is on


### Replace Mode ###
Overwrites the existing text as you type

R		go from command mode to replace mode
		on the cursor


### Visual Mode ###
Allows you to highlight text

v go from command mode to visual mode

you can then use any of the commands below to do what you want with the highlighted text

v, move the cursor, :w FILENAME
	this selects text and saves it to a file


### Command Completion ###

in command mode type a letter followed by CTRL+D and it will show you the list of options
TAB to complete
it can complete arguments also, filenames is all I've seen so far

### Help ###

F1 / :help

### Substitution ###

:s/a/b
	replace the next occurance of a with b in the file
:s/a/b/g
	replace all occurance of a with b on this line
:#,#s/a/b/g
	replace all occurances of a with b on the lines specified by the #'s
:%s/a/b/g(c)
	replace every occurance of a with b in the file
	add in the c to prompt whether to substitute each one



### move the cursor around the text: ###

:100
	go to line one hundred

		k
	h		l
		j
		
0(zero)
	move to the start of the line

$
	move to the end of the line
	
w
	move the cursor to the start of the next word
	
e
	move the cursor to the end of the next word

**Enter a number before any of ther 2 commands above to jump that number of words**	

### edit the text while in command mode: ###

A way to think about the shortcut keys for editing is as follows:

mode (number) action, when working within a line
(number) mode, when working across lines

x
	delete a character

dw
	delete a word

de 
	delete to the end of the current word
	
**Enter a number in the middle of the 2 above commands to delete that number of words**

d$
	delete to the end of the line

dd
	delete a line

yy
	cut a line
	
**Enter a number before any of the 2 commands above and that number of lines will be cut/deleted**

p
	paste what you have just cut
	if it's a full line it goes on the next line
	if it's part of a line it goes after the cursor
	
ce
	change to the end of the word
	this command deletes the text to the end of the word and places you in edit mode
	

### Un ###

u
	undo the last operation
	
U
	undo all operations on this line since moving to it


### Cursor location & File Status ###

CTRL+G
	show location in file and file status
	
G to move to the end of the file
gg to move to the start of the file

<number> + G to go back to the line you were on when you pressed CTRL+G

### Searching ###

/ to start search query going forward
? to start a search query going backwards
enter to search
n to search for next occurance
N to search for previous occurance
* highlights all occurances of the word the cursor is on

% to find matching closing bracket
	cursor needs to be on opening bracket
	
There are options you can set before searching
Looks like these can be set in the .vimrc file!]

:set ic -> ignore case when searching
:set hls -> highlight 
:set is -> incrementally search, as you type the search results are refined

prepend no to switch a setting off
:Set noic for example
	

### Executing external shell commands in VI ###

:!<command>
	command can take any amount of arguments
	

### Writing to files ###

:w (filename)
	write the contents of the file to disk, optionally specifying the file name
	If you are working in a file that already exists and you specify a file name it works like "Save As"

v, move the cursor, :w FILENAME
	this selects text and saves it to a file
	
	in visual mode you can also copy or cut the text!


### vimrc file

nnoremap
	change the operation of a key in normal mode
	example:
		nnoremap: <Up> <Nop>
			change the operation of the up arrow key to no operation
inoremap
	change the operation of a key in insert mode
	example:
		inoremap: rty <esc>
			change the operation of the sequence of keys rty to esc, which moves into normal mode
