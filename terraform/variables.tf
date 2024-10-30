variable "project_id" {
  description = "The ID of the Google Cloud project."
  type        = string
}

variable "project_name" {
  description = "The project name of the Google Cloud Run project."
  type        = string
}


variable "region" {
  description = "The Google Cloud region."
  type        = string
}


variable "redis_url" {
  description = "The URL for the Redis instance."
  type        = string
}


variable "client_url" {
  description = "The URL for the client application."
  type        = string
}

variable "gcp_account_key" {
  description = "Path to the Google Cloud service account key file."
  type        = string
}

variable "firebase_account_key_location" {
  description = "Firebase account key location in Docker container."
  type        = string
}