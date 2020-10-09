# Github-repo-crawler

This REST-full API connects to the Github API and provides functionality for fetching data from public repositorys. Caching is implemented for quicker response times and also for relieving the Github API. PostgreSQL is used for persistant data storage.

# Endpoints

<details><summary>GET /repositories/{username}</summary>
<p>

#### Description:

Returns list of public repositories from user.

#### Parameters:

Content-Type: **application/json**<br/>
path: {username} \*required

##### Example Response Body:

```json
["project1", "project3", "project3"]
```

##### Responses:

200 OK<br/>
400 Bad Request<br/>
404 Not Found

</p>
</details>

<details><summary>GET /repositories/{username}/commits/{reponame}</summary>
<p>

#### Description:

Returns list of up to 20 commits related to the username and name of the repository. Results can be filtered by commit message when passing a search keyword as query parameter.

#### Parameters:

Content-Type: **application/json** <br/>
path: {username} \*required <br/>
path: {reponame} \*required <br/>
query: {search} optional

##### Example Response Body:

```json
[
  {
    "author": {
      "date": "2019-10-08T01:27:22Z",
      "email": "test@gmail.com",
      "name": "test test"
    },
    "comment_count": 0,
    "committer": {
      "date": "2019-10-08T01:27:22Z",
      "email": "test@github.com",
      "name": "GitHub"
    },
    "message": "Merge pull request #110 from test/patch-1\n\nHandle anchors in multiline mode",
    "tree": {
      "sha": "3723ec4f47f5f4fccdd9e53dcdd8b0739f1231f4",
      "url": "https://api.github.com/repos/test/test/git/trees/3723ec4f47f5f4fccdd9e53dcdd8b0739f1231f4"
    },
    "url": "https://api.github.com/repos/test/test/git/commits/41d6eabad7b055a83923150efd5518813831c9a5",
    "verification": {
      "payload": "tree 3723ec4f47f5f4fccdd9e53dcdd8b0739f1231f4\nparent 78bb627792fc8a5253baa9cd9d8160533b16fd85\nparent eab427817c819676cedf2d8998f571a10a8a703e\nauthor Brian Gesiak <test@gmail.com> 1570498042 -0400\ncommitter GitHub <noreply@github.com> 1570498042 -0400\n\nMerge pull request #110 from test/patch-1\n\nHandle anchors in multiline mode",
      "reason": "valid",
      "signature": "-----BEGIN PGP SIGNATURE-----\n\nwsBcBAABCAAQBQJdm+X6CRBK7hj4Ov3rIwAAdHIIAEhvV8HKAECVK+rMApAePuzi\n7HWbOf1EVZ+Tu1jKVI9klEQB5yJBeRng7RhORKM820MUqDkRsnohSjTBVZO/Qk0+\nWGlICe5qEoUVg4DkRX+Gr76qvtE1qkaOD1nE7N6yGnRVcJuilb1cLKMD9p2zoE1N\nWjsngJy2S3HiNwkhtEn/qKtuFFcDYymrlj2aOC3lLLbzUPRmgK1NocrciYu698va\n28Wf5AoYI6Sv7I/ep8SBFrOySiSBTqVyHE4rnVRElTI36MTbSptMAsKAo3CyyfUX\nUWqneG0Vz599zpyjSZOp/znMJE2Nfhtyto0bnXWWBazhWqAaCAdnrg0Ul1K4X80=\n=1Xbc\n-----END PGP SIGNATURE-----\n",
      "verified": true
    }
  }
]
```

##### Responses:

200 OK<br/>
400 Bad Request<br/>
404 Not Found

</p>
</details>

<details><summary>GET /recentrepositories</summary>
<p>

#### Description:

Returns list of the 20 recently requested repositories.

#### Parameters:

Content-Type: **application/json**<br/>

##### Example Response Body:

```json
[
  {
    "id": "172581071",
    "username": "username",
    "name": "repositoryName"
  }
]
```

##### Responses:

200 OK

</p>
</details>

### Requirements

- Go v1.14.4
- Docker v19.03.12
- Docker-Compose v1.26.2
- PostgreQSL Docker-Image

---

### Run the github-repo-crawler

To use persistant data store the PostgreSQL files location is needed. The path is working on MacOS & Linux, for Windows systems you may need to change the Postgres volume path within the docker-compose.yml.

```
mdkir $HOME/docker/volumes/postgres
```

Copy PostgreSQL files into the directory above or create a new database. Use the migration.sql script to create the nessessary tables.

Starting the API & PostgreSQL in Docker containers:

```
docker-compose up
```

Stop both containers:

```
docker-compose down
```

---

## Development Environment:

For development purposes it is recommended to start the PostgreSQL Database manually using docker.

```
cd $GOPATH/src/github-repo-crawler

docker run --name dev --env-file ./config/dbConfig.env -p 5432:5432 -v $HOME/docker/volumes/postgres:/var/lib/postgresql/data postgres
```

### Access Database within Docker container:

1. Open new bash window: `docker container ls`
2. `docker exec -it *POSTGRES_CONTAINER_ID* psql -U postgres -W postgres`
3. Run SQL Querys: Example `SELECT * FROM Repositories;`

The command `go run main.go` will start the Rest-API listen to http://localhost:8080. The API will try to connect with the, make sure the Database is already running.

`go run main.go -h`

```
NAME:
   Github-Repo-Crawler - A new cli application

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   v1.0.0

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --port value        Port the Rest-API will listen on. (default: 8080)
   --configPath value  Path to *.env postgres config file. (default: ./config/dbConfig.env)
   --help, -h          show help (default: false)
   --version, -v       print the version (default: false)
```
