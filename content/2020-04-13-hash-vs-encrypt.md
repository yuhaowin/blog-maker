# HASH & ENCRYPT
Hash 算法 和 Encrypt 算法的区别和对应的使用场景。

### 哈希（Hash）与加密（Encrypt）的区别

> Hash 是将目标文本转换成具有相同长度的、不可逆的杂乱字符串，这个杂乱字符串也叫做消息摘要；Encrypt 是将目标文本转换为具有不同长度的、可逆的密文。

1、任何目标文本经过 Hash 算法后得到的都是相同长度的消息摘要，而 Encrypt 得到的密文的长度是和明文的长度有关。

2、Hash 是不可逆的，Encrypt 是可逆的。

> 一个哈希算法 $R = H(S)$ 是一个多对一映射，给定目标文本 S，H 可以将其唯一映射为 R，并且对于所有 S，R 具有相同的长度。由于是多对一映射，所以H不存在逆映射 $S = H^{-1}(R)$ 使得 R 转换为唯一的 S。

> 一个加密算法 $R=E(S,K_E)$ 是一个一一映射，其中第二个参数叫做加密密钥，E 可以将给定的明文 S 结合加密密钥 Ke 唯一映射为密文 R，并且存在另一个一一映射 $S=D(R,K_D)$，可以结合 Kd 将密文 R唯一映射为对应明文 S，其中 Kd 叫做解密密钥。

### Hash 算法

#### Hash 算法的特点：

+ 正向快速：给定明文和 hash 算法，在有限时间和有限资源内能计算出 hash 值。
+ 逆向困难：给定（若干） hash 值，在有限时间内很难（基本不可能）逆推出明文。
+ 输入敏感：原始输入信息修改一点信息，产生的 hash 值看起来应该都有很大不一样。
+ 冲突避免：很难找到两段内容不一样的明文，使得它们的 hash 值一致（发生冲突）。即对于任意两个不一样的数据块，其hash值相同的可能性极小；对于一个给定的数据块，找到和它hash值相同的数据块极为困难。

#### Hash 算法的主要应用

> Hash 算法主要应用于消息摘要和签名，主要是对消息的完整性进行校验。

#### 常用的 Hash 算法的实现

+ MD4(RFC 1320)， MD = Message Digest 其输出为 128 位。MD4 已证实不够安全。
+ MD5(RFC 1321)其输出是 128 位。MD5 比 MD4 复杂，而且计算速度要慢一点，更安全一些。MD5 已被证实不具有「强抗碰撞性」。
+ SHA(Secure Hash Algorithm)是一个 Hash 函数族， SHA-1 在 1995 年面世，它的输出为长度 160 位的 hash 值，所以抗穷举性更好。SHA-1 设计时基于和 MD4 相同原理，而且模仿了该算法。SHA-1 已被证实不具”强抗碰撞性”。
+ SHA-224、SHA-256、SHA-384，和 SHA-512 算法（统称为 SHA-2），跟 SHA-1 算法原理相似。SHA-3 相关算法也已被提出。

### 对称加密算法

> 对称加密是指，在 **加密** 和 **解密** 的过程中使用的密钥是相同的。

#### 常用的对称加密算法

+ AES（AES128、AES192、AES256）默认安装的 `JDK` 尚不支持 `AES256`，需要安装对应的 `jce` 补丁进行升级 `jce1.7`，`jce1.8`。
+ DES - 加密算法是一种 **分组密码**，以 `64` 位为 **分组对数据** 加密，它的 **密钥长度** 是 `56` 位，**加密解密** 用 **同一算法**。
+ 3DES - 是基于 `DES` 的 **对称算法**，对 **一块数据** 用 **三个不同的密钥** 进行 **三次加密**，**强度更高**。

### 非对称加密算法

> 非对称加密是指，在 **加密** 和 **解密** 的过程中是使用两个不同的密钥，**加密** 过程使用的是 **共钥** ，**解密** 的过程中使用的是 **私钥** 。

#### 常用的非对称加密算法

+ RSA 算法
  >`RSA` 加密算法是目前最有影响力的 **公钥加密算法**，并且被普遍认为是目前 **最优秀的公钥方案** 之一。`RSA` 是第一个能同时用于 **加密** 和 **数字签名** 的算法，它能够 **抵抗** 到目前为止已知的 **所有密码攻击**，已被 `ISO` 推荐为公钥数据加密标准。

+ ECC 算法
  > `ECC` 也是一种 **非对称加密算法**，主要优势是在某些情况下，它比其他的方法使用 **更小的密钥**，比如 `RSA` **加密算法**，提供 **相当的或更高等级** 的安全级别。不过一个缺点是 **加密和解密操作** 的实现比其他机制 **时间长** (相比 `RSA` 算法，该算法对 `CPU` 消耗严重)。

###  HMAC算法 ？？？？

[参考1](https://www.shangmayuan.com/a/047a517e76154c099213091d.html)
[参考2](https://juejin.im/post/5b48b0d7e51d4519962ea383)
