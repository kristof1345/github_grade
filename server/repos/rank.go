package repos

import (
	"fmt"
	"time"

	"github.com/google/go-github/github"
)

type Grades struct {
	Star    int
	Fork    int
	Issue   int
	Update  int
	Watcher int
	Prs     int
	Commits int
	Overall int
}

func Rank(repoStats Repository) []Grades {
	ret := make([]Grades, 0)
	starGrade := rateStars(repoStats.StarsCount)
	forkGrade := rateForks(repoStats.ForksCount)
	issuesGrade := rateIssues(repoStats.OpenIssuesCount)
	updatedGrade := rateUpdated(repoStats.LastUpdatedAt)
	watchersGrade := rateWatchers(repoStats.WatchersCount)
	prsGrade := ratePrs(repoStats.PullRequestsAYear)
	commitsGrade := rateCommits(repoStats.CommitsAYear)

	average := (starGrade + forkGrade + issuesGrade + updatedGrade + watchersGrade + prsGrade + commitsGrade) / 7

	fmt.Printf("stars: %d\n", starGrade)
	fmt.Printf("fork: %d\n", forkGrade)
	fmt.Printf("issue: %d\n", issuesGrade)
	fmt.Printf("updated: %d\n", updatedGrade)
	fmt.Printf("watchers: %d\n", watchersGrade)
	fmt.Printf("prs: %d\n", prsGrade)
	fmt.Printf("commits: %d\n", commitsGrade)

	grades := &Grades{
		Star:    starGrade,
		Fork:    forkGrade,
		Issue:   issuesGrade,
		Update:  updatedGrade,
		Watcher: watchersGrade,
		Prs:     prsGrade,
		Commits: commitsGrade,
		Overall: average,
	}

	ret = append(ret, *grades)

	return ret
}

func rateStars(starsCount int) int {
	switch {
	case starsCount >= 45000:
		return 10
	case starsCount < 45000 && starsCount >= 40000:
		return 9
	case starsCount < 40000 && starsCount >= 35000:
		return 8
	case starsCount < 35000 && starsCount >= 30000:
		return 7
	case starsCount < 30000 && starsCount >= 25000:
		return 6
	case starsCount < 25000 && starsCount >= 20000:
		return 5
	case starsCount < 20000 && starsCount >= 15000:
		return 4
	case starsCount < 15000 && starsCount >= 10000:
		return 3
	case starsCount < 10000 && starsCount >= 5000:
		return 2
	case starsCount < 5000 && starsCount >= 500:
		return 1
	default:
		return 0
	}
}

func rateForks(forkCount int) int {
	switch {
	case forkCount >= 4500:
		return 10
	case forkCount < 4500 && forkCount >= 4000:
		return 9
	case forkCount < 4000 && forkCount >= 3500:
		return 8
	case forkCount < 3500 && forkCount >= 3000:
		return 7
	case forkCount < 3000 && forkCount >= 2500:
		return 6
	case forkCount < 2500 && forkCount >= 2000:
		return 5
	case forkCount < 2000 && forkCount >= 1500:
		return 4
	case forkCount < 1500 && forkCount >= 1000:
		return 3
	case forkCount < 1000 && forkCount >= 500:
		return 2
	case forkCount < 500 && forkCount >= 50:
		return 1
	default:
		return 0
	}
}

func rateUpdated(lastTime github.Timestamp) int {
	timeString := lastTime.String()
	modDate := timeString[:10]
	date, err := time.Parse("2006-01-02", modDate)
	if err != nil {
		fmt.Println("Can't convert string back to time")
	}
	sinceUpdate := time.Since(date)
	days := sinceUpdate.Hours() / 24
	switch {
	case days >= 180:
		return 0
	case days < 180 && days >= 162:
		return 1
	case days < 162 && days >= 144:
		return 2
	case days < 144 && days >= 124:
		return 3
	case days < 124 && days >= 108:
		return 4
	case days < 108 && days >= 90:
		return 5
	case days < 90 && days >= 72:
		return 6
	case days < 72 && days >= 54:
		return 7
	case days < 54 && days >= 36:
		return 8
	case days < 36 && days >= 18:
		return 9
	case days < 18:
		return 10
	}
	return 0
}

func rateIssues(issues int) int {
	switch {
	case issues >= 90:
		return 10
	case issues < 90 && issues >= 80:
		return 9
	case issues < 80 && issues >= 70:
		return 8
	case issues < 70 && issues >= 60:
		return 7
	case issues < 60 && issues >= 50:
		return 6
	case issues < 50 && issues >= 40:
		return 5
	case issues < 40 && issues >= 30:
		return 4
	case issues < 30 && issues >= 20:
		return 3
	case issues < 20 && issues >= 10:
		return 2
	case issues < 10 && issues >= 5:
		return 1
	default:
		return 0
	}
}

func rateWatchers(watchers int) int {
	switch {
	case watchers >= 90:
		return 10
	case watchers < 90 && watchers >= 80:
		return 9
	case watchers < 80 && watchers >= 70:
		return 8
	case watchers < 70 && watchers >= 60:
		return 7
	case watchers < 60 && watchers >= 50:
		return 6
	case watchers < 50 && watchers >= 40:
		return 5
	case watchers < 40 && watchers >= 30:
		return 4
	case watchers < 30 && watchers >= 20:
		return 3
	case watchers < 20 && watchers >= 10:
		return 2
	case watchers < 10 && watchers >= 5:
		return 1
	default:
		return 0
	}
}

func rateCommits(count int) int {
	switch {
	case count >= 40:
		return 10
	case count < 40 && count >= 36:
		return 9
	case count < 36 && count >= 32:
		return 8
	case count < 32 && count >= 28:
		return 7
	case count < 28 && count >= 24:
		return 6
	case count < 24 && count >= 20:
		return 5
	case count < 20 && count >= 16:
		return 4
	case count < 16 && count >= 12:
		return 3
	case count < 12 && count >= 8:
		return 2
	case count < 8 && count >= 2:
		return 1
	}
	return 0
}

func ratePrs(count int) int {
	switch {
	case count >= 20:
		return 10
	case count < 20 && count >= 18:
		return 9
	case count < 18 && count >= 16:
		return 8
	case count < 16 && count >= 14:
		return 7
	case count < 14 && count >= 12:
		return 6
	case count < 12 && count >= 10:
		return 5
	case count < 10 && count >= 8:
		return 4
	case count < 8 && count >= 6:
		return 3
	case count < 6 && count >= 4:
		return 2
	case count < 4 && count >= 2:
		return 1
	}
	return 0
}
