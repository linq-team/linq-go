---
ticket: linqg-301
concepts: []
new_concepts: []
paths:
  - Makefile
  - .github/workflows/ci.yml
  - plans/linqg-301.md
contracts_changed:
  - "CI: new aggregate check `ci` (needs build, lint, test); new `test` job"
  - "Root Makefile targets: lint, format, test"
acceptance:
  - "`make lint`, `make format` and `make test` run from the repo root and exit 0 on a clean tree (they call scripts/lint, scripts/format, scripts/test)."
  - "ci.yml's lint job runs `make lint` and checks `make format` leaves no diff; a test job runs `make test`."
  - "ci.yml has a job named `ci` with `needs: [build, lint, test]` and `if: always()`, which exits non-zero when any needed job's result is failure or cancelled."
  - "lint and test run on pull requests to main from the same repo (not only forks), so the `ci` check on a PR reflects real results."
tests: []
decisions:
  - q: "lint (and the new test job) only run on push or fork PRs today (Stainless default), so on a same-repo PR they are skipped and `ci` would pass vacuously. How should they run on PRs?"
    options:
      - "Run lint/test on every pull_request too (drop the fork-only condition). Same-repo branches get a push run and a PR run."
      - "Keep Stainless conditions; `ci` relies on the push-triggered run attached to the head commit."
    recommend: "Run on every pull_request"
    why: "The required `ci` check must come from the PR event to be meaningful; a skipped needs chain reports success. The duplicate push run costs about a minute of CI."
epic:
---

Stainless generated `ci.yml` and `scripts/*`; keep the scripts as the single
source of truth and add a thin root `Makefile` (`lint`, `format`, `test`
targets calling `./scripts/<name>`). The 81 mock-server tests are skipped by a
hard-coded `t.Skip("Mock server tests are disabled")` (and would fail via
`testutil.CheckTestServer` without `SKIP_MOCK_TESTS=true` if unskipped), so
`make test` covers the unit tests only.

In `ci.yml`: the lint job calls `make lint`, then `make format` and
`git diff --exit-code`; add a `test` job calling `make test`; add `ci` with
`needs: [build, lint, test]`, `if: always()`, failing when
`contains(needs.*.result, 'failure') || contains(needs.*.result, 'cancelled')`.
Skipped needs (e.g. build on a same-repo PR) do not fail `ci`.
