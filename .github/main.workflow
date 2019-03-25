workflow "ci" {
  on = "push"
  resolves = "ci/*"
}

action "ci/static" {
  uses = "docker://ddollar/static"
  runs = "static"
}
