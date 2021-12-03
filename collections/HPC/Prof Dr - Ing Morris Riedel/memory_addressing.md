





we offer details in this chapter on how Intel 80x86 microprocessors address memory chips and how Linux makes use of the available addressing circuits.





This is the first of three chapters related to memory management: Chapter 6, discusses how the kernel allocates main memory to itself, while Chapter 7, considers how linear addresses are assigned to processes.







# 2.1 Memory Addresses









Programmers casually refer to a memory address as the way to access the contents of a memory cell. But when dealing with Intel 80x86 microprocessors, we have to distinguish among three kinds of addresses:









## Logical address



Included in the machine language instructions to specify the address of an operand or of an instruction. This type of address embodies the well-known Intel segmented architecture that forces MS-DOS and Windows programmers to divide their programs into segments. Each logical address consists of a segment and an offset (or displacement) that denotes the distance from the start of the segment to the actual address.







## Linear address



A single 32-bit unsigned integer that can be used to address up to 4 GB, that is, up to 4,294,967,296 memory cells. Linear addresses are usually represented in hexadecimal notation; their values range from 0x00000000 to 0xffffffff.



## Physical address



Used to address memory cells included in memory chips. They correspond to the electrical signals sent along the address pins of the microprocessor to the memory bus. Physical addresses are represented as 32-bit unsigned integers.









The CPU control unit transforms a logical address into a linear address by means of a hardware circuit called a segmentation unit; successively, a second hardware circuit called a paging unit transforms the linear address into a physical address (see Figure 2-1).





![image-20211202163928286](/Users/user/playground/share/nrookie.github.io/collections/HPC/Prof Dr - Ing Morris Riedel/image-20211202163928286.png)





## 2.2 Segmentation in Hardware



Starting with the 80386 model, Intel microprocessors perform address translation in two different ways called real mode and protected mode. Real mode exists mostly to maintain processor compatibility with older models and to allow the operating system to bootstrap (see Appendix A, for a short description of real mode). We shall thus focus our attention on protected mode.



### 2.2.1 Segmentation Registers



A logical address consists of two parts: a segment identifier and an offset that specifies the relative address within the segment. The segment identifier is a 16-bit field called Segment Selector, while the offset is a 32-bit field.



To make it easy to retrieve segment selectors quickly, the processor provides segmentation registers whose only purpose is to hold Segment Selectors; these registers are called cs, ss, ds, es, fs, and gs. Although there are only six of them, a program can reuse the same segmentation register for different purposes by saving its content in memory and then restoring it later.





Three of the six segmentation registers have specific purposes:



cs

The code segment register, which points to a segment containing program instructions



ss

The stack segment register, which points to a segment containing the current program stack



ds



The data segment register, which points to a segment containing static and external data



The remaining three segmentation registers are general purpose and may refer to arbitrary segments.





The cs register has another important function: it includes a 2-bit field that specifies the Current Privilege Level (CPL) of the CPU. The value denotes the highest privilege level, while the value 3 denotes the lowest one. Linux uses only levels and 3, which are respectively called Kernel Mode and User Mode.



### 2.2.2 Segment Descriptors





Each segment is represented by an 8-byte Segment Descriptor (see Figure 2-2) that describes the segment characteristics. Segment Descriptors are stored either in the Global Descriptor Table (GDT ) or in the Local Descriptor Table (LDT ).



![image-20211202164413014](/Users/user/playground/share/nrookie.github.io/collections/HPC/Prof Dr - Ing Morris Riedel/image-20211202164413014.png)





Usually only one GDT is defined, while each process may have its own LDT. The address of the GDT in main memory is contained in the gdtr processor register and the address of the currently used LDT is contained in the ldtr processor register.



Each Segment Descriptor consists of the following fields:





- A 32-bit Base field that contains the linear address of the first byte of the segment.

- A G granularity flag: if it is cleared, the segment size is expressed in bytes; otherwise,

  it is expressed in multiples of 4096 bytes.

- A 20-bit Limit field that denotes the segment length in bytes. If G is set to 0, the size

  of a non-null segment may vary between 1 byte and 1 MB; otherwise, it may vary

  between 4 KB and 4 GB.

- An S system flag: if it is cleared, the segment is a system segment that stores kernel

  data structures; otherwise, it is a normal code or data segment.

- A 4-bit Type field that characterizes the segment type and its access rights. The

  following Segment Descriptor types are widely used:





#### Code Segment Descriptor



Indicates that the Segment Descriptor refers to a code segment; it may be included either in the GDT or in the LDT. The descriptor has the S flag set.



#### Data Segment Descriptor



Indicates that the Segment Descriptor refers to a data segment; it may be included either in the GDT or in the LDT. The descriptor has the S flag set. Stack segments are implemented by means of generic data segments.





#### Task State Segment Descriptor (TSSD)



Indicates that the Segment Descriptor refers to a Task State Segment (TSS), that is, a segment used to save the contents of the processor registers (see Section 3.2.2 in Chapter 3); it can appear only in the GDT. The corresponding Type field has the value 11 or 9, depending on whether the corresponding process is currently executing on the CPU. The S flag of such descriptors is set to 0.





#### Local Descriptor Table Descriptor (LDTD)





Indicates that the Segment Descriptor refers to a segment containing an LDT; it can appear only in the GDT. The corresponding Type field has the value 2. The S flag of such descriptors is set to 0.



- A DPL (Descriptor Privilege Level ) 2-bit field used to restrict accesses to the segment. It represents the minimal CPU privilege level requested for accessing the segment. Therefore, a segment with its DPL set to is accessible only when the CPL is 0, that is, in Kernel Mode, while a segment with its DPL set to 3 is accessible with every CPL value.
- A Segment-Present flag that is set to if the segment is currently not stored in main memory. Linux always sets this field to 1, since it never swaps out whole segments to disk.
- An additional flag called D or B depending on whether the segment contains code or data. Its meaning is slightly different in the two cases, but it is basically set if the addresses used as segment offsets are 32 bits long and it is cleared if they are 16 bits long (see the Intel manual for further details).
- A reserved bit (bit 53) always set to 0.
- An AVL flag that may be used by the operating system but is ignored in Linux.



### 2.2.3 Segment Selectors



