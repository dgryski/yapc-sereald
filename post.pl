#!/usr/bin/perl

use MIME::Base64;
use Sereal::Encoder qw(encode_sereal);
use LWP::UserAgent;
my $ua = LWP::UserAgent->new();
$ua->env_proxy(1);

my $data = { h => [1,2,3,4] };
my $e = encode_base64(encode_sereal($data));
my $r = $ua->post('http://yapc-sereald.herokuapp.com/', { sereal => $e } );
print $r->decoded_content;
