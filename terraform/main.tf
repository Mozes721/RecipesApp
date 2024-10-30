terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = ">= 4.0.0"
    }
  }
  required_version = ">= 0.12"
}

provider "google" {
  credentials = file(var.gcp_account_key)
  project     = var.project_id
  region      = var.region
}

# Get project information
data "google_project" "project" {
  project_id = var.project_id
}


module "docker" {
  source      = "./modules/docker"
  project_id  = var.project_id
}

module "cloud_run" {
  source      = "./modules/gcp"
  project_id  = var.project_id
  region      = var.region
  redis_url   = var.redis_url
  client_url  = var.client_url
  firebase_account_key_location = var.firebase_account_key_location
  cloudrun_image = "gcr.io/${var.project_id}/recipe-server:latest"

  depends_on = [
    module.docker
  ]
}