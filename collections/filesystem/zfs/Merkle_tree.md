# Merkle tree 

In cryptography and computer science, a hash tree or Merkle tree is a tree in which every leaf node is labelled with the cryptographic hash of a data block, and every non-leaf node is labelled with the cryptographic hash of the labels of its child nodes. Hash trees allow efficient and secure verification of the contents of large data structures. Hash trees are a generalization of hash lists and hash chains.


Demonstrating that a leaf node is a part of a given binary hash tree requires computing a number of hashes proportional to the logarithm of the number leaf nodes of the tree; this contrasts with hash lists, where the number is proportional to the number of leaf nodes itself. Merkle trees are therefore an efficient example of a cryptographic commitment scheme, in which the root of the Merkle tree is seen as a commitment and leaf nodes may be revealed and proven to be part of the original commitment.



# Uses 

Hash trees can be used to verify any kind of data stored, handled and transferred in and between computers. They can help ensure that data blocks received from other peers in a peer-to-peer network are received undamaged and unaltered. and even to check that the other peers do not lie and send fake blocks.

Hash trees are used in hash-based cryptography. Hash trees are also used in the IPFS, Btrfs and ZFS file systems (to counter data degradation); Dat protocol; Apache Wave protocol; Git and Mercurial distributed revision control system; the Tahoe-LAFS backup system; Seronet; 