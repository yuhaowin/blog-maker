# MySQL InnoDB Buffer Pool Size 调优

> 背景：项目提测，发现数据库性能不佳，内存占用并不高，经过排查，发现是数据库的缓存池配置不合理，导致数据库频繁 IO，影响性能。

### MySQL 内存分配

MySQL 的内存可以分为 global 级的共享内存和 session 级的私有内存两部分，共享内存是实例启动时分配的是所有连接共享的。
私有内存用于每个连接到 MySQL 服务器时才分配各自的缓存，一些特殊的 SQL 或字段类型会导致单个线程可能分配多次缓存，因此当出现
OOM 异常，都是由各个连接私有内存造成的。

### InnoDB Buffer Pool

该部分缓存是 InnoDB
引擎最重要的缓存区域，是通过内存来弥补物理数据文件的重要手段，MySQL [参考手册推荐该部分配置为机器内存规格配置的50% - 75%](https://dev.mysql.com/doc/refman/8.0/en/memory-use.html)
。其中主要包含数据页、索引页、undo 页、insert buffer、自适应哈希索引、锁信息以及数据字典等信息。在进行 SQL
读和写的操作时，首先并不是对物理数据文件操作，而是先对 buffer_pool 进行操作，再通过 checkpoint
等机制写回数据文件。该空间的优点是可以提升数据库的性能、加快 SQL 运行速度，缺点是故障恢复速度较慢。

### InnoDB Buffer Pool Size 调优

#### 1. 查看当前 Buffer Pool Size

> 默认大小为 128M，支持配置的单位时 K、M、G，不支持小数点，如果配置的值不是这些单位，会被转换为字节。

```shell
mysql> show variables like '%innodb_buffer_pool_size%';
```

#### 2. 查看当前 Buffer Pool 使用情况

```shell
mysql> show engine innodb status
```

#### 3. 查看当前 Buffer Pool 命中率

```shell
mysql> show global status like 'innodb_buffer_pool_read%';
```

### MySQL 基准测试

> 修改配置前后需要进行基准测试，以便对比优化效果。

#### 使用 [sysbench](https://github.com/akopytov/sysbench) 做基准测试

##### 1. 安装 sysbench

```shell
curl -s https://packagecloud.io/install/repositories/akopytov/sysbench/script.rpm.sh | sudo bash
sudo yum -y install sysbench
```

##### 2. 查看版本

```shell
sysbench --version
```

##### 3. 准备数据

```shell
sysbench oltp_common.lua --time=300 --mysql-host=10.120.68.91 --mysql-port=3306 --mysql-user=root --mysql-password=xkjc_123 --mysql-db=sbtest --table-size=1000000 --tables=10 --threads=32 --events=999999999 prepare
```

##### 4. 运行测试

```shell
sysbench oltp_read_write.lua --time=300 --mysql-host=10.120.68.91 --mysql-port=3306 --mysql-user=root --mysql-password=xkjc_123 --mysql-db=sbtest --table-size=1000000 --tables=10 --threads=16 --events=999999999  --report-interval=10  run
```

##### 5. 清理数据

```shell
sysbench oltp_read_write.lua --time=300 --mysql-host=10.120.68.91 --mysql-port=3306 --mysql-user=root --mysql-password=xkjc_123 --mysql-db=sbtest --table-size=1000000 --tables=10 --threads=16 --events=999999999 --report-interval=10 cleanup
```

##### 6. 效果对比

> 机器物理内存 16G。

buffer_pool_size 默认为 128M 时，性能如下：

![150923](https://image.yuhaowin.com/2023/06/06/150923.png)

buffer_pool_size 调整为 6G 时，性能如下：

![151022](https://image.yuhaowin.com/2023/06/06/151022.png)

### innodb_buffer_pool_size 无法动态调整到 1G 以下问题

如果需要调整 innodb_buffer_pool_size 至 1G 以下，innodb_buffer_pool_instances 参数值必须是 1，否则会报错。

如果配置的 innodb_buffer_pool_size 大于 1G，[innodb_buffer_pool_instances 没有手动配置
则会自动配置为 8](https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_buffer_pool_instances)
。此时无法把 innodb_buffer_pool_size 动态调整到 1G 以下。

### 参考

+ https://cloud.tencent.com/document/product/236/32534
+ https://cloud.tencent.com/developer/article/1536046
+ https://help.aliyun.com/document_detail/35264.html?spm=a2c4g.53629.0.0.30392049WYKPId