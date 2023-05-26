package services

import (
	"net/url"
	"server/entities"
	pb "server/proto"
	"server/repository"
	adminservices "server/services/admin"
	"time"
)

type ArticleCache struct {
	ID          uint     `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Link        string   `json:"link"`
	Tags        []string `json:"tags"`
}

//go:generate mockery --name AdminServices
type AdminServices interface {
	GetAdminUsername(username string) (string, error)
	CheckAdminUsernameToken(username string) error
	ChangePassword(admin *adminservices.AdminWithConfirmPassword, usernameToken string) error
	LoginWithUsername(admin *adminservices.Admin) (string, error)
	LoginWithEmail(admin *adminservices.Admin) (string, error)
	Register(admin *adminservices.RegisterUserInput) error
	GoogleOAuth(googleUser *adminservices.GoogleUserResult) (string, error)
}

//go:generate mockery --name LeaguesServices
type LeaguesServices interface {
	CreateLeague(newLeaguesName string) error
	ListLeagues() (*[]entities.League, error)
	GetLeaguesNameActive() ([]string, error)
	GetLeaguesName() ([]string, error)
	ChangeStatus(leagueName string) (bool, error)
}

//go:generate mockery --name TagsServices
type TagsServices interface {
	ListTagsName() ([]string, error)
	AddTag(newTags string) error
	DeleteTag(tagName string) error
	GetTagsByTagNames(tagNames []string) (*[]entities.Tag, error)
	Get(tagName string) (*entities.Tag, error)
}

//go:generate mockery --name ArticleServices
type ArticleServices interface {
	SearchArticles(keyword string, formatedTags []string, from int) ([]entities.Article, int64, error)
	AddTagForAllArticle(tag string) error
	GetFirstPageOfLeagueRelatedArticle(leagueName string) ([]ArticleCache, error)
	RefreshCache()
	GetArticles(keywords []string)
	GetArticleCount() (total int64, today int64, err error)
	DeleteArticle(id uint) error
	StoreArticles(respArticles []*pb.Article, league string)
}

//go:generate mockery --name SchedulesServices
type SchedulesServices interface {
	GetAllScheduleLeagueOnDay(date time.Time) (repository.ScheduleOnDay, error)
	GetScheduleLeagueOnDay(date time.Time, league string) (repository.ScheduleOnDay, error)
	GetSchedules(date string)
	SignalMatchDetailServiceToCrawl(matchURLs repository.AllMatchURLsOnDay) []*pb.MatchDetail
	GetMatchURLsOnTime() repository.MatchURLsWithTimeOnDay
	ClearMatchURLsOnTime()
	GetAllMatchURLs() repository.AllMatchURLsOnDay
	ClearAllMatchURLs()
}

//go:generate mockery --name ClubServices
type ClubServices interface {
	GetClubByName(clubName string) (*entities.Club, error)
	GetOrCreate(clubName string, logo string) (*entities.Club, error)
}

//go:generate mockery --name MatchServices
type MatchServices interface {
	StoreMatch_MatchDetailCrawl(resp *pb.MatchDetail, date time.Time) error
	StoreMatch_ScheduleCrawl(match entities.Match, ID uint, date time.Time) error
	GetMatch(date time.Time, club1Name string, club2Name string) (*entities.Match, error)
	GetMatchDetailsOnDayFromCrawler(matchURLs repository.AllMatchURLsOnDay) []*pb.MatchDetail
	GetLineUps(id1 uint, id2 uint) (*entities.MatchLineUp, *entities.MatchLineUp, error)
}

//go:generate mockery --name StatsItemServices
type StatsItemServices interface {
	FirstOrCreate(statsItem *entities.StatisticsItem) error
}

//go:generate mockery --name EventServices
type EventServices interface {
	FirstOrCreate(event *entities.MatchEvent) error
}

//go:generate mockery --name OverviewItemServices
type OverviewItemServices interface {
	FirstOrCreate(overviewItem *entities.OverviewItem) error
}

//go:generate mockery --name PlayerServices
type PlayerServices interface {
	FirstOrCreate(player *entities.Player) error
}

//go:generate mockery --name LineUpServices
type LineUpServices interface {
	GetOrCreate(lineup *entities.MatchLineUp) (*entities.MatchLineUp, error)
	GetLineUps(id1 uint, id2 uint) (*entities.MatchLineUp, *entities.MatchLineUp, error)
}

//go:generate mockery --name CrawlerServices
type CrawlerServices interface {
	GetHtmlPage(url *url.URL) error
	Upsert(configCrawler *Crawler) error
	List() ([]Crawler, error)
	Get(url string) (*entities.Crawler, error)
	Delete(urlInput string) error
	TestCrawler(configCrawler *Crawler) ([]entities.Article, error)
	GetArticles(configCrawler *Crawler) ([]entities.Article, error)
	UpdateRunEveryTime(crawler *entities.Crawler) error
	CreateCustomCrawlerCronjob() error
	ChangeScheduleCronjob(cronjobIn CronjobChangeTimeRequestPayload) error
}

//go:generate mockery --name CronjobServices
type CronjobServices interface {
	CreateCrawlerCronjob(crawler *entities.Crawler)
	GetCronjobRuntime() []CronjobResponse
	CronjobOnHour(timeString string) (*[]ChartHourResponse, error)
	CreateCronjobGetArticleFromGoogle()
	CreateCronjobRefreshCache()
	CronjobOnDay(timeString string) (*[24]ChartDay, error)
}

type Crawler struct {
	Url                string `json:"url" validate:"required"`
	ArticleDiv         string `json:"article_div" validate:"required"`
	ArticleTitle       string `json:"article_title" validate:"required"`
	ArticleDescription string `json:"article_description"`
	ArticleLink        string `json:"article_link" validate:"required"`
	NextPage           string `json:"next_page"`
	NetxPageType       string `json:"next_page_type" validate:"required"`
}

type CronjobResponse struct {
	Name        string `json:"name"`
	Url         string `json:"url"`
	RunEveryMin int    `json:"run_every_min"`
}

type CronjobChangeTimeRequestPayload struct {
	Name           string `json:"name"`
	Url            string `json:"url"`
	RunEveryMinOld int    `json:"run_every_min_old"`
	RunEveryMinNew int    `json:"run_every_min_new"`
}

type ChartHourResponse struct {
	Time        time.Time `json:"time"`
	AmountOfJob int       `json:"amount_of_jobs"`
	// MemoryUsage  int       `json:"memory_usage"`
	Cronjobs []CronjobInChart `json:"cronjobs"`
}

type ChartDay struct {
	Hour        int            `json:"hour"`
	AmountOfJob int            `json:"amount_of_jobs"`
	Cronjobs    map[string]int // map[cronjob_name]runnng_times
}

type CronjobInChart struct {
	Name    string    `json:"name"`
	StartAt time.Time `json:"start_at"`
	EndAt   time.Time `json:"end_at"`
}
