{{define "_style"}}
<style>
/* ── Reset ── */
*, *::before, *::after { box-sizing: border-box; }

/* ── Base ── */
body {
    margin: 0;
    padding: 0;
    background: #fff;
    color: #1a1a1a;
    font-family: "LXGW WenKai Screen", Georgia, "Noto Serif SC", serif;
    font-size: 17px;
    line-height: 1.85;
    -webkit-font-smoothing: antialiased;
}

/* ── Layout ── */
.site-header {
    position: sticky;
    top: 0;
    z-index: 100;
    background: rgba(255, 255, 255, 0.92);
    backdrop-filter: blur(8px);
    -webkit-backdrop-filter: blur(8px);
    border-bottom: 1px solid #ebebeb;
}

.nav-inner {
    max-width: 720px;
    margin: 0 auto;
    padding: 0 1.5rem;
    display: flex;
    align-items: center;
    justify-content: space-between;
    height: 56px;
}

.site-name {
    font-size: 1rem;
    font-weight: 700;
    color: #1a1a1a;
    text-decoration: none;
    letter-spacing: -0.01em;
}

.site-name:visited { color: #1a1a1a; }

main {
    max-width: 720px;
    margin: 0 auto;
    padding: 2.5rem 1.5rem 5rem;
}

/* ── Desktop Nav ── */
.nav-links {
    display: flex;
    align-items: center;
    gap: 0.25rem;
}

.nav-dropdown {
    position: relative;
}

.nav-dropdown-toggle {
    display: flex;
    align-items: center;
    gap: 5px;
    padding: 6px 10px;
    border-radius: 6px;
    background: none;
    border: none;
    cursor: pointer;
    font-family: inherit;
    font-size: 0.9rem;
    color: #555;
    transition: background 0.15s, color 0.15s;
}

.nav-dropdown-toggle:hover {
    background: #f5f5f5;
    color: #1a1a1a;
}

.nav-arrow {
    display: inline-block;
    width: 0;
    height: 0;
    border-left: 4px solid transparent;
    border-right: 4px solid transparent;
    border-top: 5px solid currentColor;
    margin-top: 2px;
    transition: transform 0.2s;
}

.nav-dropdown:hover .nav-arrow {
    transform: rotate(180deg);
}

.nav-dropdown-menu {
    display: none;
    position: absolute;
    right: 0;
    top: calc(100% + 6px);
    background: #fff;
    border: 1px solid #e5e5e5;
    border-radius: 8px;
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);
    min-width: 120px;
    padding: 5px 0;
    z-index: 200;
}

.nav-dropdown:hover .nav-dropdown-menu {
    display: block;
}

.nav-dropdown-menu a {
    display: block;
    padding: 8px 16px;
    color: #444;
    text-decoration: none;
    font-size: 0.875rem;
    white-space: nowrap;
}

.nav-dropdown-menu a:hover {
    background: #f5f5f5;
    color: #1a1a1a;
}

.nav-dropdown-menu a:visited {
    color: #444;
}

/* ── Hamburger ── */
.nav-hamburger {
    display: none;
    flex-direction: column;
    justify-content: center;
    gap: 5px;
    background: none;
    border: none;
    cursor: pointer;
    padding: 6px;
    border-radius: 6px;
}

