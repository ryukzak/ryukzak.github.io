.PHONY: build serve lint clean

# Build public/ and check that every URL in urls.txt exists.
build:
	go run .

# Preview at http://localhost:4000, refresh to see edits.
serve:
	go run . serve

lint:
	test -z "$$(gofmt -l .)" || (gofmt -l . && false)
	go vet ./...

clean:
	rm -rf public
