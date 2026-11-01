package flashdecoder

import (
	flashinfo "github.com/Kasesiki/flashDecoder-go/info"
	"github.com/Kasesiki/flashDecoder-go/utils"
)

func GenItoString(info flashinfo.Flashinfo) string {
	if info.Generation > 0 {
		switch info.Vendor {
		case "Intel":
			return utils.IntelGenString[info.Generation-1]
		case "Micron":
			return utils.MicronGenString[info.Generation-1]
		default:
			return ""
		}
	} else {
		return ""
	}

}
