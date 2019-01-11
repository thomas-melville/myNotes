# TLS SSL

https://dzone.com/articles/tlsssl-explained-examples-of-a-tls-vulnerability-a
https://dzone.com/articles/what-is-ssl-how-do-ssl-certificates-work

TLS & SSL are 2 separate protocols

Secure Socket Layer
Transport Layer Security

They can be defined as cryptographic protocols that provide secure communication.

Behind these protocols there is a complicated network of cryptographic functions, algorithms and structures

Provide privacy and integrity, identification and perfect forward secrecy

There can be 1-way and 2-way SSL
Minimize man in the middle attacks and eavesdropping
SSL is standard security technology for establishing encrypted links between clients and servers. Can be browser to webserver, or another example is docker daemon to a registry
**Encryption** uses a private/public key pair
data can be encrypted by one key, but only decrypted by the other.
This is refered to as Public Key Infrastructure (PKI) Scheme.
Public key is shared while private key is kept locally
**Trust** is achieve through the use of certificate trust.
Chain that starts with Certificate Authority.
Entity that issues SSL Certificates.
Certificates can also be self signed, for testing purposes.

## What is the difference between TLS & SSL
SSL is older
both are cryptographic protocols
designed to provide communications security over a computer network
SSL 3.0 served as the basis for TLS 1.0

## How do SSL certificates work

based on public key infrastructure (PKI)
This method involves two distinct cryptographic keys, Private Key & Public Key
The public key is used for encryption and the private key for decryption.
Data encrypted by a public key can only be decrypted by it's corresponding private key
The public key is shared with everyone who receives the digital certificate.
The digital certificate ensures the public key belongs to the user in question

## 2 way SSL
client verifies server identity, then vice-versa.
also referred to as client or mutual authentication
Establishing the encrypted channel using certificate-based 2-Way SSL involves:
    A client requests access to a protected resource.
    The server presents its certificate to the client.
    The client verifies the server’s certificate.
    If successful, the client sends its certificate to the server.
    The server verifies the client’s credentials.
    If successful, the server grants access to the protected resource requested by the client.

## 1 way SSL
client verifies the server identity, but server does not verify client identity

## Privacy & Integrity
allows connection between 2 mediums to be encrypted. Ensures no third party is able to read / tamper with the data en route
Unencrypted an attacker could see sensitive data if they intercepted it.

## Identification
Public key cryptyography provides identification between 2 mediums.
This means both parties can know who they are communicating with.

Once a secure connection is established, the server will send its SSL cert to the client.
The client will verify this against a trusted certificate authority, which will validate the servers identity

## Perfect Forward Secrecy
In the event of a servers private key being compromised an attacker will not be able to decrypt any previous TLS communications.
Diffie-Hellman ephemeral key exchange, new key for every session and only valid as long as the session is valid.

##### Encryption
convert human readable message into encrypted not human readable format
plaintext -> ciphertext
Purpose is to ensure only authorized receivers will be able to decrypt message.

without encryption there is a loss of confidentiality (anyone can intercept and read it), integrity (anyone can intercept & change it), authenticity (anyone can intercept & change it)
Unencrypted data is susceptable to Man in the Middle attacks which change the data

##### Symmetric Encryption
When the same key is used for encryption and decryption by both parties
Biggest problem is both parties must have the same key, if it was exposed the attacker could attack both sides
shared key distribution must be done over an already established secure encrypted communication channel.
Not possible to authenticate sender of message.

##### Asymmetric Encryption
Also called Public Key Cryptography
uses a pair of keys, public key and private key
the keys are uniquely related which means that what is encrypted with one key can be decrypted with the other.
Public key is shared with everyone, private key must be known only to the server
Sender can be authenticated.

##### Ciphers
methods/algorithms used to encrypt/decrypt data.

#### Block Ciphers
split data into fixed length chunks and encrypt.
pad last chunk with random data if it is smaller than size because most of the block ciphers require that the input data is a multiple of their size
Popular ones are: AES, Blowfish, 3DES, DES, RC5

##### Certificate Authority
An entity which issues TLS/SSL or Digital Certificates
These authorities have their own certificate for which they user their private key to sign the issued TLS/SSL or Digital Certs
	This is called the Root Certificate
CAs Root certificate, and therefore public key is installed and trusted by default on browsers.

##### Types of Certificates

#### Single-Domain
secures one hostname or subdomain

#### Wildcard
secures an entire domain and all its subdomains

#### Multi-domain
does what it says on the tin

##### Certificate Validation

#### Domain Validation
only validates that the person who applies for the cert is the owner of the domain name

#### Organization Validation
CA not only validates the domains ownership and identity

#### Extended Validation
Highest level of validation

##### Certificate Generation
For a CA to issue a cert it needs our servers CSR (Certificate Signing Request)
We first create a private key which will be used to decrypt our certificate and the we generate CSR

##### Establishing a TLS connection

https://dzone.com/articles/establishing-a-tls-connection-part-4

Handshake is the most important part of connection establishment
exchange important info regarding properties
Its a hybrid of asymmetric and symmetric encryption

The SSL Handshake involves three steps:

* hello
  * Client sends message which contains some SSL Certificate information
  * Server can make sure the clients cert is valid using a CA
  * Server responds with some SSL Certificate information
* server verification
  * Client makes sure a valid certificate authority (CA) has validated the certificate
* transfer of keys
  * Clients uses the server public key to generate a pre-master key
  * Clients sends this pre-master key to the server
  * Server decrypts this pre-master key using it's private key
  * This way a new key is computed by the client and server
    * This is an example of asymmetric encryption
  * This master key is used to encrypt and decrpyt the info transferred between the client and server.
    * This is an example of symmetric encryption
