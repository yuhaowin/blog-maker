# Bob 单词本插件

[Bob](https://github.com/ripperhe/Bob) 是一款 Mac 平台上的聚合翻译软件，支持用户编写插件完成特定功能，本插件主要是将输入的英文单词通过接口的方式同步到 `有道词典` 、`欧路词典` 等单词本中，以方便对查询过的单词进行集中复习。

## 特性

1、`有道词典` 仅实现了通过 cookie 的方式添加单词到单词本；无法使用账号登录，[因为 Bob 提供的 API 无法获取到请求返回的
cookie](https://github.com/ripperhe/Bob/issues/115)。

2、`欧路词典` 是通过开放 api 添加到指导单词本，但是需要指定单词本的 id，可以通过 api 获取单词本 id。

3、`扇贝单词` 是通过 api 添加到指导单词本，需要登录后从网页获取 auth_token

## 设置

![233336](https://image.yuhaowin.com/2023/05/04/233336.png)

## 有道词典获取 cookie

1、[登录有道词典](https://dict.youdao.com/)

2、如下图操作：

![233401](https://image.yuhaowin.com/2023/05/04/233401.png)

3、在控制台执行 `document.cookie` 获取到 cookie

## 欧路词典获取 token

1、[登录欧路词典](https://dict.eudic.net/)

2、在右上角「账户管理」获取授权

![233416](https://image.yuhaowin.com/2023/05/04/233416.png)

3、获取单词本ID

![233437](https://image.yuhaowin.com/2023/05/04/233437.png)

## 扇贝单词获取 token
1. [登录扇贝](https://www.shanbay.com/)
2. 在浏览器中按F12
3. 随便点击一下一面，查看 Header 中的 auth_token

## 效果

![233449](https://image.yuhaowin.com/2023/05/04/233449.png)


[Bob 单词本插件 Github 地址](https://github.com/yuhaowin/youdao-wordbook-bob-plugin)