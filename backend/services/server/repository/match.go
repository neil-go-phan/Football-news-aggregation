package repository

import (
	"errors"
	"server/entities"
	"time"

	"gorm.io/gorm"
)

type matchRepo struct {
	DB *gorm.DB
}

func NewMatchRepo(db *gorm.DB) *matchRepo {
	return &matchRepo{
		DB: db,
	}
}

func (repo *matchRepo) Create(match *entities.Match) error {
	err := repo.DB.Create(match).Error
	if err != nil {
		return err
	}
	err = repo.DB.Save(match).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *matchRepo) UpdateWhenScheduleCrawl(match *entities.Match) error {
	err := repo.DB.Model(&match).
		Updates(entities.Match{
			Time:            match.Time,
			Round:           match.Round,
			Scores:          match.Scores,
			MatchDetailLink: match.MatchDetailLink}).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *matchRepo) UpdateWhenMatchDetailCrawl(match *entities.Match) error {
	err := repo.DB.Model(&match).
		Updates(entities.Match{
			MatchStatus: match.MatchStatus,
			Scores:      match.Scores,
			LineupClub1ID: match.LineupClub1ID,
			LineupClub2ID: match.LineupClub2ID,}).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *matchRepo) GetIDWithDateAndClubName(date time.Time, clubName1 string, clubName2 string) (*entities.Match, error) {
	match := new(entities.Match)
	dateString := date.Format("2006-01-02")
	err := repo.DB.Where("time_start BETWEEN ? AND ?", dateString+" 00:00:00", dateString+" 23:59:59").
		Joins("JOIN clubs AS club1 ON matches.club1_id = club1.id").
		Joins("JOIN clubs AS club2 ON matches.club2_id = club2.id").
		Where("club1.name = ?", clubName1).
		Where("club2.name = ?", clubName2).
		First(&match).Error
	if err != nil {
		return match, err
	}

	return match, nil
}

func (repo *matchRepo) GetMatch(match *entities.Match) (*entities.Match, error) {
	err := repo.DB.
		Preload("Statistics").
		Preload("Club1").
		Preload("Club2").
		Preload("Events").
		Preload("Club1Overview").
		Preload("Club2Overview").
		First(&match).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Không tìm thấy record
			return match, err
		}
		return match, err
	}

	return match, nil
}
