# keys

## Manually generate public-private key pair for authentication
means no need for a password.
And is more secure than password authentication

public key is placed on all compters that must allow access to the owner of the matching private key (kept secret by the owner)
	So, if I want no password authentication I must put my public key on a server I want to access.

## Public Key Authentication
Authenticating entity has a public key and a private key.
Each key is a large number with special mathematical properties.
Private Key is kept on computer you are logging in from
Public key is stored in .ssh/authorized_keys file on all computers you want to log into.

Once you're logged in the public key "locks"messages in a way that only the private key can "unlock"
	provides end-to-end security

Added security:
	most SSH programs store private key in a passphrase-protected format.


#### Authorized keys ####
file in ssh folder.
Contains public keys of users who are allowed to remotely login

#### Private keys ####
private keys can be locked with a passphrase
	That's what's happeneing for the one I generated from trystack
