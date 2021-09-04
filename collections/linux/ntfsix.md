## What is NTFS and How does it Work ?



NT file system (NTFS), which is also sometimes called the New Technology File System, is a process
that the Window NT operating uses for storing, organizing, and finding files on a hard disk efficiently.


NTFS was first introduced in 1993, as apart of the windows NT 3.1 release.

The benefits of NTFS are that , compared to other similar file systems like File Allocation Table (FAT) 
and High-Performance File System (HPFS), NTFS focuses on;

- Performance: NTFS allows file compression so your organization can enjoy increased storage space on a disk.

- Security access control: NTFS will enable you to place permissions on files and folders so you can restrict access to mission-critical data.

- Reliability: NTFS focuses on the consistency of the file system so that in the event of a disaster (such as a power loss or system failure), you can quickly restore your data.


- Disk space utilization: In addition to file compression, NTFS also allows disk quotas. This feature enables businesses to have even more control over storage space.

- File system journaling: This means that you can easily keep a log of-and audit-the files added, modified,or deleted on a drive. This log is called the Master File Table(MFT).


## How Does NTFS Work

The technical breakdown of NTFS is as follow.

1. A hard disk is formatted.
2. A file gets divided into partitions within the hard disk.
3. Within each partition, the operating system tracks every file stored in a specific operating system
4. Each file is distributed and stored in one or more clusters or disk spaces of a predefined uniform size (on the hard disk).
5. The size of each cluster will range from 512 bytes to 64 kilobytes.


You can control the size of a cluster based on what's most important to your organization:

- Efficient use of disk space
- The number of disk accesses required to access a file.





# ntfsfix

fix common errors and force windows to check NTFS

## SYNOPSIS

``` shell
ntfsfix [options] device
```

## description

``` shell
       ntfsfix is a utility that fixes some common NTFS problems.  ntfsfix is NOT a Linux version
       of chkdsk.  It only repairs some fundamental NTFS inconsistencies, resets the NTFS journal
       file and schedules an NTFS consistency check for the first boot into Windows.

       You may run ntfsfix on an NTFS volume if you think it was damaged by Windows or some other
       way and it cannot be mounted.
```

# mkntfs

create an NTFS file systm

## SYNOPSIS

``` shell
       mkntfs [options] device [number-of-sectors]

       mkntfs  [  -C  ]  [  -c  cluster-size  ]  [  -F  ]  [ -f ] [ -H heads ] [ -h ] [ -I ] [ -L
       volume-label ] [ -l ] [ -n ] [ -p part-start-sect ] [ -Q ] [ -q ] [ -S sectors-per-track ]
       [  -s  sector-size  ]  [  -T ] [ -U ] [ -V ] [ -v ] [ -z mft-zone-multiplier ] [ --debug ]
       device [ number-of-sectors ]

```




In computer file systems, a cluster(sometimes also called allocation unit or block) is a unit of disk space allocation for files and directories. To reduce the overhead of managing on-disk data structures, the filesystem does not allocate individual disk sectors by default, but contiguous groups of sectors, called clusters.

On a disk that uses 512-byte sectors, a 512-byte cluster contains one sector, whereas a 4-kibibyte(KiB) cluster contains eight sectors.

A cluster is the smallest logical amount of disk space that can be allocated to hold a file. Storing small files on a filesystem with large clusters will therefore waste disk space; such wasted disk space is called slack space. For cluster sizes which are small versus the average file size, the wasted space per file will be statistically about half of the cluster size; for large cluster sizes , the wasted space will become greater.