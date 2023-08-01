module github.com/theshaunwalker/eventsourcing

go 1.20

require (
	github.com/aklinkert/go-gorm-repository v1.2.0
	github.com/google/uuid v1.3.0
	github.com/sirupsen/logrus v1.9.3
	gorm.io/driver/sqlite v1.5.2
	gorm.io/gorm v1.25.2
)

require (
	github.com/aklinkert/go-logging v1.2.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.17 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
)

replace (
	github.com/theshaunwalker/eventsourcing/eventsource v0.0.0 => "./eventsource"
)