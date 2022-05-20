---
title: 'JUC | READ_WRITE_LOCK'
date: 2020-04-26 22:31:56
tags: []
published: true
hideInList: false
feature: 
isTop: false
---

<!-- more -->



### READ_WRITE_LOCK 读写锁

> 读写锁的出现是为了提高锁的效率，实现在 `读与读` 之间锁是共享的，`写与写` 、`读与写` 之间锁是互斥的。这样在 **读多写少** 的场景中十分适合使用读写锁。



### 源码

```java
public interface ReadWriteLock {
    Lock readLock();
    Lock writeLock();
}
```



在 JDK 中 `ReentrantReadWriteLock` 是 `ReadWriteLock` 的实现，同时具有可重入性，



### 读写锁的升级 & 降级

> 对于对写锁而言，锁升级是指：从读锁变成写锁，锁降级是指：从写锁变成读锁。`ReentrantReadWriteLock` 锁可以降级，但是不可以升级。



**读锁无法升级为写锁**

```java
public class T10_ReadWriteLock2 {
    public static void main(String[] args) {
        ReentrantReadWriteLock lock = new ReentrantReadWriteLock();
        lock.readLock().lock();
        System.out.println("read");
        lock.writeLock().lock();
        System.out.println("write");
    }
}
```



**写锁可以降级为读锁**

```java
public class T10_ReadWriteLock3 {
    public static void main(String[] args) {
        ReentrantReadWriteLock lock = new ReentrantReadWriteLock();
        lock.writeLock().lock();
        System.out.println("write");
        lock.readLock().lock();
        System.out.println("read");
    }
}
```

**注意：虽然写锁降级成读锁，但并不会自动释放当前线程获取的写锁，因此仍然需要显示的释放，否则别的线程永远也获取不到写锁。**

