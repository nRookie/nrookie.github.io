## CR


LFS has just such a fixed place on disk for this, known as the checkpoint region (CR). The checkpoint region contains pointers to (i.e., addresses of) the latest pieces of the inode map, and thus the inode map pieces can be found by reading the CR first. Note the checkpoint region is only updated periodically (say every 30 seconds or so), and thus performance is not ill-affected. Thus, the overall structure of the on-disk layout contains a checkpoint region (which points to the latest pieces of the inode map). The inode map pieces each contain addresses of the inodes and the inodes point to files (and directories) just like typical UNIX file systems.

Here is an example of the checkpoint region (note it is all the way at the beginning of the disk, at address 0), and a single imap chunk, inode, and data block. A real file system would of course have a much bigger CR (indeed, it would have two, as we’ll come to understand later), many imap chunks, and of course many more inodes, data blocks, etc.





## Reading a File from Disk: A Recap
To make sure you understand how LFS works, let us now walk through what must happen to read a file from disk. Assume we have nothing in memory to begin. The first on-disk data structure we must read is the checkpoint region.

The checkpoint region contains pointers (i.e., disk addresses) to the entire inode map, and thus LFS then reads in the entire inode map and caches it in memory. After this point, when given an inode number of a file, LFS simply looks up the inode-number to inode-disk-address mapping in the imap, and reads in the most recent version of the inode. To read a block from the file, at this point, LFS proceeds exactly like a typical UNIX file system, by using direct pointers or indirect pointers or doubly-indirect pointers as need be. In the common case, LFS should perform the same number of I/Os as a typical file system when reading a file from the disk; the entire imap is cached and thus the extra work LFS does during a read is to look up the inode’s address in the imap.


##  Recursive update problem

There is one other serious problem in LFS that the inode map solves, known as the recursive update problem. The problem arises in any file system that never updates in place (such as LFS), but rather moves updates to new locations on the disk.

Specifically, whenever an inode is updated, its location on disk changes. If we hadn’t been careful, this would have also entailed an update to the directory that points to this file, which then would have mandated a change to the parent of that directory, and so on, all the way up the file system tree.

LFS cleverly avoids this problem with the inode map. Even though the location of an inode may change, the change is never reflected in the directory itself. Rather, the imap structure is updated while the directory holds the same name-to-inode-number mapping. Thus, through indirection, LFS avoids the recursive update problem.


## garbage collection
 a technique that arises in programming languages that automatically free unused memory for programs.

Earlier we discussed segments as important as they are the mechanism that enables large writes to disk in LFS. As it turns out, they are also quite integral to effective cleaning. Imagine what would happen if the LFS cleaner simply went through and freed single data blocks, inodes, etc., during cleaning. The result: a file system with some number of free holes mixed between allocated space on the disk. Write performance would drop considerably, as LFS would not be able to find a large contiguous region to write to disk sequentially and with high performance.

Instead, the LFS cleaner works on a segment-by-segment basis, thus clearing up large chunks of space for subsequent writing. The basic cleaning process works as follows. Periodically, the LFS cleaner reads in a number of old (partially-used) segments, determines which blocks are live within these segments, and then write out a new set of segments with just the live blocks within them, freeing up the old ones for writing. Specifically, we expect the cleaner to read in MMM existing segments, compact their contents into NNN new segments (where N<MN < MN<M ), and then write the NNN segments to disk in new locations. The old MMM segments are then freed and can be used by the file system for subsequent writes.

We are now left with two problems, however. The first is the mechanism: how can LFS tell which blocks within a segment are live, and which are dead? The second is policy: how often should the cleaner run, and which segments should it pick to clean?



## Determining Block Liveness

Given a data block DDD within an on-disk segment SSS, LFS must be able to determine whether DDD is live. To do so, LFS adds a little extra information to each segment that describes each block. Specifically, LFS includes, for each data block DDD, its inode number (which file it belongs to) and its offset (which block of the file this is). This information is recorded in a structure at the head of the segment known as the segment summary block.

Given this information, it is straightforward to determine whether a block is live or dead. For a block DDD located on disk at address AAA, look in the segment summary block and find its inode number NNN and offset TTT. Next, look in the imap to find where NNN lives and read NNN from disk (perhaps it is already in memory, which is even better). Finally, using the offset TTT, look in the inode (or some indirect block) to see where the inode thinks the Tth block of this file is on disk. If it points exactly to the disk address AAA, LFS can conclude that the block DDD is live. If it points anywhere else, LFS can conclude that DDD is not in use (i.e., it is dead) and thus know that this version is no longer needed.

