package registry

type Payload struct {
	Status string
	Result []CertificateAction
}

type CertificateAction struct {
	ActionType        string
	CompletedTime     string
	CertificateRanges []CertificateRange
}

type CertificateRange struct {
	CertificateType        string
	RegisteredPersonNumber int
	AccreditationCode      string
	GenerationYear         int
	GenerationState        string
	StartSerialNumber      int
	EndSerialNumber        int
	FuelSource             string
	OwnerAccount           string
	OwnerAccountID         int
	Status                 string
}
