# certificates

## API Server

## Ingress

When an ingress is HTTPS then it requires a certificate and private key.
This is provided to the ingress object through a secret, a special kind of secret.

```bash

kubectl create secret tls my-tls-secret --key tls.key --cert tls.crt

```

1. The secret must contain a PEM encoded certificate and a private key
2. Intermediate certificates must be bundled with public key certificate
    This is because all verification done by the client is local, i.e. no need to make request to get other certificates.
    So, the only two places certificates (including intermediate) can be is in the local store CA store or in the certificates provided by the server
3. TLS secret and ingress need to be in the same namespace

The secret is defined for each hosts section of the ingress

```yaml

# ...
spec:
  ingressClassName: nginx
  tls:
    - hosts:
      - my.example.com
      secretName: my-tls-secret
  rules:
    - host: my-example.com
      http:
        paths:
          # ...

```

When the ingress object is created the certificate and private key are loaded from the secret into the Ingress Controller

Ingress Controller relies on server name indication extension to figure out which x509 certificate to present by way of identity.
Or is it the CN or Subject Alternative Name???


## cert-manager

is a kubernetes operator

define, request, apply, renew and remove are all parts of certificate management

goal: automate the mgmt of the lifecycle of TLS certificates within a K8S Cluster

benefits:
1. Better automation
    allowing us to focus on more important tasks
2. Minimize risk
    removing human element reduces risk
3. Reliable Outcomes
    increase confidence in achieving our goals

CRDs:
1. Issuer
    CA from which x.509 certs can be obtained
    namespace scoped
    There is also ClusterIssuer if cluster wide is required.
    types:
    * Self-signing
    * CA
    * Vault, hashicorp
    * ACME (Automatic Certificate Management Environment)
        is a protocol
2. CertificateRequest
    PEM encoded certificate request, CSR,  which is sent to a CA issuer that it references
3. Challenge
    Manages the lifecycle of the challenge presented by an ACME server for authorization
4. Certificate
    Defines the attributes of an X.509 cert to be retrieved from the CA associated with an issuer.
    namespace scoped
    use issuerRef field to refer back to Issuer, kind and name
    secretName fields refers to the tls secret which cert manager will create.
      This secretName is then referenced in the Ingress Object
    This can all be automated for us using annotations on the Ingress object.
      cert-manager watches Ingress Objects and if it has the annotation it will provision the certificates
      cert-manager.io/(cluster)-issuer: letsencrypt
5. Order
    Used to manage the order process for an X.509 certificate placed on an ACME server.
