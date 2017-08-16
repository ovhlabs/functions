deploy:
	cd hipchat    && funk-cli deploy
	cd kafka-http && funk-cli deploy
	cd redis      && funk-cli deploy

exec: exec-hipchat exec-kafka-http exec-redis

exec-hipchat:
	echo '{"message":"'$$(date)'"}' | \
		funk-cli exec hipchat.notify

exec-kafka-http:
	echo '{"message":"'$$(date)'"}' | \
		funk-cli exec kafka.pub

exec-redis:
	echo '{"key":"date","value":"'$$(date)'"}' | funk-cli exec redis.set
	echo '{"key":"date"}' | funk-cli exec redis.get