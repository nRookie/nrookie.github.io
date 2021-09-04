## ZFS

This article presents the notion of ZFS and the concepts that underlie it. ZFS stands for Zettabyte File System and is a next generation file system originally developed by Sun Microsystems for building next generation NAS solutions with better security, reliability and performance. ZFS was designed in 2001 by Matthew Ahrens and Jeff Bonwick and it was supposed to be a next generation file system for another Sun Microsystems’ system called OpenSolaris. A port for FreeBSD was made in 2008. Unlike other systems on the market, ZFS is a 128-bit file system offering virtually unlimited capacity. In turn, ZFS on Linux (ZOL) was created in 2013. Brian Behlendorf, Jorgen Lundman, Aron Xu, and Richard Yao were among those who helped to create and maintain ZOL.

After purchase of Sun Microsystems by Oracle, OpenSolaris was no longer an open-source project. Also, two-thirds of the system developers left Oracle at that time. In September 2013, OpenZFS project was founded by the developers who previously worked on OpenSolaris. They continued to work on an open-source implementation of ZFS.

ZFS is licensed under the Common Development and Distribution License that is incompatible with the GNU General Public License and this is why it cannot be included in the Linux kernel. The reason why the license is kept unchanged is the fact that it is nearly impossible to contact all of the contributors of the OpenZFS implementation.


ZFS comprises functions of a copy-on write file system, logical volume manager and software RAID for serving the purposes of highly scalable storage. To understand how ZFS works, you need to get familiarized with the basic concepts displayed on the diagram below.
