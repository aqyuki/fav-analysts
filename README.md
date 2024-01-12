# README

**You should be updated to the README that suits the project**

## How to create new Go project with `gonew`?

If you do not have gonew installed, you can install it from the following command.

```bash
go install golang.org/x/tools/cmd/gonew@latest
```

You can create new project by `gonew` from the following command.
In this time, `<your_project_name> is used to package name when initialize Go.

```bash
gonew github.com/aqyuki/template-go <your-project-name> [output directory]
```

## Recommended configuration for GitHub

- Don't allow direct push for default branch

  If you want to activate it, You should create **Branch Protection Rule** on your repository.\
  Settings -> Branches -> Activate **Require a pull request before merging**
  and **Do not allow bypassing the above settings** from Branch Protection Rules

- Enable **Automatically delete head branches** option\
  GitHub deletes head branches when it was merged.

- Configure Pull Request
  > When merging pull requests, you can allow any combination of merge commits, squashing, or rebasing. At least one option must be enabled. If you have linear history requirement enabled on any protected branch, you must enable squashing or rebasing.

## Optional

- Enable Renovate or Dependabot
- Add a webhook to the repository
