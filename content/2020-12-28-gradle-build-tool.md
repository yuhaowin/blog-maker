# Gradle 项目构建工具

### Gradle

+ Gradle 核心概念
  + Distribution
  + Wrapper
  + Gradle UserHome
  + Daemon
+ Groovy 基础
  + 动态调用和MOP
  + 闭包
+ Gradle 构建
  + Project
  + Task
  + Lifecycle、Hook
+ Gradle Plugin 编写
  + 构建逻辑的复用
  + 简单插件
  + script 插件
  + buildSrc 插件
  + 发布插件

***********

#### [Distribution](https://services.gradle.org/distributions/)

#### Wrapper

> 通过 gradlew 脚本，下载和具体项目相关的 gradle 版本。下载后的 gradle 默认存放在 `~/.gradle/wrapper/dists/`

![162211](https://image.yuhaowin.com/2020/12/27/162222.png)

#### Gradle UserHome

当前用户的 `.gradle`  目录就是 Gradle UserHome。

#### Daemon

> Gradle 和 Maven 不同的一个地方是，Gradle 默认会启动一个 Daemon 进程，来加快 Gradle 的启动速度。通过 Client 进程和 Daemon 进程进行通信。



#### 一个 demo

一个简单的 gradle 项目的结构如下：

![170852](https://image.yuhaowin.com/2020/12/27/170926.png)



+ build.gradle -- 构建的核心说明文件，类似 maven 的 pom.xml 文件。




#### [Gradle Lifecycle](https://docs.gradle.org/current/userguide/build_lifecycle.html)





![Gradle教程](https://image.yuhaowin.com/2021/01/18/155222.jpg)

gradle 的使用

gradle 的脚本使用的是 groovy 语言编写的。groovy 是一种 java 虚拟机的动态语言，最终编译完成后会生成java字节码。

groovy vs java

+ groovy 完全兼容 java 语法，完全可以将一个.java 文件改为.groovy 文件。同时在groovy中可以混合java语法。
+ groovy中 代码行尾的分号可以不写。
+ groovy的类和方法默认都是public的。没有default作用域。
+ 编译器会自动给属性添加 getter 和 setter 方法。
+ 属性可以直接点出来。
+ 方法的最后一个表达式的值会作为返回值，（可以省略 return 关键字）。
+ == 和 equals() 是等价的。

groovy 高级用法
+ 随处可以的 assert 断言操作。
+ 无需指定变量类型，相当于弱类型无语如：javascript def a = 12。
+ 调用方法是的括号不是必须的。
+ 闭包

```groovy
def a = 3;
def b = 3;
def list = new ArrayList<>()
list.add(1);
list.add(2)
list << '你好'
println list
assert list.size() == 3
println list.getClass();
assert list.getClass() == ArrayList

//闭包
//一个代码块，可以赋值给一个变量，也可以做为一个方法的参数。
def c1 = {
    v1, v2 -> //v1,v2 是闭包的参数，省略的参数的类型 箭头后面是方法体。
        println(v1 + v2)
}
def c2 = {
    println("你好");
}

def method1(Closure closure) {
    closure('hello ', 'word');
}

def method2(Closure closure) {
    closure();
}

method1 c1
method2 c2

method2({
    println("你好1");
}
)

method2 {
    println("你好2");
}
```

```
plugins {
    id 'java'
}
apply(plugin:"java")  //project对象的方法
apply(plugin:"war")
group 'com.yuhaowin'  //project对象的属性
version '1.0-SNAPSHOT'
sourceCompatibility = 1.8
repositories { //repositories 方法接收一个闭包，这里省略了方法的括号
    mavenCentral()
}
dependencies {//dependencies 方法接收一个闭包，这里省略了方法的括号
    testCompile group: 'junit', name: 'junit', version: '4.12'
}
```


![简单gradle项目示例](https://image.yuhaowin.com/2021/01/18/155222.jpg)

#### 构建脚本 (bulid.gradle)

##### 构建块

gradle 构建中有两个基本概念：项目-project、任务-task，每一个构建至少包含一个project，一个project包含一个或多个task。在多项目构建中，一个项目可以依赖其他的项目，任务也可以依赖其他的任务，以确定任务的执行顺序。

##### 构建脚本-project

一个项目代表一个正在构建的组件如：一个jar文件，当构建启动后，gradle会基于build.gradle实例化一个`org.gradle.api.Project`类的对象，并且可以通过project变量使其隐式可用。

##### 构建脚本-task

任务对应的是`org.gradle.api.Task`类，主要包括任务的动作和任务的依赖，任务定义了一个最小的工作单元，可以定义依赖的其他任务、动作和执行条件。