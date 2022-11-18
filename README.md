# parse http params toke=jwt payload.id to x-xxx-userid
- urllike: http://xxxx.com/headers?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWQiOjExMTEyMjIyLCJpYXQiOjE1MTYyMzkwMjJ9.Lp-EEKsLfOUrHlkvUNskrRJDg4UU1Wt4P45xFEO-OvU&name=aaa
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
### 性能简单测试结论：
相对来说，使用了wasmplugin的性能损耗很小（单网关(8C4G)，单pod(2C1G),），完全可以接受。
rust性能比go好一点点，rust的资源利用也比go好一点点。
并发3000时，测得（rust|go|nowasm,cpu:1494m,1464m,1340m;lat:0.0541s|0.0562s|0.0502s）
