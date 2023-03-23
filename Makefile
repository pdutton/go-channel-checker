

GO?=go
MOCKGEN?=mockgen

build: unit_test
generate: checker/testing_mock_test.go
mockgen: $(GO) install github.com/golang/mock/mockgen
	
checker/testing_mock_test.go: checker/testing_test.go
	$(MOCKGEN) -source=$< -destination $@ -package checker

unit_test: 
	$(GO) test -cover ./checker

