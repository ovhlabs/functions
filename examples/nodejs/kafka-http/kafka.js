var request = require('request')

exports.pub = function(event, callback) {
  var topic = process.env.KAFKA_TOPIC
  var host = process.env.KAFKA_HOST
  var user = process.env.KAFKA_USER
  var password = process.env.KAFKA_PASSWORD

  if (!event.data) return callback('data not found')

  var uri = host+'/topic/'+topic+'?format=raw'
  var req = {
    method: 'POST',
    auth: {
      user: user,
      password: password
    },
    json: event.data,
  }

  request.post(uri, req, function (error, resp) {
      if (error) return callback(error)
      callback(null, JSON.stringify(resp.body))
    })
}
