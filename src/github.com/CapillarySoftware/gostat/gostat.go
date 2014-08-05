package main

import (
	"github.com/CapillarySoftware/gostat/stat"
	"github.com/CapillarySoftware/gostat/bucketer"
	"fmt"
	"time"
	"os"
	"os/signal"
	"math/rand"
	nano "github.com/op/go-nanomsg"
	log  "github.com/cihub/seelog"
)

func main () {
	stats            := make(chan *stat.Stat)   // stats received from producers
	bucketedStats    := make(chan []*stat.Stat) // raw bucketed (non-aggregated) stats are output here
	shutdownBucketer := make(chan bool)         // used to signal to the bucketer we are done
	shutdownListener := make(chan bool)         // used to signal to the socket listener we are done

  installCtrlCHandler(shutdownBucketer, shutdownListener)

	// create and start a Bucketer
	b := bucketer.NewBucketer(stats, bucketedStats, shutdownBucketer)
	go b.Run()

	// start a socket listener
	go bindSocketListener(stats, shutdownListener)

	// TODO: remove this simulation of stats when the socket listener is really processing data
	for true {
		<-time.After(time.Second * time.Duration(rand.Intn(3))) // sleep 0-3 seconds

		// create a state named "stat[1-10]" with a random value between 1-100
		stat := stat.Stat{Name : fmt.Sprintf("stat%v", (rand.Intn(9)+1)), Timestamp : time.Now().UTC(), Value : float64(rand.Intn(99)+1)}
		stats <- &stat // send it to the Bucketer
	}
}

// installCtrlCHandler starts a goroutine that will signal the workers when it's time
// to shut down
func installCtrlCHandler(shutdown ...chan<- bool) {
	c := make(chan os.Signal, 1)                                       
	signal.Notify(c, os.Interrupt)                                     

	go func() {                                                        
	  for sig := range c {                                             
	    log.Infof(" captured %v, stopping stats collection and exiting...\n", sig)
	    for _, s := range shutdown {
	    	s <-true
	    }
	    <-time.After(time.Second * 5) // wait for a clean shutdown, TODO: wait on signal from all routines
	    log.Info("Done")                                      
	    os.Exit(1)                                                     
	  }                                                                
	}()
}


func bindSocketListener(stats chan<- *stat.Stat, shutdown <-chan bool) {
	var (
		msg []byte
		err error
		done = false
	)
	socket, err := nano.NewPullSocket()

	if nil != err {
		log.Error(err)
	}

	socket.SetRecvTimeout(time.Second)

	defer socket.Close()
	_, err = socket.Bind("tcp://*:2025")
	if nil != err {
		log.Error(err)
	}
	log.Info("Ready to receive data")

	for !done {
		msg, err = socket.Recv(0) //blocking
		if nil != err {
			log.Error(err)
		}
		if nil != msg {
			// TODO: new up a Stat and send it on the stats channel
			log.Debug("Received message: ", msg)
		}

		select {
			case done = <- shutdown : break
			default                 : break
		}
	}

	log.Info("Exiting socket listener")
}



