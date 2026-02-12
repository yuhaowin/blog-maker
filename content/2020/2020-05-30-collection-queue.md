# 容器 - 集合｜Queue

java 中常用的集合 - Queue。


### Queue - 接口

在 Queue 中定义了 add/offer、remove/poll、element/peek 六个方法。这六个方法对应三组，每组对元素操作的功能一样，但是也有区别。

**add/offer 区别**

> 都是往队列中添加一个元素，对容量有限的队列，当容量达到上限时，再通过 add 添加元素会抛出异常，而通过 offer 添加元素时会返回 false，不会抛异常。

**remove/poll 区别**

> 都是移除并返回队列的头部元素，当队列为空时，remove 方法会抛异常。poll 只会返回 null。

**element/pick区别**

> 都是返回队列中的头部元素，但是不删除这个元素，当队列为空时，element 方法会抛异常，pick 只会返回 nul。 



### BlockingQueue - 接口

`BlockingQueue` 通常用于具有一个线程生产对象，另一个线程消费对象。原理如图：

![142907](https://image.yuhaowin.com/2020/05/31/142907.jpg)

生产线程将生产新对象并将其插入`BlockingQueue`，直到队列达到其所包含内容的上限。换句话说，如果阻塞队列达到其上限，则在尝试插入新对象时会阻塞生产线程。它一直保持阻塞状态，直到消费线程将对象从队列中移出为止。

消费线程不断  `BlockingQueue`  从中取出对象进行处理。如果消费线程试图将对象从空队列中取出，则消费线程将被阻塞，直到生产线程将对象放入队列中为止。



以下是 BlockingQueue 方法的 4 种处理方式：

|             | **Throws Exception** | **Special Value** | **Blocks** | **Times Out**                 |
| ----------- | -------------------- | ----------------- | ---------- | ----------------------------- |
| **Insert**  | add(o)               | offer(o)        | put(o)   | offer(o, timeout, timeunit) |
| **Remove**  | remove(o)            | poll()          | take()   | poll(timeout, timeunit)     |
| **Examine** | element()            | peek()            |         |                               |



+ [参考资料](http://tutorials.jenkov.com/java-util-concurrent/blockingqueue.html)