Here is a pseudocode summary:

Here is a diagram depicting the mechanism, in which the segment summary block (marked SSSSSS) records that the data block at address A0A0A0 is actually a part of file kkk at offset 0. By checking the imap for kkk, you can find the inode, and see that it does indeed point to that location.

There are some shortcuts LFS takes to make the process of determining liveness more efficient. For example, when a file is truncated or deleted, LFS increases its version number and records the new version number in the imap. By also recording the version number in the on-disk segment, LFS can short circuit the longer check described above simply by comparing the on-disk version number with a version number in the imap, thus avoiding extra reads.



## Crash Recovery and the Log


One final problem: what happens if the system crashes while LFS is writing to disk? As you may recall in the previous chapter about journaling, crashes during updates are tricky for file systems, and thus something LFS must consider as well.

### When do crashes happen

During normal operation, LFS buffers write in a segment, and then (when the segment is full, or when some amount of time has elapsed), write the segment to disk. LFS organizes these writes in a log, i.e., the checkpoint region points to a head and tail segment, and each segment points to the next segment to be written. LFS also periodically updates the checkpoint region. Crashes could clearly happen during either of these operations (write to a segment, write to the CR). So how does LFS handle crashes during writes to these structures?


## How crashes are handled#


Let’s cover the second case first. To ensure that the CR update happens atomically, LFS actually keeps two CRs, one at either end of the disk, and writes to them alternately. LFS also implements a careful protocol when updating the CR with the latest pointers to the inode map and other information. Specifically, it first writes out a header (with timestamp), then the body of the CR, and then finally one last block (also with a timestamp). If the system crashes during a CR update, LFS can detect this by seeing an inconsistent pair of timestamps. LFS will always choose to use the most recent CR that has consistent timestamps, and thus consistent update of the CR is achieved.

Let’s now address the first case. Because LFS writes the CR every 30 seconds or so, the last consistent snapshot of the file system may be quite old. Thus, upon reboot, LFS can easily recover by simply reading in the checkpoint region, the imap pieces it points to, and subsequent files and directories. However, the last many seconds of updates would be lost.

## Roll forward

To improve upon this, LFS tries to rebuild many of those segments through a technique known as roll forward in the database community. The basic idea is to start with the last checkpoint region, find the end of the log (which is included in the CR), and then use that to read through the next segments and see if there are any valid updates within it. If there are, LFS updates the file system accordingly and thus recovers much of the data and metadata written since the last checkpoint. See Rosenblum’s award-winning dissertation for details.



## Summary

LFS introduces a new approach to updating the disk. Instead of overwriting files in places, LFS always writes to an unused portion of the disk, and then later reclaims that old space through cleaning. This approach, which in database systems is called shadow paging and in file-system-speak is sometimes called copy-on-write, enables highly efficient writing, as LFS can gather all updates into an in-memory segment and then write them out together sequentially.

The large writes that LFS generates are excellent for performance on many different devices. On hard drives, large writes ensure that positioning time is minimized. On parity-based RAIDs, such as RAID-4 and RAID-5, they avoid the small-write problem entirely. Recent research has even shown that large I/Os are required for high performance on Flash-based SSDs. Thus, perhaps surprisingly, LFS-style file systems may be an excellent choice even for these new mediums.


The downside to this approach is that it generates garbage; old copies of the data are scattered throughout the disk, and if one wants to reclaim such space for subsequent usage, one must clean old segments periodically. Cleaning became the focus of much controversy in LFS, and concerns over cleaning costs perhaps limited LFS’s initial impact on the field. However, some modern commercial file systems, including NetApp’s WAFL, Sun’s ZFS, and Linux btrfs, and even modern flash-based SSDs, adopt a similar copy-on-write approach to writing to disk, and thus the intellectual legacy of LFS lives on in these modern file systems. In particular, WAFL got around cleaning problems by turning them into a feature. By providing old versions of the file system via snapshots, users could access old files whenever they deleted current ones accidentally.


copy-on-write (yes, COW), and is used in a number of popular file systems, including Sun’s ZFS. This technique never overwrites files or directories in place; rather, it places new updates to previously unused locations on disk. After a number of updates are completed, COW file systems flip the root structure of the file system to include pointers to the newly updated structures. Doing so makes keeping the file system consistent straightforward. We’ll be learning more about this technique when we discuss the log-structured file system (LFS) in the next chapter; LFS is an early example of a COW.