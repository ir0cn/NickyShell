export GOPROXY=https://goproxy.cn,direct

default: dist nicky.exe

nicky.exe:
	go mod tidy && go build -gcflags "-N -l" -o $@ ./

dist:
	cd frontend && npm run build && cd -

build:
	docker build -f docker/Dockerfile.nginx . -t nicky-nginx
	docker build -f docker/Dockerfile.nickyshell . -t nicky-shell

clean:
	rm -fr frontend/dist nicky.exe

