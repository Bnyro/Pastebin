package utilities

import (
	"strings"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

func Format(text string) string {
	// lexer := lexers.Analyse(text)
	lexer := lexers.Get("python")
	if lexer == nil {
		lexer = lexers.Fallback
	}
	style := styles.Get("monokai")
	if style == nil {
	  	style = styles.Fallback
	}
	formatter := html.New(html.Standalone(true))
	iterator, err := lexer.Tokenise(nil, text)
	writer := new(strings.Builder)
	err = formatter.Format(writer, style, iterator)
	if err != nil {
		return text
	}
	return writer.String()
}