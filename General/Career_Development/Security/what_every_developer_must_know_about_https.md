# what every developer must know about https

## why

Man in the middle attacks
DNS Hijacking
Phishing
Malicious host file entries

all not possible when using https

Confidentiality, integriry, authenticity

## perceived barriers

Cost
overhead on infrastructure
complexity
speed
compatibility

## HTTPS Fundamentals

### CA

Issue certificates, which we load into our websites
You can drill down in browser to who the CA is
A hugh number of CAs
every machine has a local list of CAs
more can be downloaded as required

### SSL & TLS

Secure Sockets Layer
version 2.0
version 3.0 last major release
Poodle attack in 2014

Transport Later Security
intended to be the successor to SSL 3.0
current standard
1.3 is still in draft
TLS can fall back to previous versions

just use TLS term
they're not the same but terms are used interchangeably

#### TLS Handshake

client hello (HTTPS CONNECT) (talks about creating a tunnel)
  highest level of tls
  cipher
server hello
  provides public key to client
client verifies public key
client key exchange
  clients key is encrypted with servers public key
server finished

now all comms is encrypted

### Developing & Debugging with HTTPS

Fiddler software, debugger tools in browser have replaced this I'd say

### Testing Bad HTTPS Implementations

badssl.com

## Securing the application

### redirecting from HTTP to HTTPS

client HTTP Request
server response with HTTP 301, permanent redirect
  tells the client they need to make a subsequent request
  location header with new url
client HTTPS request

the first request is still unsecured so chance of a MITM attack

HTTP Strict Transport Security (HSTS / STS) is the answer to this problem
returns a header stric-transport-security with a max-age field.
until max age expires, the browser will not make an unsecure request. It will always internally redirect before sending the request.
  So the browser handles the switch to HTTPS

Subsequent HTTP requests return a 307, internal redirect

TOFU: Trust On First Use
There's still an insecure request going out first before getting the HSTS

you can preload HSTS, using another attribute in the header.
submit your url to a site, hstspreload.org, your url is then added to a list which is in each browser.
TOFU problem is solved.

### Mixed Content

page loaded over HTTPS but it requested and insecure resource (image, script, ...)
Parts of your page can't be trusted even if they are loaded over HTTPS because you are retrieving them insecurely.
Use the relative path of resources, so it uses the scheme of the requester

active / passive content.
passive doesn't do anything on the page, so it can't do much bad

active
embed an iframe can have scripts in them

### CSP

Content Security Policy meta tag
you can upgrade the http scheme to https use using this tag in the html -> head -> meta
Can also be put into a response header
It's better supported now

You should also go through your site and update to https for all resources.

CSP can also block unsecure content

### Cookies

You can set a flag on cookies so that they are only sent over https

## over coming the preceived barriers

### impact on server performance

istlsfastyet.com

negligible!
Adam Langley, Google:

less than 1% of cpu load
less then 10kb memory per connection
less than 2% of network load

httpvshttps.com!

HTTPS is faster!

  HTTPS is over Http/2
  allows for a binary stream of data in parallel in http/2

## let's encrypt

free, automated and open CA!
Linuc foundation project

certbot, run on target system and talk back and over to let's encrypt

certificates only last 3 months
but it's automated so it doesn't matter

increasing number of hosting providers incorporate let's encrypt into their environment

cloud flare also provides same functionality as let's encrypt, and then more
acts as a revers proxy
  edge nodes
  people connect to edge nodes, and then request gets forwarded to actual website
  this allows for caching also.
  can terminate tls at the edge and host certificates there

## beyond the basics

### TLS & HTTPS Acronyms

SNI
  Server Name Indication
  multiple certs on the one IP address
SAN
  Subject Alternative Name
  multiple domain names on the one cert
PFS
  Perfect Forward Secrecy
  protects past sessions against future key compromise
DNSSEC
  DNS SECurity Extensions
  protects against forged DNS records
DANE
  DNS Based Authentication of Named Entities
  specifices certificate keys at DNS level
CAA
  Certificate Authority Authorization
  White list CAs which are allowed to sign at the DNS level
CRL
  Certificate Revocation List
  list of revoked certs
OCSP
  Online Certificate Status Protocol
  alternative to CRL
PKP
  Public Key Pinning
  predefined list of public keys

### Public Key Pinning

the certificate is valid, but is it the correct certificate

define public keys which are permissable for a domain
specifies max age
does it apply to sub domains
facilitates reporting of violations

Response Header
  contains the hash of the public keys
  max-age in seconds
  includeSubDomains
  report-uri, path where client should submit reports

if the certs public key is not in the response header reject

HPKP can't preload like HSTS

### SSL Labs

helps you test your TLS
enter a hostname
gives the hostname a rating based on how it can be communicated with

### Extended Validation Certificates

certificates say site is secure but says nothing about who is behind the site.
A phishing site could be secure
it's so easy to create a secure site

When you see a company name in the green box that means they have an Extended Validation Certificate
Go through a much more rigurous process to prove who they are

### other reasons to adopt HTTPS

SEO, search engine optimization
  Google gives tou a ranking bump if you are HTTPS

Referrer header
  don't add referrer header to non secure site
  but only uri is visible

Only use Brotli compression
  from Google

### blind spots

if there's a data breach on the site

going to a http site from https leaves you open to mitm

going to https from http, who says a mitm won't change the address of the login page
