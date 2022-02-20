1. [Functions](function.go)
1. [Methods](methods.go)
1. [Interfaces](interface.go)

## Concurrency

- Definition: Computations happening at the same time
- Even on single CPU (kernel makes that feeling)

## Communication via message queues

- Analogy of phone calls
- Phone ringing: blocking call
- Answerphone: non-blocking call (tape length = queue size)
- Multiple receivers: Call centre
- Multiple senders: Conf call

## Threads and goroutines

- Reminder: Threads share memory
- Go routines are light-weighted threads, cheap to spawn
- It should be fine to spawn thousands of goroutines
- Internally multiplexed across thread pool
- Spawned with simple go keyword

## Channels

- Way to transfer data between goroutines
- Type of data is part of the channel type
- By default blocks on both tx/rx (but can be buffered too)
- Same operator to send and receive from channel
- Created with make()

### Closing channels

- Channels can be closed with close()
- Readers can check for closed channels using optional second return argument
- Alternatively readers can detect closed channel with the loop over range

1. [Goroutines](goroutine.go)
2. [Channels](chan.go)
3. [Sync](sync.go)
4. [Atomic](atomic.go)
5. [Mutex](mutex.go)
6. [File](file.go)
7. [Context](context.go)
8. [Time](time.go)
