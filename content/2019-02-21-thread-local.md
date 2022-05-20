---
title: 'THREAD_LOCAL 学习'
date: 2020-05-07 23:54:36
tags: []
published: true
hideInList: false
feature: 
isTop: false
---

### ThreadLocal 简单使用

> ThreadLocal 准确的说应该叫 ThreadLocalVar - 线程本地变量，其目的是使得该变量是线程隔离的，对其他的线程是不可见的。
>
> ThreadLocal 类一共就 3 个公开的方法，`set(T value)`、`get()`、`remove()`。

每一个线程都有一个存放线程本地变量的容器，通过 ThreadLocal.set() 将变量的引用保存到各线程的自己的一个容器中，在执行 ThreadLocal.get() 时，各线程从自己的容器中取出放进去的对象，因此取出来的是各线程自己中的对象，ThreadLocal 实例是作为这个容器的 key 来使用的。


### 源码分析

```java
    public void set(T value) {
        // 获取当前线程
        Thread t = Thread.currentThread();
        // 获取当前线程存放线程本地变量的容器
        ThreadLocalMap map = getMap(t);
        // 如果存在就直接 set，没有则创建并 set
        if (map != null)
            map.set(this, value);
        else
            createMap(t, value);
    }

		ThreadLocalMap getMap(Thread t) {
      	// Thread 类中有一个 ThreadLocalMap 类型的成员变量
        // 但是该变量的维护是由 ThreadLocal 负责的
      	return t.threadLocals;
 		}

		void createMap(Thread t, T firstValue) {
        t.threadLocals = new ThreadLocalMap(this, firstValue);
    }
```

```java
    public T get() {
        Thread t = Thread.currentThread();
        ThreadLocalMap map = getMap(t);
        if (map != null) {
            ThreadLocalMap.Entry e = map.getEntry(this);
            if (e != null) {
                @SuppressWarnings("unchecked")
                T result = (T)e.value;
                return result;
            }
        }
      	// 如果没 set 就开始 get，上面的 map 为 null，
      	// 就会先初始化后，再获取，默认初始化的值是 null
        return setInitialValue();
    }
```

### ThreadLocal 存在内存泄露？

> 是这样的，一般来说，存放 ThreadLocalMap 和 Thread 是同生死的，ThreadLocalMap 的 key 为 ThreadLocal 实例，value 为 存放的变量。当 Thread 生命周期结束，被 GC 回收时，ThreadLocalMap 也会被回收，所谓的内存泄露是指，在当前 Thread 还在使用期间，由于 ThreadLocalMap 的 key 是 弱引用，在发生 GC 时，可能被回收（此时 Thread 还没有被回收），既然 key 被回收，那 value 也就没有存在的意义了，但是此时的 value 由于强引用的存在，没有被回收。这就是人们说的 ThreadLocal 出现的内存泄露。但是当 Thread 被回收时，ThreadLocalMap 和 其中的 value 依然会被回收的。

> 所以只有在线程池中，使用了 ThreadLocal 才需要特别注意这个问题，因为核心线程是不会被销毁的，如果这些线程的 ThreadLocalMap 中存在用不到的 value，且由于 Thread 一直存在 value 不能被 回收，可以认为是真正发生了内存泄露。解决办法就是：使用完后主动的 remove 掉。

代码如下：

```java
public class ThreadPoolTest {
    static class LocalVariable {
        private Long[] a = new Long[1024 * 1024];
    }

    final static ThreadPoolExecutor poolExecutor = new ThreadPoolExecutor(
            10, 10,
            1, TimeUnit.MINUTES,
            new LinkedBlockingQueue<>());

    public static void main(String[] args) throws InterruptedException {
        for (int i = 0; i < 50; ++i) {
            poolExecutor.execute(() -> {
                ThreadLocal<LocalVariable> localVariable = new ThreadLocal<>();
                localVariable.set(new LocalVariable());
                System.out.println("use local varaible");
                //localVariable.remove(); 为了避免 ThreadLocalMap 中不在使用的 value 不能被及时的回收，造成内存泄露，可以在使用完后主动的 remove 掉。
            });
            Thread.sleep(1000);
        }
        System.out.println("pool execute over");
    }
}
```

