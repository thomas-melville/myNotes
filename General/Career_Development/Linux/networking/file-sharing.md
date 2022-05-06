# file sharing

## ftp

one of the first protocols on the internet. Data is sent in plain text.
There is interactive and non-interactive mode.
Two data transfer modes:
* Active
  * server pushes data to the client on a random high-numbered port.
  * not compatible with firewalls and NAT.
* Passive
  * The client requests a passive connection and the server opens a random high-numbered port for data transfer.

### FTP Server

Very secure FTP Daemon, vsftpd, was suggested by SANS, IBM anf others for security reasons.
upload and download of some content not allowed.
features include:
* Virtual Ips
* Virtual users
* Stand-alone daemon or inetd-ready
* per-user config
* pre-source config
* optional SSL integration

## rsync

protocol written as a replacement for rcp. Uses an advanced algorithm to intelligently transfer files.
Only the files or parts of files which have changed are copied. Uses delta encoding and requires the source and destination to be specified (either or both may be remote)

Does not have it's own in transit security, however it can be tunneled through SSH.
Will default to use the ssh shell, unless you specify an alternative.
Can be invoked using rsync:// in the URI or with rsyncd.

## scp

"Secure copy", non-interactive.
uses machine ssh configuration.

## sftp

Simple File Transfer Protocol is interactive.
uses machines ssh configuration.

## WebDav

An extension to HTTP for r/w access to resources and is available on most OS's.
Must be enabled in the webserver configuration, mod_dav module in apache.
