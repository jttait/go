package main

import (
	"log"
	"html/template"
	"os"
	"gopl.io/ch4/github"
)

var milestoneList = template.Must(template.New("milestoneList").Parse(`
<h1>Milestones</h1>
<table>
<tr>
	<th>Number</th>
	<th>URL</th>
</tr>
{{range .}}
<tr>
	<td>{{.Number}}</td>
	<td>{{.URL}}</td>
</tr>
{{end}}
</table>
`))

var userList = template.Must(template.New("userList").Parse(`
<h1>Users</h1>
<table>
<tr>
	<th>Login</th>
	<th>HTMLURL</th>
</tr>
{{range .}}
<tr>
	<td>{{.Login}}</td>
	<td><a href='{{.HTMLURL}}'>{{.HTMLURL}}</a></td>
</tr>
{{end}}
</table>
`))

var issuesList = template.Must(template.New("issuesList").Parse(`
<h1>Issues</h1>
<table>
<tr>
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Title</th>
	<th>Milestone</th>
</tr>
{{range .Items}}
<tr>
	<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
	<td>{{.State}}</td>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
	<td><a href='{{.Milestone.URL}}'>{{.Milestone.Number}}</a></td>
</tr>
{{end}}
</table>
`))

func main() {
	milestonesResult, err := github.ListMilestones()
	if err != nil {
		log.Fatal(err)
	}
	f, _ := os.Create("milestones.html")
	if err = milestoneList.Execute(f, milestonesResult); err != nil {
		log.Fatal(err)
	}

	usersResult, err := github.ListUsers()
	if err != nil {
		log.Fatal(err)
	}
	f, _ = os.Create("users.html")
	if err = userList.Execute(f, usersResult); err != nil {
		log.Fatal(err)
	}

	issuesResult, err := github.ListIssues()
	if err != nil {
		log.Fatal(err)
	}
	f, _ = os.Create("issues.html")
	if err = issuesList.Execute(f, issuesResult); err != nil {
		log.Fatal(err)
	}
}
