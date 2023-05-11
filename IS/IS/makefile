SHELL=/bin/sh

include ../config/make.common

IS.o:	IS.go
	go build -o IS IS.go

clean:		
	go clean
	rm -f *.o *~