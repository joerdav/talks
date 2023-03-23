---
theme: eloc
lineNumbers: true
---


`don't forget about memory` 

<!--
- recent hobby/obsession
- hyper optimization
- let me tell you how it began
-->

---

`story time`

<!--
- the go community
- keep tabs on what go offers

- benchmarks
- out of the box
- run code many times
- used on templ
-->

---

`github.com/a-h/templ`

<!--
- html templating library
- templ to Go to html
- 10MB file
-->

---

`2 minutes 50 seconds`

<!--
- baseline to start improving
- go profile tools
- sys calls
-->

---

```diff
-		_, err = generator.Generate(t, w)
+		b := bufio.NewWriter(w)
+		_, err = generator.Generate(t, b)
 		if err != nil {
 			return fmt.Errorf("%s generation error: %w", fileName, err)
 		}
+		if err := b.Flush(); err != nil {
+			return fmt.Errorf("%s write file error: %w", targetFileName, err)
+		} 
```

<!--
- write to memory not to file
- not always about using the least amount of memory
- some cases allocating more is better
-->

---

`8 seconds`

<!--
- so my obsession began
- I want to tell you some things
- align our mental models
-->

---

```
|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|
|0|1|1|0|0|1|1|1|0|1|1|0|1|1|1|1|
|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|
```

<!--
- visualize memory as a line
-->

---

```go
b := true                       // 00000001

l := byte('a')                  // 01100001

r := rune('😂')                 // 00000000000000011111011000000010

num := int32(8)                 // 0000000000000000000000000000000000000000000000000000000000001000

ui := uint8(255)                // 11111111

i := int8(-127)                 // 11111111

s := struct{ num int8 }{num: 1} // 00000001
```

---

`stack`

<!--
- stack overflow
-->

---

```go
func foo() {
	a := 1
	println(a)
	if x {
		b := 2
		println(b)
		c := 3
		println(c)
	}
}
```

<!--
- example
-->

---

```go {1-2}
func foo() {
	a := 1
	println(a)
	if x {
		b := 2
		println(b)
		c := 3
		println(c)
	}
}
```

```
00000000000000000000000000000001
```

<!--
- defined at compile time
- stack overflow
-->

---

```go {4-9}
func foo() {
	a := 1
	println(a)
	if x {
		b := 2
		println(b)
		c := 3
		println(c)
	}
}
```

```
000000000000000000000000000000010000000000000000000000000000001000000000000000000000000000000011
```

---

```go {9}
func foo() {
	a := 1
	println(a)
	if x {
		b := 2
		println(b)
		c := 3
		println(c)
	}
}
```

```
00000000000000000000000000000001
```
---

`heap`

---

```go
func main() {
	a := "Hello, " + getName()
}
```

<!--
- imagine you are the compiler
- how many bytes should we add
-->

---

```go
func newUser() *User {
	user := User{ ID: 1 }
	return &user
}
```

<!--
- explain pointers
-->

---

`why is joe telling me this?`

<!--
- does he hate us
- send us to sleep
-->
---

`efficiency`

<!--
- stack memory is more efficient
-->

---

`sub esp, 0x10`
<!--
- any robots here today
- move the stack pointer 
-->

---

`malloc`

<!--
- hey operating system
- slow to allocate
- slower to access
- harder to manage
-->

---

`manual memory management`

<!--

- c

- malloc/free

- wild west
-->

---

```c{all|1|4}
char *str = malloc(sizeof(char)*4);
strcpy(str, "Joe");
printf("Hello, %s\n", str);
free(str);
```

<!--
- control
- bugs: access freed mem, free twice, forget all together
- can be quick, error prone, owness on the developer
-->

---

`automatic reference counting`

<!--
- swift
- count the references to something
-->

---

```swift{all|6|7|2|4|8}
{
	func sayHello(name: String) { // #2: +1 ref
	    print("Hello, \(name)")
	} // #3: -1 ref

	var my_name = "Joe" // #1: +1 ref
	sayHello(my_name)
} // #4: -1 ref deallocate
```

<!--
- runtime tracks heap
- dangling references are tracked
- slow, robust, owness on the runtime
-->

---

`static memory management`

<!--
- rust
- define the ownership of all memory
-->

---

```rust
fn say_hello(name: String) {
	println!("Hello, {}", name);
}

fn main() {
	let my_name = String::from("Joe");
	say_hello(my_name);
	// name no longer valid
}
```

<!--
- explain
- different scopes own variables
- setting a variable changes ownership
- can't create dangling references
- quick, safe, owness on the developer with guidance
-->

---

`garbage collection`

<!--
- go python javascipt dotnet

- the runtime handles it 
-->

---

```go
{
	sayHello := func (name string) {
		println("Hello,", name)
	}
	myName := "Joe"
	sayHello(myName)
}
```

<!--
- how does it know
- it has to stop the world

- inadvertently allocate
- quick, safe, owness on runtime, with optional dev input
-->

---

```go{all|1-3|5-7|9-11}
func caclulateStuff(add adder, a, b int) int {
	return add.Add(a, b)
}

type adder interface {
	Add(a, b int) int
}

type concreteAdder struct{}

func (ca concreteAdder) Add(a, b int) int { return a + b }
```

