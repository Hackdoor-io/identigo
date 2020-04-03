<img src="/misc/identigo.png" width="100%" />

[![Build Status](https://travis-ci.org/Hackdoor-io/identigo.svg?branch=master)](https://travis-ci.org/Hackdoor-io/identigo)


Identigo generates always the same image when given the same string. <br />
That way, we'll be able to generate a beautiful gradient to be used as a temporary avatar for **Hackdoor** users.

## Usage

Installation via Go Modules is encouraged:

```bash
go get github.com/Hackdoor-io/identigo
```

You can now use **Identigo** as follows:

```go
package main

import identigo "github.com/Hackdoor-io/identigo"

func main() {
  str := "gopher"
  img := identigo.GenerateFromString(str, 256, 256)

	file, _ := os.Create(str + ".png")
	png.Encode(file, img)
}
```
where `GenerateFromString` accepts the following arguments:

| Argument | Type     | Example   |
|----------|----------|-----------|
| `text`   | `string` | `"gopher" |
| `width`  | `int`    | `256`     |
| `height` | `int`    | `256`     |

this code generates the following image:

<img src="/misc/examples/gopher.png" width="150px" />

## Examples

```go
identigo.GenerateFromString("ada_lovelace", 256, 256)
```

<img src="/misc/examples/ada_lovelace.png" width="150px" />

```go
identigo.GenerateFromString("alonzo_church", 256, 256)
```

<img src="/misc/examples/alonzo_church.png" width="150px" />

```go
identigo.GenerateFromString("johndoe", 256, 256)
```

<img src="/misc/examples/johndoe.png" width="150px" />

## License

[GPLv3](/LICENSE.md)