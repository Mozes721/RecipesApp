variable "project_id" {
  description = "Google Cloud project ID"
  type        = string
}

variable "region" {
  description = "Google Cloud region"
  type        = string
}

variable "redis_url" {
  description = "URL for Redis instance"
  type        = string
}

variable "client_url" {
  description = "URL for the client application"
  type        = string
}

variable "cloudrun_image" {
  description = "Cloud Run Artifacts Image"
  type        = string
}


variable "firebase_account_key_location" {
  description = "Firebase account key location in Docker container."
  type        = string
}