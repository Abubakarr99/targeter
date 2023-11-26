# targeter
Target terraform file

## Installation

To install the package
```shell
go install github.com/Abubakarr99/targeter@v1.0.0
```

## Examples

```shell
targeter get -f example.tf
# Output: terraform plan target="aws_vpc.example" target="aws_instance.web"
```

