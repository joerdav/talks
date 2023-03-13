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

`sub esp, 0x10`

<!--
- any robots here today
- move the stack pointer 
-->

---

`stack`

<!--
- stack overflow
-->

---

```go
func main() {
	a := 1
	if x {
		b := 2
		if y {
			c := 3
		{
	}
}
```

<!--
- example
-->

---

```
 {       {       {
|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|
|0|1|1|0|0|1|1|1|0|1|1|0|1|1|1|1|
|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|
```

<!--
- frames
- stack overflow
-->

---

`heap`

---

```go
func main() {
	a := newRandomString()
}
```

<!--
- how much to advance
-->

---

```go
func newUser() *User {
	return &User{ Name: "Joe" }
}
```

<!--
- don't lose the memory
-->

---

`malloc`

<!--
- hey operating system
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

`manual memory management`

<!--
- c

- malloc/free

- something freeing about using malloc

- wild west
-->

---

```c
char *str = malloc(sizeof(*str));
strcpy(str, "Joe");
printf("Hello, %s\n", str);
free(str);
```

<!--
- Manual
- Don't always have to free
-->

---

`garbage collectors`

<!--
- go python java dotnet

- stop the world
- this is slow but can be tuned
-->

---

`automatic reference counting`

<!--
- swift
- count the references to something
- even slower
-->

---

`static memory management`

<!--
- rust
- define the ownership of all memory
-->
