# go-rets
RETS import processor written in Go for RethinkDB

`docker-compose up`

will bring up Elasticsearch, Kibana, Logstash, RethinkDB

`logstash-input-rethinkdb` binds changefeeds to logstash event input for rethinkdb to send documents to elasticsearch.

The custom logstash configuration is in `logstash/config/logstash.conf`

Kick off the import with:

`go run import.go types.go -url=https://retsfeeds.example/xxxxxx.xml.gz -user=XXXXX -pass=XXXXXX`
