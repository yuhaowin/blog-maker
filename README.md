## Install

Compiling and building binary executable file from source code:

```bash
go build -o blog-maker main.go
```

### Showing your website on local

```bash
./blog-maker s
```

### Render all markdown files in content folder to path

```bash
./blog-maker -o ./blog
```

### Classify your articles

If you want to classify your articles into different groups, you should create a subfolder into `content`, and then add
the folder link into `templates/partials/_nav.html.tpl`. 

For example, if you have written an article about `videos`, and want to create a `videos` group in your website.

```bash
mkdir content/videos
mv your_md_file content/videos/
```

Add `videos` into `templates/partials/_nav.html.tpl`

```html
[ <a class="nav-btn" href="/videos/">videos</a> ]
```
