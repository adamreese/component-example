.PHONY: generate
generate:
	rm -rf ./wasi ./docs
	wit-bindgen-go go --world subtractor ./wit

.PHONY: build
build:
	tinygo build -x -target=wasip2 --wit-package ./wit --wit-world subtractor -o subtractor.wasm -x -work main.go
