{{define "_style"}}
<style>
body {
    color: #000000;
    line-height: 1.6em;
    padding: 1em;
    box-sizing: border-box;
    margin: auto;
    max-width: 72em;
    background: #fefefe;
    font-family: "LXGW WenKai Screen", Fira Code, Monaco, Consolas, Ubuntu Mono, PingFang SC, Hiragino Sans GB, Microsoft YaHei, WenQuanYi Micro Hei, monospace, sans-serif;
}

blockquote {
    background: #f9f9f9;
    border-left: 10px solid #ccc;
    margin: 1.5em 10px;
    padding: 0.5em 10px;
    quotes: "\201C" "\201D" "\2018" "\2019";
}

blockquote p {
    display: inline;
}

em {
    color: grey;
}

pre {
    background-color: #eed;
}

code {
    background-color: #eed;
}

nav {
    margin-bottom: 20px;
}

small {
    font-size: smaller;
}

a {
    color: #007bff;
}

a:hover,
a:focus {
    color: #0056b3;
}

a:visited {
    color: grey;
}

article {
    padding-bottom: 1em;
}

img {
    max-width: 100%;
}

body {
    background-color: #fff;
    color: #212529;
}

.comments {
    text-align: center;
    margin-top: 20px;
    padding: 5px 10px;
    font-size: .8rem;
    text-transform: uppercase;
    text-decoration: none;
    letter-spacing: .1em;
    z-index: 1;
}

.utterances {
    max-width: 100%;
}

.hr-middle-text {
    line-height: 1em;
    position: relative;
    outline: 0;
    border: 0;
    color: black;
    text-align: center;
    height: 1.5em;
    opacity: .5;
}

.hr-middle-text:before {
    content: '';
    background: linear-gradient(to right, transparent, black, transparent);
    position: absolute;
    left: 0;
    top: 50%;
    width: 100%;
    height: 1px;
}

.hr-middle-text:after {
    content: attr(data-content);
    position: relative;
    display: inline-block;
    color: black;
    padding: 0 .5em;
    line-height: 1.5em;
    background-color: #fcfcfa;
}

strong {
    color: #212529;
    font-weight: 900;
}

table {
    display: flow-root;
    border-collapse: collapse;
    margin: 0 auto;
    text-align: center;
}

table td, table th {
    border: 1px solid #cad9ea;
    color: #666;
    height: 30px;
}

table thead th {
    background-color: #CCE8EB;
    width: 100px;
}

table tr:nth-child(odd) {
    background: #fff;
}

table tr:nth-child(even) {
    background: #F5FAFA;
}


/* _nav style */
.navbar-item::before {
    content: "[ ";
}

.navbar-item::after {
    content: " ]"
}

.navbar-dropdown {
    position: relative;
    display: inline-block;
}

.navbar-menu {
    display: none;
    position: absolute;
    background-color: #fefefe;
    min-width: 120px;
    box-shadow: 0px 8px 16px 0px rgba(0,0,0,0.2);
    z-index: 1;
    border: 1px solid #ccc;
    margin-top: 5px;
}

.navbar-dropdown:hover .navbar-menu {
    display: block;
}

.navbar-menu-item {
    color: #007bff;
    padding: 8px 16px;
    text-decoration: none;
    display: block;
}

.navbar-menu-item:hover {
    background-color: #f1f1f1;
    color: #0056b3;
}

@media screen and (min-width: 300px) and (max-width: 700px) {
    .navbar {
        width: 100%;
        display: flex;
        align-items: center;
        justify-content: space-between;
    }

    .navbar-item {
        flex: 1;
        text-align: center;
        display: flex;
        align-items: center;
        justify-content: center;
        overflow: hidden;
    }

    .nav-btn {
        flex: 1;
        white-space: nowrap;
        text-overflow: ellipsis;
        overflow: hidden;
        text-align: center;
    }

    .navbar-menu {
        left: 0;
        min-width: 100px;
    }
}

/* _nav style end */

/* index style */
.posts-item {
    display: flex;
    flex-direction: column;
    margin-bottom: 20px;
}

@media screen and (min-width: 300px) and (max-width: 700px) {
    .posts-item {
        flex-direction: row;
        align-items: center;
        justify-content: space-around;
        margin-bottom: 10px;
    }

    .posts-item > a {
        display: inline-block;
        width: 70%;
        white-space: nowrap;
        text-overflow: ellipsis;
        overflow: hidden;
    }

    .posts-item > small {
        flex: 1;
        text-align: right;
    }

    article {
        letter-spacing: -0.003em;
        line-height: 28px;
        font-size: 18px;
        word-wrap: break-word;
    }
}

/* index style end */
</style>
{{end}}
