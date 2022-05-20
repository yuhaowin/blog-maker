---
title: 'JUC｜COUNTDOWNLATCH'
date: 2020-04-21 23:28:08
tags: [多线程与高并发-进程内]
published: true
hideInList: false
feature: https://image.yuhaowin.com/2020/04/22/011401.jpg
isTop: false
---
> CountDown 倒数的意思，Latch 门栓，倒数完成，门栓打开，CountDownLatch 使一个线程等待其他线程各自执行完毕后再执行。
<!-- more -->
### CountDownLatch

**CountDownLatch 是通过一个计数器来实现的，计数器的初始值一般是线程的数量。每当一个线程执行完毕后，计数器的值就减 1，当计数器的值为 0 时，表示所有线程都执行完毕，然后执行过 await 方法的线程就可以恢复工作了。计数器的初始值也可以是任务数量，任务被线程池中的线程执行，每完成一个任务计数器就减 1。当所有任务执行完成，再继续执行其他事情。**

### CountDownLacth 源码

+ 只提供了一个构造器

```java
//参数count为计数值
public CountDownLatch(int count) {} 
```

+ 主要方法

```java
//调用 await() 方法的线程会被挂起，它会等待直到count值为0才继续执行
public void await() throws InterruptedException { };   
//和await()类似，只不过等待一定的时间后count值还没变为0的话就会继续执行
public boolean await(long timeout, TimeUnit unit) throws InterruptedException { };  
//将count值减1
public void countDown() { };  
```

+ 使用案例

```java
public class T06_CountDownLatch3 {
    public static void main(String[] args) throws InterruptedException {
        CountDownLatch startSignal = new CountDownLatch(1);
        CountDownLatch doneSignal = new CountDownLatch(10);
        for (int i = 0; i < 10; ++i)  // create and start threads
            new Thread(new Worker(startSignal, doneSignal)).start();
        doSomethingElse();
        startSignal.countDown();      // let all threads proceed
        doSomethingElse();
        doneSignal.await();           // wait for all to finish
        System.out.println("finish");
    }
    private static void doSomethingElse(){
        System.out.println("doSomethingElse method");
    }
}

class Worker implements Runnable {
    private final CountDownLatch startSignal;
    private final CountDownLatch doneSignal;
    Worker(CountDownLatch startSignal, CountDownLatch doneSignal) {
        this.startSignal = startSignal;
        this.doneSignal = doneSignal;
    }

    public void run() {
        try {
            startSignal.await();
            doWork();
            doneSignal.countDown();
        } catch (InterruptedException ex) {
        }
    }

    void doWork() {
        System.out.println(Thread.currentThread().getName());
    }
}
```



> + [参考资料1](https://www.jianshu.com/p/e233bb37d2e6)
> + [CountDownLatch与thread.join()的区别](https://www.jianshu.com/p/795151ac271b)

