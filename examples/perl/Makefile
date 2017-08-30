deploy:
	cd hipchat    && ovh-functions deploy
	cd kafka-http && ovh-functions deploy
	cd redis      && ovh-functions deploy

exec: exec-hipchat exec-kafka-http exec-redis

exec-hipchat:
	echo '{"message":"'$$(date)'"}' | \
		ovh-functions exec hipchat.notify

exec-kafka-http:
	echo '{"message":"'$$(date)'"}' | \
		ovh-functions exec kafka.pub

exec-redis:
	echo '{"key":"date","value":"'$$(date)'"}' | ovh-functions exec redis.set
	echo '{"key":"date"}' | ovh-functions exec redis.get