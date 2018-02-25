all: bld

bld: frontEnd

frontEnd:
	go build -o bin/frontEnd 

clean:
	@rm -f bin/frontEnd
	@rm -f  log/*log*
	@rm -rf ./output
	@rm -rf output.tar.gz

cleanlog:
	@rm -f log/*log*
