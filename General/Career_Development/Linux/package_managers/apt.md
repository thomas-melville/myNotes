# apt

advanced packaging tool

debian distributions: ubuntu, mint

main utilities: apt-get and apt-cache

.deb file extension

## ui

synaptic, ubuntu software centre, aptitude

## apt-get

main apt command line tool for package management
install, manage, upgrade single packages or entire system.

/etc/apt/sources.list contains the list of repos to look in

you must update before you upgrade in apt-get

### commands

update          synchronize package index files with remote repos
install <>      install the package
remove <>       remove the package, without removing config
                --purge to remove config files too
upgrade         apply all available updates to packages already installed.
dist-upgrade    smart upgrade, thorough dependency resolution, remove obsolete packages and install new dependencies
autoremove      get rid of any packages not needed any more
clean           clean out the cache files

## apt-cache

queries are done using the apt-cache utility

### commands

search <package> search for the package in the repo
show <package> display basic info about the package
showpkg <package> display detailed info about the package
depends <package> list all dependent packages of package
search <file> search the repo for a file named file
list <package> list all files in the package