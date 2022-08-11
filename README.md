# semantic-githooks

These [githooks](https://git-scm.com/docs/githooks) ensure your branch & commit match the [Conventional Commits spec](https://www.conventionalcommits.org/).

The typical use case is to use this in combination with a tool like [semantic-release](https://github.com/semantic-release/semantic-release) to automate releases.

## Validation

Examples for valid commits message:
- `fix: Correct typo`
- `feat: Add support for Node 12`
- `refactor!: Drop support for Node 6`
- `feat(ui): Add Button component`

Note that since PR titles only have a single line, you have to use the `!` syntax for breaking changes.

See [Conventional Commits](https://www.conventionalcommits.org/) for more examples.

## Installation

We use `pre-commit` to enforce branch naming convention and `commit-msg` for the commit message. To enable them, please run the following commands on the root project directory:

```
curl --fail -o .git/hooks/commit-msg https://raw.githubusercontent.com/nonkronk/semantic-githooks/main/commit-msg \
  && chmod 500 .git/hooks/commit-msg

curl --fail -o .git/hooks/pre-commit https://raw.githubusercontent.com/nonkronk/semantic-githooks/main/pre-commit \
  && chmod 500 .git/hooks/pre-commit
```

Verify your installation by committing with random word. If it’s failed to do commit, then the installation is success.
