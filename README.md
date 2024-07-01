# Building a Calculator of Wasm Components

This repo contains the example from the [component model](https://github.com/bytecodealliance/component-docs/tree/main/component-model/examples/tutorial) documentation. The [adder](adder) component has been refactored to TinyGo.

## Requirements for building the TinyGo example

- TinyGo compiled with [wasip2](https://github.com/dgryski/tinygo/tree/dgryski/wasi-preview-2).
- [wit-bindgen-go](https://github.com/ydnar/wasm-tools-go)

## Building and running the example

To compose a calculator component with an add operator, run the following:

```sh
make build
make compose
```

Now, run the component with wasmtime:

```sh
wasmtime run final.wasm 1 2 add
1 + 2 = 3
```
