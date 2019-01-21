FROM alpine

RUN apk add socat gzip qemu-img qemu qemu-system-x86_64

RUN mkdir -p /opt/qemu/{sock,images,volumes}
RUN mkdir -p /opt/qemu/local/images

CMD qemu-system-x86_64 -enable-kvm -m 4096 -cpu Penryn,kvm=off,vendor=GenuineIntel,vmx,rdtscp \
          -vnc 0.0.0.0:0 \
	  -machine pc-q35-2.4 \
	  -smp 4,cores=2 \
	  -usb -device usb-kbd -device usb-mouse \
	  -device isa-applesmc,osk="ourhardworkbythesewordsguardedpleasedontsteal(c)AppleComputerInc" \
	  -kernel /opt/qemu/images/macos10.d/macos10kernel \
	  -smbios type=2 \
	  -device ich9-intel-hda -device hda-duplex \
	  -device ide-drive,bus=ide.2,drive=MacHDD \
	  -drive id=MacHDD,if=none,snapshot=on,file=/opt/qemu/images/macos10.qcow2 \
	  -netdev user,id=net0,hostfwd=tcp::22-:22,hostfwd=tcp::873-:873,hostfwd=tcp::3389-:3389 \
	  -device e1000-82545em,netdev=net0,id=net0,mac=52:54:00:c9:18:27 \
	  -monitor stdio

