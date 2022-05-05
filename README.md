# gogram-bot

**A super light-weight library for [Telegram bot API](https://core.telegram.org/bots/api).**
---

It's very easy to pick up:

* bot.go: Contains Bot struct, which represents a bot and methods related to it.
* data.go: Contains the majority of structs that will be used to send something to telegram
(text, pic...), we call those structs data. Each data has two methods,Send and Check. 
Send sends the data to Request function and Check checks if required fields are empty or not
(also might check validation of data fields).
* utils.go: You don't really need it. In general, it adds data fields to request, and contains a very 
important function called Request (which you also won't use directly).
* data_test.go: All test are here. tests must be done on a real bot token and most tests need more than
just bot token. you must provide them by flags.

[//]: # (* types: )
