---
theme: eloc
---


`don't forget about memory` 

<!--
- recent hobby/obsession
- what triggered it
- start with a story
-->

---

`story time`

<!--
- the go community
- wasn't always that way
- roots in c#
- not here to talk about go

- benchmarks
- out of the box
- run code many times
- used on templ
-->

---

`github.com/a-h/templ`

<!--
- for stongly typed html templates in Go
- templ to Go to html
- 10MB file
-->

---

`2 minutes 50 seconds`

<!--
- baseline to start improving
- go profile tools
- write to memory not to a file
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
- not always about using the least amount of memory
- some cases allocating more is better
-->

---

`8 seconds`

<!--
- so my obsession began
- let me tell you some things I learned
- not a realistic use case
- but can help you write better programs
-->

---

`remind me...`

<!--
- back to basics
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
		if y {
			c := 3
			println(c)
		}
	}
}
```

<!--
- example
-->

---

```go {1}
func foo() {
	a := 1
	println(a)
	if x {
		b := 2
		println(b)
		if y {
			c := 3
			println(c)
		}
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

```go {4}
func foo() {
	a := 1
	println(a)
	if x {
		b := 2
		println(b)
		if y {
			c := 3
			println(c)
		}
	}
}
```

```
0000000000000000000000000000000100000000000000000000000000000010
```

<!--
- frames
- stack overflow
-->
---

```go {7}
func foo() {
	a := 1
	println(a)
	if x {
		b := 2
		println(b)
		if y {
			c := 3
			println(c)
		}
	}
}
```

```
000000000000000000000000000000010000000000000000000000000000001000000000000000000000000000000011
```
---

```go {10}
func foo() {
	a := 1
	println(a)
	if x {
		b := 2
		println(b)
		if y {
			c := 3
			println(c)
		}
	}
}
```

```
0000000000000000000000000000000100000000000000000000000000000010
```
---

```go {11}
func foo() {
	a := 1
	println(a)
	if x {
		b := 2
		println(b)
		if y {
			c := 3
			println(c)
		}
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
	return &User{ Name: "Joe" }
}
```

<!--
- can I just pop the stack?
- when should I remove this data?
- dangling reference
-->

---

`why is joe telling me this?`

<!--
- does he hate us
- trying to be the most tedious talk
-->
---

`efficiency`

<!--
- speed
- heap slower, stack faster
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

```c
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

```swift
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

```go
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
