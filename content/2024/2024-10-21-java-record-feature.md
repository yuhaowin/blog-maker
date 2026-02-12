# java record class feature

## 使用 jackson 反序列化 record 类型遇到的问题

### 定义一个 record 类

> 包含两个组件(字段), 一个是基本类型 age, 一个是集合类型 nickNames。

```java
public record Person(int age, List<String> nickNames) {
}
```

上述代码没有声明构造器，java 编译器会自动生成一个标准构造器，参数是 age 和 nickNames。

```java
public record Person(int age, List<String> nickNames) {
    // java 编译器会自动生成一个标准构造器
    public Person(int age, List<String> nickNames) {
        this.age = age;
        this.nickNames = nickNames;
    }
}
```

> 通过 jackson 反序列化 json 字符串到 Person 类型对象。

```json5
{
  "age": 18,
  "nick_names": [
    "Tom",
    "Jerry"
  ]
}
```

### 问题 1

> 通过 jackson 反序列化 json 字符串到 Person 类型对象时，遇到如下异常。

```java
public record Person(int age, List<String> nickNames) {

    public static void main(String[] args) throws Exception {
        ObjectMapper mapper = new ObjectMapper();
        Person person = mapper.readValue("""
                {
                  "age": 18,
                  "nick_names": ["Tom", "Jerry"]
                }
                """, Person.class);
        System.out.println(mapper.writeValueAsString(person));
    }
}
```

```text
Exception in thread "main" com.fasterxml.jackson.databind.exc.UnrecognizedPropertyException: Unrecognized field "nick_names"
```

这是因为 jackson 默认使用驼峰命名法，而 json 字符串中的字段是下划线命名法，所以需要在 record 类中使用 @JsonProperty 注解。

```java
public record Person(int age, @JsonProperty("nick_names") List<String> nickNames) {

    public static void main(String[] args) throws Exception {
        ObjectMapper mapper = new ObjectMapper();
        Person person = mapper.readValue("""
                {
                  "age": 18,
                  "nick_names": ["Tom", "Jerry"]
                }
                """, Person.class);
        System.out.println(mapper.writeValueAsString(person));
    }
}
```

同样会生成一个标准构造器。

```java
public record Person(int age, @JsonProperty("nick_names") List<String> nickNames) {
    // java 编译器会自动生成一个标准构造器
    public Person(int age, @JsonProperty("nick_names") List<String> nickNames) {
        this.age = age;
        this.nickNames = nickNames;
    }
}
```

### 对 nickNames 字段进行校验

> 使用简化构造器（Compact Constructor）对 nickNames 字段进行校验，如果为 null，则初始化为一个空的集合。
> 这种构造器可以省略构造器参数列表，直接在构造器主体中编写逻辑。编译器会自动推断出所有定义的字段，并进行赋值。

```java
public record Person(int age, @JsonProperty("nick_names") List<String> nickNames) {

    // 使用简化构造器对 nickNames 字段进行校验
    public Person {
        if (Objects.isNull(nickNames)) {
            nickNames = List.of();
        }
    }

    public static void main(String[] args) throws Exception {
        ObjectMapper mapper = new ObjectMapper();
        Person person = mapper.readValue("""
                {
                  "age": 18,
                  "nick_names": ["Tom", "Jerry"]
                }
                """, Person.class);
        System.out.println(mapper.writeValueAsString(person));
    }
}
```

上述简化构造编译后等价于下面代码实现

```java
public record Person(int age, @JsonProperty("nick_names") List<String> nickNames) {
    // 上述简化构造编译后等价于下面的标准构造器
    public Person(int age, @JsonProperty("nick_names") List<String> nickNames) {
        if (Objects.isNull(nickNames)) {
            nickNames = List.of();
        }
        this.age = age;
        this.nickNames = nickNames;
    }
}
```

同样我们也开始手动实现一个标准构造器，标准构造器。

```java
public record Person(int age, @JsonProperty("nick_names") List<String> nickNames) {

    // 手动实现标准构造器，对 nickNames 字段进行校验
    public Person(int age, @JsonProperty("nick_names") List<String> nickNames) {
        this.age = age;
        this.nickNames = Objects.isNull(nickNames) ? List.of() : nickNames;
    }

    public static void main(String[] args) throws Exception {
        ObjectMapper mapper = new ObjectMapper();
        Person person = mapper.readValue("""
                {
                  "age": 18,
                  "nick_names": ["Tom", "Jerry"]
                }
                """, Person.class);
        System.out.println(mapper.writeValueAsString(person));
    }
}
```

> 需要注意的是，手动实现的标准构造器的 nickNames 字段必须用 @JsonProperty("nick_names") 修饰。不然会出现一下错误。

```java
public record Person(int age, @JsonProperty("nick_names") List<String> nickNames) {

    // nickNames 字段没有使用 @JsonProperty("nick_names") 修饰，导致反序列化失败
    public Person(int age, List<String> nickNames) {
        this.age = age;
        this.nickNames = Objects.isNull(nickNames) ? List.of() : nickNames;
    }

    public static void main(String[] args) throws Exception {
        ObjectMapper mapper = new ObjectMapper();
        Person person = mapper.readValue("""
                {
                  "age": 18,
                  "nick_names": ["Tom", "Jerry"]
                }
                """, Person.class);
        System.out.println(mapper.writeValueAsString(person));
    }
}
```

```text
Exception in thread "main" com.fasterxml.jackson.databind.JsonMappingException: Can not set final java.util.List field com.wosai.crm.qa.Person.nickNames to java.util.ArrayList
```

这是因为 jackson 无法通过构造器完成 Person 对象的反序列化，构造器参数 nickNames 无法和 nick_names 对应，然后使用反射机制直接修改了
nickNames 字段，而 nickNames 字段是 final 的，所以会抛出异常。

*********

### 问题 2

> 如果使用普通类代替 record 类，需要使用 @JsonCreator 修改构造器和使用 @JsonProperty 修饰构造器参数。

```java
public class Person {

    @JsonProperty("age")
    private final Long age;

    @JsonProperty("nick_names")
    private final List<String> nickNames;

    // 使用 @JsonCreator 修改构造器和使用 @JsonProperty 修饰构造器参数。
    @JsonCreator
    public Person(@JsonProperty("age") Long age, @JsonProperty("nick_names") List<String> nickNames) {
        this.age = age;
        this.nickNames = nickNames;
    }

    public static void main(String[] args) throws JsonProcessingException {
        ObjectMapper mapper = new ObjectMapper();
        Person metaData = mapper.readValue("""
                {
                   "age": 18,
                   "nick_names": ["Tom", "Jerry"]
                 }
                """, Person.class);
        System.out.println(mapper.writeValueAsString(metaData));
    }
}
```
***********

 + [JEP 395: Records](https://openjdk.org/jeps/395)
 + [Why when a constructor is annotated with @JsonCreator, its arguments must be annotated with @JsonProperty?](https://stackoverflow.com/questions/21920367/why-when-a-constructor-is-annotated-with-jsoncreator-its-arguments-must-be-ann)