![142454](https://image.yuhaowin.com/2020/05/27/142454.jpg)


### 补充知识 - 强、软、弱、虚引用
>Java中提供这四种引用类型主要有两个目的：第一是可以通过代码的方式决定某些对象的生命周期；第二是有利于JVM进行垃圾回收。

```java
public class M {
    @Override
    protected void finalize() {
        // 重新该方法是为了观察实例有没有被垃圾回收器回收
        System.out.println("finalize");
    }
}
```
#### 强引用
> 对象被引用时，不会被 GC 回收。

```java
public class NormalReference {
    public static void main(String[] args) throws IOException {
        M m = new M();
        m = null;
        System.gc(); //DisableExplicitGC
        System.in.read();
    }
}
```
控制台打印 finalize ，说明到 m 置为 null 后，new M() 这个对象没有被引用，所以 GC 会清理掉。

#### 软引用
>软引用在内存充足时可能不会被 GC 回收，在内存不够时会被回收。

```java
/**
 * 软引用是用来描述一些还有用但并非必须的对象。
 * 对于软引用关联着的对象，在系统将要发生 OOM 异常之前，将会把这些对象列进回收范围进行第二次回收。
 * 如果这次回收还没有足够的内存，才会抛出内存溢出异常。
 * 软引用非常适合缓存使用
 * VM options -Xms20M -Xmx20M
 */
public class SoftReference {
    public static void main(String[] args) {
        SoftReference<byte[]> m = new SoftReference<>(new byte[1024 * 1024 * 10]); //10M
        System.out.println(m.get());
        System.gc();
        try {
            Thread.sleep(500);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        System.out.println(m.get()); //经历  GC 后对象仍然存在，不会被回收。

        //再分配一个数组，heap将装不下，这时候系统会垃圾回收，先回收一次，如果不够，会把软引用干掉
        byte[] b = new byte[1024 * 1024 * 10]; //10M 此时堆内存不足 10M 会回收软引用，如果此时堆内存依然不足，则抛 OOM。
        System.out.println(m.get());
    }
}
```

#### 弱引用
>发生 GC 时就会被回收。

```java
/**
 * 弱引用遭到 gc 就会回收
 * 一般使用在容器中，弱引用的主要用途是，一旦指向对象
 * 的强引用断开，就无需关心该对象不会被回收。
 * <p>
 * WeakHashMap
 */
public class WeakReference {
    public static void main(String[] args) throws IOException {
        WeakReference<M> m = new WeakReference<>(new M());

        System.out.println(m.get());
        System.gc();
        System.out.println(m.get());

        System.in.read();
    }
}
```
![001723](https://image.yuhaowin.com/2020/05/28/001723.png)

```java
public class WeakReference1 {
    public static void main(String[] args) throws IOException {
        ThreadLocal<M> tl = new ThreadLocal<>();
        tl.set(new M());
        tl.remove(); // 用完养成 remove 的好习惯。
        System.gc();
        System.in.read();
    }
}
```
#### 虚引用
> 未知，使用很少。
------
<br/>
<details>
<summary>参考资料</summary>

+ [参考资料1](https://www.jianshu.com/p/3c5d7f09dfbd)

+ [参考资料2](https://www.jianshu.com/p/b74de925cd7a)

+ [参考资料3](https://www.jianshu.com/p/6fc3bba12f38)

+ [参考资料4](https://www.iteye.com/topic/103804)

+ [参考资料5](https://blog.csdn.net/qjyong/article/details/2158097)

+ [参考资料6](http://ifeve.com/%E4%BD%BF%E7%94%A8threadlocal%E4%B8%8D%E5%BD%93%E5%8F%AF%E8%83%BD%E4%BC%9A%E5%AF%BC%E8%87%B4%E5%86%85%E5%AD%98%E6%B3%84%E9%9C%B2/)

+ [参考资料7](https://www.jianshu.com/p/1a5d288bdaee)

+ [参考资料8](https://juejin.im/post/5ba9a6665188255c791b0520)
</details>





