all: zlint

zlint: cmd/cmd
	cp cmd/cmd zlint

cmd/cmd:
	cd cmd && go build

clean:
	rm -f cmd/cmd zlint

test:
	go test ./...

.PHONY: clean cmd/cmd zlint test
