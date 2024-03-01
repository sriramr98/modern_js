package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/sriramr98/modern_js/lexer"
	"github.com/sriramr98/modern_js/token"
)

func Start(in io.Reader, out io.Writer) {
    scanner := bufio.NewScanner(in)
    for {
        fmt.Fprintf(out, ">> ")
        scanned := scanner.Scan()
        if !scanned {
            return
        }

        line := scanner.Text()
        l := lexer.New(line)
        for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
            fmt.Fprintf(out, "Type: %s, Literal: %s\n",tok.Type, tok.Literal)
        }
    }
}
