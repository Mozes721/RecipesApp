#!/bin/bash

#server 
export TF_VAR_redis_url="redis_url"
export TF_VAR_firebase_account_key="your_account_key.json"
export TF_VAR_client_url="client_url"
export TF_VAR_gcp_account_key="client_url"

echo "Environment variables for Terraform GCP set."

#client
