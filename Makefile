build:
	@go build -o ./bin/gpm

cp-bin:
	@cp ./bin/gpm $$HOME/go/bin

.PHONY: build cp-bin