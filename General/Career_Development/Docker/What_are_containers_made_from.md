# What are containers made from

cgroups, namespaces - isolation technologies

What is a container
	feels like a vm
	you could ssh into it, but don't!
	chroot on steriods
		normal processes running on a normal kernel
	a container is a process running on the host kernel

How are they implemented
	cgroups(control groups) & namespaces in the kernel
	cgroups - impl metering, limiting, usage on the resources used by processes (cpu, ram io, mem, pages, ...)

Put one service per container
If the process in the container exceeds any of it's limits only it is killed
And not some random process somewhere else
