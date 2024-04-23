Fundamentals of I/O in Go

Everything is a bunch of bytes
Every piece of data you manipulate, whether strings, images, or files, is just a bunch of bytes.

Go's ```len()``` function gives you the number of bytes, not characters, of a string.

```
package main

import "fmt"

func main() {
    str := "ana"
    fmt.Println(len(str))  // 3
    
    str = "世界"
    fmt.Println(len(str)   // 6 not 2
}

```

If you want the length of something, you need the number of bytes, not the number of interpreted things in the data.

If you load an image in Go and run ```len()``` on the data, you the number of bytes (e.g. 293931), not 1.

Readers & Writers
Data movement is comprised of two fundamental parts: reading and writing.

You read from somewhere, apply some transformations, and subsequently write it somewhere.

Because reading and writing are fundamental operations, Go offers two abstractions: ```io.Reader``` and ```io.Writer```.

Remember, a Reader is something you can read from (data source), and a Writer is something you can write to (e.g. STDOUT).

When working with I/O in Go, you don't care what something is but whether you can read from it or write to it. 

As long as it implements the Reader interface, you can read from it or the Writer interface, you can write to it.

This is why we don't have a File interface with the Read and Write methods but rather Reader and Writer interfaces implemented by a concrete file type.


I/O Tools

bytes.Buffer
A bunch of bytes can be represented as ```[]byte``` but this doesn't implement the Reader and Writer interfaces.

Rather, ```bytes.Buffer``` implements both interfaces.


io.Copy
This copies the contents from a Reader (data source) to a Writer (something you write to).

```
package main

import (
    "bytes"
    "io"
    "os"
)

func main() {
    var buf bytes.Buffer
    buf.WriteString("hello world\n")
    
    io.Copy(os.Stdout, &buf)
}

```

fmt.Fprint
Use fmt.Fprint% if you just wan tto print some strings to a Writer.

```
package main

import (
    "bytes"
    "fmt"
    "os"
)

func main() {
    var buf bytes.Buffer

    // Fprintf can be called to print a string to any io.Writer

    fmt.Fprintf(&buf, "hello world!")         // on bytes.Buffer
    fmt.Println(buf.String())

    fmt.Fprintf(os.Stdout, "hello world!\n")  // on a *os.File
}

```


io.ReadAll
Useful when you want to read all bytes from a certain Reader

```
package main

import (
    "fmt" 
    "io" 
    "strings"
)

func main() {
    r := strings.NewReader("Hello World!")
    
    data, err := io.ReadAll(r)
    // handle error
    
    fmt.Println(data)   // [72 101 108 108 111 32 87 111 114 108 100 33]
    fmt.Println(string(data))    // Hello World!
}

```

However, io.ReadAll is not always a good idea if you're dealing with a huge file.