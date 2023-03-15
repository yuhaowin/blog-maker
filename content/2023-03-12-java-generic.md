# Java generic & wildcard

### 泛型类

泛型类的定义

```text
    class name<T1, T2, Tn> { /* ... */ }
```

把 T1, T2, and Tn 称为泛型类型(Generic Types)，或者类型参数(type parameters), 类型参数可以在类中任何地方使用。
T1, T2, and Tn 在使用时可以是任何的类，接口，数组，甚至是另一个泛型类。

```text
    public class Box<T> {
        // T stands for "Type"
        private T t;

        public void set(T t) { this.t = t; }
        public T get() { return t; }
    }
```

### 泛型接口

和泛型类一样，泛型接口也可以定义一个或多个类型参数，将类型参数置于接口名的后面，跟着类似于类声明的尖括号。

```text
    public interface Pair<K, V> {
        public K getKey();
        public V getValue();
    }
```

```text
    public class OrderedPair<K, V> implements Pair<K, V> {
        private K key;
        private V value;
    
        public OrderedPair(K key, V value) {
	    this.key = key;
	    this.value = value;
        }
    
        public K getKey()	{ return key; }
        public V getValue() { return value; }
    }
```

OrderedPair<K, V> 是泛型的定义，implements Pair<K, V> 中的K和V是泛型的使用，而不是定义，因为OrderedPair<K, V>已经定义了K和V。

```text
    OrderedPair<String, Box<Integer>> p = new OrderedPair<>("primes", new Box<Integer>(...));
```

### 泛型方法

泛型方法有自己声明的类型参数，这些类型参数的声明放在返回值之前。但类型参数的作用域只限于方法本身。适用于静态方法和实例方法以及泛型类的构造器。
可以被用到形参声明、方法返回值、方法定义中的变量声明和类型转换。 泛型方法使得该泛型方法的类型参数独立于类而产生变化。泛型方法和泛型类没有关系

```text
    public class Util {
        public static <K, V> boolean compare(Pair<K, V> p1, Pair<K, V> p2) {
            return p1.getKey().equals(p2.getKey()) &&
                   p1.getValue().equals(p2.getValue());
        }
    }
```

```text
    Pair<Integer, String> p1 = new Pair<>(1, "apple");
    Pair<Integer, String> p2 = new Pair<>(2, "pear");
    boolean same = Util.<Integer, String>compare(p1, p2);
    // 借助类型推断，可以省略类型参数
    boolean same = Util.compare(p1, p2);
```

### 有界的类型参数

可以限制类型参数的范围，限制上界使用 extends 关键字。这里可以 extends 类和接口。

```text
    public class NaturalNumber<T extends Integer> {

        private T n;
    
        public NaturalNumber(T n)  { this.n = n; }
    
        public boolean isEven() {
            return n.intValue() % 2 == 0;
        }
    }
```

可以在类中使用泛型上界类的方法，如这里 Integer#intValue() 方法。

### 多个边界的类型参数

类型参数可以有多个上界，用 & 分隔。

```text
    <T extends A & B & C>    
```

这里类型参数 T 拥有 A B C 三个上界，这三个上界可以是类，也可以是接口，如果是类的话，只有有一个类，其他的都是接口，且类必须放在第一个。

```text
    Class A { /* ... */ }
    interface B { /* ... */ }
    interface C { /* ... */ }
    
    class D <T extends A & B & C> { /* ... */ }
```

泛型方法中的多个边界

```text
    public static <T extends Comparable & Serializable> T max(T[] a) { ... }
```

### 通配符类型

无边界通配符 Unbounded Wildcards
通常用于不依赖类型参数的方法中，如下面的 printList 方法。

```text
    public static void printList(List<?> list) {
        for (Object elem: list)
            System.out.print(elem + " ");
        System.out.println();
    }
```

上界通配符 Upper Bounded Wildcards，

```text
    List<? extends Number> list = new ArrayList<Integer>();
```

```text
    public static double sumOfList(List<? extends Number> list) {
        double s = 0.0;
        for (Number n : list)
            s += n.doubleValue();
        return s;
    }
```

下界通配符 Lower Bounded Wildcards

```text
    public static void addNumbers(List<? super Integer> list) {
        for (int i = 1; i <= 10; i++) {
            list.add(i);
        }
    }
```

### 类型参数和通配符的区别

+ 通配符有且仅有一个边界，而类型参数可以有多个边界。
+ 通配符可以有一个上界或有一个下界，而类型参数不支持下界。
+ 通配符不能在类、接口中使用，只能在方法中使用。

### 类型参数和通配符的选择

### 在方法中可以同时使用类型参数和通配符

```text
    public <T> void func(List<? extends T> list, T t) {
        list.add(t); // compile error 不可以修改 list
    }
```

```text
    public <T> void func(List<? super T> list, T t) {
        list.add(t); // 可以修改 list
    }
```

```text
    public <T, E extends T> void func(List<E> list, T t){
        list.add((E) t); // 可以修改，但是需要强制类型转换
    }
```

### 通配符中 extents 和 super 的选择

遵循 PECS 原则，Producer Extends Consumer Super。

+ 频繁往外读取内容的，适合用上界Extends，无法往里插入任何元素。
+ 经常往里插入的，适合用下界Super，取出来的元素都是 Object 类型。


```text
    public static <T> void copy(List<? extends T> src, List<? super T> dest) {
        for (int i = 0; i < src.size(); i++)
            dest.set(i, src.get(i));
    }
```

### 参考

https://blog.51cto.com/u_3664660/3213120
https://docs.oracle.com/javase/tutorial/java/generics/index.html
https://stackoverflow.com/questions/18176594/when-to-use-generic-methods-and-when-to-use-wild-card
