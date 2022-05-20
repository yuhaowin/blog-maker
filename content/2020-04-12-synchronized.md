---
title: 'SYNCHRONIZED IN JAVA'
date: 2020-04-12 16:48:35
tags: [多线程与高并发-进程内]
published: true
hideInList: false
feature: 
isTop: false
---
JAVA 中 synchronized 的常用方法，和锁升级的过程。
<!-- more -->
### synchronized 锁定的是什么？

> 我们常说使用 synchronized 关键字锁定一个方法，使得这个方法变成同步方法；使用 synchronized 关键字锁定一段代码块，使得这段代码变成同步代码块，其用法如下：

```java
    public synchronized void add() {
        // do something
    }
    public void count() {
        synchronized (this) {
            // do something
        }
    }
```

实际上是 synchronized 使得一个对象成为了一把锁，因此 synchronized 锁是叫对象锁，当某一个线程获得了这个对象锁后，才能够执行被其包含的代码。至于是哪个对象充当这把锁，都是可以的。如果 synchronized 修饰的是普通的方法，如上面的 add 放法，那么充当这个锁的对象就是 this，如果 synchronized 修饰的是一个代码片段，可以指定任意的一个对象，当然指定 this 这个对象也是可以的。如果 synchronized 修饰的是 static 方法，这个对象锁就是当前类的 class 对象

```java
public class T {
    private static int count = 10;
    public synchronized static void m() { //这里等同于synchronized(T.class)
        count--;
        System.out.println(Thread.currentThread().getName() + " count = " + count);
    }
    public static void mm() {
        synchronized (T.class) { //考虑一下这里写synchronized(this)是否可以？
            count--;
        }
    }
}
```



### synchronized 保证了线程的互斥性和线程的可见性

#### 互斥性

> 即在同一时间只允许一个线程持有某个对象锁，通过这种特性来实现多线程中的协调机制，这样在同一时间只有一个线程对需同步的代码块(复合操作)进行访问。互斥性我们也往往称为操作的原子性。

#### 可见性

> 必须确保在锁被释放之前，对共享变量所做的修改，对于随后获得该锁的另一个线程是可见的（即在获得锁时应获得最新共享变量的值），否则另一个线程可能是在本地缓存的某个副本上继续操作从而引起不一致。

### synchroniezd 的可重入性

> 在一个同步方法 A 中又调用另一个同步方法 B，（当然需要保证这个两个同步方法的锁是同一个）如果当一个线程 T 获得了方法 A 的对象锁，那么这个线程就一定能成功的获取到方法 B 的对象锁。

可重入锁：**一个线程在获取某个锁后且在释放这个锁之前，依然可以再次获取这个锁。**

```java
public class Parent {
    public synchronized void m() {
        System.out.println("Parent m method is called");
        System.out.println("当前的对象锁是："+this);
    }
}
```

```java
public class Son extends Parent {
    @Override
    public synchronized void m() {
        super.m();
        System.out.println("Son method is called");
        System.out.println("当前的对象锁是："+this);
    }
    public static void main(String[] args) {
        Parent son = new Son();
        son.m();
    }
}
```

```java
public class WhatReentrant {
    public static void main(String[] args) {
        new Thread(new Runnable() {
            @Override
            public void run() {
                synchronized (this) {
                    System.out.println("第1次获取锁，这个锁是：" + this);
                    int index = 1;
                    while (true) {
                        synchronized (this) {
                            System.out.println("第" + (++index) + "次获取锁，这个锁是：" + this);
                        }
                        if (index == 5) {
                            break;
                        }
                    }
                }
            }
        }).start();
    }
}
```

### 子类继承父类的 synchronized 方法时不保留 synchronized
> 如果父类某个方法被 synchronized 修饰，其子类在重写父类方法时，如果也需要时同步的，子类需要
> 显示的指定为 synchronized。否则重写的子类方法不是同步方法。



### synchronized 的 4 种锁状态

> 无锁、偏向锁、轻量级锁、重量级锁

先了解两个概念，**对象头、Monitor**

#### 对象头

对象头的方式有以下两种(以32位JVM为例)，主要包含两部分 Mark Word 和 Klass Word

+ Mark Word ：默认存储对象的HashCode，分代年龄和锁标志位信息。这些信息都是与对象自身定义无关的数据，所以Mark Word被设计成一个非固定的数据结构以便在极小的空间内存存储尽量多的数据。它会根据对象的状态复用自己的存储空间，也就是说在运行期间Mark Word里存储的数据会随着锁标志位的变化而变化。
+ Klass Word：对象指向它的类元数据的指针，虚拟机通过这个指针来确定这个对象是哪个类的实例。

