# Create a VPC
resource "aws_vpc" "example" {
  cidr_block = "10.0.0.0/16"
}

resource "aws_instance" "web" {
  ami           = "ami"
  instance_type = "t3.micro"

  tags = {
    Name = "HelloWorld"
  }
}

