package repository

import (
	"server/entities"
	"time"
)

//go:generate mockery --name AdminRepository
type AdminRepository interface {
	UpdatePassword(admin *entities.Admin) error 
	Get(username string) (*entities.Admin, error)
	Create(admin *entities.Admin) (error) 
	GetWithEmail(email string) (*entities.Admin, error)
	Upsert(admin *entities.Admin) (error)
}

//go:generate mockery --name LeaguesRepository
type LeaguesRepository interface {
	Create(league *entities.League) (error)
	Update(league *entities.League) error 
	List() (*[]entities.League, error) 
	GetLeaguesNameActive() (*[]entities.League, error)
	GetLeaguesName() (*[]entities.League, error) 
	GetByName(leagueName string) (*entities.League, error) 
}

//go:generate mockery --name TagRepository
type TagRepository interface {
	List() (*[]entities.Tag, error)
	Create(tag *entities.Tag) error 
	Delete(tagName string) error
	Get(tagName string) (*entities.Tag,error)
	GetTagsByTagNames(tagNames []string) (*[]entities.Tag, error)
}

//go:generate mockery --name ArticleRepository
type ArticleRepository interface {
	FirstOrCreate(article *entities.Article) error 
	Delete(id uint) error 
	AddTag(ids []uint, newTag *entities.Tag) error 
	RemoveTag(ids []uint, newTag *entities.Tag) error
	GetArticlesByIDs(ids []uint) ([]entities.Article, error)
	GetCrawledArticleToday() (int64, error)
	GetTotalCrawledArticle() (int64, error)
}

//go:generate mockery --name SchedulesRepository
type SchedulesRepository interface {
	FirstOrCreate(schedule *entities.Schedule) error

	GetScheduleOnLeague(leagueName string, date time.Time) (*entities.Schedule, error) 
	GetScheduleOnDay(date time.Time) (*[]entities.Schedule, error)
}

//go:generate mockery --name MatchRepository
type MatchRepository interface {
	Create(match *entities.Match) error
	UpdateWhenScheduleCrawl(match *entities.Match) error
	UpdateWhenMatchDetailCrawl(match *entities.Match) error 
	GetIDWithDateAndClubName(date time.Time, clubName1 string, clubName2 string) (*entities.Match, error)
	GetMatch(match *entities.Match) (*entities.Match, error) 
}

//go:generate mockery --name ClubRepository
type ClubRepository interface {
	GetByName(clubName string) (*entities.Club,error)
	FirstOrCreate(clubName string, logo string) (*entities.Club,error)
}

//go:generate mockery --name StatsItemlRepository
type StatsItemlRepository interface {
	FirstOrCreate(statsItem *entities.StatisticsItem) error 
}

//go:generate mockery --name EventRepository
type EventRepository interface {
	FirstOrCreate(eventRepo *entities.MatchEvent) error 
}

//go:generate mockery --name OverviewItemRepository
type OverviewItemRepository interface {
	FirstOrCreate(overviewItem *entities.OverviewItem) error
}

//go:generate mockery --name PlayerRepository
type PlayerRepository interface {
	FirstOrCreate(player *entities.Player) error 
}

//go:generate mockery --name LineupRepository
type LineupRepository interface {
	FirstOrCreate(lineup *entities.MatchLineUp) (*entities.MatchLineUp, error)
	Get(id uint) (*entities.MatchLineUp, error) 
}

//go:generate mockery --name ConfigCrawlerRepository
type ConfigCrawlerRepository interface {
	Upsert(configCrawler *entities.ConfigCrawler) (error)
	List() (*[]entities.ConfigCrawler, error) 
	Get(url string) (*entities.ConfigCrawler,error) 
	Delete(url string) error
}

type AllMatchURLsOnDay struct {
	Date time.Time
	Urls []string
}

// The matchs takes place in a spectific time on a day
type MatchURLsOnTime struct {
	Date time.Time
	Urls []string
}

type MatchURLsWithTimeOnDay struct {
	MatchsOnTimes []MatchURLsOnTime
}

type ScheduleElastic struct {
	Date       time.Time `json:"date"`
	LeagueName string    `json:"league_name"`
	Matches     []Match   `json:"matches"`
}

type ScheduleOnDay struct {
	Date              time.Time          `json:"date"`
	DateWithWeekday   string             `json:"date_with_weekday"`
	ScheduleOnLeagues []ScheduleOnLeague `json:"schedule_on_leagues"`
}

type ScheduleOnLeague struct {
	LeagueName string  `json:"league_name"`
	Matches     []Match `json:"matches"`
}

type Match struct {
	Time            string `json:"time"`
	Round           string `json:"round"`
	Club1           Club   `json:"club_1"`
	Club2           Club   `json:"club_2"`
	Scores          string `json:"scores"`
	MatchDetailLink string `json:"match_detail_link"`
}

type Club struct {
	Name string `json:"name"`
	Logo string `json:"logo"`
}
