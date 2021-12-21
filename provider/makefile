include ${GOPATH}/src/github.com/NETkiddy/nft-svr/makefile.go.comm

TARGET=bee_provider
DESTDIR=/usr/local/services/${TARGET}

all: ${TARGET}
${TARGET}:main
	mv main ${TARGET}
main:main.go

clean:
	rm -f ${TARGET}

install:
	sudo rm -Rf $(DESTDIR)
	sudo rm -Rf $(DESTDIR).tar.gz
	sudo -uuser_00 -gusers mkdir -p $(DESTDIR)/{bin,conf}
	sudo install -c -ouser_00 -gusers -m0777 ./app-release.toml $(DESTDIR)/conf/
	sudo install -c -ouser_00 -gusers -m0777 ./app-beta.toml $(DESTDIR)/conf/
	sudo install -c -ouser_00 -gusers -m0777 $(TARGET) $(DESTDIR)/bin/

.PHONY : all clean debug install
