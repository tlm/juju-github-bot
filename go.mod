module github.com/tlm/juju-github-bot

go 1.13

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.2.0+incompatible
	k8s.io/api => k8s.io/api v0.17.3
	k8s.io/apimachinery => k8s.io/apimachinery v0.17.3
	k8s.io/client-go => k8s.io/client-go v0.17.3
)

require (
	github.com/sirupsen/logrus v1.4.2
	k8s.io/test-infra v0.0.0-20200315172834-05821564b6e4
)
