package repository

import (
	"server/entities"
	"time"

	"gorm.io/gorm"
)

type CronjobRepo struct {
	DB *gorm.DB
}

func NewCronjobRepo(db *gorm.DB) *CronjobRepo {
	return &CronjobRepo{
		DB: db,
	}
}

func (repo *CronjobRepo) Create(cronjob *entities.Cronjob) (*entities.Cronjob, error) {
	err := repo.DB.Create(cronjob).Error
	if err != nil {
		return cronjob, err
	}
	return cronjob, nil
}

func (repo *CronjobRepo) Get(lastTrack time.Time, now time.Time) (*[]entities.Cronjob, error) {
	cronjobs := make([]entities.Cronjob, 0)
	lastTrackString := lastTrack.Format("2006-01-02 15:04:05")
	nowString := now.Format("2006-01-02 15:04:05")

	err := repo.DB.
		Where("start_at BETWEEN ? AND ?", lastTrackString, nowString).
		Find(&cronjobs).Error
	if err != nil {
		return &cronjobs, err
	}
	return &cronjobs, nil
}

func (repo *CronjobRepo) UpdateEndAt(cronjob *entities.Cronjob) error {
	err := repo.DB.Model(&entities.Cronjob{}).Where("id = ?", cronjob.ID).Update("end_at", cronjob.EndAt).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *CronjobRepo) GetRuning() (*[]entities.Cronjob, error) {
	cronjobs := make([]entities.Cronjob, 0)

	err := repo.DB.
		Where("end_at IS NULL").
		Find(&cronjobs).Error
	if err != nil {
		return &cronjobs, err
	}
	return &cronjobs, nil
}
