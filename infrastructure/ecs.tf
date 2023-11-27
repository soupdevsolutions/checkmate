resource "aws_ecs_cluster" "healthchecker-be" {
  name = "healthchecker-be"

  setting {
    name  = "containerInsights"
    value = "enabled"
  }
}

resource "aws_ecs_task_definition" "healthchecker-be" {
  family = "healthchecker-be"
  container_definitions = jsonencode([
    {
      name      = docker_registry_image.healthchecker-be-handler.name
      image     = docker_registry_image.healthchecker-be-handler.name
      cpu       = 10
      memory    = 512
      essential = true
      portMappings = [
        {
          containerPort = 80
          hostPort      = 80
        }
      ]
    },
  ])

  placement_constraints {
    type       = "memberOf"
  }
}

resource "aws_ecs_service" "healthchecker-be" {
  name            = "healthchecker-be"
  cluster         = aws_ecs_cluster.healthchecker-be.id
  task_definition = aws_ecs_task_definition.healthchecker-be.arn
  desired_count   = 3
  iam_role        = aws_iam_role.healthchecker-be.arn
  depends_on      = [aws_iam_policy.healthchecker-be]

  ordered_placement_strategy {
    type  = "binpack"
    field = "cpu"
  }

  placement_constraints {
    type       = "memberOf"
  }
}