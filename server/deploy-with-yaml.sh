#!/bin/bash

source env-variables.sh

#Comment if correctly deployed
docker build -t gcr.io/$PROJECT_ID/recipe-server:latest .
docker push gcr.io/$PROJECT_ID/recipe-server:latest

#Uncomment if json needs to be added to GCP 
# gcloud secrets create firebase-account-key --data-file=/mnt/c/own_dev/RecipesApp/server/config/account_key.json --project=recepies-6e7c0

#Add permission IAM
# gcloud projects add-iam-policy-binding recepies-6e7c0 \
#     --member="serviceAccount:service-988443547488@serverless-robot-prod.iam.gserviceaccount.com" \
#     --role="roles/artifactregistry.reader"


gcloud run deploy recipe-service \
  --image gcr.io/$PROJECT_ID/recipe-server:latest \
  --region $REGION \
  --platform managed \
  --set-env-vars REDIS_URL=$REDIS_URL,CLIENT_URL=$CLIENT_URL,FIREBASE_ACCOUNT_KEY=$FIREBASE_ACCOUNT_KEY







