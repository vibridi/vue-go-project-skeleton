.PHONY: deps test dev-live

deps:
	go mod download && go mod tidy

test:
	go test -cover ./...

dev-live:
	FRONTEND_ADDR=":30777" \
	gin -p 30776 -a 30777 --all -x node_modules -x view -x build -b build/gin-bin
