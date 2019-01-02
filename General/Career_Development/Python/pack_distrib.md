# packaging and distribution

## 15. modules

**Solve name clashes**

```python
import <filename_minus_extension>
import <filename_minus_extension> as <new_name>
from <filename_minus_extension> import <class_name>
from <filename_minus_extension> import <class_name> as <new_name>
```

/* is used as a wildcard for imports in python too

a module is a python file!!!

when defining a package the folder must contain an __init__.py file

the root of the package should define a requirements.txt file to specify what versions of other packages it requires.
This can lead to issues if you have multiple dependencies who depend on different versions of a dependency that don't overlap

## 21. scripting

to execute code in a module only if it's executed as a script do the following

```python
if __name__ == '__main__':
	# code to execute
```

This will not be executed if the module is imported into another module

## 12. pipenv

dependency manager for Python projects
similar to npm

manages dependencies on a per project basis

```bash
cd project
pipenv install <package>
```

this installs the package and creates/updates a pipfile
others can use this pipfile to install all the required dependencies

run the module with pipenv too!

pipenv creates a virtualenv for the module

to use the virtualenv execute "pipenv shell"

Looks to be more for creating the environment locally from a setup.py file

## 13. virtualenv

create independent working environments for each application with specific versions of required modules.
create a virtual environment with the required modules installed to run your python script
lower level than pipenv

install with pip

create a virtual environment:

```bash
virtualenv <env_name>
```

you can specify python version using --python=

a virtual environment needs to be activated

source <path_to_created_venv>/bin/activate

at this point installed packages will only be installed to this virtualenv

deactivate to leave the virtualenv

```bash
pip freeze > requirements.txt
```
this generates your requirements.txt file with all the packages installed in the virtual env

then you can do

```bash
pip install -r requirements.txt on a new environment
```

## 17. installing python packages

pip, package manager for python
package index at pypi.python.org

## 18. creating an executable file

you can package the interpreter with the code

pyinstaller is the tool, install it using pip

pyinstaller main.py
and the executable file is created
--onefile flag packages everything required into one file

## 19. creating a setup file

the tool he uses is only on windows

### 20. building a whl file

The setup.py script is used to create the whl file.
The naming of the whl file is as follows: <name>-<version>-<python-version>-<abi-tags???>-<architecture>.whl
This is enforced, if you try and install a file with a different name pip fails
