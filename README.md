# Team Inviter

A tool helps you invite members to a team.

## Usage

```sh
go build
./team-inviter -token {your token} -user {invited person} -team {team name} -org {org name}
```

**The token must have access for orgs, unless you will get `401 Bad credentials`.**

![token access](./images/access.png)
