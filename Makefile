

bench:
	@go test -benchmem -run=^$ -bench . ./... 