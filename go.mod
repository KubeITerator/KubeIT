module kubeIT

go 1.15

replace sigs.k8s.io/controller-tools => sigs.k8s.io/controller-tools v0.2.9

require (
	github.com/argoproj/argo v0.0.0-20201102205842-bf3fec176cf6
	github.com/aws/aws-sdk-go v1.33.16
	github.com/coreos/go-oidc/v3 v3.0.0
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.14.6
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.4.0
	github.com/sirupsen/logrus v1.6.0
	github.com/srikrsna/protoc-gen-gotag v0.5.0
	go.mongodb.org/mongo-driver v1.3.1
	golang.org/x/oauth2 v0.0.0-20210427180440-81ed05c6b58c
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	google.golang.org/genproto v0.0.0-20210426193834-eac7f76ac494
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.26.0
	k8s.io/api v0.17.8
	k8s.io/apimachinery v0.17.8
	k8s.io/client-go v0.17.8
	k8s.io/utils v0.0.0-20201104234853-8146046b121e // indirect
)
