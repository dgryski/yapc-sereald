#!/usr/bin/perl

use MIME::Base64;
use Sereal::Encoder qw(encode_sereal);
use LWP::UserAgent;
my $ua = LWP::UserAgent->new();

my $data = { h => [1,2,3,4] };
my $e = encode_base64(encode_sereal($data));
my $r = $ua->post('http://localhost:8080/', { sereal => $e } );
print $r->decoded_content;
