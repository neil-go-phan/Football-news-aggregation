docker_prepare:
	docker pull elasticsearch:7.17.9
	docker network create football-network
docker_network:
	docker network create football-network
docker_search:
	docker run --name elasticsearch7179 --network football-network -p 9200:9200 -e "discovery.type=single-node" -e "ES_JAVA_OPTS=-Xms512m -Xmx512m" -d elasticsearch:7.17.9
docker_redis:
	docker run --name redis72 --network football-network -p 6379:6379 -d redis:7.2-rc-alpine

.PHONY: docker_network docker_search docker_prepare
