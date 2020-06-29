# depends on golang.org/x/perf/cmd/benchstat

bench:
	go test --bench=. > benchmarks/out.txt

results:
	benchstat benchmarks/*.txt