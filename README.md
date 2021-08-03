[![GoDoc](https://godoc.org/github.com/uudashr/structfield?status.svg)](https://godoc.org/github.com/structfield/gocognit)
# Structfield
Find struct literals using non-labeled fields.

The structfield analysis reports the usage of struct literal using non-labeled fields more than defined limit (default: 2). The non-labeled field above the limit considered has higher cognitive load (harder to understand and rememeber).

## Understanding Struct Literal

Given code, variable assigned using struct literal:
```go
acc := Account{
    Name: "John Smith",
    Email: "john.smith@example.com",
    []Permission{
        Permission{"account", "read"},
        Permission{"account", "write"},
    },
    false,
}
```

Above code is harder to remember the field name without looking the `Account` type declaration.

Suggestion is to refactor the code to:
```go
acc := Account{
    Name: "John Smith",
    Email: "john.smith@example.com",
    []Permission{
        Permission{"account", "read"}, // Non-labeled here is still ok
        Permission{"account", "write"},
    },
    Deactivated: false,
}
```

The limit set to 2 (default value), which considered easy to understand and remember.

## Installation
```
$ go get github.com/uudashr/structfield/cmd/structfield
```

## Usage

```
$ structfield -limit 2 testdata/src/a/*.go
testdata/src/a/a.go:20:9: Found 4 non-labeled fields on struct literal (> 2)
```