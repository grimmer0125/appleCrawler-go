# appleCrawler-go

It is a Line Bot and also an automation Bot to grab the updated information from http://www.apple.com/tw/shop/browse/home/specialdeals/mac then send to its Line Friends. 

It will send the updated information to its Line Friend. 

Scan this QR Code to add it as your Line Friend. 

![Image of LineBotID](https://grimmer.io/images/qr-code-apple-line-bot.png)

## How to run
1. setup the needed enviornment, such as postgres and Line bot API,  follow the [Steps for development](https://github.com/grimmer0125/appleCrawler-node#steps-for-development) without node/NPM parts. 
2. Install go and setup its PATH
3. Install dependency tool, Godep, `go get github.com/tools/godep`
4. Install dependencies, `godep restore`
5. Use either the below ways to startup
    1. Use Visual Studio Code
    2. Use command line    
        1. `go install`(binary file will be the same folder) or `go build` (binary will be saved in the GOPATH/bin)
        2. `appleCrawler-go`

## License

appleCrawler-go is released under the [Apache 2.0 license][license].

[license]: LICENSE.md


