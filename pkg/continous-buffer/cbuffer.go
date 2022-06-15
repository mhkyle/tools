package continous_buffer

import (
	"bytes"
	"fmt"
	"io"
	"sync"
)

type MyDataBucket struct {
	br     *bytes.Buffer
	gmutex *sync.RWMutex
	rcond  *sync.Cond //读操作需要用到的条件变量
}

func NewDataBucket() *MyDataBucket {
	buf := make([]byte, 0)
	db := &MyDataBucket{
		br:     bytes.NewBuffer(buf),
		gmutex: new(sync.RWMutex),
	}
	db.rcond = sync.NewCond(db.gmutex.RLocker())
	return db
}

//读的goroutine需要与锁合用，通过锁进行封锁自己，然后收到cond的信号之后才解封自己
func (db *MyDataBucket) Read(i int) {
	db.gmutex.RLock()
	defer db.gmutex.RUnlock()
	var data []byte
	var d byte
	var err error
	for {
		//读取一个字节
		if d, err = db.br.ReadByte(); err != nil {
			if err == io.EOF {
				if string(data) != "" {
					fmt.Printf("reader-%d: %s\n", i, data)
				}
				db.rcond.Wait()
				data = data[:0]
				continue
			}
		}
		data = append(data, d)
	}
}

//写数据的goroutine,写完毕之后发出broadcast信号，来通知所有人解除封锁。
func (db *MyDataBucket) Put(d []byte) (int, error) {
	db.gmutex.Lock()
	defer db.gmutex.Unlock()
	//写入一个数据块
	n, err := db.br.Write(d)
	db.rcond.Broadcast()
	return n, err
}
