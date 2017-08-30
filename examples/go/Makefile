deploy:
	cd hipchat && ovh-functions deploy
	cd kafka-http   && ovh-functions deploy
	cd kafka   && ovh-functions deploy
	cd mail    && ovh-functions deploy
	cd redis   && ovh-functions deploy

exec: exec-hipchat exec-kafka-http exec-kafka exec-mail exec-redis

exec-hipchat:
	echo '{"message":"'$$(date)'"}' | \
		ovh-functions exec hipchat.notify

exec-kafka-http:
	echo $(date) | \
		ovh-functions exec kafka-http.pub

exec-kafka:
	echo $(date) | \
		ovh-functions exec kafka.pub

exec-mail:
	echo '{"to":["hello@functions.ovh"], "subject":"Time", "text":"'$$(date)'"}' | \
		ovh-functions exec mail.send

exec-redis:
	echo '{"k":"date","v":"'$$(date)'"}' | ovh-functions exec redis.set
	echo '{"k":"date"}' | ovh-functions exec redis.get