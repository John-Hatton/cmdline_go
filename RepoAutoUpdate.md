# HOWTO : Github Actions


---


## Version Number Repo Auto-Update

***See Code below:


You can grant the token the required permissions by following these steps:

    Go to the repository settings page on Github.
    Click on "Secrets" in the left-hand sidebar.
    Click on "New repository secret".
    Enter a name for the secret (e.g. "GITHUB_TOKEN").
    Generate a new Github Token with the "repo" scope.
    Copy the token value and paste it into the "Value" field in the "New repository secret" form.
    Save the secret.


--- 

(NOTE Here Tokens are referred to as Tokens (classic))

To generate a new Github token with the "repo" scope, follow these steps:

    Go to your Github account settings
    Click on "Developer settings"
    Click on "Personal access tokens"
    Click on "Generate new token"
    Give the token a name and select the "repo" scope
    Click "Generate token"
    Copy the token




---


name: Create Release
on:
  push:
    branches:
      - main
jobs:
  release:
    name: Create Release
    runs-on: ubuntu-latest
    steps:
      - name: Checkout Code
        uses: actions/checkout@v2
      - name: Setup Go
        uses: actions/setup-go@v2
        with:
          go-version: 1.16
      - name: Get Version
        id: get_version
        run: echo "::set-output name=version::$(grep -Eo '[0-9]+\.[0-9]+\.[0-9]+' cmdline_go.go)"
      - name: Check Release
        id: check_release
        uses: actions/github-script@v4
        with:
          script: |
            const { data: releases } = await github.repos.listReleases({
              owner: context.repo.owner,
              repo: context.repo.repo,
            });
            core.setOutput('release_exists', releases.some(release => release.tag_name === `${{ steps.get_version.outputs.version }}`));
          token: ${{ secrets.MYTOKEN }}
      - name: Create Release
        id: create_release
        if: steps.check_release.outputs.release_exists != 'true'
        uses: actions/create-release@v1
        env:
          GITHUB_TOKEN: ${{ secrets.MYTOKEN }}
        with:
          tag_name: ${{ steps.get_version.outputs.version }}
          release_name: Release ${{ steps.get_version.outputs.version }}
          draft: false
          prerelease: false
          
          
          
          
--- 

Be Sure to change cmdline_go.go to <filename>.go and the same with the new token 
