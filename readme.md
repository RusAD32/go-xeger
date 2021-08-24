go-xeger
=====

go-xeger is a golang module that generates random strings from a regular expression.
Inspired by Java library [Xeger](https://code.google.com/archive/p/xeger/) 
and Python library [xeger](https://pypi.org/project/xeger/)

### Installation

To install go-xeger:

```bash
$ go get github.com/RusAD32/go-xeger
```

### Usage

```go
package main

import (
    "fmt"
    "regexp/syntax"
    "math/rand"
    
    "github.com/RusAD32/go-xeger"
)

func main() {
    // for one time generation
    resultStr, err := xeger.Generate("[0-9]+")
    if err != nil {
        panic(err)
    }
    fmt.Println(resultStr)
    
    // using a Xeger object 
    x, err := xeger.NewXeger("[0-9]+")
    if err != nil {
        panic(err)
    }
    // Optionally, set custom upper limit
    x.Limit = 15
    fmt.Println(x.Generate())
    
    // using a set seed
    x1, err := xeger.NewXegerWithSeed("[0-9]+", 123456)
    if err != nil {
        panic(err)
    }
    fmt.Println(x1.Generate())
    
    // creating a Xeger object directly
    myRegex, _ := syntax.Parse("[0-9]+", syntax.Perl) // handle this error in the real code
    myXeger := &xeger.Xeger{
    	myRegex,
    	rand.NewSource(123456),
    	15,
    }
    fmt.Println(myXeger.Generate())
}
```

### Contributing

Contributions are very welcome. Please open a tracking issue or pull request and we can work to get things merged in.
