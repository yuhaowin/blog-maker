# Changelog

## 2026-02-25

### Added
- Video navigation now supports "all videos" and per-year dropdowns, matching the blog navigation behavior
  - `/videos/` — aggregated index of all videos
  - `/videos/<year>/` — per-year video index pages
  - Nav bar videos entry changed from a plain link to a dropdown menu
- `VideoYears []string` added to `Post` and `IndexPageData` structs and propagated through all render/generate function signatures
- Both `post.html.tpl` and `index.html.tpl` now pass full page context to `_nav` partial (`{{template "_nav" .}}`)

## 2026-02-12

### Changed
- Navigation year links are now dynamic — extracted from the content folder structure at render time, no hardcoded years in templates
- `IndexPageData` struct introduced to wrap posts list together with years for index pages
- Content reorganized into year subdirectories (`content/<year>/`, `content/videos/<year>/`)
- Render logic updated to handle 4-part paths (`/videos/year/file`) in addition to 3-part paths (`/year/file`)
- Root index (`/`) aggregates all blog posts across all years, excluding videos

## 2026-02-XX

### Changed
- CI: build `blog-maker` binary during CI instead of committing it to the repo
- CI: pin `ssh-deploy` action to v4, use `with` instead of `env` for configuration