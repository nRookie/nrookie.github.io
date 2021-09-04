## Understanding RAID


RAID is a form of data management that spreads your data across multiple drives. With a RAID, you can use multiple
drives, but your computer will recognize the RAID as one disk volume.

There are  various reasons why you may want to use RAID:


- Data Recovery: You can copy your data across multiple drives so that in the case of a drive failure you will be able to easily and quickly recover your data. To do this, you should use RAID 5.

- Backup: You can mirror your drives with RAID 1 so that you can remove one copy for safekeeping, preferably in a storage place off-site.

- Speed: For speed, you should choose RAID 0. The data added to a CRU enclosure will be split into parts and spread computer across all drives in the enclosure at the same time, which increases the speed of your data transfers. Your computer will see a single volume with roughly the capacity of all the drives in the enclosure.

## Introduction to different common RAID levels


### RAID 0
striping

Combines two or more hard drives together and treats them as one large volume. For example, two 250GB drives combined in a RAID 0 configuration creates a single 500GB volume. RAID 0 is used by those wanting the most speed out of two or more drives.


Because the data is split across both drives, the speed of data reading and writing increases as more disks are added.


Every drive has a limited lifespan and each disk adds another point of failure to the RAID. Every disk in a RAID 0 is critical – losing any of them means the entire RAID (and all of the data) is lost.

### RAID 1
mirroring


Mirroring creates an exact duplicate of a disk. Every time you write information to one drive, the exact information is written to the other drive in your mirror. Important files (accounting, financial, personal records) are commonly backed up with a RAID 1.

This is the safest option for your data. If one drive is lost, your data still exists in its complete form, and takes no time to recover.

Your investment in data safety increases your drive costs since each RAID 1 volume requires two drives.


### RAID 5

parity striping

A common RAID setup for volumes that are larger, faster, and more safe than any single drive. Your data is spread across all the drives in the RAID along with information that will allow your data to be recovered in case of a single drive failure. At least three drives are required for RAID 5. No matter how many drives are used, an amount equal to one of them will be used for the recovery data and cannot be used for user data.

You can lose any one disk and not lose your data. Just replace the disk with a new one.

Your investment in data safety increases your drive costs since at least three drives are needed.


## HARDWARE vs software RAID

RAID can be implemented in hardware, in the form of special disk controllers that are typically built into a multi-drive enclosure, or in software, with an operating system module that takes care of the housekeeping required for data to be written properly to the disks used in the RAID configuration.

The Windows, Mac OS X, and Linux operating systems all offer the ability to create a RAID configuration without any additional software. The drawback to using your operating system or other software to create a RAID is that it will add to the computational load on your computer, which will likely slow your computer’s performance.

Using a hardware RAID system, in an external drive enclosure or an expansion card installed in the computer, would not slow down your computer’s performance.



## HOW DO I RAID?

You need at least two hard drives. Some RAID levels require at least 3 disks, but some need 4 or 5. You’ll want to buy matching drives for your RAID, so plan accordingly.

If you attempt to RAID disks of different sizes together, most RAID methods treat each of the disks as if they are the same size as the smallest disk in the RAID.

You will also need a way to RAID your drives together, whether via hardware or software. Many CRU drive enclosures include RAID capability that can be configured on the enclosure itself so you don’t need additional software.

However, CRU does provide the CRU Configurator which is compatible with many of our RAID enclosures, which provides SMS and email notifications when a drive fails and allows you to view or update your device’s firmware as well as configure what kinds of events cause your enclosure’s audible alarm to sound.

