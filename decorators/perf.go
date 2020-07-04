package decorators

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/streadway/amqp"
)

// RestFunc type for a typical REST handler
type RestFunc func(http.ResponseWriter, *http.Request)

// RestPerf measures endpoint performance and logs execution time
func RestPerf(restHandler RestFunc, rmqChannel *amqp.Channel) RestFunc {

	return func(res http.ResponseWriter, req *http.Request) {

		defer func(t time.Time) {
			elapsed := fmt.Sprintf("--- Time Elapsed: %fs ---\n", time.Since(t).Seconds())
			queue := os.Getenv("RMQ_QNAME")
			rmqChannel.Publish(
				"",
				queue,
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(elapsed),
				},
			)
		}(time.Now())

		restHandler(res, req)
	}
}
