# email

## protocols

### SMTP

Simple Message Transfer Protocol.
TCP/IP protocol used as an internet standard.
Uses plain english syntax.

### POP3

Post Office Protocol.
Used by email clients to fetch email.
by default, the protocol downloads the messages and then deletes them from the server.
Simpler but less flexible.

### IMAP

Internet Message Access Protocol.
Other main protocol used by email clients.
emails are managed on the server and left there.
Copies are downloaded to the client.
More complex and flexible than POP3

## email program roles and protocols

MUA
  Mail User agent
  main role of your email client
  read and compose email.
  implementations: Thunderbird, Mutt, Evolution, Outlook

MSP
  Mail Submission Program
  the role your email client has when you press send.
  submites your message to the MTA

MTA
  Mail Transfer Agent
  main role of your email server
  responsible for starting the processing of sending the message to the recipient
  sends the message to their MTA
  implementations: sendmail, exim, postfix

MDA
  Mail Delivery Agent
  role of your email server when it receives an email
  stores the message for future retrieval
  implementations: sendmail, postfix, procmail, sieve, cyrus, spam assassin

## Email life cycle

1. compose the email using your client, MUA
2. MUA connects to your outbound MTA via SMTP
3. Your outbound MTA connects to the inbound MTA of the recipient
4. Final MTA delivers message to MDA
5. MDA stores the message
6. recipient connects to their email server and retrieves the email.
