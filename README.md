# apple-server
## Usage
 To get the version
```bash
curl -X GET  "http://127.0.0.1:3000/api/v1/version"
```
 To post Hello
```bash
curl -X POST -d '{"name":"Your Name"}' "http://127.0.0.1:3000/api/v1/hello"
```