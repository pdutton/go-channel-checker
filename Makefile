

GO?=go
MOCKGEN?=mockgen

build: unit_test
generate: checker/T_mock_test.go
mockgen: $(GO) install github.com/golang/mock/mockgen
	
checker/T_mock_test.go: checker/T.go
	$(MOCKGEN) -source=$< -destination $@ -package checker

unit_test: 
	$(GO) test -cover ./checker

