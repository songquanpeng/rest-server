# Rest Server
## Description
+ Rest server implemented in GoLang.
+ Out-of-the-box feature.
+ Use it as your toy project's backend, so you can focus on the frontend.
+ Quite easy to deploy, only one binary file.
+ [Demo](https://iamazing.cn/page/rest-server-demo).

## TODO List
- [x] Implement basic REST api (GET, POST, PUT & DELETE).
- [ ] Support token authentication.
- [ ] API call limit.
- [ ] Support http proxy.

## Usage
```
Usage of rest-server:
  -port int
        specify the listening port. (default 3000)
  -token string
        specify the private token. (default "token")
```

Example: `./server.exe -port 80 -token private`

You can use pm2 to start the binary: `pm2 start ./rest-server -- -port 3002 -token private`