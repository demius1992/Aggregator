package jobs

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

func Collect(db *pgxpool.Pool) error {

	message1 := make(chan string)
	message2 := make(chan string)
	message3 := make(chan string)
	message4 := make(chan string)
	message5 := make(chan string)
	message6 := make(chan string)

	go func() {
		if err := flatCollector(db); err != nil {
			logrus.Printf("failed to collect data about flats %s", err.Error())
			return
		}
		message6 <- "flatCollector is ready"
	}()

	go func() {
		if err := cottageCollector(db); err != nil {
			logrus.Printf("failed to collect data about cottages %s", err.Error())
			return
		}
		message1 <- "cottageCollector is ready"
	}()

	go func() {
		if err := dachaCollector(db); err != nil {
			logrus.Printf("failed to collect data about cottages %s", err.Error())
			return
		}
		message2 <- "dachaCollector is ready"
	}()

	go func() {
		if err := areaCollector(db); err != nil {
			logrus.Printf("failed to collect data about areas %s", err.Error())
			return
		}
		message3 <- "areaCollector is ready"
	}()

	go func() {
		if err := flatLongRentCollector(db); err != nil {
			logrus.Printf("failed to collect data about flats for long rent %s", err.Error())
			return
		}
		message4 <- "flatLongRentCollector is ready"
	}()

	go func() {
		if err := flatDayRentCollector(db); err != nil {
			logrus.Printf("failed to collect data about flats for long rent %s", err.Error())
			return
		}
		message5 <- "flatDayRentCollector is ready"
	}()

	msg1 := <-message1
	msg2 := <-message2
	msg3 := <-message3
	msg4 := <-message4
	msg5 := <-message5
	msg6 := <-message6

	logrus.Println(msg1)
	logrus.Println(msg2)
	logrus.Println(msg3)
	logrus.Println(msg4)
	logrus.Println(msg5)
	logrus.Println(msg6)

	return nil
}
