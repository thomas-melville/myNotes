# openssl

https://www.openssl.org/

a toolkit for TLS and SSL
general purpose cryptography library

https://www.feistyduck.com/books/openssl-cookbook/

consists of:
* high performance implementation of key cryptographic algorithms
* a complete SSL/TLS and PKI stack
* command line toolkit

## key generation

key algorithm
  lots of algorithms, which one suits your needs
  webservers use RSA
key size
  default size, 512 bits, may not be secure enough
  easy to brute force attack the key out
  2.048-bit RSA is minimum
passphrase
  optional, not vital.
  only useful when not installed on production systems

```bash

# generate public key
openssl genrsa -aes128 -out fd.key 2048

```

-aes128 is the algorithm used

## Certificate Signing Request generation

contains the public key, and is signed with the private key

```bash

# generate CSR
openssl req -new -key fd.key -out fd.csr

# view CSR
openssl req -text -in fd.csr -noout

# create a new CSR from an existing certificate
openssl x509 -x509toreq -in fd.crt -out fd.csr -signkey fd.key
```

you can automate the interactive part of the command by creating a .cnf file and populating all the relevant fields:

[req]
prompt = no
distinguished_name = dn
req_extensions = ext
input_password = PASSPHRASE

[dn]
CN = www.feistyduck.com
emailAddress = webmaster@feistyduck.com
O = Feisty Duck Ltd
L = London
C = GB

[ext]
subjectAltName = DNS:www.feistyduck.com,DNS:feistyduck.com

```bash

# generate CSR with configuration
openssl req -new -config fd.cnf -key fd.key -out fd.csr

```
