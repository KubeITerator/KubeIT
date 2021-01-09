module kubeIT

go 1.15

replace sigs.k8s.io/controller-tools => sigs.k8s.io/controller-tools v0.2.9

require (
	github.com/argoproj/argo v0.0.0-20201102205842-bf3fec176cf6
	github.com/aws/aws-sdk-go v1.33.16
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/imdario/mergo v0.3.11 // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	k8s.io/api v0.17.8
	k8s.io/apimachinery v0.17.8
	k8s.io/client-go v0.17.8
	k8s.io/utils v0.0.0-20201104234853-8146046b121e // indirect
)
