resource "google_project_service" "container_registry_api" {
  project = var.project_id
  service = "containerregistry.googleapis.com"
  disable_on_destroy = false
}

resource "null_resource" "docker_build_push" {
  triggers = {
    always_run = timestamp()
  }

  provisioner "local-exec" {
    command = <<-EOT
      # Build the Docker image
      docker build -t gcr.io/${var.project_id}/recipe-server:latest .
      
      # Configure docker to authenticate with GCP
      gcloud auth configure-docker --quiet
      
      # Push the image
      docker push gcr.io/${var.project_id}/recipe-server:latest
    EOT
  }

  depends_on = [
    google_project_service.container_registry_api
  ]
}
