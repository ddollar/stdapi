workflow "ci/static" {
  on = "push"
  resolves = "static"
}

action "static" {
  uses = "docker://ddollar/static"
}
