# OVH Functions examples

Go to [docs.functions.ovh](https://docs.functions.ovh/) if you don't have the ovh-functions cli yet.

## Environment variables

These examples are using environment variables.
See corresponding functions.yml config to know which ones.

## Examples

### Redis SET/GET

Set and get Redis key/value.
[code](nodejs/redis)

```
echo '{"k":"date","v":"'$(date)'"}' | ovh-functions exec redis.set
OK

echo '{"k":"date"}' | ovh-functions exec redis.get
vendredi 21 avril 2017, 01:23:36 (UTC+0200)
```

### Mail

Send an email.
[code](nodejs/mail)

```
echo '{"to":"<email>", "subject":"Time", "text":"'$(date)'"}' | \
  ovh-functions exec mail.send
{"accepted":["<email>"],"rejected":[],"response":"250 2.0.0 OK ..."}
```

### Kafka HTTPS

Produce a message in Kafka using the OVH Kafka HTTPS proxy.
[code](nodejs/kafka)

```
echo '{"message":"'$(date)'"}' | \
  ovh-functions exec kafka.pub
[{"Value":"\"vendredi 21 avril 2017, 01:13:49 (UTC+0200)\""}]

```

### Hipchat

Send an Hipchat notification.
[code](nodejs/hipchat)

```
echo '{"message":"'$(date)'"}' | \
  ovh-functions exec hipchat.notify
{"statusCode":204,"location":"https://api.hipchat.com/v2/room/123456/history/7a254146-...}
```
