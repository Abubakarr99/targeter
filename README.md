# targeter
Target terraform file

## Installation

To install the package
```shell
go install github.com/Abubakarr99/targeter@latest
```

## Examples

```shell
targeter get -f examples/example.tf
# Output: terraform plan target="aws_vpc.example" target="aws_instance.web"
```

