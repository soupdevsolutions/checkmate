resource "aws_iam_role" "healthchecker-be" {
  assume_role_policy = data.aws_iam_policy_document.healthchecker-be-assume-policy.json
}

data "aws_iam_policy_document" "healthchecker-be-assume-policy" {
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

data "aws_iam_policy_document" "healthchecker-be" {
}

resource "aws_iam_policy" "healthchecker-be" {
  name   = "healthchecker-be"
  policy = data.aws_iam_policy_document.healthchecker-be.json
}

resource "aws_iam_policy_attachment" "healthchecker-be" {
  name       = "healthchecker-be"
  roles      = [aws_iam_role.healthchecker-be.name]
  policy_arn = aws_iam_policy.healthchecker-be.arn
}

