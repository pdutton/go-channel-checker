

GO?=go
MOCKGEN?=mockgen

build: unit_test
generate: mock/mock_testing.go
mockgen: $(GO) install github.com/golang/mock/mockgen
	
mock/mock_testing.go: 
	$(MOCKGEN) -destination $@ -package mock testing TB

unit_test: 
	$(GO) test -cover 

