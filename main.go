package eazymail

import (
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	//Just testing
	var (
		RecipientName = os.Getenv("RECNAME")
		RecpientEmail = os.Getenv("RECMAIL")
		Subject       = os.Getenv("SUBJECT")
		Message       = os.Getenv("MESSAGE")
	)
	var e Emailer
	var wg sync.WaitGroup
	e = NewLazymail(os.Getenv("SMTPSERVER"), os.Getenv("SMTPUSER"), os.Getenv("SMTPPASSWORD"), SendUpToEightEmailsAndThenDelay(4000))
	// e = NewEmail(os.Getenv("SMTPSERVER"), os.Getenv("SMTPUSER"), os.Getenv("SMTPPASSWORD"))
	wg.Add(1)
	go func() {
		f := 0
		for {
			f++
			err := e.Send(RecipientName, RecpientEmail, Subject+"_"+strconv.Itoa(f), Message)
			if err != nil {

				log.Println(err)
			}
			log.Println("Sende in Main")
			time.Sleep(10 * time.Millisecond)
			if f >= 100 {
				f = 0
				log.Println("8 Sekunden Sendepause")
				time.Sleep(20000 * time.Millisecond)
				wg.Done()
			}
		}
	}()
	wg.Wait()
}
