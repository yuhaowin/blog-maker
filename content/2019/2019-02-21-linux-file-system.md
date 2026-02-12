# Linux 文件系统

### 虚拟文件系统

![134318](https://image.yuhaowin.com/2020/08/27/134318.jpg)

https://www.ibm.com/developerworks/cn/linux/l-cn-vfs/index.html

inode 号
pagecache：多个程序读取通一个文件，它们的操作系统中是，共享一份 pagecache 的。
pagecache 有 dirty 的概念，当某个应用程序修改了 pagecache 时，此时该pagecache 会被标记为 dirty，这些脏页是需要被写入（flush）到磁盘中去的，这个
flush 的时机可以交由操作系统处理。也可以是用户程序通过系统调用立刻进行 flush。

不同应用程序在对同一份文件进行读取的时候，每个应用程序持有一个 FD ，这个 FD 有自己读取文件的指针，标记当前应用程序读取文件的位置。

命令： df：显示每个文件所在的文件系统的信息,或全部默认文件系统。

mount unmount

一切皆文件，文件是为了 I/O 读写。

文件类型分类：

+ : - 普通文件，可以是图片、文本、可执行文件等。
+ ：d 目录
+ ：l 链接 软连接、硬链接，相对于是引用，修改任意一方，其他都可见，软连接不会增加 inode 号
+ ：b 块设备 如硬盘
  +：c 字符设备 如键盘，无法读取到过去和未来的。
  +：s socket
  +：p pipeline

dd if=/dev/zero of=mydisk.img bs=1048576 count=100
losetup /dev/loop0 mydisk.img
mke2fs /dev/loop0

lsof - list open files

lsof -p $PID

FD 文件描述符
0 文件标准输入 0u u表示可读可写
1 文件标准输出
2 文件错误输出

exec 8<  yuhao.txt

exec 8<> /dev/tcp/www.baidu.com/80
cd /proc/$$/fd
lsof -po $$

重定向：不是命令，是一种机制。
ls / xx 1> test 2>&1

管道

{ echo $BASHPID; read x; } | { cat; }
{ echo $BASHPID; read x; echo $x; } | { cat; }

### strace 命令

应用程序写数据，是先写在操作系统的 pagecache 中，pagecache 会不定期写入磁盘中的。