use MIME::Lite;

my $msg = MIME::Lite->new(
         'From'     => 'yo@functions.ovh',
         'To'       => 'couderthomas@gmail.com',
         'Subject'  => 'sujet',
         'Type'     => 'text/html',
         'Data'     => 'heyhey',
    );
    print $msg->send('smtp', 'ssl0.ovh.net',
        AuthUser=>'yo@functions.ovh',AuthPass=>'houbahouba' );