all:
	$(MAKE) -C handlers

	mkdir -p client
	gohandlers client \
		-dir handlers/pets \
		-out client/client.gh.go \
		-pkg client \
		-import "github.com/ufukty/gohandlers-petstore/handlers/pets" 
	gohandlers mock \
		-dir handlers/pets \
		-out client/mock.gh.go \
		-pkg client \
		-import "github.com/ufukty/gohandlers-petstore/handlers/pets" 