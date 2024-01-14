
Install mergerfs and snapraid
```
mkdir nas
cd nas
sudo apt-get -y install gcc
sudo apt-get -y install make
sudo mkdir /var/lib/snapraid
sudo chmod a+w /var/lib/snapraid
cd /var/lib/snapraid
wget -O snapraid.tar.gz https://github.com/amadvance/snapraid/releases/download/v12.2/snapraid-12.2.tar.gz
tar -xzf snapraid.tar.gz
cd snapraid
./configure
make
make check
sudo make install
cd .. & cd .. && rm /var/lib/snapraid/*.tar.gz
cd ~/nas

wget -O mergerfs.deb https://github.com/trapexit/mergerfs/releases/download/2.36.0/mergerfs_2.36.0.ubuntu-kinetic_amd64.deb
sudo dpkg -i mergerfs.deb
rm mergerfs.deb
```

Steps to create snapraid and mergerfs
1. sudo gdisk /dev/sda 
use the following entires to proceed o n w
2. sudo mkfs.ext4 /dev/sda1
3. sudo  nano /etc/fstab
get the drive id
  3.1 blkid
  3.2 

# mount drives
UUID="1614edf0-af9e-4879-b01c-e9c6934f955c"  /mnt/disk1 ext4 nofail 0 0
UUID="21cc6b7e-e2da-40bb-b670-22b73e2e7e13" /mnt/disk2 ext4 nofail 0 0

# mount the parity 
UUID="3b140e17-e513-4cb6-9b1a-f5385187a94a" /mnt/parity ext4 nofail 0 0

# merferfs
/mnt/disk* /mnt/pool fuse.mergerfs defaults,allow_other,use_ino,hard_remove 0 0



4. change the the snapraid.conf
    sudo nano /etc/snapraid.conf


    `
    # Example configuration for snapraid

# Defines the file to use as parity storage
# It must NOT be in a data disk
# Format: "parity FILE [,FILE] ..."
parity /mnt/parity/snapraid.parity

# Defines the files to use as additional parity storage.
# If specified, they enable the multiple failures protection
# from two to six level of parity.
# To enable, uncomment one parity file for each level of extra
# protection required. Start from 2-parity, and follow in order.
# It must NOT be in a data disk
# Format: "X-parity FILE [,FILE] ..."
#2-parity /mnt/diskq/snapraid.2-parity
#3-parity /mnt/diskr/snapraid.3-parity
#4-parity /mnt/disks/snapraid.4-parity
#5-parity /mnt/diskt/snapraid.5-parity
#6-parity /mnt/disku/snapraid.6-parity

# Defines the files to use as content list
# You can use multiple specification to store more copies
# You must have least one copy for each parity file plus one. Some more don't hurt
# They can be in the disks used for data, parity or boot,
# but each file must be in a different disk
# Format: "content FILE"
content /mnt/disk1/snapraid.content
content /mnt/disk2/snapraid.content

# Defines the data disks to use
# The name and mount point association is relevant for parity, do not change it
# WARNING: Adding here your /home, /var or /tmp disks is NOT a good idea!
# SnapRAID is better suited for files that rarely changes!
# Format: "data DISK_NAME DISK_MOUNT_POINT"
data d1 /mnt/disk1/
data d2 /mnt/disk2/

# Excludes hidden files and directories (uncomment to enable).
#nohidden

# Defines files and directories to exclude
# Remember that all the paths are relative at the mount points
# Format: "exclude FILE"
# Format: "exclude DIR/"
# Format: "exclude /PATH/FILE"
# Format: "exclude /PATH/DIR/"
exclude *.unrecoverable
exclude /tmp/
exclude /lost+found/

# Defines the block size in kibi bytes (1024 bytes) (uncomment to enable).
# WARNING: Changing this value is for experts only!
# Default value is 256 -> 256 kibi bytes -> 262144 bytes
# Format: "blocksize SIZE_IN_KiB"
#blocksize 256

# Defines the hash size in bytes (uncomment to enable).
# WARNING: Changing this value is for experts only!
# Default value is 16 -> 128 bits
# Format: "hashsize SIZE_IN_BYTES"
#hashsize 16

# Automatically save the state when syncing after the specified amount
# of GB processed (uncomment to enable).
# This option is useful to avoid to restart from scratch long 'sync'
# commands interrupted by a machine crash.
# It also improves the recovering if a disk break during a 'sync'.
# Default value is 0, meaning disabled.
# Format: "autosave SIZE_IN_GB"
#autosave 500

# Defines the pooling directory where the virtual view of the disk
# array is created using the "pool" command (uncomment to enable).
# The files are not really copied here, but just linked using
# symbolic links.
# This directory must be outside the array.
# Format: "pool DIR"
#pool /pool

# Defines a custom smartctl command to obtain the SMART attributes
# for each disk. This may be required for RAID controllers and for
# some USB disk that cannot be autodetected.
# In the specified options, the "%s" string is replaced by the device name.
# Refers at the smartmontools documentation about the possible options:
# RAID -> https://www.smartmontools.org/wiki/Supported_RAID-Controllers
# USB -> https://www.smartmontools.org/wiki/Supported_USB-Devices
#smartctl d1 -d sat %s
#smartctl d2 -d usbjmicron %s
#smartctl parity -d areca,1/1 /dev/sg0
#smartctl 2-parity -d areca,2/1 /dev/sg0
    `



## add following to cron
`crontab -e`
`@reboot /home/jupiter/api`