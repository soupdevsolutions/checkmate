resource "aws_vpc" "healthchecker" {
  cidr_block = var.VPC_CIDR_BLOCK

  tags = {
    Name = "healthchecker"
  }
}

resource "aws_internet_gateway" "healthchecker-gw" {
  vpc_id = aws_vpc.healthchecker.id

  tags = {
    Name = "healthchecker-gw"
  }
}

resource "aws_subnet" "healthchecker" {
  count                   = var.AZ_COUNT
  vpc_id                  = aws_vpc.healthchecker.id
  availability_zone       = data.aws_availability_zones.available.names[count.index]
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

  depends_on = [aws_lb.healthchecker-be-lb]
}

resource "aws_lb_listener" "healthchecker" {
  load_balancer_arn = aws_lb.healthchecker-be-lb.arn
  port              = "80"
  protocol          = "HTTP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.healthchecker.arn
  }
}
