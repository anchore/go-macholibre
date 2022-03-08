all:
	cd fixture-ls && make

non-mach-o:
	GOOS=linux GOARCH=amd64 go build -o linux_amd64 fixture-non-mach-o/main.go
	mv linux_amd64 assets/