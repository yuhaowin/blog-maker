---
title: 'Java Normal Thread & Daemon Thread'
date: 2021-08-26 15:49:08
tags: []
published: true
hideInList: false
feature: https://image.yuhaowin.com/2021/12/08/000031.png
isTop: false
---

> 有一天我日常启动我的项目后，我陷入了深深的沉思，我搞不清楚，为什么我的项目可以运行而没有终止。
<!-- more -->

程序的入口方法是一个 main 方法，我突然想到，这就是一个 main 方法，为什么当这个 main 方法执行完成后，程序没有退出？这和我无数次写的测试 demo 不一样。于是我决定搞清楚原委。

![224903](https://image.yuhaowin.com/2021/12/07/224903.png)

main 方法执行结束，进程退出。

![225716](https://image.yuhaowin.com/2021/12/07/225716.png)

但是我的程序，在 main 是有启动线程的（启动的是非守护线程）,只要有一个以上的非守护线程在运行，jvm 就不会退出

![230844](https://image.yuhaowin.com/2021/12/07/230844.png)

如果上面启动的线程是守护线程的话，则当 main 线程退出后，进程也会随之退出；

![231330](https://image.yuhaowin.com/2021/12/07/231330.png)

Thread#setDaemon 注释表明：

> The Java Virtual Machine exits when the only threads running are all daemon threads.

当此刻运行的所有线程都是 daemon thread 的时候，jvm 会退出。daemon thread 可以在 jvm 退出的时候，自己结束自己。使用场景是：希望在 jvm 退出时，线程可以自动关闭。如：垃圾回收线程。

________


+ [参考资料一](https://www.cnblogs.com/quanxiaoha/p/10731361.html)

+ [参考资料二](https://www.twle.cn/c/yufei/javatm/javatm-basic-daemon-thread.html)
