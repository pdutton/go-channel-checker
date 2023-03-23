

GO?=go
MOCKGEN?=mockgen

build: unit_test
generate: mocks/mock_testing.go
mockgen: $(GO) install github.com/golang/mock/mockgen
	
mocks/mock_testing.go: 
	$(MOCKGEN) -destination $@ testing TB

unit_test: 
	$(GO) test -cover 

