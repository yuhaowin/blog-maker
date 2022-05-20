# 一个统计用户会话数量的案例

基本情况是我有两个接口：

+ config 接口
+ upload 接口

1. 客户端先请求 config 接口，config 会返回当前服务器时间戳 `timestamp` 和 `sessionid`;
2. 客户端再请求一次或多次 upload 接口，请求的参数有 config 接口返回的 `timestamp` 和 `sessionid`.

流程如下：

```
+---------------+                                                                  +----------------+
|               |                                                                  |                |
|               |            client request config api                             |                |
|               |----------------------------------------------------------------->|                |
|               |                                                                  |                |
|               |                                                                  |                |
|               |            config api response server timestamp & sessionid      |                |
|               |<-----------------------------------------------------------------|                |
|               |                                                                  |                |
|    client     |                                                                  |     server     |
|               |                                                                  |                |
|               |  1:client request upload api-> param: timestamp & sessionid      |                |
|               |----------------------------------------------------------------->|                |
|               |                                                                  |                |
|               |  2:client request upload api-> param: timestamp & sessionid      |                |
|               |----------------------------------------------------------------->|                |
|               |                                                                  |                |
|               |  n:client request upload api-> param: timestamp & sessionid      |                |
|               |----------------------------------------------------------------->|                |
+---------------+                                                                  +----------------+

```


我需要每隔 5min 统计 sessionid 持续时间在 30min 内的数量。

意思是某个 sessonid 如：s1 持续时间在 30min 内算一次，在 60min 内算两次。

我现在有一个方案是: 对每一个 upload 请求的 sessionid 做为 redis 的 key，setNX，TTL 为 30min， 如果设置成功计数器 +1，5min 后，计数器置 0 。

这个方案的问题是，同一时刻有大量的 redis key 过期。



后来，我采用了另外的一种方式：

每一天用一个 redis key 统计当天不重复的 session 个数，数据结构为 hyperloglog

假设某次 upload 上 报的 sessionid = s1

s1 该次上报的时间减去该会话首次上报的时间和 30min 比较，

小于 30min s1 记为 s1-0
，大于 30min 小于 60min s1 记为 s1-1
。大于 60min 小于 90min s1 记为 s1-2，以此类推，

这个就把同一个 sessionid 按照持续时间的不同，标记为新的不同的 new-sessionid，放到 redis 定时统计出不重复的个数

这样的好处的，key 数量可以控制，一天就一个，另外， upload 的数据会存在延迟上报的问题 - 本来是昨天上报的 upload 数据，到今天上报，这样就会出现这个 session 在昨天统计了一次，再今天又被统计一次，为此，我在保留了昨天的 redis key，在统计今天的 session 的不重复个数的时候：使用 pfcount(今天的rediskey，昨天的rediskey) - pf(昨天的rediskey) - 就是把昨天和今天的并集 - 今天的数量。

弊端也有：

+ 就是 hyperloglog 本质是概率算法，有 0.81% 的标准误差



参考资料：

[参考一](https://www.v2ex.com/t/298920)

[参考二](https://www.v2ex.com/t/465067)

[参考三](https://mp.weixin.qq.com/s/AvPoG8ZZM8v9lKLyuSYnHQ)
