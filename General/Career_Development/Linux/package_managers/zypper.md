# zypper

SUSE based systems

also works with rpm packages
functionality is very similar to yum

install, update, remove, resolved dependencies, ...

## commands

list-updates
repos
search <string>
info <package>
search --provides <file>
install <package>
update
remove
shell   open interactive shell so you can execute multiple zypper commands without reading the database everything
addrepo <uri> <alias>
removerepo <alias>

### flags

--non-interactive don't ask for confirmation

## ui

YaST
    Yet another setup tool
    it's more than package management, all system administration