.PHONY: build
build: build-adder
build: build-subtractor
build: build-calculator
build: build-command

.PHONY: build-adder
build-adder:
	cd adder && make generate build

.PHONY: build-subtractor
build-subtractor:
	cd subtractor && make generate build

.PHONY: build-calculator
build-calculator:
	cd calculator && cargo component build --release

.PHONY: build-command
build-command:
	cd command && cargo component build --release

.PHONY: compose
compose:
	wasm-tools compose calculator/target/wasm32-wasi/release/calculator.wasm -d adder/adder.wasm -d subtractor/subtractor.wasm -o composed.wasm
	wasm-tools compose command/target/wasm32-wasi/release/command.wasm -d composed.wasm -o final.wasm

.PHONY: run
run:
	wasmtime run final.wasm 1 2 add
