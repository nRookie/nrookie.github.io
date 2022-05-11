# [Device assignment](https://libvirt.org/pci-addresses.html#id6)

When using VFIO to assign host devices to a guest, an additional caveat to keep in mind that the guest OS will base its decisions upon the *target address* (guest side) rather than the *source address* (host side).

For example, the domain XML snippet

```
<hostdev mode='subsystem' type='pci' managed='yes'>
  <driver name='vfio'/>
  <source>
    <address domain='0x0001' bus='0x08' slot='0x00' function='0x0'/>
  </source>
  <address type='pci' domain='0x0000' bus='0x00' slot='0x01' function='0x0'/>
</hostdev>
```

will result in the device showing up as 0000:00:01.0 in the guest OS rather than as 0001:08:00.1, which is the address of the device on the host.

Of course, all the rules and behaviors described above still apply.



https://diego.assencio.com/?index=649b7a71b35fc7ad41e03b6d0e825f07