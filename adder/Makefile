.PHONY: generate
generate:
	rm -rf ./wasi ./docs
	wit-bindgen-go go --world adder ./wit

.PHONY: build
build:
	tinygo build -x -target=wasip2 --wit-package ./wit --wit-world adder -o adder.wasm -x -work main.go
