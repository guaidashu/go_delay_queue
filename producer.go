/**
 * @Author: yy
 * @Author: 1023767856@qq.com
 * @Date: 28/06/2021
 * @Desc: 生产者
 */

package go_delay_queue

import "time"

type (
	Producer interface {
		At(body []byte, at time.Time) (string, error)
		Close() error
		Delay(body []byte, delay time.Duration) (string, error)
		Revoke(ids string) error
	}

	producerCluster struct {
		nodes []Producer
	}
)