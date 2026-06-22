# Claude

⛔ **Never git push, open a PR, or cut a release without explicit confirmation from the user.** Local commits are fine.

Cross-repo context, conventions, history, and backlog live in
**github.com/universal-field-robots/claude-context**. Ensure it's cloned at
`~/claude-context` (clone it if missing — our devcontainer homes are volumed, so
it's a one-time clone), `git pull` to refresh, then read its `CLAUDE.md`.

Keep entries brief. This is a fork RTSPtoWeb depends on — changes here need
RTSPtoWeb's `go.mod` re-pinned.
