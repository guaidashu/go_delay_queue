/**
 * @Author: yy
 * @Author: 1023767856@qq.com
 * @Date: 28/06/2021
 * @Desc: 消费者
 */

package go_delay_queue

import "github.com/go-redis/redis"

type (
	Consume func(body []byte)

	Consumer interface {
		Consume(consume Consume)
	}

	consumerCluster struct {
		nodes []*consumerNode
		red   *redis.Client
	}
)