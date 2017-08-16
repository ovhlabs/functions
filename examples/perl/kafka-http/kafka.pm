use LWP::UserAgent;
use Time::HiRes qw(time);

sub pub {
  my ($event)  = @_;
  my $topic    = $ENV{'KAFKA_TOPIC'};
  my $host     = $ENV{'KAFKA_HOST'};
  my $user     = $ENV{'KAFKA_USER'};
  my $password = $ENV{'KAFKA_PASSWORD'};
  my $message  = $event->{'data'};
  
  if ($message eq "") {
    die "data not found";
  }
  
  my $req   = HTTP::Request->new('POST' => $host . '/topic/' . $topic . '?format=raw');
  $req->authorization_basic($user, $password);
  $req->content($message);

  my $r = LWP::UserAgent->new;
  my $start = time();
  my $resp = $r->request($req);

  my $transactionTime = (time() - $start) * 1000;
  print "Kafka production status: " . $resp->code . " " . $resp->decoded_content . " in " . $transactionTime . " ms\n";
  if ($resp->is_success) {
    return $resp->decoded_content;
  } else {
    die $resp->decoded_content;
  }
}
1;