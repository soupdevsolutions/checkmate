terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.26.0"
    }
  }

  backend "s3" {
    bucket = "healthchecker-tf-state"
    key    = "terraform.tfstate"
    region = "eu-west-1"
  }

  required_version = ">= 1.5.4"
}

provider "aws" {
  region     = var.AWS_REGION
  access_key = var.AWS_ACCESS_KEY_ID
  secret_key = var.AWS_SECRET_ACCESS_KEY
}