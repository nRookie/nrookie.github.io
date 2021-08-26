# Data Transport through gob



## What is gob ?

The gob is Go’s format for serializing and deserializing program data in binary format. It is found in the encoding package. Data in this format is called a gob (short for Go binary format). It is similar to Python’s pickle or Java’s Serialization.

It is typically used in transporting arguments and results of remote procedure calls (RPCs) (see the rpc package Chapter 13).

This package works with the language in a way that an externally-defined, language-independent encoding cannot. That’s why the format is binary in the first place, not a text-format like JSON or XML. Gobs are not meant to be used in other languages than Go because, in the encoding and decoding process, Go’s reflection capability is used.

## Explanation

Gob files or streams are entirely self-describing. For every type, they contain a description of that type, and they can always be decoded in Go without any knowledge of the file’s contents. Only exported fields are encoded; zero values are not taken into account. When decoding structs, fields are matched by name and compatible type, and only fields that exist in both are affected. In this way, a gob decoder client will still function when in the source datatype fields have been added. The client will continue to recognize the previously existing fields. Also, there is excellent flexibility provided, e.g., integers are encoded as unsized, and variable-length, regardless of the concrete Go type at the sender side.


So if we have at the sender side a struct T:

``` golang

type T struct { X, Y, Z int }


var t = T{X: 7, Y: 0, Z: 8}

```



This can be captured at the receiver side in a variable u of type struct U:

``` golang
type U struct { X, Y *int 8}

var u U

```

At the receiver, X gets the value 7, and Y the value 0 (which was not transmitted). In the same way as json, gob works
by creating an encoder object with a NewEncoder() function and calling Encode(), again completely generalized by using io.Writer. The inverse is done with a decoder object with a NewDecoder() function and calling Decode(), generalized by using io.Reader.



## decode from a gob file.
``` golang
package main
import (
	"bufio"
	"fmt"
	"encoding/gob"
	"log"
	"os"
)

type Address struct {
	Type             string
	City             string
	Country          string
}

type VCard struct {
	FirstName	string
	LastName	string
	Addresses	[]*Address
	Remark		string
}

var content	string
var vc VCard

func main() {
		// using a decoder:
	file, _ := os.Open("vcard.gob")
	defer file.Close()
	inReader := bufio.NewReader(file)
	dec := gob.NewDecoder(inReader)
	err := dec.Decode(&vc)
	if err != nil {
		log.Println("Error in decoding gob")
	}
	fmt.Println(vc)
}
```