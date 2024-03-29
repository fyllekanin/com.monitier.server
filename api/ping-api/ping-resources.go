package ping_api

import (
	"github.com/fyllekanin/com.monitier.server/database/entity"
	"time"
)

func GetOverviewHours(pings []entity.PingEntity) []*Overview {
	year, month, day := time.Now().Date()
	var twentyFourHours = time.Now().Hour()
	var currentHour = time.Date(year, month, day, twentyFourHours, 0, 0, 0, time.Now().Location()).Unix()

	var response []*Overview
	for i := 0; i < 24; i++ {
		var activeHour = currentHour - (int64(i) * 86000)
		var lastOfHour = activeHour + 859999
		var responseTimeTotal = 0
		var count = 0
		for _, item := range pings {
			if item.CreatedAt >= activeHour && item.CreatedAt <= lastOfHour {
				responseTimeTotal += item.ResponseTime
				count++
			}
		}
		if responseTimeTotal > 0 {
			response = append(response, &Overview{
				Day:                 i,
				AverageResponseTime: responseTimeTotal / count,
				DataPoints:          count,
			})
		} else {
			response = append(response, &Overview{
				Day:                 i,
				AverageResponseTime: 0,
				DataPoints:          0,
			})
		}
	}
	return response
}

func GetOverviewDays(pings []entity.PingEntity) []*Overview {
	year, month, day := time.Now().Date()
	var twentyFourHours = time.Now().Add(24 * time.Hour).Unix()
	var todayMidnight = time.Date(year, month, day, 0, 0, 0, 0, time.Now().Location()).Unix()
	var todayLast = todayMidnight + (twentyFourHours - 1)

	var response []*Overview
	for i := 0; i < 90; i++ {
		var activeStart = todayMidnight - (int64(i) * twentyFourHours)
		var activeLast = todayLast - (int64(i) * twentyFourHours)
		var responseTimeTotal = 0
		var count = 0
		for _, item := range pings {
			if item.CreatedAt >= activeStart && item.CreatedAt <= activeLast {
				responseTimeTotal += item.ResponseTime
				count++
			}
		}
		
		if responseTimeTotal > 0 {
			response = append(response, &Overview{
				Day:                 i,
				AverageResponseTime: responseTimeTotal / count,
				DataPoints:          count,
			})
		} else {
			response = append(response, &Overview{
				Day:                 i,
				AverageResponseTime: 0,
				DataPoints:          0,
			})
		}
	}
	return response
}
