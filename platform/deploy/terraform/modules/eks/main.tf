variable "name" { type = string }
variable "region" { type = string }

resource "aws_kms_key" "cluster" {
  description             = "${var.name} envelope encryption"
  deletion_window_in_days = 30
  enable_key_rotation     = true
}

resource "aws_cloudwatch_log_group" "cluster" {
  name              = "/aws/eks/${var.name}/cluster"
  retention_in_days = 365
  kms_key_id        = aws_kms_key.cluster.arn
}

output "ingress_hostname" {
  value = "${var.name}.elb.${var.region}.amazonaws.com"
}
