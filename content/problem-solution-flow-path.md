> 线上应急问题的核心目标：**尽快恢复服务，消除影响**。然后**保留部分现场，再去定位问题、解决问题和复盘**。

+ 优先恢复服务
+ 保留现场异常信息（内存dump、线程dump、gc log等）

#### 常见问题一：CPU 利用率飙升

> CPU飙升只是一种现象。

##### 常见原因

+ 频繁 gc
+ 有死循环、线程阻塞、io wait 等

##### 处理步骤

+ `top` 定位 CPU 使用最高的进程，获取 pid。
+ `top -Hp pid` 定位 CPU 使用最高的线程，获取 tid。
+ ``printf '0x%x' tid` 线程 id 转化 16 进制。
+ `jstack pid | grep tid` 获取线程栈信息。

可能是 gc 线程，或者是业务线程，然后对应处理。

#### 常见问题二：线程池异常

> 当任务队列满了并且已开辟了最大线程数，此时又来了新任务，ThreadPoolExecutor 会拒绝服务。

常见原因

+ 上游服务响应时间（RT）过长
+ 数据库慢 sql 或者数据库死锁
+ `jstack –l pid | grep -i –E 'BLOCKED | deadlock'` 查看线程栈信息，是否有死锁

##### 常见问题三：频繁 Full gc 

> 如果 gc 线程占用 cpu 特别高，导致服务不可用，可能是在频繁的 full gc。

处理步骤

+ `jstat -gcutil pid 1000` 通过 jstat查看 gc 状态,如果内存占用较高，并且 full gc 频繁，可以看看 dump 文件，进行分析。
+ `jmap -dump:format=b,file=dump.bin pid` 使用 mat 工具分析是否有内存泄露问题。





线上出问题，首先定位是否是由于新版本引起的，能否直接回滚（跟其他业务是否存在依赖）。 如果是内部微服务，看你们使用的框架是否提供业务间调用的监控，帮助定位到具体出问题的功能。如果是 http 服务的话就看是哪些 url 出错。 定位到具体出错的地方，接下来就得看具体问题分析了。基本手段就是查日志，看出出错的位置，再根据代码逻辑推理



可以从函数设计入手 一个函数出现问题只有三种情况：输入错误，输出错误，实现错误 输入错误：输入的值不符合标准导致出错 输出错误：输入的值是正确的，但输出的值并不符合预期 实现错误：输入的值是正确的，但函数执行过程出现了错误导致直接 panic 了，导致整体直接出错 前两个是整条执行链里中其中一环，前面出错会让后面一直错误下去,直到遇到后者直接 panic 才发现错误



从工程角度来说，出现 Bug ，无非就是 2 个原因： A.设计问题：一开始的设计就是错的。 B.实现问题：在实现时出错，或者图省事没去做压力测试导致性能差，等等。 围绕这两个部分，从先易后难的思路去分析： 1.先看 2 的性能问题。从运维那里，拿到应用的 CPU 、内存、网络、存储 IO 的图表，看看有没有明显不合理的指标，比如 CPU 长时间高负载，比如存储 IO 的某个挂载点长时间的使用率（%util / 活动时间）很高，等等。 2.再看 2 的实现是否出错。此时，把日志的级别，从 INFO 或 WARNING 级别，切换到最详细的 DEBUG 级别，来分析模块与函数的调用顺序、执行时间、入参、返回结果等等，来观察与预期是否一致。 3.如果上述都没问题，就得开始检查设计问题了。这一步比较麻烦，因为需要找到比当前系统的设计者，水平更高的设计者，才能检查出问题。



https://cloud.tencent.com/developer/article/1600345