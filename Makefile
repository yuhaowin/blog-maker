APP=blog-maker

.PHONY: run
run:
# @ 的目的是在执行命令前不进行 echo
	@go run main.go


.PHONY: build
build:
	@go build -o ${APP} main.go


clean:
	@rm -f ${APP} .meta
	@rm ${APP}_*


platforms := $(windows linux darwin)
release:
	@for v in windows linux darwin ; do \
		GOOS=$$v GOARCH=amd64 go build -o ${APP}_$${v}_amd64 *.go ; \
		zip -ur ${APP}_$${v}_amd64.zip ${APP}_$${v}_amd64 templates ; \
		GOOS=$$v GOARCH=386 go build -o ${APP}_$${v}_386 *.go ; \
		zip -ur ${APP}_$${v}_386.zip ${APP}_$${v}_386 templates ; \
    done
