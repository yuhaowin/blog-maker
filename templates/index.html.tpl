<!DOCTYPE html>

<html lang="zh-cn">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="icon" href="/favicon.svg" type="image/svg+xml" />
    <link rel="stylesheet" href="https://cdn.staticfile.org/lxgw-wenkai-screen-webfont/1.6.0/style.css" />
    <title>yuhaowin</title>
    {{template "_style"}}
    {{template "_script"}}
</head>

<body>
    {{template "_nav" .}}
    <main>
        <article class="posts-list">
            {{range .Posts}}
            <div class="posts-item">
                <a href="{{.Link}}">{{.Title}}</a>
                <span class="posts-item-date">{{.CreateDateStr}}</span>
            </div>
            {{end}}
        </article>
    </main>
</body>
</html>
