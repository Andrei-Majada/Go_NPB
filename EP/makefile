SHELL=/bin/sh
CLASS=S

default: header
	@ $(SHELL) guide

EP: ep
ep: header
	cd EP; make CLASS=$(CLASS)

init: 
	go mod init NPB-GO

clean:
	rm -f core
	rm -f config/definition

header:
	@ $(SHELL) start

help:
	@ $(SHELL) guide
