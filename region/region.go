package region

const (
	SG = sgRegion("BYTEPLUS_REGION_SG")
)

type sgRegion string

func (s sgRegion) GetHosts() []string {
	return []string{"rec-api-sg1.recplusapi.com"}
}

func (s sgRegion) GetAuthRegion() string {
	return "ap-singapore-1"
}
