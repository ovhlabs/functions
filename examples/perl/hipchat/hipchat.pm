use JSON;
use LWP::UserAgent;

sub notify {
  my ($event)   = @_;
  my $apiUrl    = 'https://api.hipchat.com/v2/room/'
  my $roomID    = $ENV{'HIPCHAT_ROOM_ID'};
  my $authToken = $ENV{'HIPCHAT_AUTH_TOKEN'};
  my $message   = $event->{'data'}{'message'};
  my $color     = $event->{'data'}{'color'} || 'green';

  if ($message eq "") {
    die "message not found";
  }

  my $r   = LWP::UserAgent->new;
  my $req = HTTP::Request->new('POST' => $apiUrl . $roomID . '/notification?auth_token=' . $authToken);
  $req->header( 'Content-Type' => 'application/json' );

  my $content = {
    'color' => $color,
    'message_format' => 'html',
    'message' => $message
  };

  $req->content(encode_json($content));

  my $resp = $r->request($req);
  if ($resp->is_success) {
    return $resp->decoded_content;
  } else {
    die $resp->decoded_content;
  }
}
1;