resource "aws_vpc" "healthchecker" {
  cidr_block = var.VPC_CIDR_BLOCK

  tags = {
    Name = "healthchecker"
  }
}

resource "aws_subnet" "healthchecker" {
  vpc_id     = aws_vpc.healthchecker.id
  cidr_block = var.SUBNET_CIDR_BLOCK
  map_public_ip_on_launch = true
}

resource "aws_lb_target_group" "healthchecker" {
  name     = "healthchecker"
  port     = 80
  protocol = "HTTP"
  vpc_id   = aws_vpc.healthchecker.id
}