To speed up the translation of logical addresses into linear addresses, the Intel processor provides an additional nonprogrammable register—that is, a register that cannot be set by a programmer—for each of the six programmable segmentation registers. Each nonprogrammable register contains the 8-byte Segment Descriptor (described in the previous section) specified by the Segment Selector contained in the corresponding segmentation register. Every time a Segment Selector is loaded in a segmentation register, the corresponding Segment Descriptor is loaded from memory into the matching nonprogrammable CPU register. From then on, translations of logical addresses referring to that segment can be performed without accessing the GDT or LDT stored in main memory; the processor can just refer directly to the CPU register containing the Segment Descriptor. Accesses to the GDT or LDT are necessary only when the contents of the segmentation register change (see Figure 2-3). Each Segment Selector includes the following fields:



- A 13-bit index (described further in the text following this list) that identifies the corresponding Segment Descriptor entry contained in the GDT or in the LDT
- A TI (Table Indicator) flag that specifies whether the Segment Descriptor is included in the GDT (TI = 0) or in the LDT (TI = 1)
- An RPL (Requestor Privilege Level ) 2-bit field, which is precisely the Current Privilege Level of the CPU when the corresponding Segment Selector is loaded into the cs register[1]

![image-20211202165604176](/Users/user/playground/share/nrookie.github.io/collections/HPC/Prof Dr - Ing Morris Riedel/image-20211202165604176.png)





Since a Segment Descriptor is 8 bytes long, its relative address inside the GDT or the LDT is obtained by multiplying the most significant 13 bits of the Segment Selector by 8. For instance, if the GDT is at 0x00020000 (the value stored in the gdtr register) and the index specified by the Segment Selector is 2, the address of the corresponding Segment Descriptor is 0x00020000 + (2 x 8), or 0x00020010.

The first entry of the GDT is always set to 0: this ensures that logical addresses with a null Segment Selector will be considered invalid, thus causing a processor exception. The maximum number of Segment Descriptors that can be stored in the GDT is thus 8191, that is, 213-1.





### 2.2.4 Segmentation Unit





1. Figure 2-4 shows in detail how a logical address is translated into a corresponding linear address. The segmentation unit performs the following operations:

- Examines the TI field of the Segment Selector, in order to determine which Descriptor Table stores the Segment Descriptor. This field indicates that the Descriptor is either in the GDT (in which case the segmentation unit gets the base linear address of the GDT from the gdtr register) or in the active LDT (in which case the segmentation unit gets the base linear address of that LDT from the ldtr register).
- Computes the address of the Segment Descriptor from the index field of the Segment Selector. The index field is multiplied by 8 (the size of a Segment Descriptor), and the result is added to the content of the gdtr or ldtr register.
- Adds to the Base field of the Segment Descriptor the offset of the logical address, thus obtains the linear address.

![image-20211202170253229](/Users/user/playground/share/nrookie.github.io/collections/HPC/Prof Dr - Ing Morris Riedel/image-20211202170253229.png)







## 2.3 Segmentation in Linux





Segmentation has been included in Intel microprocessors to encourage programmers to split their applications in logically related entities, such as subroutines or global and local data areas. However, Linux uses segmentation in a very limited way. In fact, segmentation and paging are somewhat redundant since both can be used to separate the physical address spaces of processes: **segmentation can assign a different linear address space to each process while paging can map the same linear address space into different physical address spaces**. Linux prefers paging to segmentation for the following reasons:

• Memory management is simpler when all processes use the same segment register values, that is, when they share the same set of linear addresses.

• One of the design objectives of Linux is portability to the most popular architectures; however, several RISC processors support segmentation in a very limited way.





The 2.2 version of Linux uses segmentation only when required by the Intel 80x86 architecture. In particular, all processes use the same logical addresses, so the total number of segments to be defined is quite limited and it is possible to store all Segment Descriptors in the Global Descriptor Table (GDT). This table is implemented by the array gdt_table referred by the gdt variable. If you look in the Source Code Index, you can see that these symbols are defined in the file arch/i386/kernel/head.S. Every macro, function, and other symbol in this book is listed in the appendix so you can quickly find it in the source code.



Local Descriptor Tables are not used by the kernel, although a system call exists that allows processes to create their own LDTs. This turns out to be useful to applications such as Wine that execute segment-oriented Microsoft Windows applications.



Here are the segments used by Linux:

- A kernel code segment. The fields of the corresponding Segment Descriptor in the GDT have the following values:

  o Base=0x00000000
   o Limit=0xfffff
   o G (granularity flag) = 1, for segment size expressed in pages o S (system flag) = 1, for normal code or data segment
   o Type = 0xa, for code segment that can be read and executed o DPL (Descriptor Privilege Level) = 0, for Kernel Mode
   o D/B (32-bit address flag) = 1, for 32-bit offset addresses

Thus, the linear addresses associated with that segment start at and reach the addressing limit of 232 - 1. The S and Type fields specify that the segment is a code segment that can be read and executed. Its DPL value is 0, thus it can be accessed only in Kernel Mode. The corresponding Segment Selector is defined by the __KERNEL_CS macro: in order to address the segment, the kernel just loads the value yielded by the macro into the cs register



- A kernel data segment. The fields of the corresponding Segment Descriptor in the GDT have the following values:

  o Base=0x00000000
   o Limit=0xfffff
   o G (granularity flag) = 1, for segment size expressed in pages o S (system flag) = 1, for normal code or data segment
   o Type = 2, for data segment that can be read and written
   o DPL (Descriptor Privilege Level) = 0, for Kernel Mode
   o D/B (32-bit address flag) = 1, for 32-bit offset addresses



1. This segment is identical to the previous one (in fact, they overlap in the linear address space) except for the value of the Type field, which specifies that it is a data segment that can be read and written. The corresponding Segment Selector is defined by the __KERNEL_DS macro.





- A user code segment shared by all processes in User Mode. The fields of the corresponding Segment Descriptor in the GDT have the following values:

  o Base=0x00000000
   o Limit=0xfffff
   o G (granularity flag) = 1, for segment size expressed in pages o S (system flag) = 1, for normal code or data segment
   o Type = 0xa, for code segment that can be read and executed o DPL (Descriptor Privilege Level) = 3, for User Mode
   o D/B (32-bit address flag) = 1, for 32-bit offset addresses



