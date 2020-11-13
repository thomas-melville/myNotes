# Useful Cryptography

## Hashing

what is a hash function?

takes in a string, outputs another string
always deterministic
unique input will give unqiue output
irreversible

example: don't store a password, store the hash of the password.
  That way there is no risk to the password getting into the wrong hands
  but you can still verify login requests by hashing the password in the login request and comparing it to the stored hash

2 types of hashing algorithms

1. cryptopgrahic hash functions
  really fast
  a lot of package managers use this in the background
  useful for verifying the integrity of data

  sha1sum is a command line tool to generate hash/digest

  MD5, 1991, don't use it
    two separate inputs can give the same output

  SHA-1, 1995, don't use it
    two separate inputs can give the same output

  SHA-2, 2001
    no collision

  SHA-3, 2015
    no collisions
    faster than sha-2

  BLAKE, 2012
    fastest

  SHA-2 is the recommended choice


2. password hash functions
  opposite of crypto
  slow
  don't want hackers to be able to brute force through variants until it matches a hash

  PBKDF2, 2000
    fast today
    people now hash the hash multiple times

  bcrypt, 1999
    variable work factor to slow the algorithm down

  scrypt, 2009
    2nd parameter, specify memory
    so cpu & ram is required

  argon2, 2015
    adds a 3rd dimension
    specify threads
    3 flavours

## Randomness

the best way to generate random numbers is /dev/urandom

OS Kernel, pulls in sources of randomness
  keyboard timings
  mouse movements
  storage timings

into a randonm pool
  CSPRNG

/dev/random
  blocks until there is enough bytes
/dev/urandom
  does not block
  continues to generate random numbers

## encryption

encrypted text is known as cipher text

2 primary types of encryption

symmetric
asymmetric

### symmetric encryption

generate secret key

key + data into algorihm = cipher text

same key is used for encryption and decryption

useful in cases where you can keep the trusted secret safe!!

Amazon KMS
a lot of the cloud providers have their own

KMS
  create master key for yourself
  create data key for a piece of data
  data + data key into algorithm = cipher text
  master key + data key into algorithm = encrypted data key
  encrypted data key + cipher text = encrypted message

  called enveloped encryption
  no need to keep data key safe

### asymmetric encryption

2 keys for each party involved
  public key
  private key
    can never leave your computer

used in tls

handshake
  exchange public certs
  aysmmetrically generate a shared key
  symmetric with the shared key after that
    sym is quicker than asym

NaCl/libsodiou
  Box API
