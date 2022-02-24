module github.com/kobehaha/machinery

go 1.16

require (
	cloud.google.com/go/pubsub v1.18.0
	github.com/RichardKnop/logging v0.0.0-20190827224416-1a693bdd4fae
	github.com/RichardKnop/machinery v1.10.6
	github.com/aws/aws-sdk-go v1.43.5
	github.com/bradfitz/gomemcache v0.0.0-20220106215444-fb4bf637b56d
	github.com/go-redis/redis/v8 v8.11.4
	github.com/go-redsync/redsync/v4 v4.5.0
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/google/uuid v1.3.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/robfig/cron/v3 v3.0.1
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.7.0
	github.com/urfave/cli v1.22.5
	go.mongodb.org/mongo-driver v1.8.3
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/RichardKnop/machinery => github.com/kobehaha/machinery v1.10.6
