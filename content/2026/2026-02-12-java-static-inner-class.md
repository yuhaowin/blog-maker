# Java Static Inner Class

```java
import java.lang.reflect.Field;

public class Outer {

    class Inner {
    }

    static class StaticInner {
    }

    public static void main(String[] args) {
        System.out.println("Inner fields:");
        for (Field f : Inner.class.getDeclaredFields()) {
            System.out.println("  " + f.getName() + " : " + f.getType());
        }

        System.out.println("StaticInner fields:");
        for (Field f : StaticInner.class.getDeclaredFields()) {
            System.out.println("  " + f.getName() + " : " + f.getType());
        }
    }
}
```

#### 运行结果：

```text
Inner fields:
  this$0 : class Outer
StaticInner fields:
```

#### 结论：
- Java 的静态内部类（Static Inner Class）不会隐式地持有对外部类实例的引用，因此它没有 `this$0` 字段。
- 非静态内部类（Inner Class）会隐式地持有对外部类实例的引用，因此它有一个 `this$0` 字段，指向外部类的实例。