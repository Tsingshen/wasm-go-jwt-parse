.PHONY: build build-image clean run

build: clean
	tinygo build -scheduler=none -target=wasi ./main.go


build-image: build
	docker build -t ccr.ccs.tencentyun.com/xxx/wasm:http-add-userid-header-v0.l.0 .


clean: 
	go mod tidy
	rm -f main.wasm

run:
	envoy -c ./envoy.yaml --concurrency 2 --log-format '%v'
