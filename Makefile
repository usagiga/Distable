xxx:
	@echo "Please select optimal option."

build:
	@go build -o Distable .

clean:
	@rm -f ./Distable

run:
	@go run .

test:
	@go test -v "./..."
