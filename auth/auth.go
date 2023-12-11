package auth

// DOC: https://www.okx.com/docs-v5/zh/#overview-v5-api-key-creation

type Auth struct {
	APIKey, SecretKey, Passphrase string
}

func NewAuth(apiKey, secretKey, passphrase string) Auth {
	return Auth{
		APIKey:     apiKey,
		SecretKey:  secretKey,
		Passphrase: passphrase,
	}
}

func (a Auth) Signature(timestamp, method, requestPath, body string) *Signature {
	return &Signature{
		Timestamp:   timestamp,
		Method:      method,
		RequestPath: requestPath,
		Body:        body,
		SecretKey:   a.SecretKey,
	}
}