The S and DPL fields specify that the segment is not a system segment and that its privilege level is equal to 3; it can thus be accessed both in Kernel Mode and in User Mode. The corresponding Segment Selector is defined by the __USER_CS macro.



- A user data segment shared by all processes in User Mode. The fields of the corresponding Segment Descriptor in the GDT have the following values:

  o Base=0x00000000
   o Limit=0xfffff
   o G (granularity flag) = 1, for segment size expressed in pages o S (system flag) = 1, for normal code or data segment
   o Type = 2, for data segment that can be read and written
   o DPL (Descriptor Privilege Level) = 3, for User Mode
   o D/B (32-bit address flag) = 1, for 32-bit offset addresses







This segment overlaps the previous one: they are identical, except for the value of Type. The corresponding Segment Selector is defined by the __USER_DS macro.





- A Task State Segment (TSS) segment for each process. The descriptors of these segments are stored in the GDT. The Base field of the TSS descriptor associated with each process contains the address of the tss field of the corresponding process descriptor. The G flag is cleared, while the Limit field is set to 0xeb, since the TSS segment is 236 bytes long. The Type field is set to 9 or 11 (available 32-bit TSS), and the DPL is set to 0, since processes in User Mode are not allowed to access TSS segments.



- A default LDT segment that is usually shared by all processes. This segment is stored in the default_ldt variable. The default LDT includes a single entry consisting of a null Segment Descriptor. Each process has its own LDT Segment Descriptor, which usually points to the common default LDT segment. The Base field is set to the address of default_ldt and the Limit field is set to 7. If a process requires a real LDT, a new 4096-byte segment is created (it can include up to 511 Segment Descriptors), and the default LDT Segment Descriptor associated with that process is replaced in the GDT with a new descriptor with specific values for the Base and Limit fields.





For each process, therefore, the GDT contains two different Segment Descriptors: one for the TSS segment and one for the LDT segment. The maximum number of entries allowed in the GDT is 12+2xNR_TASKS, where, in turn, NR_TASKS denotes the maximum number of processes. In the previous list we described the six main Segment Descriptors used by Linux. Four additional Segment Descriptors cover Advanced Power Management (APM) features, and four entries of the GDT are left unused, for a grand total of 14.

As we mentioned before, the GDT can have at most 213 = 8192 entries, of which the first is always null. Since 14 are either unused or filled by the system, NR_TASKS cannot be larger than 8180/2 = 4090.



1. The TSS and LDT descriptors for each process are added to the GDT as the process is created. As we shall see in Section 3.3.2 in Chapter 3, the kernel itself spawns the first process: process running init_task . During kernel initialization, the trap_init( ) function inserts the TSS descriptor of this first process into the GDT using the statement:

   ```
   set_tss_desc(0, &init_task.tss);
   ```

2. The first process creates others, so that every subsequent process is the child of some existing process. The copy_thread( ) function, which is invoked from the clone( ) and fork( )

3. system calls to create new processes, executes the same function in order to set the TSS of the new process:

4. ```
   set_tss_desc(nr, &(task[nr]->tss));
   ```



Since each TSS descriptor refers to a different process, of course, each Base field has a different value. The copy_thread( ) function also invokes the set_ldt_desc( ) function in order to insert a Segment Descriptor in the GDT relative to the default LDT for the new process.



The kernel data segment includes a process descriptor for each process. Each process descriptor includes its own TSS segment and a pointer to its LDT segment, which is also located inside the kernel data segment.



As stated earlier, the Current Privilege Level of the CPU reflects whether the processor is in User or Kernel Mode and is specified by the RPL field of the Segment Selector stored in the cs register. Whenever the Current Privilege Level is changed, some segmentation registers must be correspondingly updated. For instance, when the CPL is equal to 3 (User Mode), the ds register must contain the Segment Selector of the user data segment, but when the CPL is equal to 0, the ds register must contain the Segment Selector of the kernel data segment.



A similar situation occurs for the ss register: it must refer to a User Mode stack inside the user data segment when the CPL is 3, and it must refer to a Kernel Mode stack inside the kernel data segment when the CPL is 0. When switching from User Mode to Kernel Mode, Linux always makes sure that the ss register contains the Segment Selector of the kernel data segment.





## 2.4 Paging in Hardware



The paging unit translates linear addresses into physical ones. It checks the requested access type against the access rights of the linear address. If the memory access is not valid, it generates a page fault exception (see Chapter 4, and Chapter 6).



For the sake of efficiency, linear addresses are grouped in fixed-length intervals called pages; contiguous linear addresses within a page are mapped into contiguous physical addresses. In this way, the kernel can specify the physical address and the access rights of a page instead of those of all the linear addresses included in it. Following the usual convention, we shall use the term "page" to refer both to a set of linear addresses and to the data contained in this group of addresses.





The paging unit thinks of all RAM as partitioned into fixed-length page frames (they are sometimes referred to as physical pages). Each page frame contains a page, that is, the length of a page frame coincides with that of a page. A page frame is a constituent of main memory, and hence it is a storage area. It is important to distinguish a page from a page frame: the former is just a block of data, which may be stored in any page frame or on disk.



The data structures that map linear to physical addresses are called page tables; they are stored in main memory and must be properly initialized by the kernel before enabling the paging unit.







In Intel processors, paging is enabled by setting the PG flag of the cr0 register. When PG = 0, linear addresses are interpreted as physical addresses.



### 2.4.1 Regular Paging



Starting with the i80386, the paging unit of Intel processors handles 4 KB pages. The 32 bits of a linear address are divided into three fields:



#### Directory

The most significant 10 bits



#### Table

The intermediate 10 bits



#### Offset

The least significant 12 bits





The translation of linear addresses is accomplished in two steps, each based on a type of translation table. The first translation table is called Page Directory and the second is called Page Table.



The physical address of the Page Directory in use is stored in the **cr3** processor register. The **Directory** field within the linear address determines **the entry in the Page Directory that points to the proper Page Table**. The address's **Table** field, in turn, determines **the entry in the Page Table that contains the physical address of the page frame containing the page**. The **Offset** field determines **the relative position within the page frame** (see Figure 2-5). Since it is 12 bits long, each page consists of 4096 bytes of data.



