# GOTMAAZD

This is my Twitch bot Botmaazd developped in Go.

I already developped Botmaazd in JS using the tmi.js module, but the functionnalities were not enough for me and I want to get better at Go.

## Configuration

- Put your vars in the go/.env file
- Change the channels you want to look at in the `channels` splice at the top of main.go

## Goals

- Basic `!commands` (already done in JS)
- Channel points interaction
- Visual elements (last subscriber, last bit giver...)
- Visual alerts
- Sound alerts
- Dockerization

## nicklaw5/helix

Many thanks to nicklaw5 that developped this nice Twitch Helix API client written in Go. It will be very useful : https://github.com/nicklaw5/helix