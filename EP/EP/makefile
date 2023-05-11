SHELL=/bin/sh

include ../config/make.common

EP.o:	EP.go
	go build -o EP EP.go

clean:		
	go clean
	rm -f *.o *~
