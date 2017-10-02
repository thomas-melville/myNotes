# Locate

if you want to find a file and you know parts of it's name but not where it is then locate is a good command

locate [pattern]

This command doesn't search the file system, it searches an index.
once a day all the files are indexed as part of a process called updatedb.

Since the db is indexed once a day, it might not be up to date.
A new file may not be there and an existing file may have been moved/deleted.
you can update the db/index by executing 'updatedb'