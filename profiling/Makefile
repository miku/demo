cpu.pprof:
	go run main.go -cpuprofile cpu.pprof


mem.mprof:
	go build -o main main.go
	./main -memprofile mem.mprof

.PHONY: pprof
pprof: cpu.pprof
	 go tool pprof cpu.pprof

.PHONY: mem
mem: mem.mprof
	go tool pprof main mem.mprof
