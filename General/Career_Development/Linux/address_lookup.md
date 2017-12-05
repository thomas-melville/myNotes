# ip address resolution in linux

1. look at /etc/nsswitch.conf for where to look
2. follow these in order, usually:
	files (/etc/hosts)
	dns
	dis
3. for dns look at /etc/resolv/conf for dns

nslookup is the command to lookup up the ip address of a url