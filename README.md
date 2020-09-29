# Github-repo-crawler

This REST-full API connects to the Github API and provides functionality for fetching data from public repositorys.

# Endpoints

<details><summary>GET /repositories/{username}</summary>
<p>

#### Description:

Returns list of public repositories from user.

#### Parameters:

Content-Type: **application/json**
path: {username} \*required

##### Example Response Body:

```json

```

##### Responses:

200 OK
400 Bad Request
404 Not Found

</p>
</details>

<details><summary>GET /repositories/{username}/commits/{reponame}</summary>
<p>

#### Description:

Returns list of up to 20 commits related to the username and name of the repository.

#### Parameters:

Content-Type: **application/json**
path: {username} \*required
path: {reponame} \*required

##### Example Response Body:

```json

```

##### Responses:

200 OK
400 Bad Request
404 Not Found

</p>
</details>

<details><summary>GET /repositories/{username}/commits/{reponame}/search/{query} </summary>
<p>

#### Description:

Returns list of commits related to the username, name of the repository and the search query. Note: This is just a simple text search.

#### Parameters:

Content-Type: **application/json**
path: {username} \*required
path: {reponame} \*required
path: {query} \*required

##### Example Response Body:

```json

```

##### Responses:

200 OK
400 Bad Request
404 Not Found

</p>
</details>