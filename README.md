# Telegom
Telegram bot written in Go

### Install Telegom
```bash
go get -u github.com/vicmanbrile/telegom
```

### Quick start

```go
package main

func main(){
	TLG := telegom.InitTelegom("TELEGRAM_BOT_TOKEN", "MONGODB_URL_CONNECTION")
	
	// Add a command telegram /start
	
	TLG.Handle("/start",  func(response server_response.ServerResponse, update *api.Update) {
		response.SendText("Welcome to my telegram bot written in telegom")
	})
}
```