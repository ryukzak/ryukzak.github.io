// Static site generator. `go run .` builds public/, `go run . serve` previews it.
package main

import (
	"bytes"
	"fmt"
	"html"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	gmhtml "github.com/yuin/goldmark/renderer/html"
)

const out = "public"

var (
	frontMatter = regexp.MustCompile(`(?s)\A---\n(.*?)\n---\n`)
	markdown    = goldmark.New(
		goldmark.WithExtensions(extension.GFM, extension.Footnote, extension.Typographer),
		goldmark.WithParserOptions(parser.WithAutoHeadingID()),
		goldmark.WithRendererOptions(gmhtml.WithUnsafe()),
	)
	monthsRu = []string{"января", "февраля", "марта", "апреля", "мая", "июня",
		"июля", "августа", "сентября", "октября", "ноября", "декабря"}
)

// Page is one language version of a source file. Versions of the same
// document share key, which is also the URL of the original.
type Page struct {
	Layout, Title, Description, Lang, URL string
	Date                                  time.Time
	Content                               template.HTML
	Alternates                            []*Page
	Autotranslated                        bool

	key        string
	collection string
	translated bool
	source     []byte
}

// DateText is the publication date in the language of the page.
func (p *Page) DateText() string {
	if p.Lang == "ru" {
		return fmt.Sprintf("%d %s %d", p.Date.Day(), monthsRu[p.Date.Month()-1], p.Date.Year())
	}
	return p.Date.Format("2 January 2006")
}

type langLink struct {
	Label, Lang, URL string
	Current, Missing bool
}

// LangLinks feeds the language toggle. A language without a version of this
// page links to its home page; a machine translation is labelled "(auto)".
func (p *Page) LangLinks() []langLink {
	var links []langLink
	for _, lang := range []string{"en", "ru"} {
		l := langLink{Label: strings.ToUpper(lang), Lang: lang, URL: "/", Missing: true}
		if lang == "ru" {
			l.URL = "/ru/"
		}
		for _, alt := range append([]*Page{p}, p.Alternates...) {
			if alt.Lang == lang {
				l.URL, l.Missing = alt.URL, false
				if alt.Autotranslated {
					l.Label += " (auto)"
				}
			}
		}
		l.Current = lang == p.Lang
		links = append(links, l)
	}
	return links
}

func load(path, collection string) (*Page, error) {
	src, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	m := frontMatter.FindSubmatch(src)
	if m == nil {
		return nil, fmt.Errorf("%s: no front matter", path)
	}
	meta := map[string]string{}
	for _, line := range strings.Split(string(m[1]), "\n") {
		k, v, _ := strings.Cut(line, ":")
		meta[k] = strings.Trim(strings.TrimSpace(v), `"`)
	}

	name := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	lang := filepath.Ext(name)
	if lang != ".en" && lang != ".ru" {
		return nil, fmt.Errorf("%s: name must end with .en.md or .ru.md", path)
	}
	name = strings.TrimSuffix(name, lang)
	// Any auto_translated value marks a translation; true also shows a notice.
	auto, translated := meta["auto_translated"]
	p := &Page{
		Layout: meta["layout"], Title: meta["title"], Description: meta["description"],
		Lang: lang[1:], collection: collection, translated: translated,
		Autotranslated: auto == "true", source: src[len(m[0]):],
	}

	switch collection {
	case "home":
		p.key = "/"
	case "posts":
		if p.Date, err = time.Parse("2006-01-02", name[:10]); err != nil {
			return nil, fmt.Errorf("%s: %w", path, err)
		}
		p.key = p.Date.Format("/2006/01/02/") + name[11:] + ".html"
	case "projects":
		p.key = "/projects/" + name + "/"
	case "courses":
		p.key = "/courses/" + name + "/"
	}
	p.URL = p.key
	if p.translated {
		p.URL = "/" + p.Lang + p.key
	}
	return p, nil
}

// headingIDs keeps non-ASCII letters, so Cyrillic anchors stay as they were
// under kramdown.
type headingIDs map[string]int

func (ids headingIDs) Generate(v []byte, _ ast.NodeKind) []byte {
	var b strings.Builder
	for _, r := range strings.ToLower(string(v)) {
		switch {
		case unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' || r == '_':
			b.WriteRune(r)
		case r == ' ':
			b.WriteRune('-')
		}
	}
	id := b.String()
	if n := ids[id]; n > 0 {
		ids[id]++
		id = fmt.Sprintf("%s-%d", id, n)
	}
	ids[id]++
	return []byte(id)
}

func (ids headingIDs) Put(v []byte) { ids[string(v)]++ }

func render(src []byte) (template.HTML, error) {
	var b bytes.Buffer
	ctx := parser.NewContext(parser.WithIDs(headingIDs{}))
	err := markdown.Convert(src, &b, parser.WithContext(ctx))
	return template.HTML(b.String()), err
}