![image-20211202175914201](/Users/user/playground/share/nrookie.github.io/collections/HPC/Prof Dr - Ing Morris Riedel/image-20211202175914201.png)









Both the Directory and the Table fields are 10 bits long, so Page Directories and Page Tables can include up to 1024 entries. It follows that a Page Directory can address up to 1024 x 1024 x 4096=2^32 memory cells, as you'd expect in 32-bit addresses.

The entries of Page Directories and Page Tables have the same structure. Each entry includes the following fields:





#### Present flag



If it is set, the referred page (or Page Table) is contained in main memory; if the flag is 0, the page is not contained in main memory and the remaining entry bits may be used by the operating system for its own purposes. (We shall see in Chapter 16, how Linux makes use of this field.)



#### Field containing the 20 most significant bits of a page frame physical address



Since each page frame has a 4 KB capacity, its physical address must be a multiple of 4096, so the 12 least significant bits of the physical address are always equal to 0. If the field refers to a Page Directory, the page frame contains a Page Table; if it refers to a Page Table, the page frame contains a page of data.





#### Accessed flag



Is set each time the paging unit addresses the corresponding page frame. This flag may be used by the operating system when selecting pages to be swapped out. The paging unit never resets this flag; this must be done by the operating system.



#### Dirty flag



Applies only to the Page Table entries. It is set each time a write operation is performed on the page frame. As in the previous case, this flag may be used by the operating system when selecting pages to be swapped out. The paging unit never resets this flag; this must be done by the operating system.



#### Read/Write flag



Contains the access right (Read/Write or Read) of the page or of the Page Table (see

Section 2.4.3 later in this chapter).



#### User/Supervisor flag



Contains the privilege level required to access the page or Page Table (see Section 2.4.3).



#### Two flags called PCD and PWT



Control the way the page or Page Table is handled by the hardware cache (see Section 2.4.6 later in this chapter).



#### Page Size flag

Applies only to Page Directory entries. If it is set, the entry refers to a 4 MB long page

frame (see the following section).



If the entry of a Page Table or Page Directory needed to perform an address translation has the Present flag cleared, the paging unit stores the linear address in the cr2 processor register and generates the exception 14, that is, the "Page fault" exception.



#### 2.4.2 Extended Paging



Starting with the Pentium model, Intel 80x86 microprocessors introduce extended paging , which allows page frames to be either 4 KB or 4 MB in size (see Figure 2-6).





![image-20211202181346198](/Users/user/playground/share/nrookie.github.io/collections/HPC/Prof Dr - Ing Morris Riedel/image-20211202181346198.png)







As we have seen in the previous section, extended paging is enabled by setting the Page Size flag of a Page Directory entry. In this case, the paging unit divides the 32 bits of a linear address into two fields:



#### Directory

The most significant 10 bits

#### Offset

The remaining 22 bits



Page Directory entries for extended paging are the same as for normal paging, except that:



• The Page Size flag must be set.
 • Only the first 10 most significant bits of the 20-bit physical address field are

significant. This is because each physical address is aligned on a 4 MB boundary, so the 22 least significant bits of the address are 0.







Extended paging coexists with regular paging; it is enabled by setting the PSE flag of the cr4 processor register. Extended paging is used to translate large intervals of contiguous linear addresses into corresponding physical ones; in these cases, the kernel can do without intermediate Page Tables and thus save memory.



### 2.4.3 Hardware Protection Scheme





The paging unit uses a different protection scheme from the segmentation unit. While Intel processors allow four possible privilege levels to a segment, only two privilege levels are associated with pages and Page Tables, because privileges are controlled by the User/Supervisor flag mentioned in Section 2.4.1. When this flag is 0, the page can be addressed only when the CPL is less than 3 (this means, for Linux, when the processor is in Kernel Mode). When the flag is 1, the page can always be addressed.





Furthermore, instead of the three types of access rights (Read, Write, Execute) associated with segments, only two types of access rights (Read, Write) are associated with pages. If the Read/Write flag of a Page Directory or Page Table entry is equal to 0, the corresponding Page Table or page can only be read; otherwise it can be read and written.



### 2.4.4 An Example of Paging



A simple example will help in clarifying how paging works.

Let us assume that the kernel has assigned the linear address space between 0x20000000 and 0x2003ffff to a running process. This space consists of exactly 64 pages. We don't care about the physical addresses of the page frames containing the pages; in fact, some of them might not even be in main memory. We are interested only in the remaining fields of the page table entries.



Let us start with the 10 most significant bits of the linear addresses assigned to the process, which are interpreted as the Directory field by the paging unit. The addresses start with a 2 followed by zeros, so the 10 bits all have the same value, namely 0x080 or 128 decimal. Thus the Directory field in all the addresses refers to the 129th entry of the process Page Directory. The corresponding entry must contain the physical address of the Page Table assigned to the process (see Figure 2-7). If no other linear addresses are assigned to the process, all the remaining 1023 entries of the Page Directory are filled with zeros.



![image-20211202181900991](/Users/user/playground/share/nrookie.github.io/collections/HPC/Prof Dr - Ing Morris Riedel/image-20211202181900991.png)





0010 0000 00     00 0011 1111    1111 1111

  128                         63

  0x80.                     0x3f





The values assumed by the intermediate 10 bits, (that is, the values of the Table field) range from to 0x03f, or from to 63 decimal. Thus, only the first 64 entries of the Page Table are significant. The remaining 960 entries are filled with zeros.

Suppose that the process needs to read the byte at linear address 0x20021406. This address is handled by the paging unit as follows:

1. The Directory field 0x80 is used to select entry 0x80 of the Page Directory, which points to the Page Table associated with the process's pages.
2. The Table field 0x21 is used to select entry 0x21 of the Page Table, which points to the page frame containing the desired page.
3. Finally, the Offset field 0x406 is used to select the byte at offset 0x406 in the desired page frame.





If the Present flag of the 0x21 entry of the Page Table is cleared, the page is not present in main memory; in this case, the paging unit issues a page exception while translating the linear address. The same exception is issued whenever the process attempts to access linear addresses outside of the interval delimited by 0x20000000 and 0x2003ffff since the Page Table entries not assigned to the process are filled with zeros; in particular, their Present flags are all cleared.







