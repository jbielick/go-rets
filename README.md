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

Many thanks to the fine folks who contributed and shared in these links for help with figuring a lot of this out:

https://github.com/rethinkdb/logstash-input-rethinkdb
https://github.com/BenMann/cookiejar/commit/0bafcf406bfe25a5f986544d73be236a9a21511f
https://www.bountysource.com/issues/28139325-update-full-text-search-documentation-to-use-logstash
http://jmoiron.net/blog/limiting-concurrency-in-go/
https://www.elastic.co/guide/en/logstash/current/plugins-filters-mutate.html#plugins-filters-mutate-replace
