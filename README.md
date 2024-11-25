# beats

For all your [Swatch Internet Time](http://en.wikipedia.org/wiki/Swatch_Internet_Time) needs.

[![GoDoc](https://godoc.org/github.com/peterhellberg/beats?status.png)](https://godoc.org/github.com/peterhellberg/beats)

## Installation

```bash
$ go get -u github.com/peterhellberg/beats
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/peterhellberg/beats"
)

func main() {
	fmt.Println(beats.NowString())
}
```

## Commands

### beats-gui

![beats-gui](https://assets.c7.se/skitch/beats-gui-20140902-222030.png)

### beats

```none
Usage: beats [<timestamp>]

Arguments:
  [<timestamp>]
          Unix timestamp to convert to .beat time

Examples:

$ beats
@456

$ beats 1732500000
@125
```

## LICENSE

Copyright (c) 2014 [Peter Hellberg](http://c7.se/)

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the "Software"),
> to deal in the Software without restriction, including without limitation
> the rights to use, copy, modify, merge, publish, distribute, sublicense,
> and/or sell copies of the Software, and to permit persons to whom the
> Software is furnished to do so, subject to the following conditions:

> The above copyright notice and this permission notice shall be included
> in all copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
> OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
> IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
> DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
> TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
> OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
