# Conda

Package, dependency and environment management for any language

* install, run & update packages
* create, save, load and switch between environments

created for python but can package anything

## concepts

### environments

a dir that contains a specific collection of conda packages that you have installed
each environment is isolated from the others
easily share an environment with an environment.yaml file

When you start conda you already have a default environment named base.

### packages

a compressed tarball that contains system level libraries, Python or other modules, executable programs and other components

## configuration

.condarc file in your home directory.
yaml format

## questions

How do I get conda to search in PyPi?
    configure this in the .condarc file
    can the install take a url?
    It looks like conda is a proxy for pypi? or replicates the packages in a conda format?