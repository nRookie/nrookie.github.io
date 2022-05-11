1. check the virsh list

2. ``` shell
   virsh list --all | grep uuid
   ```

3. virsh dumpxml

4. ``` shell
   sudo virsh dumpxml 93e9cae7-af19-4a4a-bb0d-7f77587b83a6
   <domain type='kvm'>
     <name>93e9cae7-af19-4a4a-bb0d-7f77587b83a6</name>
     <uuid>93e9cae7-af19-4a4a-bb0d-7f77587b83a6</uuid>
     <memory unit='KiB'>16777216</memory>
     <currentMemory unit='KiB'>16777216</currentMemory>
     <memoryBacking>
       <hugepages/>
       <access mode='shared'/>
     </memoryBacking>
     <vcpu placement='static'>8</vcpu>
     <numatune>
       <memnode cellid='0' mode='strict' nodeset='0'/>
     </numatune>
     <os>
       <type arch='x86_64' machine='pc-i440fx-UCLOUD-O-1.0.0'>hvm</type>
       <boot dev='hd'/>
       <bootmenu enable='yes'/>
     </os>
     <features>
       <acpi/>
       <apic/>
       <pae/>
     </features>
     <cpu mode='custom' match='exact' check='partial'>
       <model fallback='forbid'>EPYC</model>
       <topology sockets='1' cores='4' threads='2'/>
       <numa>
         <cell id='0' cpus='0-7' memory='16777216' unit='KiB'/>
       </numa>
     </cpu>
     <clock offset='timezone' timezone='Asia/Shanghai'>
       <timer name='hypervclock' present='yes'/>
       <timer name='rtc' tickpolicy='catchup' track='guest'/>
     </clock>
     <on_poweroff>destroy</on_poweroff>
     <on_reboot>restart</on_reboot>
     <on_crash>restart</on_crash>
     <devices>
       <emulator>/usr/libexec/qemu-kvm</emulator>
       <disk type='network' device='disk'>
         <driver name='qemu' type='raw' cache='writethrough'/>
         <source protocol='udisk' name='InnerUDisk-yuz4cj10'>
           <host name='block-udisk' port='12000'/>
         </source>
         <target dev='vda' bus='virtio'/>
         <iotune>
           <total_bytes_sec>104857600</total_bytes_sec>
           <total_iops_sec>2400</total_iops_sec>
         </iotune>
         <serial>UCLOUD_DISK_VDA</serial>
         <address type='pci' domain='0x0000' bus='0x00' slot='0x06' function='0x0'/>
       </disk>
       <controller type='usb' index='0' model='piix3-uhci'>
         <address type='pci' domain='0x0000' bus='0x00' slot='0x01' function='0x2'/>
       </controller>
       <controller type='pci' index='0' model='pci-root'/>
       <interface type='bridge'>
         <mac address='52:54:00:da:8f:b6'/>
         <source bridge='br0'/>
         <virtualport type='openvswitch'>
           <parameters interfaceid='881b389b-be59-4ffe-960d-52251d104f0c'/>
         </virtualport>
         <model type='virtio'/>
         <driver name='vhost'/>
         <address type='pci' domain='0x0000' bus='0x00' slot='0x05' function='0x0'/>
       </interface>
       <serial type='pty'>
         <target type='isa-serial' port='0'>
           <model name='isa-serial'/>
         </target>
       </serial>
       <console type='pty'>
         <target type='serial' port='0'/>
       </console>
       <input type='tablet' bus='usb'>
         <address type='usb' bus='0' port='1'/>
       </input>
       <input type='mouse' bus='ps2'/>
       <input type='keyboard' bus='ps2'/>
       <graphics type='vnc' port='-1' autoport='yes'>
         <listen type='address'/>
       </graphics>
       <video>
         <model type='vga' vram='65536' heads='1' primary='yes'/>
         <address type='pci' domain='0x0000' bus='0x00' slot='0x02' function='0x0'/>
       </video>
       <hostdev mode='subsystem' type='mdev' managed='no' model='vfio-pci'>
         <source>
           <address uuid='00688158-0001-0004-0004-010066144074'/>
         </source>
         <address type='pci' domain='0x0000' bus='0x00' slot='0x03' function='0x0'/>
       </hostdev>
       <memballoon model='virtio'>
         <address type='pci' domain='0x0000' bus='0x00' slot='0x04' function='0x0'/>
       </memballoon>
     </devices>
   </domain>
   
   ```



![image-20220310155456047](/Users/user/playground/share/nrookie.github.io/collections/KVM/image-20220310155456047.png)

![image-20220310160204660](/Users/user/playground/share/nrookie.github.io/collections/KVM/image-20220310160204660.png)







### take it apart



