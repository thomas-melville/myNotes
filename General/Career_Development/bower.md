# bower

A package manager for the web.
Can manage components which contain HTML, CSS, Javascript, fonts or even image files..
Fetches and installs packages from all over.
Has a manifest file, bower.json
If multiple packages depend on a single package, it will be downloaded only once.
    This is a flat dependency graph

## command line

bower install       install package, by shorthand, git url, ...
    --save          install the package and put it in the bower.json file
    --save-dev      install the package and put it in the dev section of the bower.json file