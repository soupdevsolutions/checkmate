variable "AWS_REGION" {
  type    = string
  default = "eu-west-1"
}

variable "AWS_ACCESS_KEY_ID" {
    type = string
}

variable "AWS_SECRET_ACCESS_KEY" {
    type = string
}

variable "RDS_USERNAME" {
    type = string
}

variable "RDS_PASSWORD" {
    type = string
}

variable "VPC_CIDR_BLOCK" {
    type = string
    default = "10.0.0.0/16"
}

variable "SUBNET_CIDR_BLOCK" {
    type = string
    default = "10.0.1.0/24"
}