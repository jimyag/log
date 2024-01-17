
build:
	GOBIN=$(shell pwd) go install  -trimpath -v  ./...

run:
	./sample
	echo 
	cat sample.log

clean:
	rm -f sample
	rm -f sample.log