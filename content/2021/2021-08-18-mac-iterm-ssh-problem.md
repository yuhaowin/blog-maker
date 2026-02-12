# 不同的 ssh 客户端登录服务器，影响服务器编码

今天使用的 mac 的 `Terminal.app` 连接服务器，启动一个 java 进程，发现这个 java 服务的中的中文出现了乱码，于是我在启动的时候设置了一下启动参数 `-Dfile.encoding=utf-8`
试了一下，发现中文不乱码了，

但是同事说，他昨天也启动过这个 java 进程，中文并没有乱码，并且这期间也没有修改过服务器的任何配置，（同事使用的是 win 的 xshell ），我感觉非常奇怪，同时基本可以排除服务器的关系，似乎是和使用的终端有关。

于是我移除了 `-Dfile.encoding=utf-8` 参数，使用了 mac 下的 `Termius.app` 这个工具登录了服务，启动这个 java 进程观察，发生也是没有乱码的。

到此，基本可以确定，不同的终端的确是影响了服务器的一些环境变量了。

于是我写了一个小测试类。用两个终端登录服务器后，执行:

```java
public class Test {
    public static void main(String[] args) {
        System.out.println(System.getProperty("file.encoding"));
        System.out.println("测试中文是否乱码");
    }
} 
```

测试结果一(Terminal.app)：

------
```text
ANSI_X3.4-1968

????????
```

测试结果二(Termius.app)：

------
```text
UTF-8

测试中文是否乱码
```

后来查资料发现，ssh 登录的时候，默认会把本地的 locale 发送到服务端，而我本地没有配置这个变量使用的默认值是 UTF-8，服务端默认的 locale 是 en_US.UTF-8，并不认识从本地发送 UTF-8，于是就出现了中文乱码。

+ 设置本地的 locale (我用的是 zsh): vim ~/.zshrc 添加 export LC_ALL=en_US.UTF-8 和 export LANG=en_US.UTF-8
+ 发送本地的 locale 到远程的主机的配置: vim /etc/ssh/ssh_config 中的 SendEnv LANG LC_*

> 补充 Locale 知识： http://wiki.ubuntu.org.cn/Locale
>
> 可以使用 locale 命令查看当前的 locale

