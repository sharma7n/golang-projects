# Install the Go plugin for the protobuf compiler.
go get google.golang.org/protobuf/cmd/protoc-gen-go

# Generate code from the protobuf schema.
protoc donut/donut.proto --go_out=donut

# The PGHOSTADDR environment variable is not supported by the Go plugin for postgres.
unset PGHOSTADDR