package repository

import (
	"server/entities"
	"time"

	"gorm.io/gorm"
)

type ScheduleRepo struct {
	DB *gorm.DB
}

func NewScheduleRepo(db *gorm.DB) *ScheduleRepo {
	return &ScheduleRepo{
		DB: db,
	}
}

func (repo *ScheduleRepo) FirstOrCreate(schedule *entities.Schedule) error {
	err := repo.DB.Omit("Matches").Assign(schedule).FirstOrCreate(schedule, entities.Schedule{LeagueName: schedule.LeagueName, Date: schedule.Date}).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *ScheduleRepo) GetScheduleOnLeague(leagueName string, date time.Time) (*entities.Schedule, error) {
	schedule := new(entities.Schedule)
	dateString := date.Format("2006-01-02")
	err := repo.DB.Preload("Matches").
		Preload("Matches.Club1").
		Preload("Matches.Club2").
		Where("league_name = ? AND date BETWEEN ? AND ?", leagueName, dateString+" 00:00:00", dateString+" 23:59:59").
		Find(&schedule).Error

	if err != nil {
		return nil, err
	}
	return schedule, nil
}

func (repo *ScheduleRepo) GetScheduleOnDay(date time.Time) (*[]entities.Schedule, error) {
	schedules := make([]entities.Schedule, 0)
	dateString := date.Format("2006-01-02")
	err := repo.DB.Preload("Matches").
		Preload("Matches.Club1").
		Preload("Matches.Club2").
		Where("date BETWEEN ? AND ?", dateString+" 00:00:00", dateString+" 23:59:59").
		Find(&schedules).Error

	if err != nil {
		return nil, err
	}
	return &schedules, nil
}
