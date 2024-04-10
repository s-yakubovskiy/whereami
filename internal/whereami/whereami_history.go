package whereami

import (
	"fmt"

	"github.com/s-yakubovskiy/whereami/internal/common"
)

func (l *Locator) History(num int32) {
	var (
		categories = map[string][]string{
			"Info": {
				"flag", "country",
				"city", "ip", "latitude", "longitude", "date",
			},
		}

		orderedCategories = []string{
			"Info",
		}
	)

	locations, err := l.dbclient.ShowLocations(int(num))
	if err != nil {
		common.Errorln(err.Error())
	}
	for idx, location := range locations {
		common.Infoln(fmt.Sprintf("\n----- %s: #%d ------", "location", idx+1))
		location.Output(categories, orderedCategories)
	}
}
