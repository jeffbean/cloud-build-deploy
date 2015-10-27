all: push

# 0.0 shouldn't clobber any released builds
TAG = 0.0
PREFIX = jeffbean/cloud-build-deploy

build: main.go
	CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-w' -o main ./main.go

container: build
	docker build -t $(PREFIX):$(TAG) .

push: container
	docker push $(PREFIX):$(TAG)

clean:
	rm -f main
