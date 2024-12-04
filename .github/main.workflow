workflow "ci" {
  on = "push"
  resolves = [ "ci/lint", "ci/test" ]
}

action "ci/lint" {
  uses = "docker://ddollar/static"
  runs = "lint"
  secrets = [ "GITHUB_TOKEN" ]
}

action "ci/test" {
  uses = "docker://ddollar/static"
  runs = "test"
}
