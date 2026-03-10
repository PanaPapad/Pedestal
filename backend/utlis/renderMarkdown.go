package utils

import (
	"bytes"
	"html/template"

	"github.com/yuin/goldmark"
)

func RenderMarkdown(md string) template.HTML {
	var buf bytes.Buffer
	goldmark.Convert([]byte(md), &buf)
	return template.HTML(buf.String())
}
