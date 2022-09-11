sudo rmmod ram.ko
sudo rmmod cpu.ko
sudo dmesg -C
make clean
make
sudo insmod ram.ko
sudo insmod cpu.ko
sudo dmesg 