# `/pkg`

外部应用程序可以使用的库代码（如`/pkg/mypubliclib`）。其他项目将会导入这些库来保证项目可以正常运行，所以在将代码放在这里前，一定要三四而行。

请注意,`/internal`目录是一个更好的选择来确保项目私有代码不会被其他人导入，因为这是`Go`强制执行的。使用`/pkg`目录来明确表示代码可以被其他人
安全的导入仍然是一个好方式。Travis Jeffery 撰写的关于 I’ll take pkg over internal 文章很好地概述了 pkg 和 inernal 目录以及何时使用它们。

当根目录包含大量非 Go 组件和目录时，这也是一种将 Go 代码分组到一个位置的方法，从而使运行各种 Go 工具更加容易（在如下的文章中都有提到：2018 年
GopherCon Best Practices for Industrial Programming，GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps，Golab 2018 Massimiliano Pippi - Project layout patterns in Go。

`/pkg`在许多开源项目中都使用了，但未被普遍接受，并且 Go 社区中的某些人不推荐这样做。

如果项目确实很小并且嵌套的层次并不会带来多少价值（除非你就是想用它），那么就不要使用它。但当项目变得很大，并且根目录中包含的内容相当繁杂（尤其是
有很多非 Go 的组件）时，可以考虑使用`/pkg`。

例子：
- https://github.com/prometheus/prometheus/tree/master/pkg
- https://github.com/jaegertracing/jaeger/tree/master/pkg
- https://github.com/istio/istio/tree/master/pkg
- https://github.com/GoogleContainerTools/kaniko/tree/master/pkg
- https://github.com/google/gvisor/tree/master/pkg
- https://github.com/google/syzkaller/tree/master/pkg
- https://github.com/perkeep/perkeep/tree/master/pkg
- https://github.com/minio/minio/tree/master/pkg
- https://github.com/heptio/ark/tree/master/pkg
- https://github.com/argoproj/argo/tree/master/pkg
- https://github.com/heptio/sonobuoy/tree/master/pkg
- https://github.com/helm/helm/tree/master/pkg
- https://github.com/kubernetes/kubernetes/tree/master/pkg
- https://github.com/kubernetes/kops/tree/master/pkg
- https://github.com/moby/moby/tree/master/pkg
- https://github.com/grafana/grafana/tree/master/pkg
- https://github.com/influxdata/influxdb/tree/master/pkg
- https://github.com/cockroachdb/cockroach/tree/master/pkg
- https://github.com/derekparker/delve/tree/master/pkg
- https://github.com/etcd-io/etcd/tree/master/pkg
- https://github.com/oklog/oklog/tree/master/pkg
- https://github.com/flynn/flynn/tree/master/pkg
- https://github.com/jesseduffield/lazygit/tree/master/pkg
- https://github.com/gopasspw/gopass/tree/master/pkg
- https://github.com/sosedoff/pgweb/tree/master/pkg
- https://github.com/GoogleContainerTools/skaffold/tree/master/pkg
- https://github.com/knative/serving/tree/master/pkg
- https://github.com/grafana/loki/tree/master/pkg
- https://github.com/bloomberg/goldpinger/tree/master/pkg
- https://github.com/Ne0nd0g/merlin/tree/master/pkg
- https://github.com/jenkins-x/jx/tree/master/pkg
- https://github.com/DataDog/datadog-agent/tree/master/pkg
- https://github.com/dapr/dapr/tree/master/pkg
- https://github.com/cortexproject/cortex/tree/master/pkg
- https://github.com/dexidp/dex/tree/master/pkg
- https://github.com/pusher/oauth2_proxy/tree/master/pkg
- https://github.com/pdfcpu/pdfcpu/tree/master/pkg
- https://github.com/weaveworks/kured/tree/master/pkg
- https://github.com/weaveworks/footloose/tree/master/pkg
- https://github.com/weaveworks/ignite/tree/master/pkg
- https://github.com/tmrts/boilr/tree/master/pkg
- https://github.com/kata-containers/runtime/tree/master/pkg
- https://github.com/okteto/okteto/tree/master/pkg
- https://github.com/solo-io/squash/tree/master/pkg

## 注意：

对于`/pkg`目录，如果是在微服务下，笔者更建议尽量不使用它。因为微服务，每个服务都会相对简单，也就是项目都比较小,`/pkg`不会带来多大价值。
如果有公用的代码，笔者更建议大家将这类代码做成私有库（`go module`），供其他项目复用，做了物理隔离，更有利于代码的抽象。

在 Go 语言中组织代码的方式还有一种叫”平铺“的，也就是在根目录下放项目的代码。这种方式在很多框架或者库中非常常见，如果想要引入一个使用
`/pkg`目录结构的框架时，我们往往需要使用`github.com/golang/project/pkg/somepkg`，当代码都平铺在项目的根目录时只需要使用
`github.com/golang/project`，很明显地减少了引用依赖包语句的长度。所以，对于一个 Go 语言的框架或者库，将代码平铺在根目录下也很正常，
但是在一个 Go 语言的服务中使用这种代码组织方法可能就没有那么合适了。