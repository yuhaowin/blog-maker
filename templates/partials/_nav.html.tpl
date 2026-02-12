{{define "_nav"}}
<header>
    <nav class="navbar">
        <span class="navbar-item navbar-dropdown">
            <a class="nav-btn" href="#blogs">blogs <span class="arrow"></span></a>
            <div class="navbar-menu">
                <a class="navbar-menu-item" href="/">all posts</a>
                {{range .}}
                <a class="navbar-menu-item" href="/{{.}}/">{{.}}</a>
                {{end}}
            </div>
        </span>
        <span class="navbar-item"><a class="nav-btn" href="/videos/">videos</a></span>
    </nav>
</header>
{{end}}