### 2.4.5 Three-Level Paging



Two-level paging is used by 32-bit microprocessors. But in recent years, several microprocessors (such as Compaq's Alpha, and Sun's UltraSPARC) have adopted a 64-bit architecture. In this case, two-level paging is no longer suitable and it is necessary to move up to three-level paging. Let us use a thought experiment to see why.







Start by assuming about as large a page size as is reasonable (since you have to account for pages being transferred routinely to and from disk). Let's choose 16 KB for the page size. Since 1 KB covers a range of 210 addresses, 16 KB covers 214 addresses, so the Offset field would be 14 bits. This leaves 50 bits of the linear address to be distributed between the Table and the Directory fields. If we now decide to reserve 25 bits for each of these two fields, this means that both the Page Directory and the Page Tables of a process would include 225 entries, that is, more than 32 million entries.





Even if RAM is getting cheaper and cheaper, we cannot afford to waste so much memory space just for storing the page tables.







The solution chosen for Compaq's Alpha microprocessors is the following:





- Page frames are 8 KB long, so the Offset field is 13 bits long.

- Only the least significant 43 bits of an address are used. (The most significant 21 bits

  are always set 0.)

- Three levels of page tables are introduced so that the remaining 30 bits of the address

  can be split into three 10-bit fields (see Figure 2-9 later in this chapter). So the Page Tables include 210 = 1024 entries as in the two-level paging schema examined previously.





As we shall see in Section 2.5 later in this chapter, Linux's designers decided to implement a paging model inspired by the Alpha architecture.







### 2.4.6 Hardware Cache



Today's microprocessors have clock rates approaching gigahertz, while dynamic RAM (DRAM) chips have access times in the range of tens of clock cycles. This means that the CPU may be held back considerably while executing instructions that require fetching operands from RAM and/or storing results into RAM.



Hardware cache memories have been introduced to reduce the speed mismatch between CPU and RAM. They are based on the well-known locality principle, which holds both for programs and data structures: because of the cyclic structure of programs and the packing of related data into linear arrays, addresses close to the ones most recently used have a high probability of being used in the near future. It thus makes sense to introduce a smaller and faster memory that contains the most recently used code and data. For this purpose, a new unit called the line has been introduced into the Intel architecture. It consists of a few dozen contiguous bytes that are transferred in burst mode between the slow DRAM and the fast on- chip static RAM (SRAM) used to implement caches.



The cache is subdivided into subsets of lines. At one extreme the cache can be direct mapped, in which case a line in main memory is always stored at the exact same location in the cache. At the other extreme, the cache is fully associative, meaning that any line in memory can be stored at any location in the cache. But most caches are to some degree N-way associative, where any line of main memory can be stored in any one of N lines of the cache. For instance, a line of memory can be stored in two different lines of a 2-way set of associative cache.



As shown in Figure 2-8, the cache unit is inserted between the paging unit and the main memory. It includes both a hardware cache memory and a cache controller. The cache memory stores the actual lines of memory. The cache controller stores an array of entries, one entry for each line of the cache memory. Each entry includes a tag and a few flags that describe the status of the cache line. The tag consists of some bits that allow the cache controller to recognize the memory location currently mapped by the line. The bits of the memory physical address are usually split into three groups: the most significant ones correspond to the tag, the middle ones correspond to the cache controller subset index, the least significant ones to the offset within the line.s



![image-20211202195225835](/Users/user/playground/share/nrookie.github.io/collections/HPC/Prof Dr - Ing Morris Riedel/image-20211202195225835.png)

When accessing a RAM memory cell, the CPU extracts the subset index from the physical address and compares the tags of all lines in the subset with the high-order bits of the physical, address. If a line with the same tag as the high-order bits of the address is found, the CPU has

a cache hit; otherwise, it has a cache miss.



When a cache hit occurs, the cache controller behaves differently depending on access type. For a read operation, the controller selects the data from the cache line and transfers it into a CPU register; the RAM is not accessed and the CPU achieves the time saving for which the cache system was invented. For a write operation, the controller may implement one of two basic strategies called write-through and write-back. In a write-through, the controller always writes into both RAM and the cache line, effectively switching off the cache for write operations. In a write-back, which offers more immediate efficiency, only the cache line is updated, and the contents of the RAM are left unchanged. After a write-back, of course, the RAM must eventually be updated. The cache controller writes the cache line back into RAM only when the CPU executes an instruction requiring a flush of cache entries or when a FLUSH hardware signal occurs (usually after a cache miss).



When a cache miss occurs, the cache line is written to memory, if necessary, and the correct line is fetched from RAM into the cache entry.





Multiprocessor systems have a separate hardware cache for every processor, and therefore they need additional hardware circuitry to synchronize the cache contents. See Section 11.3.2 in Chapter 11.





Cache technology is rapidly evolving. For example, the first Pentium models included a single on-chip cache called the L1-cache. More recent models also include another larger and slower on-chip cache called the L2-cache. The consistency between the two cache levels is implemented at the hardware level. Linux ignores these hardware details and assumes there is a single cache.





The CD flag of the cr0 processor register is used to enable or disable the cache circuitry. The NW flag, in the same register, specifies whether the write-through or the write-back strategy is used for the caches.



Another interesting feature of the Pentium cache is that it lets an operating system associate a different cache management policy with each page frame. For that purpose, each Page Directory and each Page Table entry includes two flags: PCD specifies whether the cache must be enabled or disabled while accessing data included in the page frame; PWT specifies whether the write-back or the write-through strategy must be applied while writing data into the page frame. Linux clears the PCD and PWT flags of all Page Directory and Page Table entries: as a result, caching is enabled for all page frames and the write-back strategy is always adopted for writing.





The L1_CACHE_BYTES macro yields the size of a cache line on a Pentium, that is, 32 bytes. In order to optimize the cache hit rate, the kernel adopts the following rules:



- The most frequently used fields of a data structure are placed at the low offset within the data structure so that they can be cached in the same line.
- When allocating a large set of data structures, the kernel tries to store each of them in memory so that all cache lines are uniformly used.





### 2.4.7 Translation Lookaside Buffers (TLB)





