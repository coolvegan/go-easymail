# GO Easy Email

This package is a simple go-mail wrapper for direct or lazy send emails. The Emailer Interface
needs two Methods for initializing go-mail. It sends formatted html mails if you put html tags inside.

###### Interface

``` go
type Emailer interface {
	Send(sender, recipient, subject, body string) error
	newEmail() *mail.Client
}
```


##### Direct Mail Init
```go
var e Emailer
e = NewEmail(os.Getenv("SMTPSERVER"), os.Getenv("SMTPUSER"), os.Getenv("SMTPPASSWORD"))
e.Send(RecipientName, RecpientEmail, Subject, Message)

```

##### Lazy Mail Init
The LazyMail structure runs a background job from the start of initializing it until the main application ends. You can inject your own method to handle the background job from the start or just take basicEmailSend. It takes the interval as milliseconds how often it should dequeue a message from the email queue and send it via go-mail. For unit testing there is although lazyunittest which i used to similuate the dequeing in time.
```go
var e Emailer
e = NewLazymail(os.Getenv("SMTPSERVER"), os.Getenv("SMTPUSER"), os.Getenv("SMTPPASSWORD"), basicEmailSend(1000))
e.Send(RecipientName, RecpientEmail, Subject, Message)

```
####


##### Running from Command Line
SMTPSERVER=smtp.foo.com SMTPUSER=user@foo.com SMTPPASSWORD=secret RECNAME=foobar
ENDER RECMAIL=foobar@foobar.com SUBJECT="Subject" MESSAGE="message"  go run .
