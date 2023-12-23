package main

/*
This is a excersice of my own. Quickly designed to play arround concurrency.

Is about a queue thats get enqueued and dequeued constantly

*/
import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var mu sync.Mutex
var wg sync.WaitGroup
var logger *log.Logger = log.Default()
var r randSecure = randSecure{
	mut: sync.Mutex{},
	r:   rand.New(rand.NewSource(time.Now().UnixNano() / 3)),
}

type randSecure struct {
	r   *rand.Rand
	mut sync.Mutex
}

type Client struct {
	id         int
	name       string
	likeCookie bool
	enterTime  time.Time
	// time       time.Duration
}

type ClientQueue struct {
	length        atomic.Int32
	capacity      atomic.Int32
	arr           []Client
	indexPush     atomic.Int32
	indexPop      atomic.Int32
	waitToBeEmpty bool
}

func (c *ClientQueue) Default() {
	arr := make([]Client, 90000)
	c.length.Store(0)
	c.capacity.Store(90000)
	c.arr = arr
	c.indexPush.Store(0)
	c.waitToBeEmpty = false
}

func (c *ClientQueue) add(v *Client) {
	mu.Lock()
	defer mu.Unlock()
	if c.length == c.capacity {
		return
	}
	v.id = int(c.indexPush.Load())
	c.arr[c.indexPush.Load()] = *v
	c.indexPush.Add(1)
	if c.indexPush == c.capacity {
		c.indexPush.Store(0)
	}
	c.length.Add(1)
}

func (c *ClientQueue) remove() (Client, bool) {
	mu.Lock()
	defer mu.Unlock()
	if c.length.Load() == 0 {
		return Client{}, false
	}
	c.length.Add(-1)
	e := c.arr[c.indexPop.Load()]
	c.indexPop.Add(1)
	if c.indexPop == c.capacity {
		c.indexPop.Store(0)
	}
	if c.waitToBeEmpty && c.length.Load() == 0 {
		wg.Done()
	}
	return e, true
}

func main() {
	queue := ClientQueue{}
	queue.Default()
	fmt.Println(queue.capacity.Load())

	wg.Add(1)

	go addClient(&queue)

	for i := 0; i < 20; i++ {
		go serveClient(&queue)
	}
	wg.Wait()
	// let the current go runtines who still are manging clients finish
	time.Sleep(150 * time.Millisecond)
}

func addClient(queue *ClientQueue) {
	for i := 0; i < 90000; {
		if queue.length.Load() < queue.capacity.Load() {
			// logger.Println(i)
			client := generatedRandomClient()
			queue.add(&client)
			i++
		}
	}
	wg.Add(1)
	queue.waitToBeEmpty = true
	wg.Done()
}

func serveClient(queue *ClientQueue) {
	for {
		if queue.length.Load() != 0 {
			c, err := queue.remove()
			if !err {
				continue
			}
			r.mut.Lock()
			f := r.r.Float32()
			r.mut.Unlock()
			s := int(f * 150)
			time.Sleep(time.Duration(s) * time.Nanosecond)
			val, delay := valoracion(&c.enterTime)
			logger.Println("index", c.id, "valoracion", val, "delay", delay)
		}
	}
}

func valoracion(t *time.Time) (int8, int64) {
	elpased := time.Since(*t).Milliseconds()

	if elpased < 10 {
		return 5, int64(elpased)
	} else if elpased < 100 {
		return 4, int64(elpased)
	} else if elpased < 500 {
		return 3, int64(elpased)
	} else if elpased < 1000 {
		return 2, int64(elpased)
	} else {
		return 1, int64(elpased)
	}
}

func generatedRandomClient() Client {
	r.mut.Lock()
	f := r.r.Float32()
	r.mut.Unlock()
	return Client{name: genRandName(), likeCookie: f < 0.5, enterTime: time.Now()}
}

func genRandName() string {
	name := [4]string{"Carlos", "Alberto", "Amelia", ""}

	r.mut.Lock()
	f := r.r.Float32()
	r.mut.Unlock()
	if f < 0.25 {
		return name[0]
	} else if f < 0.5 {
		return name[1]
	} else if f < 0.75 {
		return name[2]
	} else {
		return name[3]
	}
}
