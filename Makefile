bin/syn: syn/*.go  syn/parser.go.yy
	go tool yacc -o syn/parser.go -v syn/parser.output syn/parser.go.yy && rm syn/parser.output && go build -o bin/syn .
	@chmod 700 bin/syn

.PHONY: clean
clean:
	-rm bin/syn
