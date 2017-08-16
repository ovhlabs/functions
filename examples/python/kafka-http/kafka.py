import datetime
import os
import requests

def pub(event):
  topic = os.environ['KAFKA_TOPIC']
  host = os.environ['KAFKA_HOST']
  user = os.environ['KAFKA_USER']
  password = os.environ['KAFKA_PASSWORD']

  if event['data'] != "":
    message = event['data']
  else:
    raise Exception("data not found")

  start = datetime.datetime.now()
  uri = host + '/topic/' + topic + '?format=raw'
  r = requests.post(uri, 
    auth = (user, password), 
    json = message 
  )

  print("Kafka production status:", r.status_code, r.text, "in", datetime.datetime.now() - start)

  if r.status_code != 200:
    raise Exception(r.text)

  return r.text