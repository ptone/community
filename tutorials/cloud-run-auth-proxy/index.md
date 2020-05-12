---
title: Deploying a basic authenticating proxy in Cloud Run
description: How to deploy a basic Identity Proxy in Cloud Run
author: ptone
tags: Cloud Run, identity
date_published: 2020-05-14
---

Preston Holmes | Solution Architect | Google

Compare with IAP or Pomerium


## Objectives

1. Build a container for a proxy service
1. Establish identity and access settings between the proxy and your protected service
1. Configure and deploy the authenticating proxy to Cloud Run


## Before you begin


## Setup



## Steps

```
export PROJECT=[your project-id]
export TARGET_SERVICE="[URL of target downstream service]"

export COOKIE_SECRET=$(head -c32 /dev/urandom | base64)

# You need to set at least one of these:
export VALID_DOMAINS="gmail.com"
export VALID_EMAILS="[some_other_google_email]"


gcloud builds submit --tag gcr.io/${PROJECT}/auth-proxy .
gcloud iam service-accounts create auth-proxy




gcloud run deploy auth-proxy \
--image gcr.io/${PROJECT}/auth-proxy \
--service-account auth-proxy@${PROJECT}.iam.gserviceaccount.com \
--allow-unauthenticated \
--project ${PROJECT} \
--set-env-vars TARGET_SERVICE="$TARGET_SERVICE" \
--set-env-vars COOKIE_SECRET="$COOKIE_SECRET" \
--set-env-vars "^|^VALID_DOMAINS=$VALID_DOMAINS" \
--set-env-vars "^|^VALID_EMAILS=$VALID_EMAILS" \
--set-env-vars OAUTH_CLIENT_ID=placeholder
```


Note the URL of the deployed proxy service


Configure your project for Google Signin with the wizard at this form:
https://developers.google.com/identity/sign-in/web/sign-in

Choose "web browser" and use the auth-proxy URL as the javascript origin.

When done - copy the OAuth Client ID provided and save it to an evironment variable.


export OAUTH_CLIENT_ID=[value from form]

gcloud run services update auth-proxy --update-env-vars=OAUTH_CLIENT_ID=$OAUTH_CLIENT_ID


## Cleaning up

