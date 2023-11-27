resource "aws_ecr_repository" "healthchecker-be" {
  name = "healthchecker-be"
}

data "aws_ecr_authorization_token" "token" {}

resource "docker_image" "healthchecker-be" {
  name = "${data.aws_ecr_authorization_token.token.proxy_endpoint}/healthchecker-be:latest"
  build {
    context = "../"
  }
  platform = "linux/arm64"
}

resource "docker_registry_image" "healthchecker-be-handler" {
  name = docker_image.healthchecker-be.name
}