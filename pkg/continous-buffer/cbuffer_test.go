package continous_buffer

import (
	"fmt"
	"testing"
	"time"
)

func TestNewDataBucket(t *testing.T) {
	db := NewDataBucket()

	t1 := time.Now()
	//两个消费者
	go db.Read(1)
	go db.Read(2)

	//10个生产者，每个生产者都放一个字符串，其中每一个生产者生产完之后就通知broadcast一下各位一次写完毕，然后让消费者就可以进行消费了。
	for i := 0; i < 99999; i++ {
		go func(i int) {
			d := fmt.Sprintf("data-%d", i)
			db.Put([]byte(d))
			time.Sleep(10 * time.Millisecond)
		}(i)
	}

	t2 := time.Now()
	fmt.Printf("time duration : %v", t2.Sub(t1))

}
