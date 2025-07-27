package fetchers

import (
	"fmt"
	"time"

	"groupietracker/models"
)

func TimeToLiveFetch() {
	// this comment if you want to test the update data
	// fmt.Println(time.Since(models.LastUpdateTime), " : ", models.Ttl)
	models.Mu.Lock()
	if time.Since(models.LastUpdateTime) >= models.Ttl {
		InitFetch()
		fmt.Println("🔄 Data Update in => ", time.Now())
	}
	models.Mu.Unlock()
}
