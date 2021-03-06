package lib

import (
	"fmt"
	"testing"
)

func TestParseCertificate(t *testing.T) {
	val := `
-----BEGIN CERTIFICATE-----
MIIKaDCCCVCgAwIBAgISAwA5Mwa8ZukaDrEU+7V9GEk4MA0GCSqGSIb3DQEBCwUA
MDIxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1MZXQncyBFbmNyeXB0MQswCQYDVQQD
EwJSMzAeFw0yMTA3MDUwMjUxNDFaFw0yMTEwMDMwMjUxNDBaMBMxETAPBgNVBAMT
CGJramsuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA0At5q1oI
/+8Ec1YRYyG7f/5GXl0U2PypDlYtWLRWVtR+5poyqf75siTNOSKtko4YhZ0FOvyd
U1k+ZiT72am0fyun8fLEAPy9vKslNEVr4kzOe/SUlIhYFxcxPbkPI2eRH3SeqzgD
fm2fzdLHVLwrf3ZMmKMzC+LsRhr+kL01maRpho7TirQhCJiFRiQ1QBSsK/4uREnw
Bth2nyMpHUnKerxsE+L0cnQ09/kA1IJpA7e+W8zA1z+cLz1fAmTXjcNP+g5lLrLB
U+q52dnpjJHLS70NmNpx6rFXWGMgpyj3qi95whvwV3PxD30ziMuaBnyR6Ljt8Lie
HA46gHFU37v2jQIDAQABo4IHlTCCB5EwDgYDVR0PAQH/BAQDAgWgMB0GA1UdJQQW
MBQGCCsGAQUFBwMBBggrBgEFBQcDAjAMBgNVHRMBAf8EAjAAMB0GA1UdDgQWBBS6
yuLmAWl0ws08lEIX7nKezkay2zAfBgNVHSMEGDAWgBQULrMXt1hWy65QCUDmH6+d
ixTCxjBVBggrBgEFBQcBAQRJMEcwIQYIKwYBBQUHMAGGFWh0dHA6Ly9yMy5vLmxl
bmNyLm9yZzAiBggrBgEFBQcwAoYWaHR0cDovL3IzLmkubGVuY3Iub3JnLzCCBWQG
A1UdEQSCBVswggVXghAqLmFubGliYW94aWFuLmNughEqLmFubGliYW94aWFuLmNv
bYIMKi5hbmxpYnguY29tgg4qLmJramstaW5jLmNvbYIJKi5ia2prLmNuggoqLmJr
amsuY29tgg0qLmJramsuY29tLmNughQqLmNsb3VkLmJramstaW5jLmNvbYIQKi5k
YWppYWppbmZ1LmNvbYIUKi5kZXYuYW5saWJhb3hpYW4uY26CFSouZGV2LmFubGli
YW94aWFuLmNvbYIPKi5kZXYuYW5saWJ4LmNughAqLmRldi5hbmxpYnguY29tghIq
LmRldi5ia2prLWluYy5jb22CDSouZGV2LmJramsuY26CDiouZGV2LmJramsuY29t
ghgqLmRldi5jbG91ZC5ia2prLWluYy5jb22CFCouZGV2LmRhamlhamluZnUuY29t
gg8qLmRldi5pbnNlcnMuY26CEyouZGV2Lmluc2Vycy5jb20uY26CGiouZGV2Lmxp
YW5qaWEuYmtqay1pbmMuY29tghUqLmRldi5saWFuamlhLmJramsuY26CFiouZGV2
LmxpYW5qaWEuYmtqay5jb22CGCouZGV2Lm9jZWFuLmJramstaW5jLmNvbYILKi5p
bnNlcnMuY26CDyouaW5zZXJzLmNvbS5jboIWKi5saWFuamlhLmJramstaW5jLmNv
bYISKi5saWFuamlhLmJramsuY29tghQqLm9jZWFuLmJramstaW5jLmNvbYIQKi5v
Y2Vhbi5ia2prLmNvbYISKi5wcmUuYmtqay1pbmMuY29tgg0qLnByZS5ia2prLmNu
gg4qLnByZS5ia2prLmNvbYIVKi5wcmUubGlhbmppYS5ia2prLmNughYqLnNhbmRi
b3guYmtqay1pbmMuY29tghIqLnNhbmRib3guYmtqay5jb22CFiouc3RhZ2UuYW5s
aWJhb3hpYW4uY26CFyouc3RhZ2UuYW5saWJhb3hpYW4uY29tghEqLnN0YWdlLmFu
bGlieC5jboISKi5zdGFnZS5hbmxpYnguY29tghQqLnN0YWdlLmJramstaW5jLmNv
bYIPKi5zdGFnZS5ia2prLmNughAqLnN0YWdlLmJramsuY29tghYqLnN0YWdlLmRh
amlhamluZnUuY29tghEqLnN0YWdlLmluc2Vycy5jboIVKi5zdGFnZS5pbnNlcnMu
Y29tLmNughwqLnN0YWdlLmxpYW5qaWEuYmtqay1pbmMuY29tghcqLnN0YWdlLmxp
YW5qaWEuYmtqay5jboIYKi5zdGFnZS5saWFuamlhLmJramsuY29tghUqLnRlc3Qu
YW5saWJhb3hpYW4uY26CFioudGVzdC5hbmxpYmFveGlhbi5jb22CECoudGVzdC5h
bmxpYnguY26CESoudGVzdC5hbmxpYnguY29tghMqLnRlc3QuYmtqay1pbmMuY29t
gg4qLnRlc3QuYmtqay5jboIPKi50ZXN0LmJramsuY29tghUqLnRlc3QuZGFqaWFq
aW5mdS5jb22CECoudGVzdC5pbnNlcnMuY26CFCoudGVzdC5pbnNlcnMuY29tLmNu
ghsqLnRlc3QubGlhbmppYS5ia2prLWluYy5jb22CFioudGVzdC5saWFuamlhLmJr
amsuY26CFyoudGVzdC5saWFuamlhLmJramsuY29tghIqLnRlc3QuenViZWliay5j
b22CDSouenViZWliay5jb22CDmFubGliYW94aWFuLmNugg9hbmxpYmFveGlhbi5j
b22CCmFubGlieC5jb22CCGJramsuY29tgglpbnNlcnMuY24wTAYDVR0gBEUwQzAI
BgZngQwBAgEwNwYLKwYBBAGC3xMBAQEwKDAmBggrBgEFBQcCARYaaHR0cDovL2Nw
cy5sZXRzZW5jcnlwdC5vcmcwggEDBgorBgEEAdZ5AgQCBIH0BIHxAO8AdQCUILwe
jtWNbIhzH4KLIiwN0dpNXmxPlD1h204vWE2iwgAAAXp0ywpcAAAEAwBGMEQCIHcI
rDFdmG5+RcEJ35gcZD9HJdqnSZLGSUtp01ch8iTDAiBbx1EuQ8BjLi4o8Fsr4i9B
fq3maFGW1XvTcm790ggtpgB2AH0+8viP/4hVaCTCwMqeUol5K8UOeAl/LmqXaJl+
IvDXAAABenTLCpsAAAQDAEcwRQIhANA+ceClar3eiGjmOQZHW4kX5zsZTiUD887p
redNnbJFAiB+1B333bsVDpVP9nNEf1cgQr+uXV0TSuUr3IP0VDnSBzANBgkqhkiG
9w0BAQsFAAOCAQEABXEuxvRgzVNkhZyvaAS4+xGhnavpO/RexWvKWrwr/AEAQmNr
/SIevHmYGkheYtRhyXsl4o3e/d48JQn6iG/TaRk3NdVVXVz5EI7++bl/OxRKFSLo
qSGpVG/HxokS2yb5kw4om17FppHCf3iFQ2tqGYKqMu30jANWCGTJ7/ujcRvT0kT3
jqNANFV6fB9cgynMXg7L22ZboHnCLMwMdtaDRditlKgnhDDRSiKdUGdkltgOrqs4
0DnaLktTdGvgOYSiCNT2ffoZRgEpK5kiFsN8DcXyAr7ZCvh+/fAZO04v2IwlAdOg
1GJG/FIqM/TqS25Tt/A/b8iENZbcND1jEFuISg==
-----END CERTIFICATE-----

-----BEGIN CERTIFICATE-----
MIIFFjCCAv6gAwIBAgIRAJErCErPDBinU/bWLiWnX1owDQYJKoZIhvcNAQELBQAw
TzELMAkGA1UEBhMCVVMxKTAnBgNVBAoTIEludGVybmV0IFNlY3VyaXR5IFJlc2Vh
cmNoIEdyb3VwMRUwEwYDVQQDEwxJU1JHIFJvb3QgWDEwHhcNMjAwOTA0MDAwMDAw
WhcNMjUwOTE1MTYwMDAwWjAyMQswCQYDVQQGEwJVUzEWMBQGA1UEChMNTGV0J3Mg
RW5jcnlwdDELMAkGA1UEAxMCUjMwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEK
AoIBAQC7AhUozPaglNMPEuyNVZLD+ILxmaZ6QoinXSaqtSu5xUyxr45r+XXIo9cP
R5QUVTVXjJ6oojkZ9YI8QqlObvU7wy7bjcCwXPNZOOftz2nwWgsbvsCUJCWH+jdx
sxPnHKzhm+/b5DtFUkWWqcFTzjTIUu61ru2P3mBw4qVUq7ZtDpelQDRrK9O8Zutm
NHz6a4uPVymZ+DAXXbpyb/uBxa3Shlg9F8fnCbvxK/eG3MHacV3URuPMrSXBiLxg
Z3Vms/EY96Jc5lP/Ooi2R6X/ExjqmAl3P51T+c8B5fWmcBcUr2Ok/5mzk53cU6cG
/kiFHaFpriV1uxPMUgP17VGhi9sVAgMBAAGjggEIMIIBBDAOBgNVHQ8BAf8EBAMC
AYYwHQYDVR0lBBYwFAYIKwYBBQUHAwIGCCsGAQUFBwMBMBIGA1UdEwEB/wQIMAYB
Af8CAQAwHQYDVR0OBBYEFBQusxe3WFbLrlAJQOYfr52LFMLGMB8GA1UdIwQYMBaA
FHm0WeZ7tuXkAXOACIjIGlj26ZtuMDIGCCsGAQUFBwEBBCYwJDAiBggrBgEFBQcw
AoYWaHR0cDovL3gxLmkubGVuY3Iub3JnLzAnBgNVHR8EIDAeMBygGqAYhhZodHRw
Oi8veDEuYy5sZW5jci5vcmcvMCIGA1UdIAQbMBkwCAYGZ4EMAQIBMA0GCysGAQQB
gt8TAQEBMA0GCSqGSIb3DQEBCwUAA4ICAQCFyk5HPqP3hUSFvNVneLKYY611TR6W
PTNlclQtgaDqw+34IL9fzLdwALduO/ZelN7kIJ+m74uyA+eitRY8kc607TkC53wl
ikfmZW4/RvTZ8M6UK+5UzhK8jCdLuMGYL6KvzXGRSgi3yLgjewQtCPkIVz6D2QQz
CkcheAmCJ8MqyJu5zlzyZMjAvnnAT45tRAxekrsu94sQ4egdRCnbWSDtY7kh+BIm
lJNXoB1lBMEKIq4QDUOXoRgffuDghje1WrG9ML+Hbisq/yFOGwXD9RiX8F6sw6W4
avAuvDszue5L3sz85K+EC4Y/wFVDNvZo4TYXao6Z0f+lQKc0t8DQYzk1OXVu8rp2
yJMC6alLbBfODALZvYH7n7do1AZls4I9d1P4jnkDrQoxB3UqQ9hVl3LEKQ73xF1O
yK5GhDDX8oVfGKF5u+decIsH4YaTw7mP3GFxJSqv3+0lUFJoi5Lc5da149p90Ids
hCExroL1+7mryIkXPeFM5TgO9r0rvZaBFOvV2z0gp35Z0+L4WPlbuEjN/lxPFin+
HlUjr8gRsI3qfJOQFy/9rKIJR0Y/8Omwt/8oTWgy1mdeHmmjk7j1nYsvC9JSQ6Zv
MldlTTKB3zhThV1+XWYp6rjd5JW1zbVWEkLNxE7GJThEUG3szgBVGP7pSWTUTsqX
nLRbwHOoq7hHwg==
-----END CERTIFICATE-----

-----BEGIN CERTIFICATE-----
MIIFYDCCBEigAwIBAgIQQAF3ITfU6UK47naqPGQKtzANBgkqhkiG9w0BAQsFADA/
MSQwIgYDVQQKExtEaWdpdGFsIFNpZ25hdHVyZSBUcnVzdCBDby4xFzAVBgNVBAMT
DkRTVCBSb290IENBIFgzMB4XDTIxMDEyMDE5MTQwM1oXDTI0MDkzMDE4MTQwM1ow
TzELMAkGA1UEBhMCVVMxKTAnBgNVBAoTIEludGVybmV0IFNlY3VyaXR5IFJlc2Vh
cmNoIEdyb3VwMRUwEwYDVQQDEwxJU1JHIFJvb3QgWDEwggIiMA0GCSqGSIb3DQEB
AQUAA4ICDwAwggIKAoICAQCt6CRz9BQ385ueK1coHIe+3LffOJCMbjzmV6B493XC
ov71am72AE8o295ohmxEk7axY/0UEmu/H9LqMZshftEzPLpI9d1537O4/xLxIZpL
wYqGcWlKZmZsj348cL+tKSIG8+TA5oCu4kuPt5l+lAOf00eXfJlII1PoOK5PCm+D
LtFJV4yAdLbaL9A4jXsDcCEbdfIwPPqPrt3aY6vrFk/CjhFLfs8L6P+1dy70sntK
4EwSJQxwjQMpoOFTJOwT2e4ZvxCzSow/iaNhUd6shweU9GNx7C7ib1uYgeGJXDR5
bHbvO5BieebbpJovJsXQEOEO3tkQjhb7t/eo98flAgeYjzYIlefiN5YNNnWe+w5y
sR2bvAP5SQXYgd0FtCrWQemsAXaVCg/Y39W9Eh81LygXbNKYwagJZHduRze6zqxZ
Xmidf3LWicUGQSk+WT7dJvUkyRGnWqNMQB9GoZm1pzpRboY7nn1ypxIFeFntPlF4
FQsDj43QLwWyPntKHEtzBRL8xurgUBN8Q5N0s8p0544fAQjQMNRbcTa0B7rBMDBc
SLeCO5imfWCKoqMpgsy6vYMEG6KDA0Gh1gXxG8K28Kh8hjtGqEgqiNx2mna/H2ql
PRmP6zjzZN7IKw0KKP/32+IVQtQi0Cdd4Xn+GOdwiK1O5tmLOsbdJ1Fu/7xk9TND
TwIDAQABo4IBRjCCAUIwDwYDVR0TAQH/BAUwAwEB/zAOBgNVHQ8BAf8EBAMCAQYw
SwYIKwYBBQUHAQEEPzA9MDsGCCsGAQUFBzAChi9odHRwOi8vYXBwcy5pZGVudHJ1
c3QuY29tL3Jvb3RzL2RzdHJvb3RjYXgzLnA3YzAfBgNVHSMEGDAWgBTEp7Gkeyxx
+tvhS5B1/8QVYIWJEDBUBgNVHSAETTBLMAgGBmeBDAECATA/BgsrBgEEAYLfEwEB
ATAwMC4GCCsGAQUFBwIBFiJodHRwOi8vY3BzLnJvb3QteDEubGV0c2VuY3J5cHQu
b3JnMDwGA1UdHwQ1MDMwMaAvoC2GK2h0dHA6Ly9jcmwuaWRlbnRydXN0LmNvbS9E
U1RST09UQ0FYM0NSTC5jcmwwHQYDVR0OBBYEFHm0WeZ7tuXkAXOACIjIGlj26Ztu
MA0GCSqGSIb3DQEBCwUAA4IBAQAKcwBslm7/DlLQrt2M51oGrS+o44+/yQoDFVDC
5WxCu2+b9LRPwkSICHXM6webFGJueN7sJ7o5XPWioW5WlHAQU7G75K/QosMrAdSW
9MUgNTP52GE24HGNtLi1qoJFlcDyqSMo59ahy2cI2qBDLKobkx/J3vWraV0T9VuG
WCLKTVXkcGdtwlfFRjlBz4pYg1htmf5X6DYO8A4jqv2Il9DjXA6USbW1FzXSLr9O
he8Y4IWS6wY7bCkjCWDcRQJMEhg76fsO3txE+FiYruq9RUWhiF1myv4Q6W+CyBFC
Dfvp7OOGAN6dEOM4+qR9sdjoSYKEBpsr6GtPAQw4dy753ec5
-----END CERTIFICATE-----
`
	c := &CertChain{}
	err := c.ParseCertificate(val)
	fmt.Println(err)
	for _, line := range c.CertChain {
		fmt.Println(fmt.Sprintf("%+v", line))
		fmt.Println("-----------------------------------")
	}
}
