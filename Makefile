watch_test:
	ulimit -n 1000
	reflex -s -r '\.go$$' make run_tests

run_tests:
	go test ./... -cover

run_benchmarks:
	go test ./... -bench=.

run_timed_benchmarks:
	go test ./... -bench=. -benchtime=10s