```shell
// 普通对象
|------------------------------------------------------------|
|                   Object Header (64 bits)                  |
|----------------------------------|-------------------------|
|        Mark Word (32 bits)       |    Klass Word (32 bits) |
|----------------------------------|-------------------------|

// 数组对象
|---------------------------------------------------------------------------------|
|                                 Object Header (96 bits)                         |
|--------------------------------|-----------------------|------------------------|
|        Mark Word(32bits)       |    Klass Word(32bits) |  array length(32bits)  |
|--------------------------------|-----------------------|------------------------|
```



```shell
|-------------------------------------------------------|--------------------|
|                  Mark Word (32 bits)                  |       State        |
|-------------------------------------------------------|--------------------|
| identity_hashcode:25 | age:4 | biased_lock:0| lock:01 |       Normal       |
|-------------------------------------------------------|--------------------|
|  thread:23 | epoch:2 | age:4 | biased_lock:1| lock:01 |       Biase        |
|-------------------------------------------------------|--------------------|
|             ptr_to_lock_record:30           | lock:00 | Lightweight Locked |
|-------------------------------------------------------|--------------------|
|             ptr_to_heavyweight_monitor:30   | lock:10 | Heavyweight Locked |
|-------------------------------------------------------|--------------------|
|                                             | lock:11 |    Marked for GC   |
|-------------------------------------------------------|--------------------|
```



**偏向锁**：是指一段同步代码一直被一个线程所访问，那么该线程会自动获取锁，降低获取锁的代价。在大多数情况下，锁总是由同一线程多次获得，不存在多线程竞争，所以出现了偏向锁。其目标就是在只有一个线程执行同步代码块时能够提高性能。当一个线程访问同步代码块并获取锁时，会在Mark Word里存储锁偏向的线程ID。引入偏向锁是为了在没有多线程竞争的情况下尽量减少不必要的轻量级锁执行路径，因为轻量级锁的获取及释放依赖多次CAS原子指令，而偏向锁只需要在置换ThreadID的时候依赖一次CAS原子指令即可。

**轻量级锁**：是指当锁是偏向锁的时候，被另外的线程所访问，偏向锁就会撤销升级为轻量级锁，其他线程会通过自旋的形式尝试获取锁，不会阻塞，从而提高性能。

**重量级锁**：升级为重量级锁时，锁标志的状态值变为“10”，此时Mark Word中存储的是指向重量级锁的指针，此时等待锁的线程都会进入阻塞状态。为了减少大量的自旋造成的 cpu 的消耗。



### synchronized 锁的底层实现原理

> synchronized 代码块是由一对儿 monitorenter/monitorexit 指令实现的。实际上在 Java 中，每个对象都会有一个 monitor 对象，这个对象其实就是 Java 对象的锁，通常会被称为“内置锁”或“对象锁”。类的对象可以有多个，所以每个对象有其独立的对象锁，互不干扰。

### synchronized 锁升级过程

![133630](https://image.yuhaowin.com/2022/01/25/133630.jpg)



默认情况下对象一创建出来就是可偏向的，当第一个线程 Thread1 通过 synchronized 使用的这个对象锁后，该对象的 Mark Word 被写入 Thread1 的线程 id，以后这个对象锁就偏向与 Thread1，此时如果解锁后 Mark Word 依然是 Thread1 的线程 id，接下来如果 Thread1 再次使用时，就可以直接获得锁。

当出现第二个线程 Thread2 的时候，偏向锁被撤销，升级为轻量级锁，解锁后进入 normal 状态。此时如果依然有别的线程竞争就会进入自旋等待。

当自旋达到一定次数，仍然没有获取到锁，就升级重量级锁。以减少对 cpu 的消耗。



扩展知识

> 执行时间短（同步的代码块），线程数量少，适合使用自旋锁。
>
> 执行时间长，线程数多，适合使用系统锁。

偏向锁可以通过参数禁用 `-XX:-UseBiasedLocking`，如果禁用后，对象创建出来就是 normal 的，第一次使用的就是轻量级锁，解锁后又恢复为 normal


锁升级的概念：
+ [我就是厕所所长一](https://www.jianshu.com/p/b43b7bf5e052)
+ [我就是厕所所长二](https://www.jianshu.com/p/16c8b3707436)