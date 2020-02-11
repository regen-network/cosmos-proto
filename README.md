# cosmos-proto

Extensions to [gogoprotobuf](github.com/gogo/protobuf) for Cosmos.

## `interface_type`

*See [test](test/) for a full example.*

The `cosmos_proto.interface_type` message option generates a getter and
setter from the provided go interface. Messages which use this option must
consist of a single `oneof` and no other fields.

Given this example `.proto` file:

```proto
message ABC {
    option (cosmos_proto.interface_type) = "my_package.Msg";
    oneof abc {
        A a = 1;
        B b = 2;
        C c = 3;
    }
}

message A { int32 x = 1; }

message B { uint32 y = 1; }

message C { bool z = 1; }
```

The `ABC` struct will satisfy the following interface:
```go
type MsgCodec interface {
    GetMsg() Msg
    SetMsg(value Msg) error
}
```

The types `A`, `B`, and `C` must of course have implementations of `Msg`.

### Handling Non-Pointer Interfaces

By default `interface_type` generates `Set` handlers for both pointer and non-pointer
implementations of an interface (i.e. both `*A` and `A`). This means all interfaces
must be implemented for the non-pointer version of a struct. To disable the non-pointer
cases (so that just `*A` and not `A` needs to implement the interface), the interface
type name can be prepended with a `*`, ex: `option (cosmos_proto.interface_type) = "*my_package.Msg";`.

