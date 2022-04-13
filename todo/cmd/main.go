import {
	"fmt"
	"os"
	"strings"
}

const todoFileName = ".todo.json"

func main {
	l := &todo.List{}
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
