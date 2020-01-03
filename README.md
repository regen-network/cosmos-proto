# cosmos-proto

Extensions to [gogoprotobuf](github.com/gogo/protobuf) for Cosmos.

## `interface_type`

*See [test](test/) for a full example.*

With this proto file:

```proto
message ABC {
    option (cosmos_proto.interface_type) = "my_package.MyInterface";
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

and these interface definitions:
```go
type MyInterface interface { SomeMethod() }

func (m A) SomeMethod() { }

func (m B) SomeMethod() { }

func (m C) SomeMethod() { }
```

it is possible to convert the `ABC` message to and from `MyInterface`:

```go
func ExampleFrom(x MyInterface) {
  var abc ABC 
  abc.FromInterface(x)
}

func ExampleTo(abc ABC) {
  var x MyInterface
  x = abc.ToInterface()
}
```
