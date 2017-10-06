ansible playbook for my development environment

software:
git
	install git and version I want
	my config
	my gitignore
	ssh keys (can't automate uploading them to github and gerrit)
		I will have at least one manual step :-(
mvn
	whatever version comes from apt-get, 16.04 => 3.3.9
	settings.xml??
java 7 & 8
	setting up alternatives
intellij
	install plugins (zip up .IdeaC2017.2 folder )
	import settings (this could be manual!) this is covered in .IdeaC2017.2 folder!!!
	set up alias (or could I just put /opt/intellij... on the path)
sublime
	install plugins
	settings ???
fish shell
	download Dmitry's configuration
	silver searcher
	make it default shell???


when executing on my reinstalled machine:
	/etc/ansible/hosts
		localhost and user and password to use, ethomev
	oh no wait, I can configure it to be the local machine
	localhost ansible_connection=local