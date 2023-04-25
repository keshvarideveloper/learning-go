# go run

Create a directory called learning_go_ch1, open up a text editor, enter the
following text, and save it inside learning_go_ch1 to a file named hello.go

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}
```

After the file is saved, open up a terminal or command prompt and type:

    $ go run hello.go

You should see Hello, world! printed in the console:

     Hello, world!

**The go run command** does in fact compile your code into a binary.

However, **the binary is built in a temporary directory**.
The go run command builds the binary, executes the binary from that temporary directory, and then deletes the binary after your
program finishes. This makes the **go run** command **useful for testing out small programs or using Go like a scripting language**.

## Tip
Use go run when you want to treat a Go program like a script and
run the source code immediately.