package server

import (
	"bytes"

	mathjax "github.com/litao91/goldmark-mathjax"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func (s *server) parseMarkdown(f string) (string, map[string]interface{}, error) {
	var (
		buf bytes.Buffer
		md  goldmark.Markdown
		err error
	)

	// prepare f to be read
	src, err := s.fs.ReadFile("content/" + f + ".md")
  if err != nil {
    return "", nil, err
  }

	// Prepend parser context to parser options
	var opts []parser.ParseOption
	ctx := parser.NewContext()
	opts = append(opts, parser.WithContext(ctx))

	md = goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
			mathjax.MathJax,
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)

	err = md.Convert([]byte(src), &buf, opts...)
	if err != nil {
		return "", nil, err
	}

	yaml := meta.Get(ctx)
	html := buf.String()

	return html, yaml, nil
}
