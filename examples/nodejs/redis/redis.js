var redis = require('redis')

exports.set = function(event, callback) {
  url = process.env.REDIS_URL

  if (!event.data.k) return callback('k is missing')
  if (!event.data.v) return callback('v is missing')

  var client = newRedisClient(url)

  client.set(event.data.k, event.data.v, function (err, reply) {
    if (err) return callback(err)
    client.quit()
    callback(null, reply.toString())
  })
}

exports.get = function(event, callback) {
  url = process.env.REDIS_URL

  if (!event.data.k) return callback('k is missing')

  var client = newRedisClient(url)

  client.get(event.data.k, function(err, reply) {
    if (err) return callback(err)
    if (!reply) return callback('key does not exist')
    client.quit()
    callback(null, reply.toString())
  })
}

function newRedisClient(url) {
  var client = redis.createClient({
    url: url
  })
  client.on('error', function (err) {
    console.log('error', err)
    callback('Internal error')
  })
  return client
}