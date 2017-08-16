import os
import smtplib

from email.mime.text import MIMEText

def send(event):
  url = os.environ['MAIL_SMTP_URL']
  user = os.environ['MAIL_SMTP_USER']
  password = os.environ['MAIL_SMTP_PASSWORD']

  try:
    msg = MIMEText(event['data']['text'])
  except:
    raise Exception('no message found')

  msg['Subject'] = event['data']['subject']
  msg['From'] = user
  msg['To'] = event['data']['to']

  s = smtplib.SMTP(url)
  s.login(user, password)
  s.sendmail(user, event['data']['to'], msg.as_string())
  s.quit()

  return "mail sent to " + event['data']['to']