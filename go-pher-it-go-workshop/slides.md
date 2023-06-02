---
theme: eloc
lineNumbers: true
---

`go-pher it: a go workshop` 

<!--
- intro
- goal
-->

---

`2 robs & a ken`

<!--
- rob pike, rob griesemer and ken thompson
- 2007
- dislike of c++
-->

---

`simplicity, efficiency, concurrency`

<!--
- simplicity, readability and usability
- out of the box tools
- runtime efficiency
- networking and batch processing
-->

---

`ship-it`

<!--
- 2009
-->

---

<img src="/gopher.png" class="m-40 h-40 rounded shadow" />


<!--
- the iconic mascot
-->

---

`anatomy of a go project`

<!--
- common folder structures
- terms
-->

---

```
.
├── a.go
├── a2.go
├── b
│   └── b.go
└── c
    ├── c.go
    └── d
        └── d.go
```

<!--
- packages are directories
- groups code logically
- like a namespace
-->

---

```go {all|2|4-8|10,17|13,18}
// a.go
package a

import (
	"fmt"

	"github.com/joerdav/example/otherpackage"
)

func MyPublicFunction() {
	err := myPrivateFunction()
	if err != nil {
		fmt.Println(err)
	}
}

func myPrivateFunction() error {
	err := otherpackage.DoSomething()
	return err
}
```

<!--
- name the package, usually the folder name
- import other packages
- capital letters exports things
- qualify package names in usages
-->

---

```
.
└── foo
    ├── foo.go
    └── bar.go
```

<div class="grid grid-cols-2 gap-4">
<div>

```go
// foo/foo.go
package foo

func add(a, b int) int {
	return a + b
}
```

</div>
<div>

```go
// foo/bar.go
package foo

import "fmt"

func DoCalculations() {
	fmt.Println(add(1, 2))
	fmt.Println(add(4, 2))
	fmt.Println(add(6, 1))
}
```

</div>
</div>




<!--
- files in the same package can access unexported functions
-->

---

```
.
└── foo
    ├── foo.go
    └── foo_test.go
```

<div class="grid grid-cols-2 gap-4">
<div>

```go
// foo/foo.go
package foo

func add(a, b int) int {
	return a + b
}
```

</div>
<div>

```go
// foo/foo_test.go
package foo

import "testing"

func TestAdd(t *testing.T) {
	result := add(1, 1)
	if result != 2 {
		t.Fatalf("expected: %d, got: %d",
			2, 
			result)
	}
}
```

</div>
</div>

<!--
- this is how tests are written
- underscore test shows go it's a test
-->

---

```
.
├── main.go
├── go.mod
└── foo
    └── foo.go
```

<div class="grid grid-cols-2 gap-4">
<div>

```
// go.mod
module github.com/joerdav/example

go 1.20
```

</div>
<div>

```go
// main.go
package main

import "github.com/joerdav/example/foo"

func main() {
	foo.DoCalculations()
}
```

</div>
</div>

<!--
- modules are repos
- modules are common packages that should be shipped together
-->

---

```
go get github.com/google/uuid
```

<div class="grid grid-cols-2 gap-4">
<div>

```
// go.mod
module github.com/joerdav/example

go 1.20

require (
	github.com/google/uuid v1.3.0
)
```

</div>
<div>

```go
// main.go
package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/joerdav/example/foo"
)

func main() {
	fmt.Println(uuid.NewString())
	foo.DoCalculations()
}
```

</div>
</div>

<!--
- modules are repos
- modules are common packages that should be shipped together
-->

---

```
$ go run .
b70d19c7-6f28-41c4-a145-bdf15dd86e03
3
6
7
```

---
