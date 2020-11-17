clean:
	rm -rf bin/*
compile:
	GOOS=linux GOARCH=386 go build -o bin/linux/sm
	GOOS=darwin GOARCH=amd64 go build -o bin/macos/sm
	GOOS=windows GOARCH=386 go build -o bin/windows/sm