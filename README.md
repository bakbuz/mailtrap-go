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

### Email Sandbox (Testing)
- Send an email
- Send an email with a template
- Message management
- List inboxes
- Get inbox details
- Update inbox settings
- Create, read, update, and delete projects
- Project configuration and settings

### General
- Account access management
- Permissions management
- List accounts you have access to
- Billing information and usage statistics
- Domain verification

## Quick Start
The following few simple steps will bring Mailtrap API functionality into your Go project.

### Install
The Mailtrap Go client packages are available through GitHub Packages.

Add Mailtrap package:

```console
go get -u github.com/bakbuz/mailtrap-go
```

### Configure
Add Mailtrap services to the DI container.

```csharp
using Mailtrap;
   
...
   
hostBuilder.ConfigureServices((context, services) =>
{
   services.AddMailtrapClient(options =>
   {
      // Definitely, hardcoding a token isn't a good idea.
      // This example uses it for simplicity, but in real-world scenarios
      // you should consider more secure approaches for storing secrets.
         
      // Environment variables can be an option, as well as other solutions:
      // https://learn.microsoft.com/aspnet/core/security/app-secrets
      // or https://learn.microsoft.com/aspnet/core/security/key-vault-configuration
      options.ApiToken = "<API_TOKEN>";
   });
});   
```

### Use
Now you can inject `IMailtrapClient` instance in any application service and use it to make API calls.

```csharp
using Mailtrap;
using Mailtrap.Emails.Requests;
using Mailtrap.Emails.Responses;
   

namespace MyAwesomeProject;


public sealed class SendEmailService : ISendEmailService
{
    private readonly IMailtrapClient _mailtrapClient;


    public SendEmailService(IMailtrapClient mailtrapClient)
    {
        _mailtrapClient = mailtrapClient;
    }


    public async Task DoWork()
    {
        try 
        {
            SendEmailRequest request = SendEmailRequest
                .Create()
                .From("john.doe@demomailtrap.com", "John Doe")
                .To("hero.bill@galaxy.net")
                .Subject("Invitation to Earth")
                .Text("Dear Bill,\n\nIt will be a great pleasure to see you on our blue planet next weekend.\n\nBest regards, John.");

            SendEmailResponse response = await _mailtrapClient
                .Email()
                .Send(request);
                
            string messageId = response.MessageIds.FirstOrDefault();
        }
        catch (MailtrapException mtex)
        {
            // handle Mailtrap API specific exceptions
        }
        catch (OperationCancelledException ocex)
        {
            // handle cancellation
        }
        catch (Exception ex)
        {
            // handle other exceptions
        }   
    }
}
```
