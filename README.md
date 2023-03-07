# slack-connector
Imagine a magical package that acts as a bridge between the client and Slack API, providing a seamless connection with effortless integration. This package is the embodiment of simplicity, making communication with Slack a breeze. 
This is a basic package that will serve as an intermediary layer between the client and Slack API. Its primary function is to enable easy integration and connection with the Slack platform. By utilizing this package, users can enjoy a streamlined communication experience, with minimal effort required to establish a connection to the Slack API. Its straightforward design simplifies the communication process, freeing users from unnecessary complexities and allowing them to focus on productive collaboration.




## Integration
```go
    package main

    import (
        ....
        "github.com/pcpratheesh/slack-connector"
    )

    func main(){

        // initializaing new slack connector object
        client := slack.New(
            slack.WithToken("< put your token here >"),
            slack.WithChannelID("< put your channel id here >"),
            slack.WithSlackAPIInit(), // must need this
        )

        // push message into channel
        _, _, err := client.PushMessage("pushed new message into channel")
    }

```
## multiple channel connection
With this package, it is provisioned to connect / send messages into different channels with single initialization of objects.

```go
    package main

    import (
        ....
        "github.com/pcpratheesh/slack-connector"
    )

    func main(){

        // initializaing new slack connector object
        client := slack.New(
            slack.WithToken("< put your token here >"),
            slack.WithChannelID("channel-id-#0001"),
            slack.WithSlackAPIInit(), // must need this
        )

        // 
        _, _, err := client.PushMessage("pushed new message into channel")

        // 
        
        // this line of code would push message into channel channel-id-#0002
    	_, _, err := client.Channel("channel-id-#0002").PushMessage("pushed data into channel-id-#0002")


        // push message into channel channel-id-#0001
        // this line would push message into channel channel-id-#0001, because this is the channel we configured at initializing the slack connector.
        // Which means, it will act as default channel
        _, _, err := client.PushMessage("pushed new message into channel")

    }

```


```go
     _, _, err := client.PushMessage("pushed new message into channel")
```
The above line will push message into channel channel-id-#0001


```go
    _, _, err := client.Channel("channel-id-#0002").PushMessage("pushed data into channel-id-#0002")
```
We need to push data into another channel but the next time when we use `client` object it should push the message into the default one (channel-id-#0001)

```go
    _, _, err := client.PushMessage("pushed new message into channel")
```
The above line will push message into channel channel-id-#0001. Because this is the channel we configured at initializing the slack connector



## Soon-to-be-added features
There will be more features will be added soon