---

```go
func BenchmarkCalculator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		caclulateStuff(concreteAdder{}, 1, 2)
	}
}
```

```
BenchmarkCalculator-12          191312350                6.231 ns/op           0 B/op          0 allocs/op
```

<!--
- anatomy of a benchmark test
-->

---

```{all|6}
$ go test -gcflags '-N -l' -bench=. -benchmem .
goos: darwin
goarch: amd64
pkg: examples/interfaces
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkCalculator-12          191312350                6.231 ns/op           0 B/op          0 allocs/op
PASS
ok      examples/interfaces     2.000s

```

---

```go{1}
func caclulateStuff(add concreteAdder, a, b int) int {
	return add.Add(a, b)
}

type concreteAdder struct{}

func (ca concreteAdder) Add(a, b int) int { return a + b }
```

```
 BenchmarkCalculator-12          310463655                3.881 ns/op           0 B/op          0 allocs/op
```
<!--
- 40%
-->

---

```go {all|12|1-2}
func newHand() []float32 {
	return []float32{
		6,
		6.11,
		6.20,
		6.15,
		5.8,
	}
}

func averageFingerSize() float32 {
	hand := newHand()
	return (hand[0] +
		hand[1] +
		hand[2] +
		hand[3] +
		hand[4]) / float32(len(hand))
}
```

```
BenchmarkFingerSize-12        45860188                26.40 ns/op           24 B/op          1 allocs/op
```

<!--
- no compiler optimisations
- kill it with fire
-->

---

```go {1-2}
func newHand() [5]float32 {
	return [5]float32{
		6,
		6.11,
		6.20,
		6.15,
		5.8,
	}
}

func averageFingerSize() float32 {
	hand := newHand()
	return (hand[0] +
		hand[1] +
		hand[2] +
		hand[3] +
		hand[4]) / float32(len(hand))
}
```

```
BenchmarkFingerSize-12        63423622                15.90 ns/op            0 B/op          0 allocs/op
```

<!--
- 30%
-->

---

```go{all|4}
func createAString() string {
	str := ""
	for i := 0; i < 100; i++ {
		str += "some string"
	}
	return str
}
```

```
BenchmarkRandomGuids-12           105433             10166 ns/op           58872 B/op         99 allocs/op
```

---

```go{2-6}
func createAString() string {
	var build strings.Builder
	for i := 0; i < 100; i++ {
		build.WriteString("some string")
	}
	return build.String()
}
```

```
BenchmarkRandomGuids-12           976308              1225 ns/op            3312 B/op          8 allocs/op
```
<!--
- 88%
-->

---

`xcfile.dev`

<!--
- task runner think npm scripts or makefile
-->

---

```md
# Level 1
```

or

```md
Level 1
===
```

<!--
- task runner think npm scripts or makefile
-->

---

```go{all|2,5}
func altHeadingLevel(nextLine string) int {
	if regexp.MustCompile("^-+$").MatchString(nextLine) {
		return 2
	}
	if regexp.MustCompile("^=+$").MatchString(nextLine) {
		return 1
	}
	return 0
}
```

```
BenchmarkHeading2-12              546982              2173 ns/op            2377 B/op         35 allocs/op
BenchmarkHeading1-12              275372              4334 ns/op            4754 B/op         70 allocs/op
BenchmarkHeading0-12              275092              4263 ns/op            4754 B/op         70 allocs/op
```

---

```go{1-4}
var (
	level2Heading = regexp.MustCompile("^-+$")
	level1Heading = regexp.MustCompile("^=+$")
)

func altHeadingLevel(nextLine string) int {
	if level2Heading.MatchString(nextLine) {
		return 2
	}
	if level1Heading.MatchString(nextLine) {
		return 1
	}
	return 0
}
```

```
BenchmarkHeading2-12            14360788                75.66 ns/op            0 B/op          0 allocs/op
BenchmarkHeading1-12            10675314               111.3 ns/op             0 B/op          0 allocs/op
BenchmarkHeading0-12            16979590                69.95 ns/op            0 B/op          0 allocs/op
```

<!--
- 98%
-->

---

```go{1-11}
func stringOnlyContains(input string, matcher rune) bool {
	if len(input) == 0 {
		return false
	}
	for i := range input {
		if []rune(input)[i] != matcher {
			return false
		}
	}
	return true
}

func altHeadingLevel(nextLine string) int {
	if stringOnlyContains(nextLine, '-') {
		return 2
...
```

```
BenchmarkHeading2-12            30728582                38.91 ns/op            0 B/op          0 allocs/op
BenchmarkHeading1-12            22408417                53.66 ns/op            0 B/op          0 allocs/op
BenchmarkHeading0-12            21599406                56.16 ns/op            0 B/op          0 allocs/op
```

<!--
- 99%
-->

---

```
constant sized variables

string concatination creates a new string

don't be afraid to hyper-optimize
```

<!--
go   - out of the box

rust - out of the box

node - jest-bench

java - caliper?

dotnet - BenchmarkDotNet
-->
