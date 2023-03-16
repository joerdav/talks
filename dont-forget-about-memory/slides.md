---
theme: eloc
---


`don't forget about memory` 

<!--
- talk about memory
- recent obsession
- what brought on this obsession
- hyper optimisation
-->

---

`story time`

<!--
- learning go
- learned about benchmarks
- out of the box
- run code many times
- used on templ
-->

---

`github.com/a-h/templ`

<!--
- for stongly typed html templates in Go
- templ to Go
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

```
TODO: examples
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
- [frames](20230316132529 FRAMES.md)
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
|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|
|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|1|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|1|0|
|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|
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
|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|
|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|1|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|1|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|1|1|
|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|
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
|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|
|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|1|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|1|0|
|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|
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
|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|
|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|0|1|
|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|
```
---

`heap`

---

```go
func main() {
	a := readString()
}
```

<!--
- how much to advance stack
-->

---

```go
func newUser() *User {
	return &User{ Name: "Joe" }
}
```

<!--
- don't lose the memory
- hanging reference
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

- something freeing about using malloc

- wild west
- 
-->

---

```c
char *str = malloc(sizeof(char)*4);
strcpy(str, "Joe");
printf("Hello, %s\n", str);
free(str);
```

<!--
- Manual
- But in control
- Don't always have to free
-->

---

`garbage collection`

<!--
- go python java dotnet

- the runtime handles it 
-->

---

```go
sayHello := func (name string) {
	println("Hello,", name)
}
myName := "Joe"
sayHello(myName)
```
<!--
- how does it know
- it has to stop the world

- stop the world
- inadvertently allocate
- this is slow but can be tuned
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
	let sayHello = { (name: String)  in // #2: +1 ref
	    print("Hello, \(name)")
	} // #3: -1 ref

	var my_name = "Joe" // #1: +1 ref
	sayHello(my_name)
} // #4: -1 ref deallocate
```

<!--
- slow
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
-->
