run:
	@go build -o generate
	@./generate

clean:
	@rm generate