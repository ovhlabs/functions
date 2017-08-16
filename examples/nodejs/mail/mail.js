var nodemailer = require('nodemailer');

exports.send = function(event, callback) {
  var url = process.env.MAIL_SMTP_URL
  var from = process.env.MAIL_FROM

  if (!event.data.to) return callback('email recipient ("to") not found')
  if (!event.data.subject) return callback('subject not found')
  if (!event.data.text) return callback('text not found')

  var mail = {
    from: from,
    to: event.data.to,
    subject: event.data.subject,
    text: event.data.text
  }

  nodemailer
    .createTransport(url)
    .sendMail(mail, function(error, resp) {
      if (error) return callback(error)
      callback(null, JSON.stringify(resp))
  });

}