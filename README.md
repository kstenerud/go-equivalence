Equivalence
===========

A go library for comparing objects.

The equivalence library compares objects without regard to their types, checking to see if they effectively contain the same values, even if their types don't match.

It was designed to be used in unit test code, so it's not super fast. For up to 100,000 comparisons per second it should be fine (depending on the complexity of the objects you are comparing).


Usage
-----

There's only one function: `equivalence.IsEquivalent()`

#### Example

```golang
import (
	"fmt"

	"github.com/kstenerud/go-equivalence"
)

type MyStruct struct {
	IntVal    int
	StringVal string
}

type ComplexStruct struct {
	Map     map[interface{}]interface{}
	Struct  MyStruct
	StructP *MyStruct
}

func DemonstrateEquivalence() {
	a := ComplexStruct{
		Map:     map[interface{}]interface{}{1: "a"},
		Struct:  MyStruct{1, "a"},
		StructP: &MyStruct{100, "test"},
	}

	b := &ComplexStruct{
		Map:     map[interface{}]interface{}{int8(1): "a"},
		Struct:  MyStruct{1, "a"},
		StructP: &MyStruct{100, "test"},
	}
	fmt.Printf("Equivalent: %v\n", equivalence.IsEquivalent(a, b))
}
```


License
-------

Copyright 2019 Karl Stenerud

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.