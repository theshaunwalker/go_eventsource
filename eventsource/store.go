package eventsource

import (
	"github.com/aklinkert/go-gorm-repository"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func getRepo() gormrepository.TransactionRepository {
	logger := logrus.New()
	db, err := gorm.Open(sqlite.Open("./dev_mess/db.sqlite"), &gorm.Config{})

	if err != nil {
		panic("get repo failed")
	}

	repo := gormrepository.NewGormRepository(db, logger)

	return repo
}

func StoreEvent(event EventEnvelope) {
	repo := getRepo()

	storeRecord := event.ToStoreRecord()

	err := repo.Save(&storeRecord)

	if err != nil {
		panic("store failed")
	}
}

func StoreAggregateRoot(root RecordsEvents) {
	events := root.ReleaseEvents()
	if len(events) == 0 {
		return
	}

	for _, event := range events {
		StoreEvent(event)
	}
}
