# Workflow for Go

[![GoDoc](https://godoc.org/github.com/neonxp/workflow?status.svg)](https://godoc.org/github.com/neonxp/workflow)

Simple state machine. Inspired by [Symfony Workflow](https://github.com/symfony/workflow).

## Example usage

```go
o := new(ObjectImplementedPlaceer)

w := NewWorkflow("Start")
w.AddTransition("Start", "A")
w.AddTransition("Start", "B")
w.AddTransition("A", "C")
w.AddTransition("B",  "D")
w.AddTransition( "C", "D")
w.AddTransition("C", "Finish")
w.AddTransition("D", "Finish")

w.Can(o, "A") // == nil
w.Can(o, "C") // == ErrTransitionNotFound

w.GetEnabledTransitions(o) // []Place{"A", "B"}
w.Apply(o, "A") // o now at "A" place
w.GetEnabledTransitions(o) // []Place{"C"}

w.DumpToDot() // See above
```

## Dump result

```
digraph {
    Start[color="blue"]
    Start -> A[label="Start → A"];
    Start -> B[label="Start → B"];
    A -> C[label="A → C"];
    B -> D[label="B → D"];
    C -> D[label="C → D"];
    C -> Finish[label="C → Finish"];
    D -> Finish[label="D → Finish"];
}
```

Visualization:

![Workflow visualization](images/example.png)

