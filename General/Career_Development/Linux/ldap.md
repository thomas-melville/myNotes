# ldap

Lightweight Directory Access Protocol

industry standard for using and administering distributed directory services over the network
is open and vendor neutral

each system/client connects to a centralized LDAP server for authentication
using TLS makes it a secure option and is recommended.

LDAP uses PAM and system-config-authentication or authconfig-tui
One has to specify the server, search base DN (Domain Name), and TLS.
Also required is openldap-clients, pam ldap, nss-pam-ldapd