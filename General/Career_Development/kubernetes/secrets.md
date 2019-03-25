# secrets

Secrets are used to make sensitive data available to Pods, passwords for example, in a secure way.
i.e. not plain text.

```bash
kubectl get secrets

kubectl create secret ...
```

You can define it in a yaml file too.

K8S automatically creates a secret for the default service account in a namespace to allow it access to the API server (read only)

Secrets are not encrypted by default, just base64 encoded.
There's some command you can run on the terminal to decode the secret.

```bash
base64 --decode vnmjkrtvwtnmtjk
```

To encrypt secrets you must create an EncryptionConfiguration object with a key & proper identity.
You then need to configure the kube-apiserver to encrypt secrets then re-create every secret. As they are encrypted on write.

A secret can be used as an environment variable in a pod

```yaml
...
  env:
  - name: PASSWORD
    valueFrom:
      secretKeyRef:
        name: mysql
        key: password

```

A secret is limited to 1MB in size.
They are stored in the tmpfs storage on the host node.

Secrets can also be mounted as files using a volume definition in a pod manifest.

imagePullSecret is a way to pass a secret that contains a Docker image registry password
