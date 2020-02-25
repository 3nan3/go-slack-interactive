# Description

This is a copy of https://github.com/tcnksm/go-slack-interactive

The changes are as follows:

- Use github.com/slack-go/slack
- Change bot mention to slash command
- Change RTM API to Event API
  - > New bot user API access tokens may not access RTM  
    https://github.com/slack-go/slack/issues/654#issuecomment-578946919
    
    :thinking_face:

![](/screencapture_suimasen.gif)

# Set up

## Create app

Create an app with any name.

![](/screenshot_create_app.png)

Copy `Verification Token` at `Basic Infomation > App Credentials`

![](/screenshot_varification_token.png)

!!! CAUTION !!!
> This deprecated Verification Token can still be used to verify that requests come from Slack, but we strongly recommend using the above, more secure, signing secret instead.

## Run Server

```sh
go build -o bot
VERIFICATION_TOKEN=kYBhXXXXXXXXXXXXXXXXXXXX ./bot
```

## Set up the slack app

Settings > Basic Information  
![](/screenshot_basic_info.png)

Features > Oauth & Permissions > Scopes  
![](/screenshot_scopes.png)

Features > Slash Commands  
![](/screenshot_slach_commands.png)

Features > Interacive Components  
![](/screenshot_interactive_components.png)

## Add the app to Slack

Add the app you just created on a channel.
Then execute `/sumimasen` command.
Cheers :beer:
