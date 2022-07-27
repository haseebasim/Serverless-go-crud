.PHONY: build clean deploy

build:
	cd middleware && env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o ../bin/middleware/middleware middleware.go && cd ..
	cd models && env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o ../bin/models/models models.go && cd ..
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o ./bin ./main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose
