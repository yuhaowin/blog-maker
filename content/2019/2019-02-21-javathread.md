# JAVA THREAD

### 单线程 🆚 多线程

> 常见的方法调用，就是单线程的，以 JAVA 的 main 方法为例，在 main 方法中调用 a 方法，在 a 方法中调用 b 方法。当运行 main 方式时，会启动一个 main 线程，在整个方法的调用链中都一个 main 线程在执行。

```java
    public static void main(String[] args) {
        a();
        System.out.println("main");
    }
    static void a() {
        b();
        System.out.println("a");
    }
    static void b() {
        System.out.println("b");
    }
```

> 同一个时间点，有多个线程运行，有不同的程序执行路径，就是多线程程序。以下代码就是「同时」有 main 线程和 thread one 线程在运行。

```java
    public static void main(String[] args){
        new Thread(() -> {
            System.out.println(Thread.currentThread().getName());
        },"thread one").start();
        System.out.println("main");
    }
```

### 显示创建线程的两种方法

1、实现 Runnable 接口

```java
    static class MyRun implements Runnable {
        @Override
        public void run() {
            System.out.println("Hello MyRun!");
        }
    }
    public static void main(String[] args) {
       new Thread(new MyRun()).start();
    }
```

2、继承 Thread 类，重写 Thread 的 run 方法

```java
    static class MyThread extends Thread {
        @Override
        public void run() {
            System.out.println("Hello MyThread!");
        }
    }
    public static void main(String[] args) {
        new MyThread().start();
    }
```

> 有这两种方法的原因是：被新线程执行的业务代码是在 Thread 类的 run 方法中。即 run 方法是在线程启动后会被虚拟机调用，因此可以通过继承并重写 run 方法的方式自定义一个线程；在 Thread 类中持有一个 Runnable 的成员变量，且 Thread 类中 run 方法的实现是直接调用成员变量 Runnable 的 run 方法，因此可以通过实现 Runnable 接口，并将其通过 Thread 的构造器传入，实现定义一个新线程。

### Thread 类的 run 方法和 start 方法的区别

> 如果调用一个线程的 run 方法，那就是一个普通的方法调用过程，并不会创建一个新的线程，run 方法中的代码是被调用它的线程执行的；而 start 方法会通过 native 的 start0 方法创建一个新的线程。**run 方法不需要也不应该显示调用，该方法是由虚拟机回调的。** 

两者都是 Thread 类的方法，start 是用来启动一个线程，run 表示线程启动后要运行的代码。



interrupt() 可以打断一个线程，被打断的线程回获得一个打断标记。通过 isInterrupted() 获得。       

t.join(): 等 t 线程运行结束，才继续执行调用 t 线程的线程。



#### two phase termination



### 线程状态转换图
![151226](https://image.yuhaowin.com/2020/04/08/151226.jpg)


[参考资料](https://blog.csdn.net/pange1991/article/details/53860651)
### Thread.join 方法
>  join 方法内部的实现是 synchronized + object.wait 实现，主线程会拿到调用join方法的线程的锁，导致主线程进入 waiting 状态，然后调用 join 方法的线程继续执行，执行完成只会调用 notifyAll 方法，主线程从 waiting 状态中被调度出来，继续执行。
```java
public static void main(String[] args) throws InterruptedException {
        Thread thread = new Thread(() -> {
            for (int i = 0; i < 200; i++) {
                System.out.println(i);
            }
        }, "Thread one");
        thread.start();
        thread.join();
        for (int i = 0; i < 50; i++) {
            System.out.println("main : " + i);
        }
}
```
注意：**join 并不是合并线程，调用 join 方法前有 n 个线程，那么在调用 join 方法后依然有 n 个线程，只是通过阻塞线程的方式，使得程序同步执行。**

### 多个线程按顺序依次执行

> 可以通过 join 方法实现，多个线程按指定顺序依次执行。

```java
public class JoinThread extends Thread {
    int count;
    Thread previousThread; //上一个线程
    public JoinThread(Thread previousThread, int count) {
        this.previousThread = previousThread;
        this.count = count;
    }
    @Override
    public void run() {
        try {
            previousThread.join();
            System.out.println("previousThreadName: " + previousThread.getName() + " currentThreadName: " + getName() + " num: " + count);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }
    public static void main(String[] args) throws InterruptedException {
        Thread previousThread = Thread.currentThread();
        for (int i = 0; i < 10; i++) {
            JoinThread joinThread = new JoinThread(previousThread, i);
            joinThread.start();
            previousThread = joinThread;
        }
        Thread.sleep(1000);
        System.out.println("main");
        Thread.sleep(1000);
    }
}
```

分析：在 thread-0 线程的 run 方法中执行了 `previousThread.join()` 此时的 previousThread 为 main 线程，这时发生的事情是：thread-0 线程被阻塞，直到 main 线程执行完才能继续执行 thread-0；在 thread-1 线程的 run 方法中执行的是 `thread-0.join()`，此时thread-1 线程被阻塞，即当 thread-0 执行完后才能继续执行 thread-1。

[参考资料](https://blog.csdn.net/u010983881/article/details/80257703)

### 线程死锁案例

> 两个线程相互等待对方释放锁，造成死锁。

```java
    public static void main(String[] args) {
        Object o1 = new Object();
        Object o2 = new Object();
        new Thread(() -> {
            synchronized (o1){
                System.out.println("A");
                try {
                    Thread.sleep(10000);
                    synchronized (o2){
                        System.out.println("B");
                    }
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
        }).start();
        new Thread(() -> {
            synchronized (o2){
                System.out.println("C");
                try {
                    Thread.sleep(10000);
                    synchronized (o1){
                        System.out.println("D");
                    }
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
        }).start();
    }
```
**更新中 ··· ···**

