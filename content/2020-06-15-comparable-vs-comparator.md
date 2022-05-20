---
title: 'Comparable & Comparator'
date: 2020-06-15 20:38:03
tags: []
published: true
hideInList: false
feature: 
isTop: false
---

> Java 中 Comparable 接口和 Comparator 接口的区别、分别在什么场景使用。

#### Comparable 接口

> Comparable 可比较的，表示具有比较的能力。如果某个类如：Animal 实现了 Comparable 接口，这表明 Animal 这个类是支持比较的。需要重写 Comparable 的 compareTo 方法。

但是这种方式是在这个类定义的时候其比较的策略就已经是确定的了，而且只有这一种比较策略。无法动态的更改比较策略。这是一种内部比较。

在 jdk 中实现 Comparable 接口的类有 String 。因此，String 是具有比较的能力的，但是无法修改比较的规则。

```java
    public static void main(String[] args) {
        List<String> list = new ArrayList<>();

        list.add("zhang");
        list.add("li");
        list.add("wang");

        System.out.println("排序前： " + list);
        Collections.sort(list);
        System.out.println("排序后：" + list);
    }
```

#### Comparator 接口

> Comparator 是比较器，可以对那些自身没有比较能力或是自身有比较能力，但是能力不符合要求的类提供一个比较的策略。

```java
    public static void main(String[] args) {
        List<String> list = new ArrayList<>();

        list.add("zhang");
        list.add("li");
        list.add("wang");

        System.out.println("排序前： " + list);
        Collections.sort(list, (str1, str2) -> {

            if (str1.length() > str2.length()) {
                return -1;
            } else if (str1.length() < str2.length()) {
                return 1;
            }
            return 0;

        });

        System.out.println("排序后：" + list);
    }
```
在 jdk 8 以前要使用 Comparator 有两种方法：
方法一、自己写一个类实现 Comparator 接口，实现 compare 方法。
方法二、通过匿名内部类的方式

在 jdk 8 之后可以通过 lambda 表达式的方式。

可以使用 lambda 的原因是，Comparator 是一个 @FunctionalInterface 接口。

##### 什么是函数式接口
> 一个接口只有一个抽象方法，或者有多个方法，但是只有一个方法是抽象的，其余的方法都是提供了默认实现，就是一个函数式接口

@FunctionalInterface 注解并不是必须的，和 @Override 都是为了在编译时提供校验是否符合要求用的。


#### 函数式接口与策略模式

> 函数式接口就是策略模式的运用。只是 jdk 8 以后提供了 lambda 这种更简单的使用方式。



+ [参考资料1](https://www.cnblogs.com/skywang12345/p/3324788.html)
+ [参考资料2](https://www.cnblogs.com/wangbin2188/p/10330231.html)
+ [参考资料3](https://mp.weixin.qq.com/s/8ZMqI6jJEOvqCk01pj05vA)