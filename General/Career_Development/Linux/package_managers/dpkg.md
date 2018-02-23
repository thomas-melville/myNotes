# dpkg

Debian Package Management
install, remove and manage software package on Debian distributions

package files have .deb file extension

database resides in /var/lib/dpkg
only knows whats installed on the system and what it's given on the command line.

## package file names

<name>_<version>_<rev_num>_<arch>.deb

a debian package consists of at least 3 files

* upstream tar ball
** unmodified source from package maintainers
**.tar.gz
* description file
** package name. meta data such as arch & dependencies.
** .dsc
* second tar ball
** patches to upstream source
** .debian.tar.gz

## dpkg command

-l              list all installed packages (--list)
-L <package>    list file installed in the specified package (--listfiles)
-s <package>    show info about installed package
-I <pack_file>  show info about package file
-c <pack_file>  list all files in a package file
-S <file>       show what package owns this file
-s <package>    show the status of the package
-V <package>    verify the integrity of the package
-i <pack_file>  install or upgrade package file
-r <package>    remove all of a package bar config files (--remove) will fail if other packages depend on this package
-P <package>    remove all of a package including the config files (--purge)