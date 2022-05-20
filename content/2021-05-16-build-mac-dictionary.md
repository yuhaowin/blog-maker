---
title: '编译自己的 Mac 词典'
date: 2021-04-16 10:28:28
tags: []
published: true
hideInList: false
feature: https://image.yuhaowin.com/2021/12/08/001502.png
isTop: false
---

+ 下载词典文件

下载 mdict 类型的词典，以朗文5的 mdict 词典为例，将.mdx与.mdd放在同一个文件夹中。

+ [安装Command Line Tools for Xcode](https://link.jianshu.com/?t=https%3A%2F%2Fblog.csdn.net%2Fchenyufeng1991%2Farticle%2Fdetails%2F47007979)

+ 下载  Additional Tools for Xcode（用的 Dictionary Development Kit）

把其放入自定义的一个目录中

![105055](https://image.yuhaowin.com/2021/04/16/105055.png)

+ 从github上克隆 [pyglossary](https://link.jianshu.com/?t=https%3A%2F%2Fgithub.com%2Filius%2Fpyglossary) 项目到自己的目录中

![105432](https://image.yuhaowin.com/2021/04/16/105432.png)

+ 安装python-lzo

  ```
  sudo pip3 install python-lzo
  ```

+ 安装BeautifulSoup4

  ```
  sudo pip3 install lxml beautifulsoup4 html5lib
  ```

+ 转化词典

  ```
  cd ~/Desktop/LDOCE5
  python3 ~/Documents/pyglossary/main.py --write-format=AppleDict LDOCE5.mdx LDOCE5-apple
  ```

+ 运行结束，修改 makefile

  ```
  cd LDOCE5-apple
  vim makefile
  ```

  修改DICT_BUILD_TOOL_DIR为自己的Dictionary Development Kit文件夹路径。

+ 编译

  ```
  make && make install
  ```

+ 把编译好的词典文件放入系统词典目录中

![110811](https://image.yuhaowin.com/2021/04/16/110811.png)


![110933](https://image.yuhaowin.com/2021/04/16/110933.png)

+ 重新启动「词典」软件

![111042](https://image.yuhaowin.com/2021/04/16/111042.png)

________________

+ [参考资料一](https://www.jianshu.com/p/e279d4a979fa)

+ [参考资料二](https://www.jianshu.com/p/0be3e3de8f84)

+ [参考资料三](https://www.jianshu.com/p/c57be986589b)

+ [参考资料四](https://kaihao.io/2018/mdict-to-macos-dictionary/)
