# github_grade
The github grade api in golang. You can access the api by making a post request to https://github-grade.herokuapp.com/api
The body of the post request must contain a JSON object as follows:
```json
{
  "owner": "owner-of-the-repo",
  "name": "name-of-the-repo"
}
```
