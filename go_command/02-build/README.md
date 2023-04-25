# go build

Most of the time you want to build a binary for later use. That’s where you use the go
build command.

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

    $ go build hello.go

This creates an executable called hello (or *hello.exe* on Windows) in the current
directory.

Run the executable file, and you unsurprisingly see **Hello, world!** printed on the screen.

## Flag -o

The name of the binary matches the name of the file or package that you passed in.  

**use the **-o** flag when:**

*   *If you want a different name for your application*
    
*   *if you want to store it in a different location*


For example, if we wanted to compile our code to a binary
called “**hello_world**”, we would use:

    $ go build -o hello_world hello.go

## Tip 
Use go build to create a binary that is distributed for other people
to use. Most of the time, this is what you want to do. Use the -o flag
to give the binary a different name or location.