deploy:
	cd hipchat    && ovh-functions deploy
	cd kafka-http && ovh-functions deploy
	cd mail       && ovh-functions deploy
	cd redis      && ovh-functions deploy

exec: exec-hipchat exec-kafka-http exec-mail exec-redis

exec-hipchat:
	echo '{"message":"'$$(date)'"}' | \
		ovh-functions exec hipchat.notify

exec-kafka-http:
	echo $(date) | \
		ovh-functions exec kafka.pub

exec-mail:
	echo '{"to":"hello@functions.ovh", "subject":"Time", "text":"'$$(date)'"}' | \
		ovh-functions exec mail.send

exec-redis:
	echo '{"key":"date","value":"'$$(date)'"}' | ovh-functions exec redis.set
	echo '{"key":"date"}' | ovh-functions exec redis.get
