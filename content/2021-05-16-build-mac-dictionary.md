# 编译自己的 Mac 词典

+ 下载词典文件

下载 mdict 类型的词典，以朗文5的 mdict 词典为例，将.mdx与.mdd放在同一个文件夹中。

+ [安装Command Line Tools for Xcode](https://link.jianshu.com/?t=https%3A%2F%2Fblog.csdn.net%2Fchenyufeng1991%2Farticle%2Fdetails%2F47007979)

+ 下载 Additional Tools for Xcode（用的 Dictionary Development Kit） 把其放入自定义的一个目录中

+ 从github上克隆 [pyglossary](https://link.jianshu.com/?t=https%3A%2F%2Fgithub.com%2Filius%2Fpyglossary) 项目到自己的目录中

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
+ 重新启动「词典」软件

![232436](https://image.yuhaowin.com/2023/05/04/232436.png)

________________

+ [Mdict 词典分享](https://www.jianshu.com/p/e279d4a979fa)

+ [Mdict to macOS Dictionary 转换笔记](https://kaihao.io/2018/mdict-to-macos-dictionary/)

+ [Mdic 字典文件转换 mac 原生字典](https://blog.i-ll.cc/archives/582/)

+ [柯林斯双解 for macOS](https://placeless.net/blog/macos-dictionaries)

+ [mac dictionary 开发文档](https://developer.apple.com/library/archive/documentation/UserExperience/Conceptual/DictionaryServicesProgGuide/Introduction/Introduction.html#//apple_ref/doc/uid/TP40006152-CH1-SW1)