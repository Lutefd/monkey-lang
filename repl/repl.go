package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/lutefd/monkey-lang/lexer"
	"github.com/lutefd/monkey-lang/token"
)

const PROMPT = ">> "

func Start (in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		// Despite the warning using the prompt is actually formatting it so we need the Fprintf
		// You can confirm that by replacing the PROMPT var with ">> " in the Fprintf call
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned{
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}

}