// listing renders one home page section in the given language, falling back
// to the original when a page has no translation.
func listing(pages []*Page, collection, lang string, small func(*Page) string) string {
	var items []*Page
	for _, p := range pages {
		if p.collection != collection || p.translated {
			continue
		}
		item := p
		for _, alt := range p.Alternates {
			if alt.Lang == lang {
				item = alt
			}
		}
		items = append(items, item)
	}
	slices.SortFunc(items, func(a, b *Page) int {
		if collection == "posts" {
			return b.Date.Compare(a.Date)
		}
		return strings.Compare(a.Title, b.Title)
	})

	var b strings.Builder
	b.WriteString("<ul>\n")
	for _, p := range items {
		fmt.Fprintf(&b, "<li><a href=\"%s\">%s<br/><small>%s</small></a></li>\n",
			p.URL, html.EscapeString(p.Title), small(p))
	}
	b.WriteString("</ul>")
	return b.String()
}

func loadAll() ([]*Page, error) {
	var pages []*Page
	for _, collection := range []string{"home", "posts", "projects", "courses"} {
		pattern := filepath.Join("content", collection, "*.md")
		if collection == "home" {
			pattern = filepath.Join("content", "index.*.md")
		}
		files, err := filepath.Glob(pattern)
		if err != nil {
			return nil, err
		}
		for _, f := range files {
			p, err := load(f, collection)
			if err != nil {
				return nil, err
			}
			pages = append(pages, p)
		}
	}

	originals := map[string]int{}
	for _, p := range pages {
		if !p.translated {
			originals[p.key]++
		}
		for _, other := range pages {
			if other != p && other.key == p.key {
				p.Alternates = append(p.Alternates, other)
			}
		}
	}
	for _, p := range pages {
		if originals[p.key] != 1 {
			return nil, fmt.Errorf("%s: needs exactly one version without `auto_translated`, found %d",
				p.key, originals[p.key])
		}
	}
	return pages, nil
}

func renderPages() error {
	tmpl, err := template.ParseFiles("layout.html")
	if err != nil {
		return err
	}
	pages, err := loadAll()
	if err != nil {
		return err
	}

	description := func(p *Page) string { return html.EscapeString(p.Description) }
	date := func(p *Page) string { return "<tt>// " + p.DateText() + "</tt>" }
	for _, p := range pages {
		src := p.source
		if p.collection == "home" {
			src = []byte(strings.NewReplacer(
				"{{projects}}", listing(pages, "projects", p.Lang, description),
				"{{courses}}", listing(pages, "courses", p.Lang, description),
				"{{posts}}", listing(pages, "posts", p.Lang, date),
			).Replace(string(src)))
		}
		if p.Content, err = render(src); err != nil {
			return fmt.Errorf("%s: %w", p.URL, err)
		}

		var b bytes.Buffer
		if err := tmpl.Execute(&b, p); err != nil {
			return fmt.Errorf("%s: %w", p.URL, err)
		}
		file := filepath.Join(out, urlToFile(p.URL))
		if err := os.MkdirAll(filepath.Dir(file), 0o755); err != nil {
			return err
		}
		if err := os.WriteFile(file, b.Bytes(), 0o644); err != nil {
			return err
		}
	}
	return nil
}

func urlToFile(url string) string {
	if strings.HasSuffix(url, "/") {
		url += "index.html"
	}
	return filepath.FromSlash(strings.TrimPrefix(url, "/"))
}

// checkURLs fails the build if a URL from urls.txt is not generated: those
// URLs are public and must keep working.
func checkURLs() error {
	list, err := os.ReadFile("urls.txt")
	if err != nil {
		return err
	}
	var missing []string
	for _, url := range strings.Split(strings.TrimSpace(string(list)), "\n") {
		if _, err := os.Stat(filepath.Join(out, urlToFile(url))); err != nil {
			missing = append(missing, url)
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf("missing URLs:\n%s", strings.Join(missing, "\n"))
	}
	return nil
}

func build() error {
	if err := os.RemoveAll(out); err != nil {
		return err
	}
	if err := renderPages(); err != nil {
		return err
	}
	if err := os.CopyFS(filepath.Join(out, "assets"), os.DirFS("assets")); err != nil {
		return err
	}
	return checkURLs()
}

// serve re-renders pages on every request and disables browser caching, so a
// plain refresh shows edits to content, layout and CSS.
func serve(addr string) error {
	var mu sync.Mutex
	site := http.FileServer(http.Dir(out))
	http.Handle("/assets/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		if err := renderPages(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		site.ServeHTTP(w, r)
	})
	log.Printf("http://%s/", addr)
	return http.ListenAndServe(addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		http.DefaultServeMux.ServeHTTP(w, r)
	}))
}

func main() {
	var err error
	if len(os.Args) > 1 && os.Args[1] == "serve" {
		err = serve("localhost:4000")
	} else {
		err = build()
	}
	if err != nil {
		log.Fatal(err)
	}
}
