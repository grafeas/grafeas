// cert is an provides insecure keys.
package cert

import (
	"crypto/tls"
	"crypto/x509"
	"log"
)

const (

	// Generated using
	// openssl req -new -x509 -key server.key -out server.pem -days 3650
	cert = `-----BEGIN CERTIFICATE-----
MIIDwzCCAqugAwIBAgIJAPQosZG8QUelMA0GCSqGSIb3DQEBCwUAMHgxCzAJBgNV
BAYTAlVTMREwDwYDVQQIDAhOZXcgWW9yazERMA8GA1UEBwwITmV3IFlvcmsxEDAO
BgNVBAoMB2dyYWZlYXMxEjAQBgNVBAMMCWxvY2FsaG9zdDEdMBsGCSqGSIb3DQEJ
ARYOd21kQGdvb2dsZS5jb20wHhcNMTgwMTA1MjA0MzI1WhcNMjgwMTAzMjA0MzI1
WjB4MQswCQYDVQQGEwJVUzERMA8GA1UECAwITmV3IFlvcmsxETAPBgNVBAcMCE5l
dyBZb3JrMRAwDgYDVQQKDAdncmFmZWFzMRIwEAYDVQQDDAlsb2NhbGhvc3QxHTAb
BgkqhkiG9w0BCQEWDndtZEBnb29nbGUuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOC
AQ8AMIIBCgKCAQEAu2rXIEykaRZ+O+cf3BZS9Z8+bxoL7/J4kj2gcLg37NOD7Aw7
/lFch7rO0RN0oaba2U10d7CGG3f8g+mkGNiFKpl/SIm8LvJ8VeNHRqjJMiSFX+v5
g7eC+hYcxcctcX01qUNwp3e8rxpbZTVeQoFCE78o9hR1EDPNcJbFcXMSvN4w78md
+NREXVDnuruA/s6Xo42DuHT71VW0xyl5zv3pmDsPhLmCIHSJUc5SZUs8AijUYZrC
lhOYEaSKyWjrjPVKzuniWj8H0QvI3zDI8/fpffg/DOo+oI3UZ2eFRwINZTSUauhS
cknaKAe6z7uW+zFYBZYBWXnjJZ5Fp8i7DRlH+wIDAQABo1AwTjAdBgNVHQ4EFgQU
AdDHlyBTKH+jnVWlFe0wjZ4//v8wHwYDVR0jBBgwFoAUAdDHlyBTKH+jnVWlFe0w
jZ4//v8wDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAQEAbOXjxyfF+iTq
gMFopfHG+cB1bHPlIoyBkjJWMs9EXCkb25JL/v3cEvN6Ixo0b9jjvK9Azk4SlOZ+
Si6X1C+z0Er5ScpjeY8wn/+dRp3O1psEos+lriwc2MCVwPAcV2XQ5Z9my6wNssnG
6IAiHmG/IGVid66dijhCQHBb0V4mUREELACL1gXD5tASoKGAbheUU3dIX7sctH1m
W7zHLtnXW1cYrKDcqD+j8NNMYcjJGMzntgFJKuQF2Awwz9Gf88rhUskoOqpE/Bhf
dK1P1LrgZBiIjDIdFiZKAyDbxPbm8IbPJVd6hCf4GBO6xqVomuxUDx5Wvhz+KhgY
oeLwZ0lcbA==
-----END CERTIFICATE-----`

	// generated using
	// openssl genrsa -out server.key 2048
	sKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAu2rXIEykaRZ+O+cf3BZS9Z8+bxoL7/J4kj2gcLg37NOD7Aw7
/lFch7rO0RN0oaba2U10d7CGG3f8g+mkGNiFKpl/SIm8LvJ8VeNHRqjJMiSFX+v5
g7eC+hYcxcctcX01qUNwp3e8rxpbZTVeQoFCE78o9hR1EDPNcJbFcXMSvN4w78md
+NREXVDnuruA/s6Xo42DuHT71VW0xyl5zv3pmDsPhLmCIHSJUc5SZUs8AijUYZrC
lhOYEaSKyWjrjPVKzuniWj8H0QvI3zDI8/fpffg/DOo+oI3UZ2eFRwINZTSUauhS
cknaKAe6z7uW+zFYBZYBWXnjJZ5Fp8i7DRlH+wIDAQABAoIBACB1MmifnWGtyZLq
RjRBkYCEYbWwFx0pKwR4s86RuO3E+/XncIRs5s+C5MqEyhAs633y0hbgdXlQYGUg
E5FR/k4QY2DWqcafrDTbtb5hAOc0N/0SyxWqtH5HUhhWlGIxQxfbXClErWLN98Ih
af+ujxkIZDmp9VQnBI9ZLTympzoaHO3gY928ytQmRnHIqMOj2xB22Eyy2IKKUh38
JvgMSTRJfNkOyJJfkM4rgv6xC2/xBsRpZ0V84EsrmNUMo9FiCL+nf5HJs0385u+Y
iCG2dsdx5vzIOMBSDeuQi3HRM6ELQlu6gn7hpDnQdD5LGaztCe9pDX2Ylc7iexpJ
Npa5F3ECgYEA6VoEfvPYy/n+KLHWayAdL6s2jwel7Emg9VpfEZk8KR7PTGtV9poq
NFc/yHDIJTs35Wo81cj8uAlrPAfNlITOz+Yx+Y3/MTQ7AgFk7xZlmd9uhR3CSaaX
QsTTuImVtK9c8DSi8Gq2DOFFKfH3mUHj+WY6MNtAULMEApxcsVRyMpcCgYEAzZuD
rBNyo3JpBGmtIvHw58KGfBOS/evoP6qKVERMkpWZY1Fsh8EcYhybmDSRS7bTjP/v
GLfiXzJQZ2q8kQr+3tXEPe/QB+E+MXNNo5euuCQ6o3Q0Ja7huPrNYfTCP6stI32S
EKoms1mw4aRD1XwpcUm3o2nVWc0h0XCkXF1P1j0CgYEA5nQLKqGBwvhyRBhVjNhb
Wp95M0o3WCLi/kwwxX2TB30w9uSuMevQsH5WNIsFbpeMPVptGCj1RH+w0slWA04h
vPo28qGEnEBb4kAkQWbaEluxl29rWDdY/QzLl1zxZ08ktukU3eBVSGUVXDZl84o6
Li0CXQu6+bfBxx5LAKpIWaMCgYAVxDLqUpy+ROxtNSrJGkfgoS1PkVrsWr8Zjlpa
lWht1DyK0SHmNUFl+ZVXRalkFJTMxoNvYHgsj80HRbt0t29H8+V0kSC61NOatJQx
j2tFv0Ad8b1bh+oJhTOc/SZbSynaKf7+mKTEM+iP2q37uctBXQZ93ERj312HKJ+d
z5sWGQKBgCi03SV6YmyoQDEqRFnB4uUdVJ+t9jf5VWjYXkOkFEQOxjQjHN5mi+ES
ezyxljmC2RIgLko1FGacJJNsYvsYGjAIFTCE6UdKPyh7PzUtPz/VQpkylrqlUALn
zRFjHulgpC6KFKqWEOoZOQsoruAArVrsD+nHhuKS+afzS/VB25Bh
-----END RSA PRIVATE KEY-----`
)

var (
	pool *x509.CertPool
	c    *tls.Certificate
)

// Pair returns a tls certificate for the provided hard-coded certs.
func Pair() *tls.Certificate {
	if c == nil {
		got, err := tls.X509KeyPair([]byte(cert), []byte(sKey))
		if err != nil {
			log.Fatalf("Error creating cert %v", err)
		}
		c = &got
	}
	return c
}

// Pool returns a cert pool using insecure hard-coded certs
func Pool() *x509.CertPool {
	if pool == nil {
		pool = x509.NewCertPool()
		ok := pool.AppendCertsFromPEM([]byte(cert))
		if !ok {
			log.Fatal("unable to append certs to pool")
		}
	}
	return pool
}
