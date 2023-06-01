go-vendor:
	docker run -ti -v $$(pwd):/go/src/github.com/karnott/skalin-sdk golang:1.20-buster /bin/bash -c "cd /go/src/github.com/karnott/skalin-sdk && GO111MODULE=on go mod vendor"
