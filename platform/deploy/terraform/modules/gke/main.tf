variable "name" { type = string }
variable "project_id" { type = string }
variable "region" { type = string }
variable "network_cidr" { type = string default = "10.52.0.0/16" }

resource "google_compute_network" "this" {
  name                    = var.name
  auto_create_subnetworks = false
  project                 = var.project_id
}

resource "google_compute_subnetwork" "private" {
  name          = "${var.name}-private"
  project       = var.project_id
  region        = var.region
  network       = google_compute_network.this.id
  ip_cidr_range = var.network_cidr

  secondary_ip_range {
    range_name    = "pods"
    ip_cidr_range = "10.54.0.0/16"
  }
  secondary_ip_range {
    range_name    = "services"
    ip_cidr_range = "10.55.0.0/20"
  }
}

resource "google_container_cluster" "this" {
  name     = var.name
  project  = var.project_id
  location = var.region

  network    = google_compute_network.this.id
  subnetwork = google_compute_subnetwork.private.id

  remove_default_node_pool = true
  initial_node_count       = 1
  deletion_protection      = true
  networking_mode          = "VPC_NATIVE"

  private_cluster_config {
    enable_private_nodes    = true
    enable_private_endpoint = true
    master_ipv4_cidr_block  = "172.16.0.0/28"
  }

  ip_allocation_policy {
    cluster_secondary_range_name  = "pods"
    services_secondary_range_name = "services"
  }

  workload_identity_config { workload_pool = "${var.project_id}.svc.id.goog" }

  master_authorized_networks_config {}

  addons_config {
    http_load_balancing { disabled = false }
    horizontal_pod_autoscaling { disabled = false }
    gcp_filestore_csi_driver_config { enabled = true }
  }
}

resource "google_container_node_pool" "workloads" {
  name       = "workloads"
  project    = var.project_id
  location   = var.region
  cluster    = google_container_cluster.this.name
  node_count = 3

  autoscaling {
    min_node_count = 3
    max_node_count = 40
  }

  node_config {
    machine_type = "e2-standard-8"
    oauth_scopes = ["https://www.googleapis.com/auth/cloud-platform"]
    labels       = { pool = "workloads" }
    shielded_instance_config {
      enable_secure_boot          = true
      enable_integrity_monitoring = true
    }
  }
}

output "cluster_name" { value = google_container_cluster.this.name }
output "network_id" { value = google_compute_network.this.id }
