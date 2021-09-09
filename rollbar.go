package main

import (
  "errors"

  "github.com/rollbar/rollbar-go"
)

func main() {
  rollbar.SetToken("c47e5c491f394c149b1ff07464f43a0b")
  rollbar.SetEnvironment("production")                 // defaults to "development"
  rollbar.SetCodeVersion("v2")                         // optional Git hash/branch/tag (required for GitHub integration)
  rollbar.SetServerHost("web.1")                       // optional override; defaults to hostname
  rollbar.SetServerRoot("github.com/heroku/myproject") // path of project (required for GitHub integration and non-project stacktrace collapsing)

  if err := DoSomething(); err != nil {
    rollbar.Critical(err)
  }

  rollbar.Info("Message body goes here")

  rollbar.Wait()
}

func DoSomething() error {
  return errors.New("new error")
}
