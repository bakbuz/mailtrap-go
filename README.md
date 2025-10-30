# Unofficial Mailtrap Go Client
Welcome to the unofficial [Mailtrap](https://mailtrap.io/) Go Client repository.  
This client allows you to quickly and easily integrate your Go application with **v2.0** of the [Mailtrap API](https://api-docs.mailtrap.io/docs/mailtrap-api-docs/5tjdeg9545058-mailtrap-api).


## Prerequisites

To get the most out of this official Mailtrap.io Go SDK:
- [Create a Mailtrap account](https://mailtrap.io/signup)
- [Verify your domain](https://mailtrap.io/sending/domains)
- Obtain [API token](https://mailtrap.io/api-tokens)

## Supported functionality

The Mailtrap Go client provides comprehensive access to the Mailtrap API v2.0, including:

### Email API/SMTP
- Send an email (Transactional and Bulk streams)
- Send an email with a template
- Send emails with attachments

## Quick Start
The following few simple steps will bring Mailtrap API functionality into your Go project.

### Install
The Mailtrap Go client packages are available through GitHub Packages.

Add Mailtrap package:

```console
go get -u github.com/bakbuz/mailtrap-go
```

### Use
Now you can inject `EmailClient` instance in any application service and use it to make API calls.

```GO
// client
var client emails.EmailClient = emails.NewEmailClient("<API_KEY>")
   

// request
req := request.Create().
	WithFrom("john.doe@demomailtrap.com", "John Doe").
	WithTo("hero.bill@galaxy.net", "").
	WithSubject("Invitation to Earth").
	WithTextBody("Dear Bill,\n\nIt will be a great pleasure to see you on our blue planet next weekend.\n\nBest regards, John.")

// response
res, err := client.Send(context.Background(), req)
if err != nil {
	log.Fatal(err)
}

messageIds := res.MessageIds
log.Println(messageIds)

```
