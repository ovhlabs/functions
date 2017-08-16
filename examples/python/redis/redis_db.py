import os
import redis

def set(event):
	r = new_redis_client()

	resp = r.set(event['data']['key'], event['data']['value'])

	return str(resp)

def get(event):
	r = new_redis_client()
	
	resp = r.get(event['data']['key'])

	return str(resp)

def new_redis_client():
	host = os.environ['REDIS_URL']
	port = os.environ['REDIS_PORT']
	db = os.environ['REDIS_DB']
	password = os.environ['REDIS_PASSWORD']

	return redis.StrictRedis(host=host, port=port, db=db, password=password)