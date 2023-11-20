resource "aws_security_group" "healthchecker_be" {
  name        = "healthchecker-be"
  description = "Healthchecker backend security group"
  vpc_id      = aws_vpc.healthchecker.id

  tags = {
    Name = "healthchecker-be"
  }
}

resource "aws_security_group_rule" "healthchecker_be_inbound" {
  security_group_id = aws_security_group.healthchecker_be.id
  type              = "ingress"
  from_port         = 80
  to_port           = 80
  protocol          = "http"
  cidr_blocks       = ["0.0.0.0/0"]
}

resource "aws_security_group_rule" "healthchecker_be_inbound" {
  security_group_id = aws_security_group.healthchecker_be.id
  type              = "ingress"
  from_port         = 0
  to_port           = 0
  protocol          = "-1"
  cidr_blocks       = ["0.0.0.0/0"]
}
