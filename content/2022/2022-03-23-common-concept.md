## 计算机相关常见概念

### TP:分位值(top percentile)

> The TP9999 is a minimum time under which 99.99% of requests have been served.
>
> 可以认为 TP9999 的意思是保证 99.99% 请求都能被响应需要的最小耗时。

#### 计算方法

calculating TP is very simple:
1.sort all respond times in ascending order，eg: [2s, 10s, 100s, 1000s]
2.find latest item in portion you need to calculate.
	2.1 for TP50 it will be ceil `4*0.5=2` requests. You need 2nd request.
	2.2 for TP90 it will be ceil `4*0.9=4` You need 4th request.
3.we get time for the item found above. TP50=10s. TP90=1000s

+++++++++

### SLA:服务等级协议(service level agreement)

>服务提供商与用户间定义的一种双方认可的协定，用来保障服务的性能和可用性。

#### 计算方法

1年 = 365天 = 8760小时,可以计算全年可停机的时间。

SLA99.9 = 8760 * 0.1% = 8760 * 0.001 = 8.76小时

SLA99.99 = 8760 * 0.0001 = 0.876小时 = 0.876 * 60 = 52.6分钟

SLA99.999 = 8760 * 0.00001 = 0.0876小时 = 0.0876 * 60 = 5.26分钟

### Amdahl‘s law 阿姆达尔定律

> 是并行计算领略一个非常著名的定律,描述对于一个固定计算任务，在并行处理下的加速比。
>
> S=1/((1-f)+f/p)

假设一个大的任务串行处理耗时 1 个单位时间，其中有 f 部分可以进行并行化处理，现在的耗时为，(1-f)+f/p 

![115234](https://image.yuhaowin.com/2022/04/27/115234.png)

reference: [Youtube-Amdahl's Law](https://www.youtube.com/watch?v=Axx2xuB-Xuo)

