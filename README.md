# cologroni

Negroni logger middleware with color. It also doesn't make a log line when response starts.

## Usage

```go
import "github.com/cag/cologroni"

func main() {

	...

	n := negroni.New(
		cologroni.New(),
		negroni.NewRecovery(),
		negroni.NewStatic(http.Dir("public")))

	...

}

```
