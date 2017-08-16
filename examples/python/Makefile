deploy:
	cd hipchat    && funk-cli deploy
	cd kafka-http && funk-cli deploy
	cd mail       && funk-cli deploy
	cd redis      && funk-cli deploy

exec: exec-hipchat exec-kafka-http exec-mail exec-redis

exec-hipchat:
	echo '{"message":"'$$(date)'"}' | \
		funk-cli exec hipchat.notify

exec-kafka-http:
	echo $(date) | \
		funk-cli exec kafka.pub

exec-mail:
	echo '{"to":"hello@functions.ovh", "subject":"Time", "text":"'$$(date)'"}' | \
		funk-cli exec mail.send

exec-redis:
	echo '{"key":"date","value":"'$$(date)'"}' | funk-cli exec redis.set
	echo '{"key":"date"}' | funk-cli exec redis.get
