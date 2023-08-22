# JAVA 参数传递是值传递还是引用传递

**-「JAVA 参数传递是值传递还是引用传递？」**\
**-「JAVA 参数的传递是值传递。」**


> 所谓 **值传递 - Pass By Value**  是指在调用方法时将参数**复制**一份传递到方法中,这样在方法中如果对参数进行修改,将不会影响到实际参数。\
> 所谓 **引用传递 - Pass By Reference** 是指在调用方法时将实际参数直接传递到函数中,那么在方法中对参数所进行的修改,将影响到实际参数。

#### 先看一个简单的例子：

> 在方法中传递基本的数据类型

```java
public class MyTest {
    public void use(int arg) {
        arg = arg + 1;
        System.out.println(arg);
    }

    public static void main(String[] args) {
        MyTest test = new MyTest();
        int a = 1;
        System.out.println(a);      // 1
        test.use(a);                // 2
        System.out.println(a);      // 1
    }
}
```

从结果可以看出来，在方法内部修改基本数据类型如 int 的值不会对原始的参数的值产生影响，这是因为把 arg=1
拷贝到方法的内部，在方法内修改的只是原始值的拷贝，自然不会对原始值产生影响。

但是，如果由上述得出 java 是值传递的，一定会有人跳出来反驳，例如下面的例子：

```java
public class MyTest {
    public void use(Map map) {
        map.put("age", 18);
        System.out.println(map);
    }

    public static void main(String[] args) {
        MyTest test = new MyTest();
        Map map = new HashMap<>();
        map.put("name", "yuhao");
        System.out.println(map);        //{name=yuhao}
        test.use(map);                  //{name=yuhao, age=18}
        System.out.println(map);        //{name=yuhao, age=18}
    }
}
```

可以看出传入方法的是引用类型的 HashMap ，在方法内部修改的 map 的内容，在方法外面是有体现的。那是否可以据此得出对于引用类型的参数，java
是传引用的呢？

其实是把引用拷贝了一份在方法内，只是在方法内修改是这个引用指向的内容，而并没有修改引用本身，因此，在方法外面也反应出了这种变化。

**复制了就是传值， 没复制就是传引用。**

**传递的内容是引用类型 != 引用传递。**

如果，java 是引用传递的话，那么在方法的内部，对参数的指向进行的修改，那在方法的外部也应该是变化的。

java 规范中表明，java 引用类型的值，指的就是指针，并不是对象本身。通过 java 中的修改引用类型的值，是指对这个变量进行重新赋值。

另外一个简单的证明 java 不是引用拷贝的例子

```java
public class MyTest {
    private void init(MyTest objVar) {
        System.out.println(objVar);  // null
        objVar = new MyTest();
        System.out.println(objVar);  // MyTest@1b6d3586
    }

    public static void main(String[] args) {
        MyTest arg = null;
        MyTest test = new MyTest();
        test.init(arg);
        system.out.println(arg);     // null

        //After calling init method, arg still points to null
        //this is because arg is passed as value and not as reference.
    }
}
```

#### C++ 是支持引用传递的

```c++
#include <iostream>

int func(int& a, int& b)
{
    // a 和 b 均是实参 a 的别名
    a = 3; // 把实参 a 修改为 3
    b = 4; // 把实参 a 修改为 4
    return a + b; // 4 + 4 = 8
}

int main()
{
    int a = 1, b = 2, c = func(a, a);
    std::cout << a << b << c; // 428
}
```

+ [c++ - pass by reference](https://www.youtube.com/watch?v=gyIQ8YPeTuk)
+ [javascript-pass-by-value](https://www.aleksandrhovhannisyan.com/blog/javascript-pass-by-reference/#prerequisite-terminology-arguments-vs-parameters)
+ [is java "pass-by-reference" or "pass-by-value"?](https://stackoverflow.com/questions/40480/is-java-pass-by-reference-or-pass-by-value)
+ [What's the difference between passing by reference vs. passing by value?](https://stackoverflow.com/questions/373419/whats-the-difference-between-passing-by-reference-vs-passing-by-value/36208432)
+ [c++ 值传递、指针传递、引用传递详解](https://www.cnblogs.com/yanlingyin/archive/2011/12/07/2278961.html)
+ [go语言参数传递是传值还是传引用](https://www.flysnow.org/2018/02/24/golang-function-parameters-passed-by-value.html)
+ [java 参数传递是值传递还是引用传递？](https://www.cnblogs.com/9513-/p/8484071.html)
+ [java 到底是值传递还是引用传递？](https://www.zhihu.com/question/31203609)
