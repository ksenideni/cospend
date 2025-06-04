package constant

const (
	GIN_KEY       = "user"
	AUTHORIZATION = "Authorization"
	XCTO          = "X-Content-Type-Options"
	XCTO_VALUE    = "nosniff"
	HSTS          = "Strict-Transport-Security"
	HSTS_VALUE    = "max-age=31536000; includeSubDomains; preload"
	CC            = "Cache-Control"
	CC_VALUE      = "no-store"
	ACAO          = "Access-Control-Allow-Origin"
	ACAO_VALUE    = "*"
	ACAM          = "Access-Control-Allow-Methods"
	ACAM_VALUE    = "GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS"
	ACAH          = "Access-Control-Allow-Headers"
	ACAH_VALUE    = "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Origin, Cookie, Signature, Timestamp"
	ACAC          = "Access-Control-Allow-Credentials"
	ACAC_VALUE    = "true"

	LIMIT_DEFAULT = 10
	PAGE_DEFAULT  = 1
)
