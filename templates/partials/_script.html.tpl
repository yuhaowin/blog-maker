
{{define "_script"}}
<!-- Google Analytics -->
<script>
    // Your Google Analytics code can be pasted here
</script>
<!-- End Google Analytics -->

<script>
document.addEventListener('DOMContentLoaded', function () {
    var btn      = document.getElementById('nav-hamburger');
    var menu     = document.getElementById('mobile-menu');
    var closeBtn = document.getElementById('mobile-menu-close');
    if (!btn || !menu) return;

    function openMenu() {
        menu.classList.add('open');
        btn.setAttribute('aria-expanded', 'true');
        menu.setAttribute('aria-hidden', 'false');
        document.body.classList.add('menu-open');
    }

    function closeMenu() {
        menu.classList.remove('open');
        btn.setAttribute('aria-expanded', 'false');
        menu.setAttribute('aria-hidden', 'true');
        document.body.classList.remove('menu-open');
    }

    btn.addEventListener('click', openMenu);
    if (closeBtn) closeBtn.addEventListener('click', closeMenu);

    document.addEventListener('keydown', function (e) {
        if (e.key === 'Escape') closeMenu();
    });

    menu.addEventListener('click', function (e) {
        if (e.target.tagName === 'A') closeMenu();
    });
});
</script>
{{end}}
