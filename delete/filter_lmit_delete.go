package del

import (
	"log"
	"time"
)

func (data ResModels) FilterLimitDelete(PageToken string) (ok []OK) {
	limit := 50

	if len(data.Data) > limit {
		xLmt := 0
		for i := 0; i < (len(data.Data)/limit)+1; i++ {
			if i == (len(data.Data) / limit) {
				d := data.Data[xLmt:]
				newData := ResModels{d}
				ok = newData.Delete(PageToken)
				if ok[0].Code == 200 {
					log.Println("Done...✅")
				} else {
					log.Println("Some Data Failed...❌ but not all")
				}
				return ok
			}
			d := data.Data[xLmt : xLmt+limit]
			newData := ResModels{d}
			ok = newData.Delete(PageToken)
			if ok[0].Code == 200 {
				log.Println("Successfully deleted some data, please wait 30 minute to next request🕑🕑🕑")
				log.Println("Keep Enjoying🎉🎉🎉")
				time.Sleep(time.Second * 900)
			}
			xLmt += limit
		}
		return
	} else {
		ok = data.Delete(PageToken)
		return ok
	}

}
