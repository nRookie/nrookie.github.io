## Inheritance of interfaces


When a type includes (embeds) another type (which implements one or more interfaces) as a pointer, then the type can use all of the interface’s methods. For example:

``` go
type Task struct {
  Command string
  *log.Logger
}

```

A factory for this type could be:

``` go
func NewTask(command string, logger *log.Logger) *Task {
return &Task{command, logger}
}
```

When log.Logger implements a Log() method, then a value task of Task can call it:


``` go
task.Log()
```

A type can also inherit from multiple interfaces providing something like multiple inheritance:


``` go
type ReaderWriter struct {
  io.Reader
  io.Writer
}
```

The principles outlined above are applied throughout all Go-packages, thus maximizing the possibility of using polymorphism and minimizing the amount of code. This is considered an important best practice in Go-programming.

Useful interfaces can be detected when the development is already underway. It is easy to add new interfaces because existing types don’t have to change (they only have to implement their methods). Current functions can then be generalized from having a parameter(s) of a constrained type to a parameter of the interface type: often, only the signature of the function needs to be changed. Contrast this to class-based OO-languages, where, in such a case, the design of the whole class-hierarchy has to be adapted.

