SHELL=/bin/sh
CLASS=S

default: header
	@ $(SHELL) guide

IS: is
is: header
	cd IS; make CLASS=$(CLASS)
	
init: 
	go mod init NPB-GO

clean:
	rm -f core
	rm -f config/definition

header:
	@ $(SHELL) start

help:
	@ $(SHELL) guide