Besides general-purpose hardware caches, Intel 80x86 processors include other caches called translation lookaside buffers or TLB to speed up linear address translation. When a linear address is used for the first time, the corresponding physical address is computed through slow accesses to the page tables in RAM. The physical address is then stored in a TLB entry, so that further references to the same linear address can be quickly translated.





The invlpg instruction can be used to invalidate (that is, to free) a single entry of a TLB. In order to invalidate all TLB entries, the processor can simply write into the cr3 register that points to the currently used Page Directory.



Since the TLBs serve as caches of page table contents, whenever a Page Table entry is modified, the kernel must invalidate the corresponding TLB entry. To do this, Linux makes use of the flush_tlb_page(addr) function, which invokes __flush_tlb_one( ). The latter function executes the invlpg Assembly instruction:



``` shell
movl $addr,%eax
invlpg (%eax)
```





Sometimes it is necessary to invalidate all TLB entries, such as during kernel initialization. In such cases, the kernel invokes the __flush_tlb( ) function, which rewrites the current value of cr3 back into it:





``` shell
movl %cr3, %eax
movl %eax, %cr3
```









## 2.5 Paging in Linux





As we explained in Section 2.4.5, Linux adopted a three-level paging model so paging is feasible on 64-bit architectures. Figure 2-9 shows the model, which defines three types of paging tables:





- Page Global Directory
- Page Middle Directory
- Page Table



The **Page Global Directory** includes the addresses of several Page Middle Directories, which in turn **include the addresses of several Page Tables**. Each Page Table entry points to a page frame. The linear address is thus split into four parts. Figure 2-9 does not show the bit numbers because the size of each part depends on the computer architecture.



![image-20211203103619125](/Users/user/playground/share/nrookie.github.io/collections/HPC/Prof Dr - Ing Morris Riedel/image-20211203103619125.png)





Linux handling of processes relies heavily on paging. In fact, the automatic translation of linear addresses into physical ones makes the following design objectives feasible:



- Assign a different physical address space to each process, thus ensuring an efficient protection against addressing errors.
- Distinguish pages, that is, groups of data, from page frames, that is, physical addresses in main memory. This allows the same page to be stored in a page frame, then saved to disk, and later reloaded in a different page frame. This is the basic ingredient of the virtual memory mechanism (see Chapter 16).





As we shall see in Chapter 7, each process has its own Page Global Directory and its own set of Page Tables. When a process switching occurs (see Section 3.2 in Chapter 3), Linux saves in a TSS segment the contents of the cr3 control register and loads from another TSS segment a new value into cr3. Thus, when the new process resumes its execution on the CPU, the paging unit refers to the correct set of page tables.



What happens when this three-level paging model is applied to the Pentium, which uses only two types of page tables? Linux essentially eliminates the Page Middle Directory field by saying that it contains zero bits. However, the position of the Page Middle Directory in the sequence of pointers is kept so that the same code can work on 32-bit and 64-bit architectures. The kernel keeps a position for the Page Middle Directory by setting the number of entries in it to 1 and mapping this single entry into the proper entry of the Page Global Directory.''



Mapping logical to linear addresses now becomes a mechanical task, although somewhat complex. The next few sections of this chapter are thus a rather tedious list of functions and macros that retrieve information the kernel needs to find addresses and manage the tables; most of the functions are one or two lines long. You may want to just skim these sections now, but it is useful to know the role of these functions and macros because you'll see them often in discussions in subsequent chapters.



### 2.5.1 The Linear Address Fields



The following macros simplify page table handling:





PAGE_SHIFT

Specifies the length in bits of the Offset field; when applied to Pentium processors it yields the value 12. Since all the addresses in a page must fit in the Offset field, the size of a page on Intel 80x86 systems is 212 or the familiar 4096 bytes; the PAGE_SHIFT of 12 can thus be considered the logarithm base 2 of the total page size. This macro is used by PAGE_SIZE to return the size of the page. Finally, the PAGE_MASK macro is defined as the value 0xfffff000; it is used to mask all the bits of the Offset field.



PMD_SHIFT



Determines the number of bits in an address that are mapped by the second-level page table. It yields the value 22 (12 from Offset plus 10 from Table). The PMD_SIZE macro computes the size of the area mapped by a single entry of the Page Middle Directory, that is, of a Page Table. Thus, PMD_SIZE yields 222 or 4 MB. The PMD_MASK macro yields the value 0xffc00000; it is used to mask all the bits of the Offset and Table fields.



PGDIR_SHIFT



Determines the logarithm of the size of the area a first-level page table can map. Since the Middle Directory field has length 0, this macro yields the same value yielded by PMD_SHIFT, which is 22. The PGDIR_SIZE macro computes the size of the area mapped by a single entry of the Page Global Directory, that is, of a Page Directory. PGDIR_SIZE therefore yields 4 MB. The PGDIR_MASK macro yields the value 0xffc00000, the same as PMD_MASK.



PTRS_PER_PTE , PTRS_PER_PMD , and PTRS_PER_PGD



Compute the number of entries in the Page Table, Page Middle Directory, and Page

Global Directory; they yield the values 1024, 1, and 1024, respectively.





### 2.5.2 Page Table Handling





pte_t, pmd_t, and pgd_t are 32-bit data types that describe, respectively, a Page Table, a Page Middle Directory, and a Page Global Directory entry. pgprot_t is another 32-bit data type that represents the protection flags associated with a single entry.

Four type-conversion macros (pte_val( ), pmd_val( ), pgd_val( ), and pgprot_val( )) cast a 32-bit unsigned integer into the required type. Four other type-conversion macros (__ pte( ), __ pmd( ), __ pgd( ), and __ pgprot( )) perform the reverse casting from one of the four previously mentioned specialized types into a 32-bit unsigned integer.

The kernel also provides several macros and functions to read or modify page table entries:



- The pte_none( ), pmd_none( ), and pgd_none( ) macros yield the value 1 if the corresponding entry has the value 0; otherwise, they yield the value 0.
- The pte_present( ), pmd_present( ), and pgd_present( ) macros yield the value 1 if the Present flag of the corresponding entry is equal to 1, that is, if the corresponding page or Page Table is loaded in main memory.
- The pte_clear( ), pmd_clear( ), and pgd_clear( ) macros clear an entry of the corresponding page table.





