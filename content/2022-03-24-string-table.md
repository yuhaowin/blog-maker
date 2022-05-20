# 为什么要有 StringTable

> In computer science, [string interning](https://en.wikipedia.org/wiki/String_interning) is a method of storing only one copy of each distinct string value, which must be immutable. Interning strings makes some string processing tasks more time- or space-efficient at the cost of requiring more time when the string is created or interned. The distinct values are stored in a string intern pool.

为什么字符串要被设计为 Immutable？

1、便于实现 StringTable。

2、在多线程环境下，并发读取同一个字符串是安全的，不会产生竞争。

3、加快字符串处理速度，字符串不可变，hashcode 唯一，可以在对象中缓存 hashcode。

