import base64
import json
import os
from sendgrid import SendGridAPIClient
from sendgrid.helpers.mail import Mail

def send_budget(event, context):
    """Background Cloud Function to be triggered by Pub/Sub.
    Args:
         event (dict):  The dictionary with data specific to this type of
         event. The `data` field contains the PubsubMessage message. The
         `attributes` field will contain custom attributes if there are any.
         context (google.cloud.functions.Context): The Cloud Functions event
         metadata. The `event_id` field contains the Pub/Sub message ID. The
         `timestamp` field contains the publish time.
    """

    print("""This Function was triggered by messageId {} published at {}
    """.format(context.event_id, context.timestamp))

    if 'data' in event:
        payload = base64.b64decode(event['data']).decode('utf-8')
    else:
        payload = 'empty'
    print(payload)
    alert = json.loads(payload)

    if 'alertThresholdExceeded' in alert:
        message = Mail(
            from_email=os.environ.get('FROM_EMAIL'),
            to_emails=os.environ.get('TO_EMAIL'),
            subject='Billing Alert {}'.format(alert['budgetDisplayName']),
            html_content=payload)
        try:
            sg = SendGridAPIClient(os.environ.get('SENDGRID_API_KEY'))
            response = sg.send(message)
            if response.status_code > 299:
                print("ERROR ", response.body)
            else:
                print("Mail Sent")
        except Exception as e:
            print("ERROR ", e.message)
    else:
        print('alert notice only, no threshold')
    return 'ok'