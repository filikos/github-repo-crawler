# Github-repo-crawler

This REST-full API connects to the Github API and provides functionality for fetching data from public repositorys.

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
[
    "project1",
    "project3",
    "project3"
]
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
