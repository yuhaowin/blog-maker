<!DOCTYPE html>

<html lang="zh-cn">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="shortcut icon" href="/favicon.ico" />
    <link rel="stylesheet" href="https://cdn.staticfile.org/lxgw-wenkai-screen-webfont/1.6.0/style.css" />
    <title>yuhaowin</title>
    {{template "_style"}}
    {{template "_script"}}
</head>

<body>
    {{template "_nav"}}
    <article class="posts">
        {{range .}}
        <div class="posts-item">
            <a href="{{.Link}}">{{.Title}}</a>
            <small>{{.CreateDateStr}}</small>
        </div>
        {{end}}
    </article>
</body>
</html>
