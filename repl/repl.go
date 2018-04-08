// repl/repl.go
package repl

import (
	"bufio"
	"fmt"
	"github.com/GiacomoTravaglini/Golang-Interpreter/lexer"
	"github.com/GiacomoTravaglini/Golang-Interpreter/token"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		// Scanner.Scan scans line by line (split fn)
		// it is possible to change the split function and
		// provide a custom one
		scanned := scanner.Scan()
		// Scanned is false if text has finished or if an error
		// occurred
		if !scanned {
			return
		}

		// Get the scanned line in text format
		line := scanner.Text()

		// Instantiate the lexer
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}

	}
}
