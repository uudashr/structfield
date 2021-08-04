[![GoDoc](https://godoc.org/github.com/uudashr/structfield?status.svg)](https://godoc.org/github.com/uudashr/structfield)
# Structfield
Find struct literals using non-labeled fields.

The structfield analysis reports the usage of struct literal using non-labeled fields more than defined limit (default: 2). The non-labeled field above the limit considered has higher cognitive load (harder to understand and rememeber).

## Understanding Struct Literal

Given code, variable assigned using struct literal:
```go
acc := Account{
    "John Smith",
    "john.smith@example.com",
    []Permission{
        Permission{"account", "read"},
        Permission{"account", "write"},
    },
    false,
    false,
}
```

Above code is harder to understand, hard to guess the field name since we have to remember exact order of the fields. The workaround is you have to always look the declaration of the `Account` type.

Suggestion is to refactor the code to:
```go
acc := Account{
    Name: "John Smith",
    Email: "john.smith@example.com",
    Permission: []Permission{
        Permission{"account", "read"}, // Non-labeled here is still ok
        Permission{"account", "write"},
    },
    Verified: false,
    Deactivated: false,
}
```

The limit set to 2 (default value), which considered easy to understand and remember.

## Benefits
By using the labeled fields you several benefits

1. The fields doesn't have to be in order
2. You don't have to declare the value if it's a default value

Example:
```go
acc := Account{
    Name: "John Smith",
    Email: "john.smith@example.com",
    Permission: []Permission{
        Permission{"account", "read"}, // Non-labeled here is still ok
        Permission{"account", "write"},
    },
    Verified: true,
    Deactivated: false,
}
```

can be simplified into:
```go
acc := Account{
    Name: "John Smith",
    Email: "john.smith@example.com",
    Permission: []Permission{
        Permission{"account", "read"}, // Non-labeled here is still ok
        Permission{"account", "write"},
    },
    Verified: true,
    // Remove the `Deactivated: false` since it use default value
}
```

## Installation
```
$ go install github.com/uudashr/structfield/cmd/structfield@latest
```

or

```
$ go get github.com/uudashr/structfield/cmd/structfield
```

## Usage

```
$ structfield -limit 2 testdata/src/a/*.go
testdata/src/a/a.go:20:9: Found 4 non-labeled fields on struct literal (> 2)
```