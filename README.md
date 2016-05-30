# go-rets
RETS import processor written in Go for RethinkDB

Maps RETS schema to structs with xml and json tags. Entire resulting listing document is inserted into rethinkdb as json, where `logstash-input-rethinkdb` binds to the table's changefeed events and sends documents to elasticsearch via logstash. The custom logstash configuration is in `logstash/config/logstash.conf`. Only a few fields are imported into elasticsearch for indexing. Kibana and Elastic HQ are running for ES administration and visualization.

### Setup

`docker-compose up rethinkdb`

Create a database `realestate` and a table `listings` in rethinkdb before beginning.

`docker-compose up`

Runs Elasticsearch, Kibana, Logstash, RethinkDB

Kick off the import with:

`go run import.go types.go -url=https://retsfeeds.example/xxxxxx.xml.gz -user=XXXXX -pass=XXXXXX`
