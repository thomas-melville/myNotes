# Libraries

programs are built using libraries of code, which are reused by many programs

## Types

Static
    code from library is inserted into the program at compile time and does not change after

Shared
    code from library is loaded into the program at runtime

shared libraries is more efficient

Shared Libraries are also called Dynamic Link Libraries (DLLs)

## Versioning

need to be carefully versioned
if there are significant changes to a library and a program can't handle it then it's hell
programs can request a version, but normally just use the latest minor version

Shared Libraries have the extension .so.N where N is the version

Symbolic links are used in the /lib folders to give specific versions based on generic links

## How to see what libraries a program uses

ldd command with the path to the executable.

## Finding shared libraries

ldconfig is run and uses /etc/ld.so.lconf which lists the locations of the shared libraries to search