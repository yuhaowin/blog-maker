
{{define "_script"}}
<!-- Google Analytics -->
<script>
    // Your Google Analytics code can be pasted here
</script>
<!-- End Google Analytics -->

<script>
document.addEventListener('DOMContentLoaded', function () {
    var btn = document.getElementById('nav-hamburger');
    var menu = document.getElementById('mobile-menu');
    if (!btn || !menu) return;
    btn.addEventListener('click', function () {
        var open = menu.classList.toggle('open');
        btn.classList.toggle('open', open);
        btn.setAttribute('aria-expanded', open ? 'true' : 'false');
        menu.setAttribute('aria-hidden', open ? 'false' : 'true');
    });
});
</script>
{{end}}
