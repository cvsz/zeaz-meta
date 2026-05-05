variable "name" { type = string }
variable "vpc_id" { type = string }
variable "subnet_ids" { type = list(string) }
variable "postgres_instance_class" { type = string default = "db.r7g.large" }

resource "aws_kms_key" "data" {
  description         = "${var.name} data plane encryption"
  enable_key_rotation = true
}

resource "aws_db_subnet_group" "postgres" {
  name       = "${var.name}-postgres"
  subnet_ids = var.subnet_ids
}

resource "aws_rds_cluster" "postgres" {
  cluster_identifier      = "${var.name}-postgres"
  engine                  = "aurora-postgresql"
  engine_version          = "16.3"
  database_name           = "zeaz"
  master_username         = "zeaz_admin"
  manage_master_user_password = true
  db_subnet_group_name    = aws_db_subnet_group.postgres.name
  storage_encrypted       = true
  kms_key_id              = aws_kms_key.data.arn
  backup_retention_period = 35
  preferred_backup_window = "03:00-04:00"
  deletion_protection     = true
  copy_tags_to_snapshot   = true
  enabled_cloudwatch_logs_exports = ["postgresql"]
}

resource "aws_rds_cluster_instance" "postgres" {
  count              = 3
  identifier         = "${var.name}-postgres-${count.index}"
  cluster_identifier = aws_rds_cluster.postgres.id
  instance_class     = var.postgres_instance_class
  engine             = aws_rds_cluster.postgres.engine
  engine_version     = aws_rds_cluster.postgres.engine_version
}

resource "aws_elasticache_subnet_group" "redis" {
  name       = "${var.name}-redis"
  subnet_ids = var.subnet_ids
}

resource "aws_elasticache_replication_group" "redis" {
  replication_group_id       = "${var.name}-redis"
  description                = "ZeaZ Redis HA cache and rate-limit store"
  engine                     = "redis"
  engine_version             = "7.1"
  node_type                  = "cache.r7g.large"
  port                       = 6379
  transit_encryption_enabled = true
  at_rest_encryption_enabled = true
  kms_key_id                 = aws_kms_key.data.arn
  automatic_failover_enabled = true
  multi_az_enabled           = true
  num_node_groups            = 3
  replicas_per_node_group    = 1
  subnet_group_name          = aws_elasticache_subnet_group.redis.name
}

output "postgres_endpoint" { value = aws_rds_cluster.postgres.endpoint }
output "redis_primary_endpoint" { value = aws_elasticache_replication_group.redis.primary_endpoint_address }
