# Omega Validator

Omega Validator is a lightweight and intuitive validation library for Go. It allows you to compose validations declaratively and provides detailed feedback for invalid fields.

## Features

- Chain-style validation methods
- Built-in validators for common use cases
- Get detailed error messages for invalid fields
- Multiple field validation support
- Zero external dependencies
- Easy to extend

## Installation

```bash
go get github.com/matheus-gondim/omega-validator
```

## How It Works

- **Create a Validator**: Use `validator.New` to define a field and its value.
- **Apply Rules**: Chain validation rules like `.Min()`, `.Max()`, `.Email()`, etc.
- **Compose Validations**: Use `validator.Compose` to combine multiple validators.
- **Check Results**: The result contains a success flag (`bool`) and detailed errors.

## Usage Examples

### Basic Validation

```go
// Single field validation
valid, err := validator.New("age", 25).Min(18).Max(100).Validate()

// Multiple field validation
valid, err := validator.Compose(
    validator.New("username", "john_doe").Required().Min(3),
    validator.New("email", "john@example.com").Required().Email(),
)
```

### Document Validation

```go
// CPF validation
valid, err := validator.New("cpf", "00000000000").FederalDocument().Validate()

// CNPJ validation
valid, err := validator.New("cnpj", "00000000000000").FederalDocument().Validate()
```

### Number Validation

```go
valid, err := validator.New("campo_1", 12).Min(1).Validate()
```

## Error Handling

The validator returns both a boolean indicating overall validation success and an error object containing detailed validation errors for each field:

```go
isValid, err := validator.Compose(
    validator.New("email", "invalid-email").Email(),
)

if err != nil {
    for field, errors := range err.Errors {
        fmt.Printf("Field %s has errors: %v\n", field, errors)
    }
}
```

## License

This project is licensed under the [MIT License](./LICENSE).

## Support

If you encounter any problems or have suggestions, please open an issue in the GitHub repository.
