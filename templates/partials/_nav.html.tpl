{{define "_nav"}}
<header class="site-header">
    <div class="nav-inner">
        <a class="site-name" href="/">yuhaowin</a>

        <!-- Desktop nav -->
        <nav class="nav-links" aria-label="主导航">
            <div class="nav-dropdown">
                <button class="nav-dropdown-toggle" aria-haspopup="true">
                    blogs <span class="nav-arrow"></span>
                </button>
                <div class="nav-dropdown-menu">
                    <a href="/">all posts</a>
                    {{range .Years}}
                    <a href="/{{.}}/">{{.}}</a>
                    {{end}}
                </div>
            </div>
            <div class="nav-dropdown">
                <button class="nav-dropdown-toggle" aria-haspopup="true">
                    videos <span class="nav-arrow"></span>
                </button>
                <div class="nav-dropdown-menu">
                    <a href="/videos/">all videos</a>
                    {{range .VideoYears}}
                    <a href="/videos/{{.}}/">{{.}}</a>
                    {{end}}
                </div>
            </div>
        </nav>

        <!-- Hamburger (mobile) -->
        <button class="nav-hamburger" id="nav-hamburger" aria-label="展开菜单" aria-expanded="false">
            <span></span>
            <span></span>
            <span></span>
        </button>
    </div>

    <!-- Mobile menu -->
    <div class="mobile-menu" id="mobile-menu" aria-hidden="true">
        <span class="mobile-menu-section">Blogs</span>
        <a href="/">all posts</a>
        {{range .Years}}
        <a href="/{{.}}/">{{.}}</a>
        {{end}}
        <span class="mobile-menu-section">Videos</span>
        <a href="/videos/">all videos</a>
        {{range .VideoYears}}
        <a href="/videos/{{.}}/">{{.}}</a>
        {{end}}
    </div>
</header>
{{end}}
