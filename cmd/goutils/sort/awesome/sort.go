package awesome

import (
	"context"
	"sort"
	"strings"

	"github.com/google/go-github/v50/github"
	_errors "github.com/liblaf/goutils/pkg/errors"
)

func sortSection(ctx context.Context, list []string, token string) string {
	var client *github.Client
	if token != "" {
		client = github.NewTokenClient(ctx, token)
	} else {
		client = github.NewClient(nil)
	}

	repos := make([]*github.Repository, len(list))
	for i, item := range list {
		owner, repo, _ := strings.Cut(item, "/")
		repository, _, err := client.Repositories.Get(ctx, owner, repo)
		_errors.Check(err)
		repos[i] = repository
	}

	sort.Slice(repos,
		func(i, j int) bool {
			return repos[i].GetStargazersCount() > repos[j].GetStargazersCount()
		},
	)

	buffer := new(strings.Builder)
	for i, repo := range repos {
		list[i] = repo.GetFullName()
		err := templateRow.Execute(buffer, repo)
		_errors.Check(err)
		_, err = buffer.WriteString("\n")
		_errors.Check(err)
	}

	return buffer.String()
}

func sortAll(ctx context.Context, groups map[string][]string, token string) string {
	keys := make([]string, 0, len(groups))
	for k := range groups {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	buffer := new(strings.Builder)
	for _, k := range keys {
		section := titleCase(k)
		items := groups[k]
		rows := sortSection(ctx, items, token)
		groups[k] = items
		err := templateSection.Execute(buffer, map[string]string{
			"Section": section,
			"Rows":    rows,
		})
		_errors.Check(err)
	}

	return buffer.String()
}
