# Programming Assignment Solution

## Overview

- As per requirement, crypto/tls library is used to get the tls information of the URL passed.
- TLS info is parsed into format as per requirements and printed in stdout.
- When URL is not passed as CLI argument, error is printed and returned.
- When URL host is invalid, error is printed and returned.

## Run and Test

```
# Build
$ CGO_ENABLED=0 go build -trimpath -ldflags "-s -w"

# Run and Test Success Response
$ ./tls-info.exe www.tintin.com:443
Issuer: CN=R3,O=Let's Encrypt,C=US
Expiry Date: 2023-09-29 03:26:37 +0000 UTC
Common Name: R3
Issuer: CN=ISRG Root X1,O=Internet Security Research Group,C=US
Expiry Date: 2025-09-15 16:00:00 +0000 UTC
Common Name: ISRG Root X1
Issuer: CN=DST Root CA X3,O=Digital Signature Trust Co.
Expiry Date: 2024-09-30 18:14:03 +0000 UTC
Common Name: DST Root CA X3

# Run and Test URL not passed failure scenario
$ ./tls-info.exe
2023/08/05 22:07:14 Please Pass the URL

# Run and Test invalid host failure scenario
$ ./tls-info.exe www.testerorr.com:443
2023/08/05 22:13:38 Error in Dial dial tcp: lookup www.testerorr.com: no such host

# Run and Test unreachable host failure scenario
$ ./tls-info.exe www.testerror.com:443
2023/08/05 22:13:09 Error in Dial dial tcp 3.18.7.81:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.
```
