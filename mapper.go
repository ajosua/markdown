package main

// https://github.com/gomarkdown/markdown

import (
	"fmt"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

var mds = `
# header

Sample text.

[link](http://example.com)
`

func printMdToHTML() {
	md := []byte(mds)
	html := mdToHTML(md)
	fmt.Printf("--- Markdown:\n%s\n\n--- HTML:\n%s\n", md, html)
}
