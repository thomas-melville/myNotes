# disk encryption

encryption should be used when ever sensitive data is being stored and transmitted.
block device level encryption is one of the strongest protections.

## luks

Linux Unified Key Setup
installed on top of cryptsetup
dm-crypt is the kernel module

to mount an encrypted partition an entry must be placed in /etc/crypttab

## commands

### cryptsetup

apparently it's a swiss army knife program :-)

```bash
sudo cryptsetup  [options...] partition
```