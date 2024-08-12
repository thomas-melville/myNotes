def creds = com.cloudbees.plugins.credentials.CredentialsProvider.lookupCredentials(
    com.cloudbees.plugins.credentials.impl.BaseStandardCredentials.class,
    Jenkins.instance,
    null,
    null
);
println("All credentials")
for (c in creds) {
     println( ( "ID: " + c.id + ", Type: " + c))
}

println("String credentials")
creds = com.cloudbees.plugins.credentials.CredentialsProvider.lookupCredentials(
    org.jenkinsci.plugins.plaincredentials.impl.StringCredentialsImpl,
    Jenkins.instance,
    null,
    null
);
for (c in creds) {
     println( ( "ID: " + c.id + ", Secret: " + c.secret ))
}

println("Username, password credentials")
creds = com.cloudbees.plugins.credentials.CredentialsProvider.lookupCredentials(
    com.cloudbees.plugins.credentials.impl.UsernamePasswordCredentialsImpl,
    Jenkins.instance,
    null,
    null
);
for (c in creds) {
     println( ( "ID: " + c.id + ", Username: " + c.username + ", Password: " + c.password ))
}

println("File credentials")
creds = com.cloudbees.plugins.credentials.CredentialsProvider.lookupCredentials(
    org.jenkinsci.plugins.plaincredentials.impl.FileCredentialsImpl,
    Jenkins.instance,
    null,
    null
);
println("Printing")
for (c in creds) {
     println( ( "ID: " + c.id + ", Filename: " + c.fileName + ", Content: " + c.getContent().text ))
}

println("SSH User Private key credentials")
creds = com.cloudbees.plugins.credentials.CredentialsProvider.lookupCredentials(
    com.cloudbees.jenkins.plugins.sshcredentials.impl.BasicSSHUserPrivateKey,
    Jenkins.instance,
    null,
    null
);
for (c in creds) {
     println( ( "ID: " + c.id + ", PassPhrase: " + c.passphrase + ", privateKeys: " + c.privateKeys))
}
