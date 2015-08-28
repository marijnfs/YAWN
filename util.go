package main

import (
	"log"
	"github.com/knieriem/markdown"
	"bytes"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func MDToHTML(data []byte) []byte {
	p := markdown.NewParser(nil)

	var buf bytes.Buffer
	p.Markdown(bytes.NewReader(data), markdown.ToHTML(&buf))
	return buf.Bytes()
}
