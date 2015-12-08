all:	euler

euler:	euler.go
	go fmt euler.go
	go build euler.go

clean:
	rm -rf euler
	find -type f -name "*~" -print0|xargs -0 rm

check:	euler.go euler_test.go
	go fmt euler_test.go
	export GOPATH=`pwd`
	go test

test:	check