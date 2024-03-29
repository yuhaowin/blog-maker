# 常用的设计模式

## 一、设计原则

![7大设计原则](https://ws2.sinaimg.cn/large/006tKfTcgy1g0dbrlzsylj310k0l0k1l.jpg)

### 依赖倒置原则

> + 高层模块不应该依赖低层模块，两者都应该依赖抽象
>+ 抽象不应该依赖细节
>+ 细节应该依赖抽象

高层是策略，低层是细节，他们都依赖于中间的抽象层，对于 **倒置** 的理解在于低层依赖了抽象层。

+ [参考资料1](https://www.jianshu.com/p/314b67f04e50)
+ [参考资料2](https://www.cnblogs.com/yulang314/p/3551586.html)

## 二、创建型

![设计模式-创建型](https://ws2.sinaimg.cn/large/006tKfTcgy1g0dbo9fy9qj31140k4gug.jpg)

### 简单工厂

> 由工厂对象决定创建哪一种产品类的实例，强调`创建产品`，不属于GOF23种设计模式之一。

适用的场景：工厂类负责的产品类的数量较少。

优点：只需要传入正确的参数，就可以获取的需要的对象，而无需关系该对象详细的创建的过程。

缺点：工厂类的职责过重，当增加新的产品时，需要修改产品类的判断逻辑，违背开闭原则；无法形成基于继承的等级结构。

### 工厂方法

> 定义一个专门创建对象的接口,让实现了这个接口的类来决定创建哪一个类的实例， **工厂方法让类的实例化过程推迟到子类中进行。**

适用的场景：创建对象需要大量的重复代码，客户端无需依赖产品类的实例如何创建，和实现的细节等。

**产品族和产品等级的概念**

+ **产品族:一个家电公司生产的多个家电(如:格力空调,格力冰箱,格力手机)是一个产品族；**
+ **产品等级:格力手机，苹果手机，小米手机，他们是一个产品等级。**

假设有如下的背景：
如果创建一个Video对象很复杂，如果没需要一个Video对象是就需要重复写一遍创建的代码。可以考虑适应工厂方法。

工厂方法需要一个产品类，和一个创建该类型的工厂类。

抽象的产品类：

```java
public abstract class Video {
    public abstract void produce();
}
```

抽象的工厂类：

```java
//一般而言只有抽象方法时也可以使用interface代替抽象类
//实际业务中类中可能不仅仅有抽象方法还有实现方法,此时使用抽象类比较合适
public abstract class VideoFactory {
    //该方法只是定义一个规范
    public abstract Video getVideo();
}
```

假设此时需要创建一个JavaVideo这种类时，需要做的是，创建一个JavaVideo的产品类和一个JavaVideoFactory的工厂类。

JavaVideo产品类:

```java
public class JavaVideo extends Video {
    @Override
    public void produce() {
        System.out.println("录制java课程视频");
    }
}
```

JavaVideoFactory工厂类:

```java
public class JavaVideoFactory extends VideoFactory {
    @Override
    public Video getVideo() {
        return new JavaVideo();
    }
}
```

如果还有新的产品，只有再同样的增加产品类和对应的工厂类。

```java
/**
 * 客户端只需知道产品对应的产品工厂,无需知道具体的产品和创建产品实例的过程
 * 在新增一个产品的时候,只需要添加一个产品,和对应该产品的工厂,该过程无需修改已有的产品和工厂
 * 对扩展是开放的,符合开闭原则.
 */
public class Test {
    private static final Logger logger = LoggerFactory.getLogger(Test.class);

    public static void main(String[] args) {
        VideoFactory videoFactory = new PythonVideoFactory();
        Video video = videoFactory.getVideo();
        video.produce();

        //只需要跟换产生具体产品的工厂即可;
        VideoFactory videoFactory1 = new JavaVideoFactory();
        Video video1 = videoFactory1.getVideo();
        video1.produce();
    }
}
```

## 三、结构型

![设计模式-结构型](https://ws2.sinaimg.cn/large/006tKfTcgy1g0dbp8k6plj313c0lkgwp.jpg)

### 代理模式

> 代理模式：为其他对象提供一种代理,以控制对这个对象的访问;
>
代理对象在客户端和目标对象之间起到中介作用，如此便于在目标实现的基础上增加额外的功能操作，前拦截，后拦截等，以满足自身的业务需求;代理模式分为静态代理和动态代理，其中动态代理又分为jdk动态代理（基于反射实现的）和cglib `Code Generation Library`
（基于asm字节码实现的）。

#### 1. 静态代理 - JAVA

在代码中显性的创建特定的代理类，使用该类的对象对目标类的对象一些行为进行代理。并且可以根据业务场景对这些行为进行一定程度的增强，如:
在被代理的方法前后增加其他行为。

假设有以下业务场景：用户在下单时需要根据 userId 对订单数据进行分库存储，分别存储在 DB0、DB1 两个数据库中。

Order 类：

```java
public class Order {
    private Object orderInfo;
    private Integer userId;

    public Object getOrderInfo() {
        return orderInfo;
    }

    public void setOrderInfo(Object orderInfo) {
        this.orderInfo = orderInfo;
    }

    public Integer getUserId() {
        return userId;
    }

    public void setUserId(Integer userId) {
        this.userId = userId;
    }
}
```

DAO 层接口

```java
public interface OrderDao {
    int insert(Order order);
}
```

DAO 层接口实现

```java
public class OrderDaoImpl implements OrderDao {
    @Override
    public int insert(Order order) {
        System.out.println("DAO 层添加 order 成功了");
        return 1;
    }
}
```

Service 层接口

```java
public interface OrderService {
    int saveOrder(Order order);
}
```

Service 层接口实现

```java
public class OrderSercviceImpl implements OrderService {
    private OrderDao orderDao;

    @Override
    public int saveOrder(Order order) {
        iOrderDao = new OrderDaoImpl();
        System.out.println("Service 层调用 DAO 层添加 Order");
        return iOrderDao.insert(order);
    }
}
```

***********

静态代理类，对象目标对象的行为进行代理

```java
public class OrderServiceStaticProxy {

    /**
     * 需要代理的目标对象,对该对象的一些方法进行增强,这里增强的是saveOrder方法
     */
    private OrderService orderService;

    /**
     * 当客户端调用该方法时，会在目标类的orderService.saveOrder(order)方法
     * 前后执行beforeMethod(order)和afterMethod()方法，来增强目标对象的saveOrder(order)方法
     */
    public int saveOrder(Order order) {
        beforeMethod(order);
        orderService = new OrderSercviceImpl();
        int result = orderService.saveOrder(order);
        afterMethod();
        return result;
    }

    private void beforeMethod(Order order) {
        System.out.println("静态代理before code");
        Integer userId = order.getUserId();
        Integer dbRouter = userId % 2;
        System.out.println("静态代理分配到 [db" + dbRouter + "] 处理数据");
        //设置dataSource（这一步不是代理模式的重点，可以忽略）
        DataSourceContextHolder.setDBType("db" + dbRouter);
    }

    private void afterMethod() {
        System.out.println("静态代理after code");
    }
}
```

*********

客户端调用静态代理类

```java
public class Test {
    public static void main(String[] args) {
        Order order = new Order();
        order.setUserId(0);
        OrderServiceStaticProxy staticProxy = new OrderServiceStaticProxy();
        staticProxy.saveOrder(order);
    }
}
```

测试结果如下：

![测试结果](https://ws2.sinaimg.cn/large/006tKfTcgy1g0botsvjx0j30qc0asdhq.jpg)

#### 2. 基于 JDK 实现的动态代理 - JAVA

jdk实现的动态代理只能对实现了接口的类生成代理,并不能针对一个具体的实现类进行代理，并且在 jdk
动态代理中用到的代理类是在程序调用到代理类对象时,才由jvm真正创建,jvm根据传入的真正的业务实现类对象以及方法名
动态的创建了一个代理类的class文件 ,这个class文件被字节码引擎执行,然后通过该代理类的对象进行方法的调用。

引入的业务场景和上述一样，只是代理类是动态生成的。

```java
public class OrderServiceDynamicProxy implements InvocationHandler {

    /**
     * 被代理的目标对象 这里是 OrderServiceImpl
     */
    private Object target;

    public OrderServiceDynamicProxy(Object target) {
        this.target = target;
    }

    /**
     * 返回被代理的目标对象 这里是 OrderServiceImpl
     *
     * @return
     */
    public Object bind() {
        Class cls = target.getClass();
        //动态代理的核心
        return Proxy.newProxyInstance(cls.getClassLoader(), cls.getInterfaces(), this);
    }

    /**
     * @param proxy  代理类 一般自己实现invoke方法是很少使用该参数
     * @param method 要被增强的目标对象的方法对象
     * @param args   具体的method的参数
     */
    @Override
    public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {
        //在该case 中 argObject 就是 Order
        Object argObject = args[0];
        beforeMethod(argObject);
        Object resultObject = method.invoke(target, args);
        afterMethod();
        return resultObject;
    }

    private void beforeMethod(Object obj) {
        int userId = 0;
        if (obj instanceof Order) {
            userId = ((Order) obj).getUserId();
        }
        Integer dbRouter = userId % 2;
        System.out.println("动态代理分配到 [db" + dbRouter + "] 处理数据");
        //设置dataSource
        DataSourceContextHolder.setDBType("db" + dbRouter);
        System.out.println("动态代理before code");
    }

    private void afterMethod() {
        System.out.println("动态代理after code");
    }
}
```

客户端调用动态代理类

```java
public class Test {
    public static void main(String[] args) {
        /* 设置此系统属性,让JVM生成的Proxy类写入文件.保存路径为：com/sun/proxy(如果不存在请手工创建) */
        System.getProperties().put("sun.misc.ProxyGenerator.saveGeneratedFiles", "true");
        Order order = new Order();
        order.setUserId(1);
        OrderService dynamicProxy = (OrderService) new OrderServiceDynamicProxy(new OrderSercviceImpl()).bind();
        dynamicProxy.saveOrder(order);
    }
}
```

测试结果如下：

![](https://ws3.sinaimg.cn/large/006tKfTcgy1g0c3d3ilxhj30ma0aomyx.jpg)

在测试类中通过 OrderServiceDynamicProxy 类的 bind() 钩子方法
返回通过反射生成的代理类的对象，动态生成的代理类是以class文件(Proxy0.class)
的形式存在，通过对该class文件进行持久化，和通过jad反编译成java文件如下：

```java
public final class $Proxy0 extends Proxy implements OrderService {

    public $Proxy0(InvocationHandler invocationhandler) {
        super(invocationhandler);
    }

    public final boolean equals(Object obj) {
        try {
            return ((Boolean) super.h.invoke(this, m1, new Object[]{
                    obj
            })).booleanValue();
        } catch (Error _ex) {
        } catch (Throwable throwable) {
            throw new UndeclaredThrowableException(throwable);
        }
    }

    public final int saveOrder(Order order) {
        try {
            return ((Integer) super.h.invoke(this, m3, new Object[]{
                    order
            })).intValue();
        } catch (Error _ex) {
        } catch (Throwable throwable) {
            throw new UndeclaredThrowableException(throwable);
        }
    }

    public final String toString() {
        try {
            return (String) super.h.invoke(this, m2, null);
        } catch (Error _ex) {
        } catch (Throwable throwable) {
            throw new UndeclaredThrowableException(throwable);
        }
    }

    public final int hashCode() {
        try {
            return ((Integer) super.h.invoke(this, m0, null)).intValue();
        } catch (Error _ex) {
        } catch (Throwable throwable) {
            throw new UndeclaredThrowableException(throwable);
        }
    }

    private static Method m1;
    private static Method m3;
    private static Method m2;
    private static Method m0;

    static {
        try {
            m1 = Class.forName("java.lang.Object").getMethod("equals", new Class[]{
                    Class.forName("java.lang.Object")
            });
            m3 = Class.forName("com.yuhaowin.design.structure.proxy.OrderService").getMethod("saveOrder", new Class[]{
                    Class.forName("com.yuhaowin.design.structure.proxy.Order")
            });
            m2 = Class.forName("java.lang.Object").getMethod("toString", new Class[0]);
            m0 = Class.forName("java.lang.Object").getMethod("hashCode", new Class[0]);
        } catch (NoSuchMethodException nosuchmethodexception) {
            throw new NoSuchMethodError(nosuchmethodexception.getMessage());
        } catch (ClassNotFoundException classnotfoundexception) {
            throw new NoClassDefFoundError(classnotfoundexception.getMessage());
        }
    }
}
```

可以看出该动态生成的该代理类extends 了 Proxy 和 implements OrderService

所以该类可以强转为OrderService

```
    OrderService dynamicProxy=(OrderService)new OrderServiceDynamicProxy(new OrderSercviceImpl()).bind();
```

由于实现了OrderService接口，所以在该类中重写了OrderService 的saveOrder()方法

```
    public final int saveOrder(Order order){
        try{
        return((Integer)super.h.invoke(this,m3,new Object[]{
        order
        })).intValue();
        }catch(Error _ex){
        }catch(Throwable throwable){
        throw new UndeclaredThrowableException(throwable);
        }
        }
```

由于继承了Proxy，所以可以调用父类的 h 参数 即：invocationhandler 参数，这个参数是在bind()这个方法中通过
Proxy.newProxyInstance(cls.getClassLoader(), cls.getInterfaces(), this) 传入，这里的this 表示的是
OrderServiceDynamicProxy 但是
OrderServiceDynamicProxy这个类实现了Invocationhandler 所以this就是invocationhandler

所以在代理类的saveOrder方法中调用的super.h.invoke() 就是调用OrderServiceDynamicProxy类的重写的invoke();

********

#### 3. 基于 cglib 实现的动态代理 - JAVA

?> 待更新。。。

## 三、行为型

![设计模式-行为型](https://ws2.sinaimg.cn/large/006tKfTcgy1g0dbq5lx67j31440lgaqg.jpg)

### 模板方法模式

> 模板方法模式：定义了一个算法的骨架,并允许子类为一个或者多个步骤提供实现;模板方法使得子类可以在不改变算法结构的情况下,重新定义算法的某些步骤。简而言之模板方法的目的是：让子类可以实现或者扩展父类中固定算法的某些部分。

模板方法的适用场景：

+ 一次性实现一个算法的不变的部分,并将可变的行为留个子类实现
+ 各子类中公共的行为被抽取出来,并集中到一个公共的父类中,以避免代码重复.

如：把东西放入冰箱中，可以使用模板方法，打开冰箱 --> 放东西 -->
关闭冰箱；整个算法的步骤和顺序是不变的，而且打开冰箱和关闭冰箱是固定的，这两个步骤由父类实现，具体放入什么东西有具体的子类决定。

模板方法的扩展

+ 钩子方法：钩子方法提供了缺省的行为,子类可以在必要的实现进行重写,目的是模板 对子类的进一步开放

假设有以下业务场景，讲师在制作教学课程时，需要制作幻灯片、制作视频，有可能还需要提供API文档（API文档不是所有课程必须提供的）

抽象类 Acourse 由该抽象类定义算法骨架，并且实现算法中公共、不变的部分。并将变化的部分开放给子类实现。

```java
public abstract class ACourse {
    //该方法是算法的骨架,使用final修饰是不希望子类破坏算法的结构;
    protected final void makeCourse() {
        this.makePPT();
        this.makeVideo();
        if (needWriteArticle()) {
            this.writeArticle();
        }
        this.packageCourse();
    }

    final void makePPT() {
        System.out.println("制作ppt");
    }

    final void makeVideo() {
        System.out.println("制作视频");
    }

    final void writeArticle() {
        System.out.println("编写手记");
    }

    //钩子方法,并提供了一个默认的实现,不使用final修饰,子类可以根据需要重新该方法
    protected Boolean needWriteArticle() {
        return false;
    }

    // 算法中不确定的部分,只定义结构,具体实现交由各子类实现
    abstract void packageCourse();
}
```

以下是两个子类

```java
public class DesignPatternCourse extends ACourse {
    @Override
    void packageCourse() {
        System.out.println("提供java课程源代码");
    }

    //该课程需要提供手记,但是ACourse中默认是不提供手记,需要重写该钩子方法
    @Override
    protected Boolean needWriteArticle() {
        return true;
    }
}
```

```java
public class FECourse extends ACourse {
    private Boolean needWriteArticleFlag = false;

    @Override
    void packageCourse() {
        System.out.println("提供前端课程源代码");
        System.out.println("提供前端课程多媒体素材");
    }

    public void setNeedWriteArticleFlag(Boolean needWriteArticleFlag) {
        this.needWriteArticleFlag = needWriteArticleFlag;
    }

    //由于前端课程的不同,有的需要提供手记有的不需要提供,所以将是否提供的权限开放给客户端
    @Override
    protected Boolean needWriteArticle() {
        return needWriteArticleFlag;
    }
}
```

客户端类

```java
public class Test {
    public static void main(String[] args) {
        ACourse designPattern = new DesignPatternCourse();
        designPattern.makeCourse();

        System.out.println();

        ACourse fecourse = new FECourse();
        ((FECourse) fecourse).setNeedWriteArticleFlag(true);
        fecourse.makeCourse();
    }
}
```

测试结果如下：

![](https://ws4.sinaimg.cn/large/006tKfTcgy1g0cudqknuvj30qq0du40m.jpg)

### 策略模式

> 定义算法家族，分别封装起来，让它们之间可以相互替换，客户端动态的选择某种算法。此模式让算法的变化不会影响使用算法的用户。

应用场景：下单时，满减、立减，优惠券等使用不同的策略。可以优雅的处理掉 if...else...

缺点：客户端必须知道所有的策略类，并自行选择策略。

> 假设不同用户下单买东西的折扣不一样。<br>
> 1、外卖平台上的某家店铺为了促销，设置了多种会员优惠，其中包含超级会员折扣8折、普通会员折扣9折和普通用户没有折扣三种。<br>
> 2、希望用户在付款的时候，根据用户的会员等级，就可以知道用户符合哪种折扣策略，进而进行打折，计算出应付金额。<br>
> 3、随着业务发展，新的需求要求专属会员要在店铺下单金额大于30元的时候才可以享受优惠。<br>
> 4、接着，又有一个变态的需求，如果用户的超级会员已经到期了，并且到期时间在一周内，那么就对用户的单笔订单按照超级会员进行折扣，并在收银台进行强提醒，引导用户再次开通会员，而且折扣只进行一次。

定义一个接口：

```java
public interface UserPayService {
    /**
     * 计算应付价格
     */
    BigDecimal quote(BigDecimal orderPrice);
}
```

定义不同策略实现类：

```java
//专属会员
public class ParticularlyVipPayService implements UserPayService {
    @Override
    public BigDecimal quote(BigDecimal orderPrice) {
        if (消费金额大于30元) {
            return 7 折价格;
        }
    }
}

//超级会员
public class SuperVipPayService implements UserPayService {
    @Override
    public BigDecimal quote(BigDecimal orderPrice) {
        return 8 折价格;
    }
}

//普通会员
public class VipPayService implements UserPayService {
    @Override
    public BigDecimal quote(BigDecimal orderPrice) {
        if (该用户超级会员刚过期并且尚未使用过临时折扣) {
            临时折扣使用次数更新();
            returen 8 折价格;
        }
        return 9 折价格;
    }
}
```

客户端类：

```java
public BigDecimal calPrice(BigDecimal orderPrice,User user){

        String vipType=user.getVipType();

        if(vipType==专属会员){
        UserPayService strategy=new ParticularlyVipPayService();
        return strategy.quote(orderPrice);
        }
        if(vipType==超级会员){
        UserPayService strategy=new SuperVipPayService();
        return strategy.quote(orderPrice);
        }
        if(vipType==普通会员){
        UserPayService strategy=new VipPayService();
        return strategy.quote(orderPrice);
        }
        return 原价;
        }
```

由上述代码可知，在客户端必须判断并选择需要的策略类，并没有完全消灭 if else，这一点可以通过引入工厂模式解决

```java
public class UserPayServiceStrategyFactory {
    // 保存所有的策略
    private static Map<String, UserPayService> services = new ConcurrentHashMap<String, UserPayService>();

    // 获取指导类型的策略
    public static UserPayService getByUserType(String type) {
        return services.get(type);
    }

    // 注册策略
    public static void register(String userType, UserPayService userPayService) {
        Assert.notNull(userType, "userType can't be null");
        services.put(userType, userPayService);
    }
}
```

此时客户端代码可以改为：

```java
public BigDecimal calPrice(BigDecimal orderPrice,User user){
        String vipType=user.getVipType();
        UserPayService strategy=UserPayServiceStrategyFactory.getByUserType(vipType);
        return strategy.quote(orderPrice);
        }
```

还记得我们前面定义的UserPayServiceStrategyFactory中提供了的register方法吗？他就是用来注册策略服务的。

接下来，我们就想办法调用register方法，把Spring通过IOC创建出来的Bean注册进去就行了。

这种需求，可以借用Spring种提供的InitializingBean接口，这个接口为Bean提供了属性初始化后的处理方法，它只包括afterPropertiesSet方法，凡是继承该接口的类，在bean的属性初始化后都会执行该方法。

那么，我们将前面的各个策略类稍作改造即可：

```java

@Service
public class ParticularlyVipPayService implements UserPayService, InitializingBean {
    @Override
    public BigDecimal quote(BigDecimal orderPrice) {
        if (消费金额大于30元) {
            return 7 折价格;
        }
    }

    @Override
    public void afterPropertiesSet() throws Exception {
        UserPayServiceStrategyFactory.register("ParticularlyVip", this);
    }
}
```

只需要每一个策略服务的实现类都实现InitializingBean接口，并实现其afterPropertiesSet方法，在这个方法中调用UserPayServiceStrategyFactory.register即可。

这样，在Spring初始化的时候，当创建VipPayService、SuperVipPayService和ParticularlyVipPayService的时候，会在Bean的属性初始化之后，把这个Bean注册到UserPayServiceStrategyFactory中。

### 责任链模式

> 为一个请求创建一个接收此次请求对象的一个链条。

责任链模式的适用场景有：
一个请求的处理需要一个或者多个其他对象进行辅助处理。

如：

+ 在对检验信息进行交易的时候适合使用责任链模式
+ 在审批流程中适合使用责任链模式

> 假设有如下的业务场景：某视频网站需要对上架的视频课程进行审核（检查是否有手记和视频）如果都有则允许上架，没有则不允许上架，首先审核的是是否含有手记。

课程实体类：

```java
public class Course {
    private String name;
    private String article;
    private String video;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getArticle() {
        return article;
    }

    public void setArticle(String article) {
        this.article = article;
    }

    public String getVideo() {
        return video;
    }

    public void setVideo(String video) {
        this.video = video;
    }

    @Override
    public String toString() {
        return "Course{" +
                "name='" + name + '\'' +
                ", article='" + article + '\'' +
                ", video='" + video + '\'' +
                '}';
    }
}
```

责任链模式的核心类：

```java
/**
 * 课程的批准者
 *
 * 所有的批准者都需要继承这个批准者
 */
public abstract class Approver {
    //这个类的核心是需要包含一个与自己相同类型的对象
    protected Approver approver;

    /**
     * 设置下一个批准者
     *
     * @param approver
     */
    public void setNextApprover(Approver approver) {
        this.approver = approver;
    }

    /**
     * 发布课程,由子类实现.
     *
     * @param course
     */
    public abstract void deployCourse(Course course);
}
```

分别有手记的审核类：

```java
/**
 * 手记审批
 */
public class ArticleApprover extends Approver {
    @Override
    public void deployCourse(Course course) {
        if (StringUtils.isNoneEmpty(course.getArticle())) {
            System.out.println(course.getName() + "包含手记, 审批通过");
            if (approver != null) {
                approver.deployCourse(course);
            }
        } else {
            System.out.println(course.getName() + "不包含手记, 审批不通过");
            return;
        }
    }
}
```

视频的审核类：

```java
/**
 * 视频审批
 */
public class VideoApprover extends Approver {
    @Override
    public void deployCourse(Course course) {
        if (StringUtils.isNoneEmpty(course.getVideo())) {
            System.out.println(course.getName() + "包含视频, 审批通过");
            if (approver != null) {
                approver.deployCourse(course);
            }
        } else {
            System.out.println(course.getName() + "不包含视频, 审批不通过");
            return;
        }
    }
}
```

客户端测试类：

```java
public class Test {
    public static void main(String[] args) {
        Approver articleApprover = new ArticleApprover();
        Approver videoApprover = new VideoApprover();

        Course course = new Course();
        course.setName("java课程");
        course.setArticle("java课程手记");
        course.setVideo("java课程视频");

        articleApprover.setNextApprover(videoApprover);
        articleApprover.deployCourse(course);

        // 注意只有一个审核员发布
        // 回循环调用,造成内存溢出
        //videoApprover.setNextApprover(articleApprover);
        //videoApprover.deployCourse(course);
    }
}
```