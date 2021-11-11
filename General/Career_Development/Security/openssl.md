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

## Signing your own certificates

when you have a CSR execute the following command:

```bash

openssl x509 -req -days 365 -in fd.csr -signkey fd.key -out fd.crt

```

## certificates valid for multiple hostnames

By default there is only one common name

1. list all desired hostnames using n x.509 extension, Subject Alternative Name (SAN)
2. use wildcards

Recommendation is to put specific domain namd and wildcard in SAN

put SAN information in text file, and add it to the openssl command using the -extfile argument

## examining certificates

certificates don't look like a lot in a text editor
The x509 command unpacks the certificate and gives useful information

-text to print the content
-noout to reduce clutter

```bash

openssl x509 -text -in fd.crt -noout

```

### Public Certificates

There is a lot more in them, in the x.590 extensions

Basic Constraints
  does it belong to a CA
  which would give it the ability to sign other certificates
Key Usage
Extended Key Usage
  restrict what a certificate can be used for
  not present => no restrictions
CRL Distribuion Points
  where the revocation lists can be found
Certificate Policies
  type of validation used to ascertain the identity of the owner

## key & certificate conversion

can be stored in a variety of formats

ASCII (PEM) certificates
  base64 encoded DER certificate
  -----BEGIN CERTIFICATE-----
  -----END CERTIFICATE-----
  usually only 1 cert per file
  although some programs allow more

ASCII (PEM) key
  base64 encoded DER key
  sometimes with addition metadata

PKCS#7 certificates
  complex format
  for transport of signed/encrypted data
  can contain the entire cert chain

### conversion

```bash

# PEM to DER
openssl x509 -inform PEM -in fd.pem -outform DER -out fd.der

# DER to PEM
openssl x509 -inform DER -in fd.der -outform PEM -out fd.pem

# PEM to PKCS#12
openssl pkcs12 -export \
    -name "My Certificate" \
    -out fd.p12 \
    -inkey fd.key \
    -in fd.crt \
    -certfile fd-chain.crt

# PKCS#12 to PEM
openssl pkcs12 -in fd.p12 -out fd.pem -nodes
# but this results in key, cert & intermediate in one file

# you can get pkcs12 to do it for you
openssl pkcs12 -in fd.p12 -nocerts -out fd.key -nodes
openssl pkcs12 -in fd.p12 -nokeys -clcerts -out fd.crt
openssl pkcs12 -in fd.p12 -nokeys -cacerts -out fd-chain.crt

```

## Creating a private CA

### Root CA configuration

create a lengthy configuration file

[default]
name                    = root-ca
domain_suffix           = example.com
aia_url                 = http://$name.$domain_suffix/$name.crt
crl_url                 = http://$name.$domain_suffix/$name.crl
ocsp_url                = http://ocsp.$name.$domain_suffix:9080
default_ca              = ca_default
name_opt                = utf8,esc_ctrl,multiline,lname,align

[ca_dn]
countryName             = "GB"
organizationName        = "Example"
commonName              = "Root CA"

[ca_default]
home                    = .
database                = $home/db/index
serial                  = $home/db/serial
crlnumber               = $home/db/crlnumber
certificate             = $home/$name.crt
private_key             = $home/private/$name.key
RANDFILE                = $home/private/random
new_certs_dir           = $home/certs
unique_subject          = no
copy_extensions         = none
default_days            = 3650
default_crl_days        = 365
default_md              = sha256
policy                  = policy_c_o_match

[policy_c_o_match]
countryName             = match
stateOrProvinceName     = optional
organizationName        = match
organizationalUnitName  = optional
commonName              = supplied
emailAddress            = optional

[req]
default_bits            = 4096
encrypt_key             = yes
default_md              = sha256
utf8                    = yes
string_mask             = utf8only
prompt                  = no
distinguished_name      = ca_dn
req_extensions          = ca_ext

[ca_ext]
basicConstraints        = critical,CA:true
keyUsage                = critical,keyCertSign,cRLSign
subjectKeyIdentifier    = hash

[sub_ca_ext]
authorityInfoAccess     = @issuer_info
authorityKeyIdentifier  = keyid:always
basicConstraints        = critical,CA:true,pathlen:0
crlDistributionPoints   = @crl_info
extendedKeyUsage        = clientAuth,serverAuth
keyUsage                = critical,keyCertSign,cRLSign
nameConstraints         = @name_constraints
subjectKeyIdentifier    = hash

[crl_info]
URI.0                   = $crl_url

[issuer_info]
caIssuers;URI.0         = $aia_url
OCSP;URI.0              = $ocsp_url

[name_constraints]
permitted;DNS.0=example.com
permitted;DNS.1=example.org
excluded;IP.0=0.0.0.0/0.0.0.0
excluded;IP.1=0:0:0:0:0:0:0:0/0:0:0:0:0:0:0:0

[ocsp_ext]
authorityKeyIdentifier  = keyid:always
basicConstraints        = critical,CA:false
extendedKeyUsage        = OCSPSigning
keyUsage                = critical,digitalSignature
subjectKeyIdentifier    = hash

#### directory structure

$ mkdir root-ca
$ cd root-ca
$ mkdir certs db private
$ chmod 700 private
$ touch db/index
$ openssl rand -hex 16  > db/serial
$ echo 1001 > db/crlnumber

#### Root CA generation

key & CSR

$ openssl req -new \
    -config root-ca.conf \
    -out root-ca.csr \
    -keyout private/root-ca.key

self signed cert

$ openssl ca -selfsign \
    -config root-ca.conf \
    -in root-ca.csr \
    -out root-ca.crt \
    -extensions ca_ext

## s_client

allows you to connect to a secure server.
You will see the certificate information and chain up to the root.

```bash

openssl s_client -crlf \
-connect www.feistyduck.com:443 \
-servername www.feistyduck.com

```

add -showcerts to show all the certificates in the chain
