# keystore V truststore

https://www.baeldung.com/java-keystore-truststore-difference

we need these when our app needs to communicate over SSL/TLS
i.e. we are a server and want to communicate over HTTPS
password protected files which sit on the same FS.
Default format up to java8 is JKS
java9 is PKSC12
The difference is JKS is specific to java, PKSC12 is language neutral

## keystore

Stores keys. private key entries, certificates with public keys or secret keys
we can use these for various cryptographic purposes.
Each one is stored by an alias for easy lookup.
Generally, it holds keys that our application owns.
we can use these to prove the integrity of messages and the authenticity of the sender
During the SSL handshake, the server looks up the private key from the keystore, and presents the corresponding public key and certificate to the client.
The client can also have a keystore.

javax.net.ssl.KeyStore / KeyStorePassword / KeyStoreType

## truststore

holds onto certificates that identify others.
we use it to trust the 3rd party we are about to communicate with.
If the certificate or certificate authorities from the sender is not in the truststore we get an SSLHandshakeException
java has bundled a truststore called cacerts in JAVA_HOME/jre/lib/security
(on my system its a symbolic link to /etc/ssl/certs/java/cacerts)
It contains default, trusted CAs
You can use keytool to see the contents.

javax.net.ssl.TrustStorePassword / TrustStoreType