``` xml
<domain type='kvm'>
  <name>93e9cae7-af19-4a4a-bb0d-7f77587b83a6</name>
  <uuid>93e9cae7-af19-4a4a-bb0d-7f77587b83a6</uuid>
  <memory unit='KiB'>16777216</memory>
  <currentMemory unit='KiB'>16777216</currentMemory>
  <memoryBacking>
    <hugepages/>
    <access mode='shared'/>
  </memoryBacking>
  <vcpu placement='static'>8</vcpu>
  <numatune>
    <memnode cellid='0' mode='strict' nodeset='0'/>
  </numatune>
  <os>
    <type arch='x86_64' machine='pc-i440fx-UCLOUD-O-1.0.0'>hvm</type>
    <boot dev='hd'/>
    <bootmenu enable='yes'/>
  </os>
  <features>
    <acpi/>
    <apic/>
    <pae/>
  </features>
  <cpu mode='custom' match='exact' check='partial'>
    <model fallback='forbid'>EPYC</model>
    <topology sockets='1' cores='4' threads='2'/>
    <numa>
      <cell id='0' cpus='0-7' memory='16777216' unit='KiB'/>
    </numa>
  </cpu>
  <clock offset='timezone' timezone='Asia/Shanghai'>
    <timer name='hypervclock' present='yes'/>
    <timer name='rtc' tickpolicy='catchup' track='guest'/>
  </clock>
  <on_poweroff>destroy</on_poweroff>
  <on_reboot>restart</on_reboot>
  <on_crash>restart</on_crash>
  <devices>
    <emulator>/usr/libexec/qemu-kvm</emulator>
    <disk type='network' device='disk'>
      <driver name='qemu' type='raw' cache='writethrough'/>
      <source protocol='udisk' name='InnerUDisk-yuz4cj10'>
        <host name='block-udisk' port='12000'/>
      </source>
      <target dev='vda' bus='virtio'/>
      <iotune>
        <total_bytes_sec>104857600</total_bytes_sec>
        <total_iops_sec>2400</total_iops_sec>
      </iotune>
      <serial>UCLOUD_DISK_VDA</serial>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x06' function='0x0'/>
    </disk>
    
    <controller type='usb' index='0' model='piix3-uhci'>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x01' function='0x2'/>
    </controller>
    <controller type='pci' index='0' model='pci-root'/>
    <interface type='bridge'>
      <mac address='52:54:00:da:8f:b6'/>
      <source bridge='br0'/>
      <virtualport type='openvswitch'>
        <parameters interfaceid='881b389b-be59-4ffe-960d-52251d104f0c'/>
      </virtualport>
      <model type='virtio'/>
      <driver name='vhost'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x05' function='0x0'/>
    </interface>
    <serial type='pty'>
      <target type='isa-serial' port='0'>
        <model name='isa-serial'/>
      </target>
    </serial>
    <console type='pty'>
      <target type='serial' port='0'/>
    </console>
    <input type='tablet' bus='usb'>
      <address type='usb' bus='0' port='1'/>
    </input>
    <input type='mouse' bus='ps2'/>
    <input type='keyboard' bus='ps2'/>
    <graphics type='vnc' port='-1' autoport='yes'>
      <listen type='address'/>
    </graphics>
    <video>
      <model type='vga' vram='65536' heads='1' primary='yes'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x02' function='0x0'/>
    </video>
    <hostdev mode='subsystem' type='mdev' managed='no' model='vfio-pci'>
      <source>
        <address uuid='00688158-0001-0004-0004-010066144074'/>
      </source>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x03' function='0x0'/>
    </hostdev>
    <memballoon model='virtio'>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x04' function='0x0'/>
    </memballoon>
  </devices>
</domain>
```





USB, PCI and SCSI devices attached to the host can be passed through to the guest using the hosted element. 



**hostdev** 

The hostdev element is the main container for describing host devices. For each device, the mode is always "subsystem" and the type is one of the following values with additional attributes noted.



**mdev**

For mediated devices ( **Since 3.2.0** ) the model attribute specifies the device API which determines how the host's vfio driver will expose the device to the guest. Currently, model='vfio-pci', model='vfio-ccw' ( **Since 4.4.0** ) and model='vfio-ap' ( **Since 4.9.0** ) is supported. [MDEV](https://libvirt.org/drvnodedev.html#MDEV) section provides more information about mediated devices as well as how to create mediated devices on the host. **Since 4.6.0 (QEMU 2.12)** an optional display attribute may be used to enable or disable support for an accelerated remote desktop backed by a mediated device (such as NVIDIA vGPU or Intel GVT-g) as an alternative to emulated [video devices](https://libvirt.org/formatdomain.html#elementsVideo). This attribute is limited to model='vfio-pci' only. Supported values are either on or off (default is 'off'). It is required to use a [graphical framebuffer](https://libvirt.org/formatdomain.html#elementsGraphics) in order to use this attribute, currently only supported with VNC, Spice and egl-headless graphics devices. **Since version 5.10.0** , there is an optional ramfb attribute for devices with model='vfio-pci'. Supported values are either on or off (default is 'off'). When enabled, this attribute provides a memory framebuffer device to the guest. This framebuffer will be used as a boot display when a vgpu device is the primary display.

Note: There are also some implications on the usage of guest's address type depending on the model attribute, see the address element below.

**source**

The source element describes the device as seen from the host using the following mechanism to describe:



**mdev**

Mediated devices ( **Since 3.2.0** ) are described by the address element. The address element contains a single mandatory attribute uuid.



**address**

The address element for USB devices has a bus and device attribute to specify the USB bus and device number the device appears at on the host. The values of these attributes can be given in decimal, hexadecimal (starting with 0x) or octal (starting with 0) form. For PCI devices the element carries 4 attributes allowing to designate the device as can be found with the lspci or with virsh nodedev-list. For SCSI devices a 'drive' address type must be used. For mediated devices, which are software-only devices defining an allocation of resources on the physical parent device, the address type used must conform to the model attribute of element hostdev, e.g. any address type other than PCI for vfio-pci device API or any address type other than CCW for vfio-ccw device API will result in an error. [See above](https://libvirt.org/formatdomain.html#elementsAddress) for more details on the address element.



![image-20220310165623274](/Users/user/playground/share/nrookie.github.io/collections/KVM/image-20220310165623274.png)





https://libvirt.org/formatdomain.html#cpu-allocation







