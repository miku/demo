cpu.pprof:
	go run main.go -cpuprofile cpu.pprof

.PHONY: pprof
pprof: cpu.pprof
	 go tool pprof cpu.pprof
