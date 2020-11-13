# Securing data with asymmetric cryptography

Asymmetric cryptography is also called Public key cryptography
use public & private keys and cryptographic and signature operations

## PKI

Public Key Infrastructure

used in an unsecure network to securly exchange data
CA is at the core
  can be
    an organization
    an implementation on a machine
      Root/Intermediate/PolicyCA

### Trust

4 major responsibilities of a PKI

1. authentication
    look at identity information
    and use key to verify
2. integrity
    has the data been altered in root
    use the key & hash data to ensure data was not altered
3. confidentiality
    only the people we want to
    encryption
4. non-repudiation
    hold an entity accountable as the origin of data

### Certificate Authorites

Sign certificates
signing analogoues to a singing notary

#### RootCA

At the top of the hierarchy
does not sign end user certificates itself

RootCA Certificate is self signed, since there's no one above it in the hierarchy
Identified as a root CA using the Basic Constraints extension (x.509)
  binary value
How do we know a RootCA is trustworthy?
  OS's ship with a folder of RootCA

Kept offline, only used to sign Intermediate CA

#### Intermediate CA

Subordinate/Intermediate/Policy/Issuing

why have a hierarchy?
  organization
    separate concerns
    specific Intermediate for specific types of certificates or regions
  risk mitigation
    if Intermediate revoked, all cert it signed it has to be revoked and reissued

#### Policy CA

Criteria for issuing, revoking and renewing certs

low to high trust for each criteria

CA Company publishes documentation on all the criteria above.

x.509 v3 standard specifies a cert policies extension which gives all this information

Policy CA writes the above into the certificate

### Key pairs

Key Pair Purposes:

1. Encryption
    public key
2. Decryption
    private key
3. Signing
    private key
4. Verifying
    public key

Which key does what?

2 things to remember

1. who has the data
2. what will be done with the data

You make the public key available to everyone.
Keep your private key to yourself.

### getting a key signed

get public key signed
generate CSR (Certificate Signing Request)
  add public key
  info about who we are, written in distinguished name format
  any extra info

then send it to certificate authority

### revoking certificates

malicious compromise
employee separation

CA will certificate number to CRL (Cert Revoation list)
CA will have url to CRL list

OCSP Online Certiicate Status Protocol
Authoity information Access extension point.

## Keys

### Symmetric V Asymmetric cryptography?

symmetric, both sides use the same key to en/decrypt!
  problem is getting key to the other person securly!

this was a motivating factor to asymmetric

one key is public and available to everyone
one key is private and only available to the owner

asymmetric is slower
  because the data has to be broken down into chunks
  depends on the bit size of the key
  every chunk adds 11 bytes

Combine the two to get the best of both worlds.

Create a symmetric key, send it using asymmetric cryptopgrahy
use the symmetric key for all subsequent transactions

### The Math of Key pairs

Discreate Logarithm Problem

Easy to generate the key using the prime numbers we provided
Extremely hard to get the prime numbers from the key pair

Modular arithmetic

2 large prime numbers, pick really big prime numbers, 100s of digits in length
Modulus of the two numbers
totient of the modulus
encryption & decryption component

p = 3
q = 11
modulus = p * q = 33 = n
totient(n) = 20?
e =
d =

:-) so much maths!

### Storing Keys

Keystores

defined by PKCS#12 file format for storing on servers

Safe bag
  5 different types
    Private Keys
    certificates
    CRL
    User defined data
    Safe Contents bag, which holds all the above

## x.509 Certificates

### Certificate types

3 types
  Certificate Request
  Certificate
  Certificate Revocation List

Attribute Certificate?
  deals with authorization

we're dealing with authentication
  PKI
  authentication
  who the subject is
  Cert, CSR, CRL

SSL/TLS certificates are X.509v3 certificates with extensions

### Certificate structure

attach attributes and information to a public key
Trust the information when it's been signed by a 3rd party we trust

6 fields in a certificate

1. version, currently X.509v3
2. serial number, has to be unique within a CA
3. signature algorithm
    used by CA to sign cert
    preferred SHA
    used in conjunction with PKI encryption algorithm
      RSA is the most common
4. Issuer
    Distinguished name of the signer
      ex: C=US, O=Duck, OU=Security, CN=Master Mallard
    usually a CA
5. Subject
    distinguished name of key owner
6. validity period
    valid from & to

The public key is in the certificate

last field is the collection extensions
  restrict what priviliegs the key has
  encryption/signing
  ....

### Certificate trust

take the information about a public key at face value
we trust that information is correct
foundational element of PKI
has to be validated by a 3rd party
This is the CA
We submit our public key in a CSR, Certiicate Signing Request
4 parts
* Version
* Subject
* Public key
* Collection of extensions

CA adds the
* serial number
* signing algorithm
* provides info about itself
* validty period

#### When a key is compromised

CA will include the Cert in a Cert Revocation List
contains list of Certs that are revoked before their validity expires
Daily -> 90 days

all the common fields
* this update
* next update
* User certificates
* extensions

### Certificate Chain

When a cert is exported to a file all certs invovlve are written out
called a cert chain
to verify chain we calculate the signature of each cert using it's parent
decrypt the signature in the cert using it's parents public key
and compare
once at the RootCA ensure it's in our collection of trusted roots


### Certificate ID info

information comes from ITU x.500 standard series defined in the 80s
think active directory

C   Country
PC  Postal Code
ST  State or Province
L   Locality
O   Organization
OU  Organizational Unit
CN  Common Name
E   Email

combine any number of these fields together

in X.509: C ST L O OU CN

only the C and CN are required
C is two letter acronym for country

### Extensions

introduced in version 3

key value pair
key is an object identifier
OID, string of numbers which identify a unique object in a repository
needs to be registered
not common to define your own

critical v non-critical

### Example Project

ASN: Abstract Syntax Notation
  de/serialize data structures in a language independent way

PEM: text format of keys and signatures
