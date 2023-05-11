package repos

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Repository struct {
	StarsCount        int
	ForksCount        int
	LastUpdatedAt     github.Timestamp
	OpenIssuesCount   int
	WatchersCount     int
	PullRequestsAYear int
	CommitsAYear      int
}

func GetRepo(s []string) *Grades {
	owner := s[0]
	name := s[1]
	ctx := context.Background()

	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tokenClient := oauth2.NewClient(ctx, tokenService)

	client := github.NewClient(tokenClient)

	repo, _, repoErr := client.Repositories.Get(ctx, owner, name)

	if repoErr != nil {
		fmt.Printf("Problem in getting repository information %v\n", repoErr)
		os.Exit(1)
	}

	prs := getPrs(ctx, owner, name, *client)

	commits := getCommits(ctx, owner, name, *client)

	repoStats := &Repository{
		ForksCount:        *repo.ForksCount,
		StarsCount:        *repo.StargazersCount,
		LastUpdatedAt:     *repo.UpdatedAt,
		OpenIssuesCount:   *repo.OpenIssuesCount,
		WatchersCount:     *repo.WatchersCount,
		PullRequestsAYear: len(prs),
		CommitsAYear:      len(commits),
	}

	fmt.Println(repoStats)

	grade := Rank(*repoStats)

	return grade
}

func getPrs(ctx context.Context, owner string, name string, client github.Client) []*github.PullRequest {
	var pullRequests []*github.PullRequest

	opts := &github.PullRequestListOptions{
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	}

	now := time.Now()
	oneYearAgo := now.AddDate(-1, 0, 0)
	startDate := time.Date(oneYearAgo.Year(), oneYearAgo.Month(), oneYearAgo.Day(), 0, 0, 0, 0, time.UTC)
	endDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	prs, _, prsErr := client.PullRequests.List(ctx, owner, name, opts)

	if prsErr != nil {
		fmt.Printf("Problem in getting repository pull requests information %v\n", prsErr)
	}

	for _, pr := range prs {
		if pr.CreatedAt.After(startDate) && pr.CreatedAt.Before(endDate) {
			pullRequests = append(pullRequests, pr)
		}
	}

	return pullRequests
}

func getCommits(ctx context.Context, owner string, name string, client github.Client) []*github.RepositoryCommit {
	opts := &github.CommitsListOptions{
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
		Since: time.Now().AddDate(-1, 0, 0),
		Until: time.Now(),
	}

	commits, _, commitsErr := client.Repositories.ListCommits(ctx, owner, name, opts)

	if commitsErr != nil {
		fmt.Printf("Problem in getting repository pull requests information %v\n", commitsErr)
	}

	return commits
}
