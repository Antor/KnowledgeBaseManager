.PHONY: kbm clean install

DESTDIR = ~/development_environment/bin/knowledge_base_manager

kbm:
	go build -o ./build/kbm ./cmd/kbm

clean:
	rm -rf ./build

install: kbm
	mkdir -p $(DESTDIR)
	cp ./build/kbm $(DESTDIR)
	cp -R ./web $(DESTDIR)

