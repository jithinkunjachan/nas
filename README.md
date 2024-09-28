## Installation of the snapraid and mergerfs
```
mkdir nas
cd nas
sudo apt-get -y install gcc
sudo apt-get -y install make
sudo mkdir /var/lib/snapraid
sudo chmod a+w /var/lib/snapraid
cd /var/lib/snapraid
wget -O snapraid.tar.gz https://github.com/amadvance/snapraid/releases/download/v12.3/snapraid-12.3.tar.gz
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


### Setting up confiugration
1. Format drive by to proceed 

   ```sudo gdisk /dev/sda```

    then presess ```o n w```
     
3. Make file system by
   
    ```sudo mkfs.ext4 /dev/sda1```
5. Auto mount drives by
  
   ```sudo  nano /etc/fstab``` get drive id using  ```blkid /dev/sda1``` 


```bash

UUID="66f3176d-2eed-45ad-8c07-048d85080a55"    /mnt/data1    ext4    defaults    0    0
UUID="d749be2a-e83a-4e19-bc61-56546abbaa8e"    /mnt/data2    ext4    defaults    0    0
UUID="43aca3fd-9b1c-4182-bf9b-496a12a46fc0"    /mnt/parity1    xfs    defaults    0    2

/mnt/data*    /mnt/pool    fuse.mergerfs    allow_other,use_ino,cache.files=partial,dropcacheonclose=true,category.create=mfs    0    0

```

4. change the the snapraid.conf
    ```sudo nano /etc/snapraid.conf```

```bash
parity /mnt/parity1/snapraid.parity
data d1 /mnt/data1/
data d2 /mnt/data2/
content /mnt/data1/.snapraid.content
content /mnt/data2/.snapraid.content

exclude *.bak
exclude *.unrecoverable
exclude /tmp/
exclude /lost+found/
exclude .AppleDouble
exclude ._AppleDouble
exclude .DS_Store
exclude .Thumbs.db
exclude .fseventsd
exclude .Spotlight-V100
exclude .TemporaryItems
exclude .Trashes
exclude .AppleDB

autosave 100


```

5. Install smarttools
``` sudo apt-get install smartmontools```

### Server Installation
```crontab -e @reboot /home/jupiter/api```

### Installing SAMBA
```sudo apt update```

```sudo apt install samba```

```sudo nano /etc/samba/smb.conf```


At the bottom of the file, add the following lines:
```
[rinyjithin]
    comment = Samba on Ubuntu
    path = /mnt/pool
    read only = no
    browsable = yes
```

``` sudo service smbd restart```

```sudo ufw allow samba```

```sudo smbpasswd -a rj```