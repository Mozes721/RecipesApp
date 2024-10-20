provider "google" {
  credentials = var.gcp_account_key
  project     = var.project_id
  region      = var.region
}

# Google Cloud Run Service
resource "google_cloud_run_service" "recipe_service" {
  name     = var.project_name
  location = var.region

  template {
    spec {
      containers {
        image = "gcr.io/${var.project_id}/recipe-server:latest"

        env {
          name  = "REDIS_URL"
          value = var.redis_url
        }
        env {
          name  = "CLIENT_URL"
          value = var.client_url
        }
        env {
          name  = "FIREBASE_ACCOUNT_KEY"
          value = var.firebase_account_key
        }

        resources {
          limits = {
            cpu    = "1000m"
            memory = "512Mi"
          }
          requests = {
            cpu    = "500m"
            memory = "256Mi"
          }
        }
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}
