resource "aws_vpc" "healthchecker" {
  cidr_block = var.VPC_CIDR_BLOCK

  tags = {
    Name = "healthchecker"
  }
}

resource "aws_subnet" "healthchecker" {
  count                   = var.AZ_COUNT
  vpc_id                  = aws_vpc.healthchecker.id
  cidr_block              = cidrsubnet(aws_vpc.healthchecker.cidr_block, 4, count.index)
  map_public_ip_on_launch = true
}

resource "aws_lb" "healthchecker-be-lb" {
  name               = "healthchecker-be-lb"
  internal           = false
  load_balancer_type = "application"
  security_groups    = [aws_security_group.healthchecker_be.id]
  subnets            = [for subnet in aws_subnet.healthchecker : subnet.id]

  enable_deletion_protection = true
}

resource "aws_lb_target_group" "healthchecker" {
  name     = "healthchecker"
  port     = 80
  protocol = "HTTP"
  vpc_id   = aws_vpc.healthchecker.id
}
