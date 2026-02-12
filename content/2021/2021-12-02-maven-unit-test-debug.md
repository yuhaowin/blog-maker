# 调试 maven 单元测试

> 最近在修改一个 bug，修改完成后，运行 `mvn clean test` 的时候，有一个 case 没有通过，但是在 `idea` 中单独对这个 case
> 测试发现是可以通过的。于是我就想到在执行 maven 命令的时候，连接上调试器，看一看这个case。

### 有两种方式可以对一个 module 的单元测试做调试

#### 一、通过 idea 的设置

![130217](https://image.yuhaowin.com/2021/12/02/130217.png)

在启动配置中，添加如下 VM 参数：`-DforkCount=0`

![130521](https://image.yuhaowin.com/2021/12/02/130521.png)

#### 二、通过 `mvn` 命令启动，并指定 VM 参数

` mvn clean -pl client -Dmaven.surefire.debug test`

其中 -pl 是指定 moudle

执行后，会打开一个默认为 5005 的远程调试端口

![131039](https://image.yuhaowin.com/2021/12/02/131039.png)

在 idea 中连接上该调试端口并启动：

![131822](https://image.yuhaowin.com/2021/12/02/131822.png)

![132019](https://image.yuhaowin.com/2021/12/02/132019.png)

+ [参考资料 - maven 开启调试](https://maven.apache.org/surefire/maven-surefire-plugin/examples/debugging.html)

+ [参考资料 - maven 指定模块](https://www.cnblogs.com/mrld/p/14214879.html)

