/**
 * @Author: yy
 * @Author: 1023767856@qq.com
 * @Date: 28/06/2021
 * @Desc: 消费者node
 */

package go_delay_queue

type (
	consumerNode struct {
		conn *connection
		tube string
		on   *AtomicBool
	}

	consumeService struct {
		c       *consumerNode
		consume Consume
	}
)