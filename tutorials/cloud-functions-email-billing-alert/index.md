---
title: Sending custom email on Cloud Billing Alerts
description: Combine Cloud Billing alert pubsub notifications with Cloud Functions and SendGrid.
author: ptone
tags: Cloud Functions, billing, email, sendgrid
date_published: 2019-11-15
---

Preston Holmes | Solution Architect | Google

## Introduction

This tutorial demonstrates how to send an email using [SendGrid](https://sendgrid.com/) from a Cloud Function that is invoked by a Billing [Budget Alert](https://cloud.google.com/billing/docs/how-to/budgets) [Cloud Pub/Sub](https://cloud.google.com/pubsub/) message.

## Objectives

* Set up a Budget Alert and configure it to use Pub/Sub notifications.
* Deploy a Cloud Function which sends an email using the SendGrid API.


## Set up your environment

1.  Create a project in the [GCP Console][console]. 
1.  [Enable billing for your project](https://cloud.google.com/billing/docs/how-to/modify-project).
1.  Open [Cloud Shell][shell], which is a command-line interface built into the GCP Console that handles many
    environment setup tasks for you.

    If you prefer to use the Cloud SDK instead of Cloud Shell, you can install the [Cloud SDK][sdk] and run commands
    from the local command line.
    
1.  Set environment variables:

        # this is automatic in Cloud Shell
        gcloud config set project [ your project id ]

        export REGION=us-central1
        export GOOGLE_CLOUD_PROJECT=$(gcloud config list project --format "value(core.project)" )

1.  Enable APIs:

        gcloud services enable \
        cloudfunctions.googleapis.com \
        compute.googleapis.com \
        servicenetworking.googleapis.com \
        vpcaccess.googleapis.com

    Enabling APIs may take a moment.

1.  Clone the tutorial code:

        git clone https://github.com/GoogleCloudPlatform/community.git
        cd community/tutorials/cloud-functions-rate-limiting

## Register with SendGrid and get your API Key

## Create the Pub/Sub topic

create a topic named budget-alerts

## Create the Budget Alert

## Deploy the function

```
gcloud functions deploy send_budget --runtime python37 --trigger-topic budget-alerts --allow-unauthenticated --set-env-vars=SENDGRID_API_KEY=$SENDGRID_API_KEY,FROM_EMAIL=admin@example.com,TO_EMAIL=team@example.com
```

### Sample pubsub message

```
  {
    "ackId": "TgQhIT4wP ... tDCypYEQ",
    "message": {
      "attributes": {
        "billingAccountId": "005196-*******-7D3824",
        "budgetId": "a6139a72-0bcd-44f9-58a543d6a5ca",
        "schemaVersion": "1.0"
      },
      "data": "ewogICJidWRnZXREaXNwb ... iOiAiVVNEIgp9",
      "messageId": "830009655068937",
      "publishTime": "2019-11-08T21:01:15.919Z"
    }
  }
```

### Sample Payload

```
{
  "budgetDisplayName": "test-budget",
  "costAmount": 0.0,
  "costIntervalStart": "2019-11-01T07:00:00Z",
  "budgetAmount": 500.0,
  "budgetAmountType": "SPECIFIED_AMOUNT",
  "currencyCode": "USD"
}
```

## Clean up

The simplest way to clean up the resources used in the tutorial is to delete the project that you created just for this 
tutorial. 


[console]: https://console.cloud.google.com/
[shell]: https://cloud.google.com/shell/
[sdk]: https://cloud.google.com/sdk/
