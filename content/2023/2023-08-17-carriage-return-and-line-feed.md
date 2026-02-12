# 回车和换行的历史
> "回车" 和 "换行" 均是来源机械英文打字机，是两个独立的过程。

## 回车
> 回车的英文是：Carriage Return，回车的意思是回到一行的开头，但并不会换到下一行，只是回到了一行的开头。

回车中的"车"指的是纸车,带着纸一起左右移动的模块，当开始打第一个字之前，要把纸车拉到最右边，弹簧处于拉伸状态，随着打字弹簧会把纸车往左边拉，每当打完一行后，纸车就完全收回去了，
当打下一行之前需要重新把纸车拉到最右边，使得指针指向行的开头，所以叫回车。

## 换行
> 换行的英文是：Line Feed，换行的意思是换到下一行，但并不会回到行的开头，只是换到了下一行的同一位置。

机械英文打字机左边有个"把手"，往下扳动一下，纸会上移一行，所以叫换行。

## 回车换行在现代计算机中的表示
+ 回车：CR，ASCII 代码是 13，或 0x0D，CR=Carriage Return=/r
+ 换行：LF，ASCII 代码是 10，或 0x0A，LF=Line Feed=/n

## 不同操作系统对行分隔符的不同实现
> 行分隔符，Line Separator，不同操作系统对行分隔符的实现是不同的。

+ Windows 系统里，行分隔符为 **/r/n**，每行结尾是 "回车+换行"，即 "CR+LF"，ASCII 码为 0x0D 0x0A
+ Mac 9 以及以前系统里，行分隔符为 **/r**，即 "CR"，ASCII 码为 0x0D
+ Unix/Linux、Mac X 以及以后系统里，行分隔符为 **/n**，即 "LF"，ASCII 码为 0x0A

## 显示文本文件中的行分隔符

在 macOS 终端中，要正确显示文件中的回车换行字符，可以使用如下命令：

```shell
cat -e filename
```

这个命令会将文件的内容显示在终端上，并以可见的方式显示回车换行字符。回车字符会显示为^M，换行字符会显示为$。

也可以使用 od 命令，将文件的内容以十六进制形式显示，并显示特殊字符的表示。以下是一个示例命令:

```shell
od -c filename
```

这将会以十六进制和字符形式显示文件内容，可以在输出中找到回车换行字符（\r 和 \n）。


## 常用文本编辑器和 IDE 中行分隔符的设置
> 一般文本编辑器都可以选择使用不同的行分隔符。

+ Sublime Text
> View -> Line Endings
 
![110302](https://image.yuhaowin.com/2023/08/17/110302.png)

+ JetBrains IDE
> Settings -> Editor -> Code Style -> Line separator

![110945](https://image.yuhaowin.com/2023/08/17/110945.png)

## reference
+ [回车和换行的故事](https://www.ruanyifeng.com/blog/2006/04/post_213.html)
+ [英文机械打字机使用说明](https://www.youtube.com/watch?v=dz-UifvPF78)
+ [Configure line separators](https://www.jetbrains.com/help/idea/configuring-line-endings-and-line-separators.html)