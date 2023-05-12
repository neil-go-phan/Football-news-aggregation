package matchservices

import (
	"server/entities"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReadTimeStart(t *testing.T) {
	dayTime := time.Date(2022, 5, 4, 0, 0, 0, 0, time.UTC)
	assert := assert.New(t)
	match := entities.Match{
		Time: "FT - 11/05",
	}

	// Kiểm tra trường hợp match đã kết thúc
	exactTime, err := readTimeStart(match, dayTime)
	assert.Nil(err)
	assert.Equal(exactTime, dayTime, "Expected exactTime to be %s, got %s", dayTime.String(), exactTime.String())

	match = entities.Match{
		Time: "11:30 - 04/05",
	}

	// Kiểm tra trường hợp match chưa diễn ra
	exactTime, err = readTimeStart(match, dayTime)
	assert.Nil(err)

	expectedTime := time.Date(2022, 5, 4, 11, 30, 0, 0, time.UTC)
	assert.Equal(expectedTime, exactTime, "Expected expectedTime to be %s, got %s", exactTime.String(), expectedTime.String())

	// Kiểm tra trường hợp time hour không hợp lệ

	match = entities.Match{
		Time: "invalid",
	}
	_, err = readTimeStart(match, dayTime)
	assert.Error(err)

	// Kiểm tra trường hợp time min không hợp lệ

	match = entities.Match{
		Time: "11:invaild - 04/05",
	}
	_, err = readTimeStart(match, dayTime)
	assert.Error(err)
}

func TestReadTimeEvent(t *testing.T) {
	dayTime := time.Date(2022, 5, 4, 0, 0, 0, 0, time.UTC)
	assert := assert.New(t)

	// success
	timeString := "45+1"
	want := dayTime.Add(time.Minute * time.Duration(46))
	got, err := readTimeEvent(timeString,dayTime)
	assert.Nil(err)
	assert.Equal(want, got)

}