resource "aws_ecr_repository" "healthchecker-be" {
  name = "healthchecker-be"
}

data "aws_ecr_authorization_token" "token" {}
