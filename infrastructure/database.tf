resource "aws_rds_cluster" "healthchecker" {
  cluster_identifier = "healthchecker"
  engine             = "aurora-postgresql"
  engine_mode        = "provisioned"
  engine_version     = "15.2"
  database_name      = "healthchecker"
  master_username    = var.RDS_USERNAME
  master_password    = var.RDS_PASSWORD

  serverlessv2_scaling_configuration {
    min_capacity = 0.5
    max_capacity = 1.0
  }
}

resource "aws_rds_cluster_instance" "healthchecker" {
  cluster_identifier = aws_rds_cluster.healthchecker.id
  instance_class     = "db.serverless"
  engine             = aws_rds_cluster.healthchecker.engine
  engine_version     = aws_rds_cluster.healthchecker.engine_version
}
