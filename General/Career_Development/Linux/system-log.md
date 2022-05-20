# system log

## rsyslog

rocket fast system for logging.
Accepts messages from various sources, separating the process that generated the log from the process that records the event message.

Config file located in /etc/rsyslog.conf

Modules are used for input and output
* imuxsock - support for local system logging
* imjournal - provides access to the systemd journal
* imklog - provides kernel logging support

### Default configuration

* accepts all system-related messages
* records events in local files
* recorded in files in /var/log/
* /etc/rsyslog.d/50-default.conf contains the default rules for where to log what

## remote logging

### client

Different options for sending logs to a remote host.
UDP was first used, but reliable transmission is an issue. So, TCP is used now too.
Both protocols require the server to be ready to receive the messages, otherwise they may be lost.

RELP, Reliable Event Logging Protocol, is another option which can be configured to prevent loss of messages.
* temporary disk queue
* saving outstanding messages on shutdown
* retry counters

### server

Server can listen for TCP, UDP and RELP messages.
Module must be loaded for each one, and requires a configuration of what port to listen on.
