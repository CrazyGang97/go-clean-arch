# `/api`

项目对外提供和依赖的 API 文件。比如：OpenAPI/Swagger specs, JSON schema 文件, protocol 定义文件等。

比如，Kubernetes 项目的 api 目录结构如下：
```text
api
├── api-rules
    └── xxx.plist
├── openapi-spec
    └── swagger.json
```

另外：在 go 中用的比较多的 gRPC proto 文件，也比较适合放在 api 目录下。
```text
api
└── protobuf-spec
    └── test
        ├── test.pb.go
        └── test.proto
```

## 注意：
一般而言，微服务的协议定义都应该收敛集中存放【方便互相引用】，另外协议定义理应与项目代码分开解耦。故`/api`目录不推荐使用。