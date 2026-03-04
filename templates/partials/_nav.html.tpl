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
            <a href="/feed.xml" class="nav-rss-link" aria-label="RSS Feed">
                RSS
            </a>
        </nav>

        <!-- Hamburger (mobile) -->
        <button class="nav-hamburger" id="nav-hamburger" aria-label="展开菜单" aria-expanded="false">
            <span></span>
            <span></span>
            <span></span>
        </button>
    </div>

</header>

<!-- 全屏菜单覆盖层 -->
<div class="mobile-menu" id="mobile-menu" aria-hidden="true">
    <div class="mobile-menu-header">
        <a class="site-name" href="/">yuhaowin</a>
        <button class="mobile-menu-close" id="mobile-menu-close" aria-label="关闭菜单">
            <span></span><span></span>
        </button>
    </div>
    <nav class="mobile-menu-body">
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
        <span class="mobile-menu-section">Subscribe</span>
        <a href="/feed.xml">RSS Feed</a>
    </nav>
</div>
{{end}}
