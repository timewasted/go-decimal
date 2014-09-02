go-decimal
==========

A package for working with arbitrary precision decimal values in Go.

Usage is similar to `math/big`:

```
import (
	"github.com/timewasted/go-decimal"
	// ...
)

decimal, err := ParseDecimal("1234.56")
if err != nil {
	log.Fatal("ParseDecimal error:", err)
}
fmt.Println(decimal.String())
// Prints: 1234.56
fmt.Println(decimal.FormattedString())
// Prints: 1,234.56
decimal2, err := ParseDecimal("6543.21")
if err != nil {
	log.Fatal("ParseDecimal error:", err)
}
if err := decimal.Add(decimal2); err != nil {
	log.Fatal("Add error:", err)
}
fmt.Println(decimal.String())
// Prints: 7777.77
```

Current limitations:
--------------------

* Aside from parsing and printing, the only operations currently implemented are `Cmp` and `Add`. More operations will be added in time, and of course pull requests are welcomed!
* `ParseDecimal` does not parse "formatted" values, such as what `FormattedString` would return. This is unlikely to change.

License:
--------
```
Copyright (c) 2014, Ryan Rogers
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met: 

1. Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer. 
2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution. 

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
```
