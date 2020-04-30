package main

import (
	"context"

	"github.com/google/go-github/v31/github"
	"golang.org/x/oauth2"
)

var (
	client *github.Client
	o      *github.Organization
	t      *github.Team
)

func initGithub(token string) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	client = github.NewClient(tc)
}

func invite(user, org, team string) error {
	var err error
	if o == nil {
		o, _, err = client.Organizations.Get(context.Background(), org)
		if err != nil {
			o = nil
			return err
		}
	}
	if t == nil {
		t, _, err = client.Teams.GetTeamBySlug(context.Background(), org, team)
		if err != nil {
			t = nil
			return err
		}
	}

	_, _, err = client.Teams.AddTeamMembershipByID(context.Background(), o.GetID(), t.GetID(), user, &github.TeamAddTeamMembershipOptions{
		Role: "member",
	})
	return err
}
