# yum

higher level

work with databases of available software and incorporate the tools needed to find, install, update and uninstall
use both local and remote repos
handles dependencies automatically
handles installation and removal in a robust way

repos are provided by distributions and other independent sw providers.

yum provides a front end to rpm
primary task is to fetch packages from multiple remote repos and resolved dependencies among packages
most Red hat based distros use it

caches information to speed up performance

repo config files are kept in /etc/yum.repos.d/ and have a repo extension
each repo has a config file with information like: desc, url, enabled, ...
you can turn off a repo by executing --disablerep=<repo_name>

## arguments

yum search <keyword>    search for this keyword in package names.
yum list "*keyword*"    search for this keyword in package names. tells what is installed or not
yum list installed/updates/available
yum info <package>      display info about this package
yum grouplist <group>
yum groupinfo <group>
yum provides <file>     what package provides a file
yum install package...  install one or more packages
yum localinstall file   install from a local rpm file
yum update
yum remove
yum history
yum shell               launch an interactive shell to execute multiple commands, it can be given a text file of commands

### flags

--downloadonly

## plugins

yum has lot of plugins, not packages.
one is yum-plugin-verify, which verifies packages during installation.

## during installation

existing config file is renamed to .rpmsave
if it can be safely applied after it is changed to .rpmnew

## ui

yumex command launches a ui for yum