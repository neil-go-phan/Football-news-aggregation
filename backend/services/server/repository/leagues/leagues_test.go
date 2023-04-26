package leaguesrepo

import (
	"fmt"
	"server/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

var PATH = "./testJson/"
var PATH_FAIL = "./testJson/fail/"
var PATH_WRITE = "./testJson/writedTest"

func TestGetLeagues(t *testing.T) {
	assert := assert.New(t)
	contructorLeagues := entities.Leagues{
		Leagues: []entities.League{
			{
				LeagueName: "Tin tức bóng đá",
				Active:     true,
			},
			{
				LeagueName: "La Liga",
				Active:     false,
			},
		},
	}
	want := entities.Leagues{
		Leagues: []entities.League{
			{
				LeagueName: "Tin tức bóng đá",
				Active:     true,
			},
			{
				LeagueName: "La Liga",
				Active:     false,
			},
		},
	}

	leaguesRepo := NewLeaguesRepo(contructorLeagues, PATH)
	got := leaguesRepo.GetLeagues()

	assert.Equal(want, got, fmt.Sprintf("Method GetLeagues is supose to %#v, but got %#v", want, got))
}

func TestGetLeaguesName(t *testing.T) {
	assert := assert.New(t)

	contructorLeagues := entities.Leagues{
		Leagues: []entities.League{
			{
				LeagueName: "Tin tức bóng đá",
				Active:     true,
			},
			{
				LeagueName: "La Liga",
				Active:     false,
			},
		},
	}
	want := []string{"Tin tức bóng đá", "La Liga"}
	leaguesRepo := NewLeaguesRepo(contructorLeagues, PATH)
	got := leaguesRepo.GetLeaguesName()

	assert.Equal(want, got, fmt.Sprintf("Method GetLeaguesName is supose to %#v, but got %#v", want, got))
}

func TestGetLeaguesNameActive(t *testing.T) {
	assert := assert.New(t)

	contructorLeagues := entities.Leagues{
		Leagues: []entities.League{
			{
				LeagueName: "Tin tức bóng đá",
				Active:     true,
			},
			{
				LeagueName: "La Liga",
				Active:     false,
			},
		},
	}
	want := []string{"Tin tức bóng đá"}
	leaguesRepo := NewLeaguesRepo(contructorLeagues, PATH)
	got := leaguesRepo.GetLeaguesNameActive()

	assert.Equal(want, got, fmt.Sprintf("Method GetLeaguesNameActive is supose to %#v, but got %#v", want, got))
}

func TestReadleaguesJSONSuccess(t *testing.T) {
	assert := assert.New(t)
	contructorLeagues := entities.Leagues{
		Leagues: []entities.League{
			{
				LeagueName: "Tin tức bóng đá test",
				Active:     true,
			},
			{
				LeagueName: "La Liga test",
				Active:     false,
			},
		},
	}

	want := entities.Leagues{
		Leagues: []entities.League{
			{
				LeagueName: "Tin tức bóng đá",
				Active:     true,
			},
			{
				LeagueName: "La Liga",
				Active:     false,
			},
		},
	}
	leaguesRepo := NewLeaguesRepo(contructorLeagues, PATH)
	got, _ := leaguesRepo.ReadleaguesJSON()

	assert.Equal(want, got, fmt.Sprintf("Method ReadleaguesJSON is supose to %#v, but got %#v", want, got))
}

func TestReadleaguesJSONCantOpenFile(t *testing.T) {
	assert := assert.New(t)
	contructorLeagues := entities.Leagues{
		Leagues: []entities.League{
			{
				LeagueName: "Tin tức bóng đá test",
				Active:     true,
			},
			{
				LeagueName: "La Liga test",
				Active:     false,
			},
		},
	}

	want := "file json not found"

	leaguesRepo := NewLeaguesRepo(contructorLeagues, PATH_FAIL)
	_, got := leaguesRepo.ReadleaguesJSON()

	assert.Errorf(got, want, fmt.Sprintf("Method ReadleaguesJSON is supose to %#v, but got %#v", want, got))
}

func TestWriteLeaguesJSON(t *testing.T) {
	assert := assert.New(t)

	contructorLeagues := entities.Leagues{
		Leagues: []entities.League{
			{
				LeagueName: "Tin tức bóng đá test",
				Active:     true,
			},
			{
				LeagueName: "La Liga test",
				Active:     false,
			},
		},
	}

	want := entities.Leagues{
		Leagues: []entities.League{
			{
				LeagueName: "Tin tức bóng đá",
				Active:     true,
			},
			{
				LeagueName: "La Liga",
				Active:     false,
			},
		},
	}
	leaguesRepo := NewLeaguesRepo(contructorLeagues, PATH)
	leaguesRepo.WriteLeaguesJSON(want)
	got := leaguesRepo.GetLeagues()
	assert.Equal(want, got, fmt.Sprintf("Method ReadleaguesJSON is supose to %#v, but got %#v", want, got))
}

func TestWriteLeaguesJSONCantOpenFile(t *testing.T) {
	assert := assert.New(t)

	contructorLeagues := entities.Leagues{
		Leagues: []entities.League{
			{
				LeagueName: "Tin tức bóng đá test",
				Active:     true,
			},
			{
				LeagueName: "La Liga test",
				Active:     false,
			},
		},
	}
	write := entities.Leagues{
		Leagues: []entities.League{
			{
				LeagueName: "Tin tức bóng đá",
				Active:     true,
			},
			{
				LeagueName: "La Liga",
				Active:     false,
			},
		},
	}

	want := "file json not found"

	leaguesRepo := NewLeaguesRepo(contructorLeagues, PATH_FAIL)
	got := leaguesRepo.WriteLeaguesJSON(write)
	assert.Errorf(got, want, fmt.Sprintf("Method ReadleaguesJSON is supose to %#v, but got %#v", want, got))
}

func TestAddLeague(t *testing.T) {
	assert := assert.New(t)

	contructorLeagues := entities.Leagues{
		Leagues: []entities.League{
			{
				LeagueName: "Tin tức bóng đá test",
				Active:     true,
			},
			{
				LeagueName: "La Liga test",
				Active:     false,
			},
		},
	}
	newLeagueName := "Test"

	want := entities.Leagues{
		Leagues: []entities.League{
			{
				LeagueName: "Tin tức bóng đá test",
				Active:     true,
			},
			{
				LeagueName: "La Liga test",
				Active:     false,
			},			
			{
				LeagueName: "Test",
				Active:     false,
			},
		},
	}

	leaguesRepo := NewLeaguesRepo(contructorLeagues, PATH_WRITE)
	leaguesRepo.AddLeague(newLeagueName)

	got := leaguesRepo.leagues


	assert.Equal(want, got, fmt.Sprintf("Method ReadleaguesJSON is supose to %#v, but got %#v", want, got))
}