The macros pmd_bad( ) and pgd_bad( ) are used by functions to check Page Global Directory and Page Middle Directory entries passed as input parameters. Each macro yields the value 1 if the entry points to a bad page table, that is, if at least one of the following conditions applies:



- The page is not in main memory (Present flag cleared).

- The page allows only Read access (Read/Write flag cleared).

- Either Accessed or Dirty is cleared (Linux always forces these flags to be set for

  every existing page table).





No pte_bad( ) macro is defined because it is legal for a Page Table entry to refer to a page that is not present in main memory, not writable, or not accessible at all. Instead, several functions are offered to query the current value of any of the flags included in a Page Table entry:





pte_read( )

Returns the value of the User/Supervisor flag (indicating whether the page is accessible in User Mode).

pte_write( )

Returns 1 if both the Present and Read/Write flags are set (indicating whether the page is present and writable).

pte_exec( )

Returns the value of the User/Supervisor flag (indicating whether the page is accessible in User Mode). Notice that pages on the Intel processor cannot be protected against code execution.

pte_dirty( )

Returns the value of the Dirty flag (indicating whether or not the page has been modified).

pte_young( )

Returns the value of the Accessed flag (indicating whether the page has been accessed).



Another group of functions sets the value of the flags in a Page Table entry:





```
pte_wrprotect( )
```

Clears the Read/Write flag pte_rdprotect and pte_exprotect( )

Clear the User/Supervisor flag pte_mkwrite( )

Sets the Read/Write flag pte_mkread( ) and pte_mkexec( )

Set the User/Supervisor flag pte_mkdirty( ) and pte_mkclean( )

Set the Dirty flag to 1 and to 0, respectively, thus marking the page as modified or unmodified

pte_mkyoung( ) and pte_mkold( )
 Set the Accessed flag to 1 and to 0, respectively, thus marking the page as accessed

(young) or nonaccessed (old)

```
pte_modify(p,v)
```

Sets all access rights in a Page Table entry p to a specified value v set_pte

Writes a specified value into a Page Table entry





pte_ page( ) and pmd_ page( )
 Return the linear address of a page from its Page Table entry, and of a Page Table

from its Page Middle Directory entry.

```
pgd_offset(p,a)
```

Receives as parameters a memory descriptor p (see Chapter 6) and a linear address a. The macro yields the address of the entry in a Page Global Directory that corresponds to the address a; the Page Global Directory is found through a pointer within the memory descriptor p. The pgd_offset_k(o) macro is similar, except that it refers to the memory descriptor used by kernel threads (see Section 3.3.2 in Chapter 3).

```
pmd_offset(p,a)
```

Receives as parameter a Page Global Directory entry p and a linear address a; it yields the address of the entry corresponding to the address a in the Page Middle Directory referenced by p. The pte_offset(p,a) macro is similar, but p is a Page Middle Directory entry and the macro yields the address of the entry corresponding to a in the Page Table referenced by p.





The last group of functions of this long and rather boring list were introduced to simplify the creation and deletion of page table entries. When two-level paging is used, creating or deleting a Page Middle Directory entry is trivial. As we explained earlier in this section, the Page Middle Directory contains a single entry that points to the subordinate Page Table. Thus, the Page Middle Directory entry is the entry within the Page Global Directory too. When dealing with Page Tables, however, creating an entry may be more complex, because the Page Table that is supposed to contain it might not exist. In such cases, it is necessary to allocate a new page frame, fill it with zeros and finally add the entry.





Each page table is stored in one page frame; moreover, each process makes use of several page tables. As we shall see in Section 6.1 in Chapter 6, the allocations and deallocations of page frames are expensive operations. Therefore, when the kernel destroys a page table, it adds the corresponding page frame to a software cache. When the kernel must allocate a new page table, it takes a page frame contained in the cache; a new page frame is requested from the memory allocator only when the cache is empty.



The Page Table cache is a simple list of page frames. The pte_quicklist macro points to the head of the list, while the first 4 bytes of each page frame in the list are used as a pointer to the next element. The Page Global Directory cache is similar, but the head of the list is yielded by the pgd_quicklist macro. Of course, on Intel architecture there is no Page Middle Directory cache.



Since there is no limit on the size of the page table caches, the kernel must implement a mechanism for shrinking them. Therefore, the kernel introduces high and low watermarks, which are stored in the pgt_cache_water array; the check_pgt_cache( ) function checks whether the size of each cache is greater than the high watermark and, if so, deallocates page frames until the cache size reaches the low watermark. The check_ pgt_cache( ) is invoked either when the system is idle or when the kernel releases all page tables of some process.







### 2.5.3 Reserved Page Frames



The kernel's code and data structures are stored in a group of reserved page frames. A page contained in one of these page frames can never be dynamically assigned or swapped to disk.

As a general rule, the Linux kernel is installed in RAM starting from physical address 0x00100000, that is, from the second megabyte. The total number of page frames required depends on how the kernel has been configured: a typical configuration yields a kernel that can be loaded in less than 2 MBs of RAM.





Why isn't the kernel loaded starting with the first available megabyte of RAM? Well, the PC architecture has several peculiarities that must be taken into account:







- Page frame is used by BIOS to store the system hardware configuration detected during the Power-On Self-Test (POST ).
- Physical addresses ranging from 0x000a0000 to 0x000fffff are reserved to BIOS routines and to map the internal memory of ISA graphics cards (the source of the well- known 640 KB addressing limit in the first MS-DOS systems).
- Additional page frames within the first megabyte may be reserved by specific computer models. For example, the IBM ThinkPad maps the 0xa0 page frame into the 0x9f one.





In order to avoid loading the kernel into groups of noncontiguous page frames, Linux prefers to skip the first megabyte of RAM. Clearly, page frames not reserved by the PC architecture will be used by Linux to store dynamically assigned pages.





Figure 2-10 shows how the first 2 MB of RAM are filled by Linux. We have assumed that the kernel requires less than one megabyte of RAM (this is a bit optimistic).



![image-20211203105834679](/Users/user/playground/share/nrookie.github.io/collections/HPC/Prof Dr - Ing Morris Riedel/image-20211203105834679.png)



