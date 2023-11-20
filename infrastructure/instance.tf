resource "aws_instance" "healthchecker_backend" {
  ami           = "ami-07355fe79b493752d" // eu-west-1 Amazon Linux 2023
  instance_type = "t3.micro"

  tags = {
    Name = "healthchecker-backend"
  }
}
