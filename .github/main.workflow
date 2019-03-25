workflow "ci/static" {
  on = "push"
  resolves = "static"
}

action "static" {
  uses = "docker://ddollar/static"
  runs = [ "sh", "-c", "ls -la $GITHUB_WORKSPACE" ]
}
