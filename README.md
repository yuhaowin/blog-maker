## Install

Compiling and building binary executable file from source code:

```bash
go build -o blog-maker .
```

## Usage

### Preview locally

```bash
./blog-maker s
```

Opens a local server at `http://localhost:8080/`.

### Generate static files

```bash
./blog-maker -o ./blog
```

Renders all markdown files in `content/` to the specified output path (default: `./public`).

## Content Structure

Organize content by year under `content/`:

```
content/
├── 2024/
│   ├── 2024-01-10-my-post.md
│   └── 2024-03-22-another-post.md
├── 2023/
│   └── 2023-11-05-some-post.md
└── videos/
    ├── 2024/
    │   └── 2024-02-01-video-01.md
    └── 2023/
        └── 2023-08-15-video-02.md
```

### Blogs

- Place markdown files under `content/<year>/`
- The root index (`/`) shows all posts across all years
- Each year gets its own index page at `/<year>/`
- Navigation dropdown is generated dynamically from the year folders

### Videos

- Place markdown files under `content/videos/<year>/`
- `/videos/` shows all videos across all years
- Each year gets its own index at `/videos/<year>/`
- Navigation dropdown is generated dynamically from the year folders

### Filename convention

Use `YYYY-MM-DD-title.md` as the filename — the date is parsed from the filename and used for sorting.

## Navigation

The navbar is generated automatically based on the content structure. No manual edits to templates are needed when adding new year folders.