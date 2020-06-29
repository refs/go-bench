## Go Performant Code

This write down does not consider readability, only performance.

### Methodology

Comparisons using [benchstat - GoDoc](https://godoc.org/golang.org/x/perf/cmd/benchstat)


### Method Sets

On the matter of interfaces vs struct, it so seems that there is a cost on indirection when having methods on interfaces.

> If we’re really trying to micro-optimize, we should also be careful to avoid using interfaces, which have some overhead. As with other kinds of dynamic dispatch, there is a cost of indirection when performing a lookup for the method call at runtime. The compiler is unable to inline these calls.


[Dynamic dispatch - Wikipedia](https://en.wikipedia.org/wiki/Dynamic_dispatch): _process of selecting which implementation of a polymorphic operation (method or function) to call at run time_

Cost of indirection quantified:

```
name            time/op
SayStruct-8     0.27ns ± 0%
SayInterface-8  2.08ns ± 0%
```

### Extra

[Practical Go Benchmarks - StackImpact](https://stackimpact.com/blog/practical-golang-benchmarks/)
