package jobs

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"sync"
)

func Collect(db *pgxpool.Pool, wg *sync.WaitGroup) error {
	wg.Add(6)

	go func() {
		defer wg.Done()
		if err := flatCollector(db); err != nil {
			logrus.Printf("failed to collect data about flats %s", err.Error())
			return
		}
		logrus.Println("flatCollector is ready")
	}()

	go func() {
		defer wg.Done()
		if err := cottageCollector(db); err != nil {
			logrus.Printf("failed to collect data about cottages %s", err.Error())
			return
		}
		logrus.Println("cottageCollector is ready")
	}()

	go func() {
		defer wg.Done()
		if err := dachaCollector(db); err != nil {
			logrus.Printf("failed to collect data about cottages %s", err.Error())
			return
		}
		logrus.Println("dachaCollector is ready")
	}()

	go func() {
		defer wg.Done()
		if err := areaCollector(db); err != nil {
			logrus.Printf("failed to collect data about areas %s", err.Error())
			return
		}
		logrus.Println("areaCollector is ready")
	}()

	go func() {
		defer wg.Done()
		if err := flatLongRentCollector(db); err != nil {
			logrus.Printf("failed to collect data about flats for long rent %s", err.Error())
			return
		}
		logrus.Println("flatLongRentCollector is ready")
	}()

	go func() {
		defer wg.Done()
		if err := flatDayRentCollector(db); err != nil {
			logrus.Printf("failed to collect data about flats for long rent %s", err.Error())
			return
		}
		logrus.Println("flatDayRentCollector is ready")
	}()

	wg.Wait()
	return nil
}
