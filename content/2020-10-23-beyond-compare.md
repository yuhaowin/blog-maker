# Beyond Compare for Mac 试用

[官网下载地址](http://www.scootersoftware.com/download.php)

### 无限试用的原理

> BCompare是应用程序启动的程序，只要在在启动的时候删除registry.dat(Library/Application Support/Beyond Compare/registry.dat)注册信息就好了，为此可以在该目录下添加一个批处理文件用来处理这个操作。

### 无限试用的操作方法

1、打开命令行终端，进入到安装目录里面的Contents/Macos，命令行指令：
```shell
cd /Applications/Beyond\ Compare.app/Contents/MacOS/
```
2、修改可执行文件名，并创建脚本
```shell
mv BCompare BCompare.real # 把可执行文件改名
touch BCompare # 创建新的启动脚本文件
vim BCompare # 编辑脚本内容，内容如下所示
```
3、编辑 BCompare 文件，内容如下：
```shell
#!/bin/bash
rm "/Users/$(whoami)/Library/Application Support/Beyond Compare/registry.dat"
"`dirname "$0"`"/BCompare.real $@
```
4、最后修改下脚本的权限：
```shell
chmod a+x BCompare # 给脚本可执行权限
```
