all: bin/mac/timezones bin/linux/timezones bin/windows/timezones.exe

bin/mac/timezones: timezones.go
	GOOS=darwin GOARCH=amd64 go build -o bin/mac/timezones

bin/linux/timezones: timezones.go
	GOOS=linux GOARCH=amd64 go build -o bin/linux/timezones

bin/windows/timezones.exe: timezones.go
	GOOS=windows GOARCH=amd64 go build -o bin/windows/timezones.exe

clean:
	rm -rf bin

.PHONY: clean
