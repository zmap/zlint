//package defines a blocking RingBuffer that can hold []string

package ringbuff

import (
	"sync"
)

type RingBuffer struct {
	readcomplete bool
	readvar      *sync.Cond
	writevar     *sync.Cond
	ringtex      sync.Mutex
	data         [][]string
	sindex       int //start
	eindex       int //end
}

func (buff *RingBuffer) Init(size int) {
	buff.readcomplete = false
	buff.ringtex = sync.Mutex{}
	buff.readvar = sync.NewCond(&buff.ringtex)
	buff.writevar = sync.NewCond(&buff.ringtex)
	buff.data = make([][]string, size+1)
	buff.sindex = 0
	buff.eindex = 0
}

func (buff *RingBuffer) Dequeue() (ret []string) {
	//entering RingBuffer control
	buff.ringtex.Lock()

	for buff.sindex == buff.eindex && !buff.readcomplete {
		buff.readvar.Wait() //ringtex unlocked while waiting
	}
	if buff.readcomplete && buff.sindex == buff.eindex { //woken by read end or called after read end
		//no data and no more coming
		//exiting RingBuffer control
		buff.ringtex.Unlock()
		return nil //nothing to return
	}

	//read data from RingBuffer
	ret = buff.data[buff.sindex] //read next pieace of data
	buff.data[buff.sindex] = nil //remove reference to allow data to be cleared after use

	buff.sindex++                      //move up start index
	if buff.sindex >= len(buff.data) { //past end of buffer
		buff.sindex = 0 //return to start
	}
	//read is complete, wake the writer if it is waiting
	buff.writevar.Signal()

	//exiting RingBuffer control
	buff.ringtex.Unlock()

	return ret
}

/* Enqueues a new chunk of data into the RingBuffer
 * This function will block until there is enough space to enqueue if the RingBuffer is full
 *
 */
func (buff *RingBuffer) Enqueue(chunk []string) {
	//entering RingBuffer control
	buff.ringtex.Lock()

	for buff.sindex == ((buff.eindex+1)%len(buff.data)) && !buff.readcomplete {
		//no space to write into, wait for space
		buff.writevar.Wait()
	}
	if buff.readcomplete {
		goto exitControl //queue is poisoned, return
	}
	//space available

	buff.data[buff.eindex] = chunk //slices are references, copied by value

	//move end index to space to be filled next
	buff.eindex++
	if buff.eindex >= len(buff.data) { //past end of buffer
		buff.eindex = 0 //return to start
	}

	//new chunk available, wake a reader
	buff.readvar.Signal()

exitControl: //exiting RingBuffer control
	buff.ringtex.Unlock()
	//return
}

/* A call to this function poisons the ring buffer, signaling that name more data will be enqueued.
 * All waiting readers will wake up, and new calls to Enqueue/Dequeue will not block.
 * New calls to Enqueue (and blocked calls) will have no effect on the RingBuffer, but Dequeue will
 * still remove data if it is already Enqueued.
 */
func (buff *RingBuffer) Poison() {
	//entering RingBuffer control
	buff.ringtex.Lock()

	buff.readcomplete = true

	//no more data coming, wake all waiting threads so that they can exit
	buff.readvar.Broadcast()
	buff.writevar.Broadcast()

	//exiting RingBuffer control
	buff.ringtex.Unlock()
}
