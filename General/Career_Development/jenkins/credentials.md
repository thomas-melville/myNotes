def creds = com.cloudbees.plugins.credentials.CredentialsProvider.lookupCredentials(
    com.cloudbees.plugins.credentials.impl.BaseStandardCredentials.class,
    Jenkins.instance,
    null,
    null
);
for (c in creds) {
     println( ( "ID: " + c.id + ", Type: " + c))
}
