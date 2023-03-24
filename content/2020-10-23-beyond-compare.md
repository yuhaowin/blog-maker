# Beyond Compare for Mac 试用

### 原理

> 安装包装的 BCompare 文件是 Beyond Compare 的启动程序，registry.dat 是记录用户注册信息，因此只要在启动的时候删除 registry.dat 注册信息就可以继续试用，为此可以在该目录下添加一个批处理文件用来处理这个操作。

### 操作方法

1、打开命令行终端，进入到安装目录里面的 Contents/Macos，命令行指令：
```shell
cd /Applications/Beyond\ Compare.app/Contents/MacOS/
```
2、修改可执行文件名，并创建脚本
```shell
mv BCompare BCompare.real # 把启动程序改名备用
touch BCompare            # 创建新的启动脚本
vim BCompare              # 编辑脚本内容，内容如下所示
```
3、编辑 BCompare 文件，内容如下：
```shell
#!/bin/bash
rm "/Users/$(whoami)/Library/Application Support/Beyond Compare/registry.dat"
"`dirname "$0"`"/BCompare.real $@
```
4、最后修改下脚本的权限：
```shell
chmod a+x BCompare        # 给脚本可执行权限
```

+ [官网下载地址](http://www.scootersoftware.com/download.php)
+ https://bytexd.com/what-is-dirname-0-and-usage-examples/