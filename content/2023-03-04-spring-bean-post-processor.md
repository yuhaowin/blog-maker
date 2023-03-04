# Spring Bean 被提前创建，导致 BeanPostProcessor 对其失效

## 问题

在一个 Spring Boot 项目中，我们集成了 Shiro 权限框架，在我们实现 AuthorizingRealm 的 UserRealm 类中，通过 `@Autowired` 注解注入了一个 LoginService 类，用于获取用户信息。

LoginService 中有使用自定义注解用于在 Bean 实例化阶段获取配置信息。通过 BeanPostProcessor 实现的。但 LoginService 在实例化的使用，没有被 BeanPostProcessor 增强。

```text
    @Configuration
    public class ShiroConfiguration {
        @Bean
        public UserRealm userRealm() {
            UserRealm userRealm = new UserRealm();
            return userRealm;
        }    
    }
```

```text
    @Slf4j
    public class UserRealm extends AuthorizingRealm {
        @Autowired
        private LoginService loginService;
    ｝    
```

```text
    @Service
    public class LoginServiceImpl implements LoginService {
        @Config("prefix")
        private String prefix;
    }
```

## 原因

UserRealm 是通过 `@Configuration` + `@Bean` 的方式注入到 Spring 容器中的。通过这种方式创建的 UserRealm 实例时，会导致 UserRealm 依赖的 LoginService 实例被提前创建，从而导致 BeanPostProcessor 对其失效。

## 解决

通过 `@Lazy` 注解，延迟 LoginService 实例的创建。

```text
    @Slf4j
    public class UserRealm extends AuthorizingRealm {
        @Lazy
        @Autowired
        private LoginService loginService;
    ｝    
```

[参考](https://www.jianshu.com/p/916a7d8311bf)