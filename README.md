# lotto24-api-test

API tests for Wikimedia public API

## Requirements

### For local run required:

* Golang 1.21+ [installed with proper GOHOME environment variable setted](https://go.dev/doc/install)

### for run in container:   

* Docker

## How to run

Open folder with project and run one of the following

### Local run

```
go test
```

### Build docker image and run in container:

```
docker build -t lotto24 .
docker run -it lotto24
```

### Run in docker container without building image:

```
docker run -it -v  .:/app -w /app golang:1.21-alpine go test
```

### See runs in CI

Just open github-actions page and see [all CI runs](https://github.com/temagi/lotto24-api-test/actions)


## Idea and implementation

The task doesn't specify whether we should search for an exact textual match or two separate words. Therefore, I also added a test for searching by exact match.

The second part and some thoughts about tests [are here](TESTCASES.md)

#### About Authorization Header

According to [Create Page documentation](https://api.wikimedia.org/wiki/Core_REST_API/Reference/Pages/Create_page#Headers), we just need to provide a [Bearer token](https://api.wikimedia.org/wiki/Authentication#Personal_API_tokens).

For example, if I need to add any test required authorization in my solution, I would like to make following steps:
1. [Create Account](https://api.wikimedia.org/wiki/Special:CreateAccount) for wikimedia
2. [Create a personal API token](https://api.wikimedia.org/wiki/Getting_started_with_Wikimedia_APIs#2._Create_a_personal_API_token)
3. Put token into secrets into project (using .env and .env.local for example)
4. Set token to auth header for required requests

#### Assume we just created that public Wikipedia API. Describe which other aspects you would want to test and how

So, as I already mention in testcase section, the approach to testing depends on the many different factors and we need to answer some questions before choose one.

Lets assume we have a kind of a averaged structure, access to all the sources, pipelines and so on. The most obvious way is to start from bottom to top with a tipical testing triangle. We could write many low-level unit tests with a high coverage, then create some mid-level tests (component test in isolated environment for example), where we could test all endpoints parameters and results.
On a high-level we could create some tests that could be closer to real user behavior instead of enpoints testing. We could think on this level from a users side and their interaction, for a simple example:
1. User create a page
2. Another user search the page
3. Another user edit page
4. Request page history contains proper information about the page

From a non-functional side we could also think about:
* Performance testing. Assess the API's performance under various load conditions to ensure it can handle expected levels of traffic efficiently and measure response times for different API endpoints and operations to identify potential bottlenecks and optimize performance.

* Security testing. Conduct security testing to identify and address potential vulnerabilities, such as injection attacks, authorization bypass, and data exposure risks, verify that sensitive data is properly protected and access controls are enforced.
