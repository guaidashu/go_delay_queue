/**
 * @Author: yy
 * @Author: 1023767856@qq.com
 * @Date: 28/06/2021
 * @Desc: 连接管理器
 */

package go_delay_queue

import (
	"github.com/beanstalkd/go-beanstalk"
	"sync"
)

type(
	connection struct {
		lock     sync.RWMutex
		endpoint string
		tube     string
		conn     *beanstalk.Conn
	}
)