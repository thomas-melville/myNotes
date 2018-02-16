# build lifecycle

there are multiple build lifecycles in maven
    default, clean, site
each one is made up of phases
    a phase represents a stage in the lifecycle

when you execute a phase all phases before that one in the lifecycle are executed.

## default 

validate
compile
test
package
verify
install
deploy