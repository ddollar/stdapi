workflow "ci" {
  on = "push"
  resolves = [ "ci/static", "ci/test" ]
}

action "ci/static" {
  uses = "docker://ddollar/static"
  runs = "static"
  secrets = [ "GITHUB_TOKEN" ]
}

action "ci/test" {
  uses = "docker://ddollar/static"
  runs = "test"
}
