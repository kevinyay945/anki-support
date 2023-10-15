.PHONY: test
test:
	RUN_INFRASTRUCTURE=true go test ./...


.PHONY: init

init:
	go install go.uber.org/mock/mockgen@bb5901fe6e45c7c5035afb29a274b9e970c8e348
	go install github.com/google/wire/cmd/wire@0ac845078ca01a1755571c53d7a8e7995b96e40d
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@7aa85bb88223ee606c2aaeb3e536aa0ed93d4054
	go install github.com/spf13/cobra-cli@74762ac083f2c4deffef229c887ffc15beb6ce0d

#.PHONY: di
#di:
#	wire gen "./di"

.PHONY: go_generate
go_generate:
	go generate ./...

.PHONY: clean
clean:
	$(RM) ./di/wire_gen.go
	#$(RM) **/*.gen.go
	$(RM) **/*.mock.go

.PHONY: di
di:
	wire gen "./di"

.PHONY: generate
generate: clean go_generate di

