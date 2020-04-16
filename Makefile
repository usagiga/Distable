xxx:
	@echo "Please select optimal option."

build:
	@go build -o distable .

clean:
	@rm -f ./distable

run:
	@go run .

test:
	@go test -v "./..."
