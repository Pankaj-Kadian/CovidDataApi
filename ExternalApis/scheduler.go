package externalapis

import (
	"time"
)

func startFecthingData() {
	for {
		final_data := gettingData()
		insertingData(final_data)
		time.Sleep(1 * time.Second)
	}

}