.nav-hamburger:hover { background: #f5f5f5; }

.nav-hamburger span {
    display: block;
    width: 20px;
    height: 2px;
    background: #1a1a1a;
    border-radius: 2px;
    transition: all 0.25s;
}

.nav-hamburger.open span:nth-child(1) {
    transform: translateY(7px) rotate(45deg);
}

.nav-hamburger.open span:nth-child(2) {
    opacity: 0;
}

.nav-hamburger.open span:nth-child(3) {
    transform: translateY(-7px) rotate(-45deg);
}

/* ── Mobile menu ── */
.mobile-menu {
    display: none;
    flex-direction: column;
    border-bottom: 1px solid #ebebeb;
    background: #fff;
    padding: 0.5rem 0;
}

.mobile-menu.open {
    display: flex;
}

.mobile-menu-section {
    padding: 4px 1.5rem 2px;
    font-size: 0.72rem;
    font-weight: 600;
    text-transform: uppercase;
    letter-spacing: 0.08em;
    color: #aaa;
    margin-top: 8px;
}

.mobile-menu a {
    padding: 10px 1.5rem;
    color: #333;
    text-decoration: none;
    font-size: 0.95rem;
}

.mobile-menu a:hover { background: #f9f9f9; }
.mobile-menu a:visited { color: #333; }

/* ── Article typography ── */
article h1 {
    font-size: 1.75rem;
    font-weight: 700;
    line-height: 1.3;
    margin: 0 0 1.5rem;
    letter-spacing: -0.02em;
}

article h2 {
    font-size: 1.3rem;
    font-weight: 700;
    line-height: 1.35;
    margin: 2.25em 0 0.6em;
    letter-spacing: -0.01em;
}

article h3 {
    font-size: 1.1rem;
    font-weight: 700;
    margin: 1.75em 0 0.5em;
}

article h4 {
    font-size: 1rem;
    font-weight: 700;
    margin: 1.5em 0 0.4em;
}

article p {
    margin: 0 0 1.2em;
}

article ul, article ol {
    margin: 0 0 1.2em;
    padding-left: 1.5em;
}

article li {
    margin-bottom: 0.3em;
}

article img {
    max-width: 100%;
    border-radius: 8px;
    display: block;
    margin: 1.5em auto;
}

article a {
    color: #0066cc;
    text-decoration: underline;
    text-underline-offset: 2px;
}

article a:hover { color: #004499; }
article a:visited { color: #5a4fcf; }

em { color: #666; font-style: italic; }

strong { font-weight: 700; }

/* ── Blockquote ── */
blockquote {
    border-left: 3px solid #ddd;
    margin: 1.5em 0;
    padding: 0.5em 1.25em;
    color: #666;
    background: #fafafa;
    border-radius: 0 6px 6px 0;
}

blockquote p {
    display: inline;
    margin: 0;
}

/* ── Code ── */
code {
    font-family: "Fira Code", "Cascadia Code", ui-monospace, "SFMono-Regular", Menlo, monospace;
    font-size: 0.85em;
    background: #f3f3f3;
    padding: 0.15em 0.45em;
    border-radius: 4px;
}

pre {
    background: #f8f8f8;
    border: 1px solid #e8e8e8;
    border-radius: 8px;
    padding: 1.1rem 1.25rem;
    overflow-x: auto;
    line-height: 1.6;
    margin: 1.25em 0;
}

pre code {
    background: none;
    padding: 0;
    font-size: 0.85rem;
    border-radius: 0;
}

/* chroma (server-side syntax highlighting) container */
.highlight {
    margin: 1.25em 0;
    border-radius: 8px;
    overflow: hidden;
    border: 1px solid #e8e8e8;
}

.highlight pre {
    margin: 0;
    border: none;
    border-radius: 0;
}

/* ── Table ── */
table {
    width: 100%;
    border-collapse: collapse;
    margin: 1.5em 0;
    font-size: 0.9em;
    display: block;
    overflow-x: auto;
}

table th, table td {
    border: 1px solid #e5e5e5;
    padding: 8px 14px;
    text-align: left;
    color: inherit;
}

table thead th {
    background: #f5f5f5;
    font-weight: 700;
}

table tr:nth-child(even) { background: #fafafa; }

/* ── HR ── */
.hr-middle-text {
    line-height: 1em;
    position: relative;
    outline: 0;
    border: 0;
    color: #666;
    text-align: center;
    height: 1.5em;
    opacity: 0.6;
    margin: 2.5rem 0;
}

.hr-middle-text::before {
    content: '';
    background: linear-gradient(to right, transparent, #ccc, transparent);
    position: absolute;
    left: 0;
    top: 50%;
    width: 100%;
    height: 1px;
}

.hr-middle-text::after {
    content: attr(data-content);
    position: relative;
    display: inline-block;
    padding: 0 0.75em;
    line-height: 1.5em;
    background: #fff;
    font-size: 0.8rem;
    letter-spacing: 0.1em;
    text-transform: uppercase;
}

/* ── Post list (index) ── */
.posts-list {
    margin-top: 0.5rem;
}

.posts-item {
    display: flex;
    justify-content: space-between;
    align-items: baseline;
    padding: 0.65rem 0;
    border-bottom: 1px solid #f0f0f0;
    gap: 1rem;
}

.posts-item:last-child {
    border-bottom: none;
}

.posts-item a {
    color: #1a1a1a;
    text-decoration: none;
    font-size: 0.95rem;
    flex: 1;
    min-width: 0;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
}

.posts-item a:hover { color: #0066cc; }
.posts-item a:visited { color: #555; }

.posts-item-date {
    color: #aaa;
    font-size: 0.8rem;
    white-space: nowrap;
    font-variant-numeric: tabular-nums;
}

/* ── Comments ── */
.comments {
    margin-top: 3.5rem;
}

.utterances {
    max-width: 100% !important;
}

/* ── Responsive ── */
@media (max-width: 680px) {
    body { font-size: 16px; }

    .nav-links { display: none; }
    .nav-hamburger { display: flex; }

    main { padding: 1.75rem 1.25rem 4rem; }

    article h1 { font-size: 1.45rem; }
    article h2 { font-size: 1.2rem; }

    pre { padding: 0.9rem 1rem; font-size: 0.8rem; }

    .posts-item a {
        white-space: normal;
    }
}
</style>
{{end}}
