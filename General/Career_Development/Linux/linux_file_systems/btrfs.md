# btrfs

B-Tree FS

intended to address lack of pooling, snapshots, checksums and integral multi-device spanning in other FSs such as ext4
They features can be crucial for Linux to scale into large enterprise storage configurations

released in kernel version 2.6.29, but still regarded as experimental

can take frequent snapshots of the entire FS in virtually no time
because it makes extensive use of COW (Copy on write) features.