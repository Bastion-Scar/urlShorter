package models

import (
	"awesomeProject13/internal/storage"
	"gorm.io/gorm"
	"time"
)

var LogChan = make(chan storage.Logs, 100)

func SendLogs(db *gorm.DB) {
	ticker := time.NewTicker(3 * time.Second)
	batch := make([]storage.Logs, 0, 10)

	for {
		select {
		case <-ticker.C:
			if len(batch) > 0 {
				flushBatch(db, batch)
				batch = batch[:0]
			}

		case logs := <-LogChan:
			batch = append(batch, logs)
			if len(batch) >= 10 {
				flushBatch(db, batch)
				batch = batch[:0]
			}
		}
	}
}

func flushBatch(db *gorm.DB, batch []storage.Logs) {
	db.Create(&batch)
}
