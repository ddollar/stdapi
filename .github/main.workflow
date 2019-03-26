workflow "ci" {
  on = "push"
  resolves = [ "ci/static" ]
}

action "ci/static" {
  uses = "docker://ddollar/static"
  runs = "static"
}

action "ci/test" {
  uses = "docker://ddollar/static"
  runs = [ "go", "test", "-v", "./..." ]
}
