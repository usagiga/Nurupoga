xxx:
	@echo "Please select optimal option."

build:
	@go build -o nurupoga .

clean:
	@rm -f ./nurupoga

run:
	@go run .

test:
	@go test -v "./..."
