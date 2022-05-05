# gogram-bot

**A super light-weight library for [Telegram bot API](https://core.telegram.org/bots/api).**

It's very easy to pick up:

* bot.go: Contains Bot struct, which represents a bot, methods related to it.
* data.go: Contains the majority of methods (sending text, pic...), we call them data. Each data has two methods,
Send and Check. Send sends the data to Request function and Check checks if required fields are empty or not
(also might check validation of data fields).


