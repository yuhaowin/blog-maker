# 可执行 jar 打包方式

> 一般而言 maven 打出的 jar 只包含工程代码，不包含依赖的第三方 jar。这类 jar 无法直接使用。

为了简化部署，期望一个项目可以直接通过 `java -jar` 的方式启动。这种 jar 包，被称为可执行 jar。

#### 可执行 jar 打包方法

##### fat-jar(uber-jar)

> fat-jar 是将工程代码、第三方代码全都展开打到一个 jar 包中。只包含 class 文件。

可以通过 [maven-shade-plugin](https://maven.apache.org/plugins/maven-shade-plugin/) plugin 实现

参考配置:

```xml
            <plugin>
                <groupId>org.apache.maven.plugins</groupId>
                <artifactId>maven-shade-plugin</artifactId>
                <version>2.4.3</version>
                <configuration>
                    <transformers>
                        <transformer
                                implementation="org.apache.maven.plugins.shade.resource.ManifestResourceTransformer">
                            <mainClass>com.xxx.xxx.xxx</mainClass>
                        </transformer>
                    </transformers>
                </configuration>
                <executions>
                    <execution>
                        <phase>package</phase>
                        <goals>
                            <goal>shade</goal>
                        </goals>
                    </execution>
                </executions>
            </plugin>
```



![](https://image.yuhaowin.com/2021/12/15/195946.png)

##### jar in jar

> jar in jar 是将工程代码、第三方代码以 jar 的形式打到一个 jar 包中,包含了工程的 class，和第三方的 jar 包。这种打包方法需要重新实现 classloader。

可以通过 [spring-boot-maven-plugin](https://docs.spring.io/spring-boot/docs/current/maven-plugin/reference/htmlsingle/) plugin 实现

该插件已经实现了 classloader, 并且会自动寻找到启动类。个人比较喜欢 spring-boot-maven-plugin ，因为工程代码和第三方类库是区分开的。

参考配置：

```xml
            <plugin>
                <groupId>org.springframework.boot</groupId>
                <artifactId>spring-boot-maven-plugin</artifactId>
                <version>2.0.2.RELEASE</version>
                <executions>
                    <execution>
                        <goals>
                            <goal>repackage</goal>
                        </goals>
                    </execution>
                </executions>
            </plugin>
```



![](https://image.yuhaowin.com/2021/12/15/200013.png)





补充网友的原理分析：[spring-boot-maven-plugin 原理分析](https://blog.csdn.net/ttzommed/article/details/114984341)


