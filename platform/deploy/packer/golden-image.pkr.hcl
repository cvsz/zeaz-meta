packer {
  required_plugins { amazon = { source = "github.com/hashicorp/amazon", version = ">= 1.2.8" } }
}

source "amazon-ebs" "zeaz" {
  region        = "us-east-1"
  instance_type = "t3.small"
  source_ami_filter { filters = { name = "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-*" } owners = ["099720109477"] most_recent = true }
  ssh_username  = "ubuntu"
  ami_name      = "zeaz-golden-{{timestamp}}"
  encrypt_boot  = true
}

build {
  sources = ["source.amazon-ebs.zeaz"]
  provisioner "shell" { inline = ["sudo apt-get update", "sudo apt-get install -y ca-certificates curl auditd", "sudo systemctl enable auditd"] }
}
