package auth

// DOC: https://www.okx.com/docs-v5/en/#overview-rest-authentication-signature

type Signature struct {
	Timestamp, Method, RequestPath, Body string
	SecretKey                            string
}

func (s *Signature) Sign() string {
	data := s.Timestamp + s.Method + s.RequestPath + s.Body
	return HmacSHA256([]byte(data), []byte(s.SecretKey))
}
