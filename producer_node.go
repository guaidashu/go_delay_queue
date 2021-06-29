/**
 * @Author: yy
 * @Author: 1023767856@qq.com
 * @Date: 28/06/2021
 * @Desc: desc
 */

package go_delay_queue

type(
	producerNode struct {
		endpoint string
		tube     string
		conn     *connection
	}
)
