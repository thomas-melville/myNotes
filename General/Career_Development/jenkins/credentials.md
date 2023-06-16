def creds = com.cloudbees.plugins.credentials.CredentialsProvider.lookupCredentials(
    com.cloudbees.plugins.credentials.impl.BaseStandardCredentials.class,
    Jenkins.instance,
    null,
    null
);
for (c in creds) {
     println( ( "ID: " + c.id + ", Type: " + c))
}

creds = com.cloudbees.plugins.credentials.CredentialsProvider.lookupCredentials(
    org.jenkinsci.plugins.plaincredentials.impl.StringCredentialsImpl,
    Jenkins.instance,
    null,
    null
);
for (c in creds) {
     println( ( "ID: " + c.id + ", Secret: " + c.secret ))
}

creds = com.cloudbees.plugins.credentials.CredentialsProvider.lookupCredentials(
    com.cloudbees.plugins.credentials.impl.UsernamePasswordCredentialsImpl,
    Jenkins.instance,
    null,
    null
);
for (c in creds) {
     println( ( "ID: " + c.id + ", Username: " + c.username + ", Password: " + c.password))
}
