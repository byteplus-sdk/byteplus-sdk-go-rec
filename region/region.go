package region

import (
	core "github.com/byteplus-sdk/byteplus-sdk-go-rec-core"
)

func init() {
	core.RegisterRegion(SG, &core.RegionConfig{
		Hosts:                SGHosts,
		VolcCredentialRegion: "ap-singapore-1",
	})
}
