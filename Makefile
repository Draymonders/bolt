BRANCH=`git rev-parse --abbrev-ref HEAD`
COMMIT=`git rev-parse --short HEAD`
GOLDFLAGS="-X main.branch $(BRANCH) -X main.commit $(COMMIT)"


TestFunc="TestPage_dump"

default: build

race:
	@go test -v -race -test.run="TestSimulate_(100op|1000op)"

# go get github.com/kisielk/errcheck
errcheck:
	@errcheck -ignorepkg=bytes -ignore=os:Remove github.com/draymonders/bolt

test: 
#	@go test -v -cover .
#	@go test -v ./cmd/bolt
	@go test -v -count=1 -run $(TestFunc) page_test.go page.go db.go bucket.go tx.go node.go freelist.go boltsync_unix.go bolt_amd64.go cursor.go bolt_unix.go errors.go

.PHONY: fmt test
