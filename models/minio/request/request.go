package request

type EventReq struct {
	GetObjectContext struct {
		InputS3Url  string `json:"inputS3Url"`
		OutputRoute string `json:"outputRoute"`
		OutputToken string `json:"outputToken"`
	} `json:"getObjectContext"`
	ProtocolVersion string       `json:"protocolVersion"`
	UserIdentity    UserIdentity `json:"userIdentity"`
	UserRequest     struct {
		Headers struct {
			Accept    []string `json:"Accept"`
			UserAgent []string `json:"User-Agent"`
		} `json:"headers"`
		Url string `json:"url"`
	} `json:"userRequest"`
}

type UserIdentity struct {
	AccessKeyId string `json:"accessKeyId"`
	PrincipalId string `json:"principalId"`
	Type        string `json:"type"`
}
