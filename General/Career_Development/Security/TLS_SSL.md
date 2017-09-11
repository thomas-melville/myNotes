Content-Type: text/x-zim-wiki
Wiki-Format: zim 0.4
Creation-Date: 2017-05-09T15:05:39+01:00

###### TLS SSL ######
Created Tuesday 09 May 2017

https://dzone.com/articles/tlsssl-explained-examples-of-a-tls-vulnerability-a?edition#298008&utm_source#Daily%20Digest&utm_medium#email&utm_campaign#dd%202017-05-07

TLS & SSL are 2 separate protocols

Secure Socket Layer
Transport Layer Security

Behind these protocols there is a complicated network of cryptographic functions, algorithms and structures

Provide privacy and integrity, identification and perfect forward secrecy

#### Privacy & Integrity ####
allows connection between 2 mediums to be encrypted. Ensures no third party is able to read / tamper with the data en route
Unencrypted an attacker could see sensitive data if they intercepted it.

#### Identification ####
Public key cryptyography provides identification between 2 mediums.
This means both parties can know who they are communicating with.

Once a secure connection is established, the server will send it's SSL cert to the client.
The client will verify this against a trusted certificate authority, which will validate the servers identity

#### Perfect Forward Secrecy ####
In the event of a servers private key being compromised an attacker will not be able to decrypt any previous TLS communications.
Diffie-Hellman ephemeral key exchange, new key for every session and only valid as long as the session is valid.

##### Encryption #####
convert human readable message into encrypted not human readable format
plaintext -> ciphertext
Purpose is to ensure only authorized receivers will be able to decrypt message.

without encryption there is a loss of confidentiality (anyone can intercept and read it), integrity (anyone can intercept & change it), authenticity (anyone can intercept & change it)
Unencrypted data is susceptable to Man in the Middle attacks which change the data

##### Symmetric Encryption #####
When the same key is used for encryption and decryption by both parties
Biggest problem is both parties must have the same key, if it was exposed the attacker could attack both sides
shared key distribution must be done over an alreadt established secure encryoted communication channel.
Not possible to authenticate sender of message.

##### Asymmetric Encryption #####
Also called Public Key Cryptography
uses a pair of keys, public key and private key
the keys are uniquely related which means that what is encrypted with one key can be decrypted with the other.
Public key is shared with everyone, private key must be known only to the server
Sender can be authenticated.

##### Ciphers #####
methods/algorithms used to encrypt/decrypt data.

#### Block Ciphers ####
split data into fixed length chunks and encrypt.
pad last chunk with random data if it is smaller than size because most of the block ciphers require that the input data is a multiple of their size
Popular ones are: AES, Blowfish, 3DES, DES, RC5

##### Certificate Authority #####
An entity which issues TLS/SSL or Digital Certificates
These authorities have their own certificate for which they user their private key to sign the issued TLS/SSL or Digital Certs
	This is called the Root Certificate
CA's Root certificate, and therefore public key is installed and trusted by default on browsers.

##### Types of Certificates #####

#### Single-Domain ####
secures one hostname or subdomain

#### Wildcard ####
secures an entire domain and all it's subdomains

#### Multi-domain ####
does what it says on the tin

##### Certificate Validation #####

#### Domain Validation ####
only calidates that the person who applies for ther cert is the owner of the domain name

#### Organization Validation ####
CA not only validates the domains ownership and identity

#### Extended Validation ####
Highest level of validation

##### Certificate Generation #####
For a CA to issue a cert it needs our servers CSR (Certificate Signing Request)
We first create a private key which will be used to decrypt our certificate and the we generate CSR

##### Establishing a TLS connection #####
Handshake is the most important part of connection establishment
exchange important info regarding properties
It's a hybrid of asymmetric and symmetric encryption

https://dzone.com/articles/establishing-a-tls-connection-part-4
