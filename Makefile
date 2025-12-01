bench:
	go test -bench=^BenchmarkWalkFromNode$$ -count=6 ./bfs | tee 02_slice.txt

run:
	go run .

profile:
	go test -bench=^BenchmarkWalkFromNode$$ -count=6 -memprofile=02_map.out ./bfs