The symbol _text, which corresponds to physical address 0x00100000, denotes the address of the first byte of kernel code. The end of the kernel code is similarly identified by the symbol _etext. Kernel data is divided into two groups: initialized and uninitialized. The initialized data starts right after _etext and ends at _edata. The uninitialized data follows and ends up at _end.





The symbols appearing in the figure are not defined in Linux source code; they are produced while compiling the kernel.[2]





The linear address corresponding to the first physical address reserved to the BIOS or to a hardware device (usually, 0x0009f000) is stored in the i386_endbase variable. In most cases, this variable is initialized with a value written by the BIOS during the POST phase.



### 2.5.4 Process Page Tables



The linear address space of a process is divided into two parts:



- Linear addresses from 0x00000000 to PAGE_OFFSET -1 can be addressed when the process is in either User or Kernel Mode.
- Linear addresses from PAGE_OFFSET to 0xffffffff can be addressed only when the process is in Kernel Mode.



Usually, the PAGE_OFFSET macro yields the value 0xc0000000: this means that the fourth gigabyte of linear addresses is reserved for the kernel, while the first three gigabytes are accessible from both the kernel and the user programs. However, the value of PAGE_OFFSET may be customized by the user when the Linux kernel image is compiled. In fact, as we shall see in the next section, the range of linear addresses reserved for the kernel must include a mapping of all physical RAM installed in the system; moreover, as we shall see in Chapter 7, the kernel also makes use of the linear addresses in this range to remap noncontiguous page frames into contiguous linear addresses. Therefore, if Linux must be installed on a machine having a huge amount of RAM, a different arrangement for the linear addresses might be necessary.



The content of the first entries of the Page Global Directory that map linear addresses lower than PAGE_OFFSET (usually the first 768 entries) depends on the specific process. Conversely,

the remaining entries are the same for all processes; they are equal to the corresponding

entries of the swapper_ pg_dir kernel Page Global Directory (see the following section).





### 2.5.5 Kernel Page Tables



We now describe how the kernel initializes its own page tables. This is a two-phase activity. In fact, right after the kernel image has been loaded into memory, the CPU is still running in real mode; thus, paging is not enabled.



In the first phase, the kernel creates a limited 4 MB address space, which is enough for it to install itself in RAM.



In the second phase, the kernel takes advantage of all of the existing RAM and sets up the paging tables properly. The next section examines how this plan is executed.





#### 2.5.5.1 Provisional kernel page tables



Both the Page Global Directory and the Page Table are initialized statically during the kernel compilation. We won't bother mentioning the Page Middle Directories any more since they equate to Page Global Directory entries.



The Page Global Directory is contained in the **swapper_ pg_dir** variable, while the Page Table that spans the first 4 MB of RAM is contained in the **pg0** variable.



The objective of this first phase of paging is to allow these 4 MB to be **easily addressed in both real mode and protected mode**. Therefore, the kernel must create a mapping from both the **linear addresses 0x00000000 through 0x003fffff** and the **linear addresses PAGE_OFFSET through PAGE_OFFSET+0x3fffff into the physical addresses 0x00000000 through 0x003fffff.** In other words, the kernel during its first phase of initialization can address the first 4 MB of RAM (0x00000000 through 0x003fffff) either using linear addresses identical to the physical ones or using 4 MB worth of linear addresses starting from PAGE_OFFSET.



Assuming that PAGE_OFFSET yields the value 0xc0000000, the kernel creates the desired mapping by filling all the swapper_ pg_dir entries with zeros, except for entries and 0x300 (decimal 768); the latter entry spans all linear addresses between 0xc0000000 and 0xc03fffff. The and 0x300 entries are initialized as follows:



- The address field is set to the address of pg0.
- The Present, Read/Write, and User/Supervisor flags are set.
- The Accessed, Dirty, PCD, PWD, and Page Size flags are cleared.



The single pg0 Page Table is also statically initialized, so that the i th entry addresses the i th page frame.







The paging unit is enabled by the startup_32( ) Assembly-language function. This is achieved by loading in the cr3 control register the address of swapper_pg_dir and by setting the PG flag of the cr0 control register, as shown in the following excerpt:







``` shell
movl $0x101000,%eax
movl %eax,%cr3
movl %cr0,%eax
orl $0x80000000,%eax
movl %eax,%cr0
/* set the page table pointer.. */
/* ..and set paging (PG) bit */
```







#### 2.5.5.2 Final kernel page table



The final mapping provided by the kernel page tables must transform linear addresses starting from PAGE_OFFSET into physical addresses starting from 0.

The _ pa macro is used to convert a linear address starting from PAGE_OFFSET to the corresponding physical address, while the _va macro does the reverse.

The final kernel Page Global Directory is still stored in swapper_ pg_dir. It is initialized by the paging_init( ) function. This function acts on two input parameters:





start_mem

The linear address of the first byte of RAM right after the kernel code and data areas.

end_mem

The linear address of the end of memory (this address is computed by the BIOS routines during the POST phase).





Linux exploits the extended paging feature of the Pentium processors, enabling 4 MB page frames: it allows a very efficient mapping from PAGE_OFFSET into physical addresses by making kernel Page Tables superfluous.[3]





The swapper_ pg_dir Page Global Directory is reinitialized by a cycle equivalent to the following:

```
address = 0;
pg_dir = swapper_pg_dir;
pgd_val(pg_dir[0]) = 0;
pg_dir += (PAGE_OFFSET >> PGDIR_SHIFT);
while (address < end_mem) {
    pgd_val(*pg_dir) = _PAGE_PRESENT+_PAGE_RW+_PAGE_ACCESSED
           +_PAGE_DIRTY +_PAGE_4M+__pa(address);
```

pg_dir++;

```
    address += 0x400000;
}
```

As you can see, the first entry of the Page Global Directory is zeroed out, hence removing the mapping between the first 4 MB of linear and physical addresses. The first Page Table is thus available, so User Mode processes can also use the range of linear addresses between and 4194303.



The User/Supervisor flags in all Page Global Directory entries referencing linear addresses above PAGE_OFFSET are cleared, thus denying to processes in User Mode access to the kernel address space.

The pg0 provisional Page Table is no longer used once swapper_ pg_dir has been initialized.





