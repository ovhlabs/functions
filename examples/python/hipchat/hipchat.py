import datetime
import os
import requests

def notify(event):
  roomID = os.environ['HIPCHAT_ROOM_ID']
  authToken = os.environ['HIPCHAT_AUTH_TOKEN']

  try:
    message = event['data']['message']
  except:
    raise Exception("message not found")

  start = datetime.datetime.now()
  r = requests.post('https://api.hipchat.com/v2/room/' + roomID + '/notification?auth_token=' + authToken, 
    json = {
      'color': event['data'].get('color', 'green'),
      'message_format': 'html',
      'message': message
    })
  print("Hipchat notify status:", r.status_code, r.text, "in", datetime.datetime.now() - start)
  
  if r.status_code != 204:
    raise Exception("failed to send", message)

  return "message " + message + " sent"