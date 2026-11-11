# ryukzak.github.io

Personal site, built by a small Go generator (`main.go`, `layout.html`).

```console
make serve   # preview at http://localhost:4000, refresh to see edits
make build   # build public/ and check that every URL in urls.txt exists
make lint    # gofmt and go vet
make clean   # remove public/
```

## Content

- `content/index.LANG.md`, `content/posts/YYYY-MM-DD-slug.LANG.md`, `content/projects/slug.LANG.md`, `content/courses/slug.LANG.md`, where `LANG` is `en` or `ru`.
- Files with the same name except `LANG` are versions of one page.
- The original version is published without a prefix, e.g. `/projects/nitta/`. A translation has `auto_translated` in its front matter and is published under `/en/` or `/ru/`, e.g. `/ru/projects/nitta/`. Use `auto_translated: true` for a machine translation (the page shows a notice) and `auto_translated: false` for a human one.

Add new public URLs to `urls.txt` if they must never disappear.
