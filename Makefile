xml2csv:
	go build xml2csv.go
	GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o xml2csv.linux xml2csv.go
	GOOS=windows GOARCH=386 go build -o xml2csv.exe xml2csv.go
