resource "aws_iam_role" "healthchecker-be" {
    assume_role_policy = data.aws_iam_policy_document.healthchecker-be.json
}

data "aws_iam_policy_document" "healthchecker-be" {
    statement {
        actions = [
            "sts:AssumeRole",
        ]

        principals {
            type        = "Service"
            identifiers = ["ec2.amazonaws.com"]
        }
    }
}