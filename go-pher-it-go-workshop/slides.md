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
