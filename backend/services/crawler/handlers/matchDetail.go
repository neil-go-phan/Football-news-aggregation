package handlers

import (
	"crawler/entities"
	"crawler/helper"
	pb "crawler/proto"
	"crawler/services"
	log "github.com/sirupsen/logrus"
	"sync"

	jsoniter "github.com/json-iterator/go"
)
// to slow down crawl match process (TODO: research about how to control amount of go routine suitable for server power)
var AMOUNT_MATCH_PER_GOROUTINE = 20

func (s *gRPCServer) GetMatchDetail(configs *pb.MatchURLs, stream pb.CrawlerService_GetMatchDetailServer) error {
	matchUrls := configs.GetUrl()

	xPath, err := crawlerhelpers.ReadXPathClassMatchDetailJSON()
	if err != nil {
		log.Println("can not read file htmlSchedulesClass.json, err: ", err)
	}

	matchUrlsChunk := matchUrlsChunk(matchUrls, AMOUNT_MATCH_PER_GOROUTINE)
	var wg sync.WaitGroup
	log.Println("Start scrapt match detail")

	for _, matchUrl := range matchUrlsChunk {
		wg.Add(1)
		go func(matchUrl []string, wg *sync.WaitGroup, xPath entities.XPathMatchDetail) {
			for _, url := range matchUrl {
				err := crawlMatchDetailAndStreamResult(stream, url, xPath)
				if err != nil {
					log.Printf("error occurred while request to url: %s, err: %v ", matchUrl, err)
				}
			}
			defer wg.Done()
		}(matchUrl, &wg, xPath)
	} 
	wg.Wait()
	log.Println("Finish scrapt match detail")
	return nil
}

func matchUrlsChunk(matchUrls []string, chunkSize int) [][]string {
	var chunks [][]string
	for {
		if len(matchUrls) == 0 {
			break
		}
		if len(matchUrls) < chunkSize {
			chunkSize = len(matchUrls)
		}
		chunks = append(chunks, matchUrls[0:chunkSize])
		matchUrls = matchUrls[chunkSize:]
	}

	return chunks
}

func crawlMatchDetailAndStreamResult(stream pb.CrawlerService_GetMatchDetailServer, matchUrl string, xPath entities.XPathMatchDetail) error {
	log.Println("request to URL: ", matchUrl)

	matchDetailEntity, err := services.CrawlMatchDetail(matchUrl, xPath)
	if err != nil {
		log.Printf("error occurred during crawl match detail process: url: %v, err: %v ", matchUrl, err)
	}

	matchDetail := crawledMatchDetailToPbMatchDetail(matchDetailEntity, matchUrl)

	err = stream.Send(matchDetail)
	if err != nil {
		log.Println("error occurred while sending response to client: ", err)
	}

	log.Printf("crawl ended: %s", matchUrl)
	return nil
}

func crawledMatchDetailToPbMatchDetail(matchDetailEntity entities.MatchDetail, url string) *pb.MatchDetail {
	pbMatchDetail := &pb.MatchDetail{}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	matchDetailByte, err := json.Marshal(matchDetailEntity)
	if err != nil {
		log.Printf("error occrus when marshal crawled schedules: %s", err)
	}

	err = json.Unmarshal(matchDetailByte, pbMatchDetail)
	if err != nil {
		log.Printf("error occrus when unmarshal crawled schedules to proto.Schedules: %s", err)
	}
	return pbMatchDetail
}
