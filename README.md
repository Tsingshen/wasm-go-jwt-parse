# parse http params toke=jwt payload.id to x-Xxx-userid
- urllike: http://xxxx.com/header?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWQiOjExMTEyMjIyLCJpYXQiOjE1MTYyMzkwMjJ9.Lp-EEKsLfOUrHlkvUNskrRJDg4UU1Wt4P45xFEO-OvU&name=aaa
```
{
    "headers": {
        "Accept": "*/*",
        "Accept-Encoding": "gzip, deflate, br",
        "Host": "xxx.com",
        "Postman-Token": "xxx",
        "User-Agent": "PostmanRuntime/7.29.2",
        "X-B3-Parentspanid": "63d1f04e21c976c6",
        "X-B3-Sampled": "0",
        "X-B3-Spanid": "f32c6f2c34cb3ba4",
        "X-B3-Traceid": "31af265a23346a7563d1f04e21c976c6",
        "X-Envoy-Attempt-Count": "1",
        "X-Envoy-External-Address": "xxx",
        "X-Xxx-Userid": "11112222"
    }
}
```
