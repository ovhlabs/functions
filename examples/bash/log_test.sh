function hello() {
  >&2 echo "$1"

  declare log_token=$1
  echo -e '{"version":"1.1", "host": "example.org", "short_message": "A short GELF message that helps you identify what is going on '$(hostname)'", "full_message": "Backtrace here more stuff", "timestamp": '$(date +'%s')', "level": 1, "_user_id": 9001, "_some_info": "foo", "some_metric_num": 42.0, "_X-OVH-TOKEN":"'$log_token'"}\0'  | openssl s_client -quiet -no_ign_eof  -connect discover.logs.ovh.com:12202
  return 0
}
