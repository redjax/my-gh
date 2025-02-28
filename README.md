# My Github

A Go CLI app for interacting with Github's REST API.

## Usage

- Run the app once (`my-gh get stars`) to create a `config.json` file
  - Edit this file, adding your Github API token
  - You can also run this app without a `config.json` by:
    - Setting an environment variable called `MYGH_GH_TOKEN`
      - (Windows) `$env:MYGH_GH_TOKEN = "<your Github API token>"`
      - (Linux) `export MYGH_GH_TOKEN="<your Github API token>"
    - Passing CLI args
      - `my-gh get stars --gh-token <your Github API token>`
- By default the app will output starred repositories to `./starred_repositories.json`
  - This is configurable by editing `config.json`, or setting the environment variable `MYGH_OUTPUT_FILE="/path/to/filename.json`
