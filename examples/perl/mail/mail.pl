use Mail::Sendmail;

	$event->{'data'}{'to'} = "couderthomas\@gmail.com";
	$event->{'data'}{'subject'} = "coucou";
	$event->{'data'}{'text'} = "rololo";
	my $from = "yo\@functions.ovh";

	print $event->{'data'}{'to'};

	%mail = ( 
		'Smtp'    => "yo%40functions.ovh:houbahouba\@ssl0.ovh.net",
		'To'      => $event->{'data'}{'to'},
        'From'    => $from,
        'Subject' => $event->{'data'}{'subject'},
        'Message' => $event->{'data'}{'text'}
    );

	sendmail(%mail) or die $Mail::Sendmail::error;
	print $Mail::Sendmail::log."\n";