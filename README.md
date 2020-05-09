# Rest Server
## Description
+ Rest server implemented in GoLang.
+ Out-of-the-box feature.

## TODO List
- [x] Implement basic REST api (GET, POST, PUT & DELETE).
- [ ] Support token authentication.
- [ ] API call limit.
- [ ] Support http proxy.

## Usage
```
Usage of server:
  -port int
        specify the server listening port. (default 3000)
  -token string
        specify the private token. (default "token")
```

Example: `./server.exe -port 80 -token private`

You can use pm2 to start the binary.