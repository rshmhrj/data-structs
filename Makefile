watch_test:
	ulimit -n 1000
	reflex -s -r '\.go$$' make run_tests

run_tests:
	go test ./...