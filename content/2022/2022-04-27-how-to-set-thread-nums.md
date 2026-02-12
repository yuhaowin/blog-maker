# 到底如何设置线程池的核心线程数、最大线程数



#### 线程池在业务中的实践

##### 场景一：快速响应用户请求

这种场景可以将用户请求封装成任务并发执行，缩短总体响应时间，该场景需要获取最大的响应速度满足客户，应该不应该设置缓冲队列，缓冲并发任务。可以适当调高 corePoolSize 和 maxPoolSize 去尽可能创造多的线程快速执行任务。

##### 场景2：快速处理批量任务

这种场景一般是需要大量执行离线任务，是吞吐量优先，但是不是要求瞬时完成，也就是要求尽可能在单位时间内处理更多的任务，可以使用缓冲队列，缓冲任务，corePoolSize 不适合特别大，太大频繁上下文切换，反而影响吞吐量。

#### 业界的线程数配置方案一般都是比较理想化

> 并发任务的执行情况和任务类型相关，IO密集型和CPU密集型的任务运行起来的情况差异非常大，较难合理预估，这导致很难有一个简单有效的通用公式帮我们直接计算出结果。

**CPU 密集型任务**

> 比如像加解密，压缩、计算等一系列需要大量耗费 CPU 资源的任务，大部分场景下都是纯 CPU 计算。

核心线程数，可以设置为CPU核数 + 1 , +1是为了实现最优的利用率。即使当密集型的线程由于偶尔的内存页失效或其他原因导致阻塞时，这个额外的线程也能确保 CPU 的时钟周期不会被浪费，从而保证 CPU 的利用率。

**IO 密集型任务**

> 比如像 MySQL 数据库、文件的读写、网络通信等任务，这类任务不会特别消耗 CPU 资源，但是 IO 操作比较耗时，会占用比较多时间。

核心数设置的一般比较多一些，因为 IO 读写速度相比于 CPU 的速度而言是比较慢的，核心线程数=CPU 核心数 * (1 + IO 耗时/ CPU 耗时)

尽管通过严谨的评估，依然很难一次计算出合适的参数，因此，可以换一个思路，把修改参数的成本降低，这样可以在告警发生时，快速调整，尽快恢复。

#### 动态调整线程池参数

简化参数配置，关注核心参数，corePoolSize maxPoolSize workQueue，

+ 延时优先的场景，同步队列。
+ 吞吐量优先的场景，使用有界队列。

修改 corePoolSize
ThreadPoolExecutor#setCorePoolSize：在运行期间可以通过该方法修改 corePoolSize，会直接覆盖之前的 corePoolSize 值。

+ 当前值小于之前值，表示有多余的 work 线程，此时会向当前 idle 的 worker 线程发起中断请求以实现回收，其余多余的 worker 在下次 idel 的时候也会被回收。
+ 当前值大于之前值，并且队列里有任务的时候，会创建新的线程执行任务。

修改 maxPoolSize
ThreadPoolExecutor#setMaximumPoolSize

+ 当前值小于之前值，超过的，并且已经在运行的线程会在 idle 的时候停止。

修改 workQueueSize

+ LinkedBlockingQueue 没有开放修改 capacity 的方法，可以参考 LinkedBlockingQueue 自定义支持修改 capacity 的 Queue。

```java
public class DynamicThreadPool {

    ThreadPoolExecutor threadPoolExecutor;

    public static void main(String[] args) throws InterruptedException {
        DynamicThreadPool threadPool = new DynamicThreadPool();
        threadPool.buildThreadPool();
        threadPool.print(threadPool.threadPoolExecutor, "init");
        for (int i = 0; i < 15; i++) {
            int finalI = i;
            threadPool.threadPoolExecutor.submit(() -> {
                threadPool.print(threadPool.threadPoolExecutor, "创建任务: " + finalI);
                try {
                    TimeUnit.SECONDS.sleep(10);
                } catch (InterruptedException e) {
                    throw new RuntimeException(e);
                }
            });
        }
        threadPool.modifyMaxSize(10);
        threadPool.modifyCoreSize(10);
        TimeUnit.SECONDS.sleep(2);
        threadPool.modifyWorkQueueSize(100);
    }

    public ThreadPoolExecutor buildThreadPool() {
        ThreadPoolExecutor executor = new ThreadPoolExecutor(
                2,
                5,
                30,
                TimeUnit.MILLISECONDS,
                new ResizeLinkedBlockingQueue<>(10));
        this.threadPoolExecutor = executor;
        return executor;
    }

    public void modifyCoreSize(int num) {
        threadPoolExecutor.setCorePoolSize(num);
    }

    public void modifyMaxSize(int num) {
        threadPoolExecutor.setMaximumPoolSize(num);
    }

    public void modifyWorkQueueSize(int size) {
        ResizeLinkedBlockingQueue<Runnable> queue = (ResizeLinkedBlockingQueue<Runnable>) threadPoolExecutor.getQueue();
        queue.setCapacity(size);
    }

    public void print(ThreadPoolExecutor executor, String name) {
        ResizeLinkedBlockingQueue queue = (ResizeLinkedBlockingQueue) executor.getQueue();
        String message = String.format("%s 核心线程数: %s,最大线程数: %s,活动线程数: %s,完成任务数: %s,队列大小: %s,队列剩余: %s",
                name,
                executor.getCorePoolSize(),
                executor.getMaximumPoolSize(),
                executor.getActiveCount(),
                executor.getCompletedTaskCount(),
                queue.size(),
                queue.remainingCapacity());
        System.out.println(message);

    }
}
```



问题一：线程池被创建后里面有线程吗？如果没有的话，你知道有什么方法对线程池进行预热吗？

默认情况线程池被创建后如果没有任务过来，里面是不会有线程的。如果需要预热的话可以调用下面的两个方法：

+ prestartCoreThread 启动一个核心线程
+ prestartAllCoreThreads 启动所有的核心线程

问题二：核心线程数会被回收吗？需要什么设置？

核心线程数默认是不会被回收的，如果需要回收核心线程数，需要调用下面的方法：

+ allowCoreThreadTimeOut(boolean value)

reference：

[美团](https://tech.meituan.com/2020/04/02/java-pooling-pratice-in-meituan.html)

[公众号](https://mp.weixin.qq.com/s/YbyC3qQfUm4B_QQ03GFiNw)

