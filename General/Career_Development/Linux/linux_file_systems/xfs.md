# XFS

originally designed by SGI and used in the IRIX OS, subsequently ported to linux.
explicitly engineered to deal with large data sets, as well as handle i/o tasks very effectively
    16EB for the whole system
    8EB for an individual file

high perf is one of the key design elements, which implements methods for:
    Employing DMA (Direct Memory Access) i/o
    guaranteeing i/o rate
    adjusting stripe size to match underlying RAID and LVM devices

can also journal quota information

most maintenance tasks can be done while the fs is mounted, making it easier.
    defrag
    enlarge
    dump/restore

## commands

these commands can be paused and resumed later.
they are also multi threaded so can be completed quickly.

### xfsdump

dump the fs

### xfsrestore

restore the fs