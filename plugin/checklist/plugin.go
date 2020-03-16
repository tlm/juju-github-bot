package checklist

import (
	"strings"

	"github.com/sirupsen/logrus"

	"k8s.io/test-infra/prow/github"
)

type githubClient interface {
	CreateComment(org, repo string, number int, comment string) error
}

func HandlePullRequestEvent(
	log *logrus.Entry,
	ghc githubClient,
	pre *github.PullRequestEvent) error {

	if pre.Action != github.PullRequestActionOpened {
		return nil
	}

	msgBuilder := strings.Builder{}
	err := checklist.Execute(&msgBuilder, pre)
	if err != nil {
		return err
	}

	return ghc.CreateComment(pre.Repo.Owner.Login, pre.Repo.Name, pre.Number,
		msgBuilder.String())
}
