# 在 mac 下构建自己的输入法

mac 是支持构建自己的输入法的，核心是编写一个 .inputplugin 的文件。

要求这个文件是 UTF-16 的编码格式。文件的具体格式见[官网示例](https://support.apple.com/zh-cn/guide/mac-help/mchlp2866/13.0/mac/13.0)

我这里编写的是一个输入英语音标的输入法：

```shell
# https://support.apple.com/zh-cn/guide/mac-help/mchlp2866/mac
# D.J. 音标 - (IPA88)
METHOD: TABLE
ENCODE: SC
PROMPT: D.J. 音标
DELIMITER: ,
VERSION: 1.2
MAXINPUTCODE: 8
VALIDINPUTKEY: ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
TERMINPUTKEY: 123456789
BEGINCHARACTER
####### 测试内容
mac     💻
apple 	🍎

####### 短元音
a 	æ
e 	e,ə
i 	i,ɪ
o 	ɒ
u 	u,ʌ,ʊ

####### 长元音
a;	ɑː
e; 	ɜː
i; 	iː
o; 	ɔː
u; 	uː

####### 双元音
au 	aʊ
ai 	aɪ
ei 	eɪ
eu 	əʊ
ee 	eə
ie 	ɪə
oi 	ɔɪ
ue 	ʊə

####### 辅音，英音和美音没有区别，D.J.63、D.J.88 和 K.K. 音标一致
p 	p
b 	b
t 	t
d 	d
k 	k
g 	g
m 	m
n 	n,ŋ
f 	f
v 	v
s 	s
z 	z
th 	θ,ð
x 	ʃ,ʒ
tx 	tʃ
dx 	dʒ
h 	h
w 	w
r 	r
j 	j
l 	l

####### 辅音连缀，可以不认为是新的音标
ts 	ts
dz 	dz
tr 	tr
dr 	dr

####### 常见两音标组合 - 待完善
an 	ʌn,æn
am 	ʌm
ek 	ək
di 	dɪ
en 	ən
em 	əm
er 	ər
eu 	əu
fe 	fə
ga 	gæ
hi 	hɪ
ke 	kə
ik 	ɪk
in 	ɪn,ɪŋ
ir 	ɪr
kj 	kjʊ
la 	lʌ,læ
le 	lə,lɛ
li 	lɪ
me 	mə
mo 	mɒ
ne 	nə
ni 	nɪ
ou 	oʊ
pa 	pæ
ra 	ræ,rʌ
re 	rə
ri 	rɪ
sa 	sʌ
se 	sə
si 	sɪ
so 	sɒ
ta 	tæ,tʌ
te 	tə
ti 	tɪ
to 	tɒ
ue 	uə
ve 	və
xn 	ʃn
xl 	ʃl

####### 符号
/ 	/
, 	ʼ,ʻ,ː,ʽ
. 	.
ENDCHARACTER
```

编写完成后保存，双击改文件就会安装到内建的输入法中，该文件会保存在 `~/Library/Input Methods` 目录下。如果不需要可以直接删除。

![032015](https://image.yuhaowin.com/2022/12/25/032015.png)

安装或删除自定义的输入法后，可能需要重新登录系统，或重启系统。

效果如下：

![031855](https://image.yuhaowin.com/2022/12/25/031855.png)