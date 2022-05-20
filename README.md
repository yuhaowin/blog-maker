
## Install
Compiling from source code:  
```bash
go build -o blog-maker main.go
```

### Showing your site on local
```bash
./blog-maker s
```

### Classify your articles
If you want to classify your articles into different groups, you should create a subfolder into `content`, and then add the folder link into `templates/partials/_nav.html.tpl`. For example, if you have written an article about English learning, and want to create a English Learning group in your website.
```bash
mkdir content/el
mv your_md_file content/el/
```

Add `el` into `templates/partials/_nav.html.tpl`
```html
[ <a class="nav-btn" href="/el/">English</a> ]
```
