var request = require('request');

exports.notify = function(event, callback) {
  var roomID = process.env.HIPCHAT_ROOM_ID
  var authToken = process.env.HIPCHAT_AUTH_TOKEN

  if (!event.data.message) return callback(null, "message not found")

  var req = {
    method: 'POST',
    uri: 'https://api.hipchat.com/v2/room/'+roomID+'/notification?auth_token='+authToken,
    json: {
      color: event.data.color || 'green',
      message_format: 'html',
      message: event.data.message
    }
  }

  request(req, function (error, resp) {
    if (error) return callback(error)
    callback(null, JSON.stringify(resp))
  })
}