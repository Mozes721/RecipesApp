resource "google_project_service" "required_apis" {
  for_each = toset([
    "run.googleapis.com",
    "containerregistry.googleapis.com"
  ])
  
  project = var.project_id
  service = each.key
  disable_on_destroy = false
}

resource "google_cloud_run_service" "recipe_service" {
  name     = "recipe-service"
  location = var.region
  project  = var.project_id

  template {
    spec {
      containers {
        image = var.cloudrun_image
        
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
          value = var.firebase_account_key_location
        }
      }
    }
  }

  depends_on = [
    google_project_service.required_apis
  ]
}

resource "google_cloud_run_service_iam_member" "public_access" {
  location = google_cloud_run_service.recipe_service.location
  project  = google_cloud_run_service.recipe_service.project
  service  = google_cloud_run_service.recipe_service.name
  role     = "roles/run.invoker"
  member   = "allUsers"
}