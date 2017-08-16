use Redis;

sub get { 
	my ($event) = @_;
	$r = new_redis_client();
	
	my $resp = $r->get( $event->{'data'}{'key'} );
	 
	$r->quit;

	return $resp;
}

sub set { 
	my ($event) = @_;
	$r = new_redis_client();

	my $resp = $r->set( $event->{'data'}{'key'} => $event->{'data'}{'value'} );
	 
	$r->quit;

	return $resp;
}

sub new_redis_client{
	my $host = $ENV{'REDIS_URL'};

	my $r = Redis->new (server => $host, encoding => undef) || die "Redis connection failed";
	$r->auth ($ENV{'REDIS_PASSWORD'});

	return $r;
}
1;