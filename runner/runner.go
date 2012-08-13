package main

import (
  "fmt"
  "runlib/sub32"
)

func sptr(s string) *string {
  return &s
}

type foo1 int

func (s *foo1) vava(f string) error {
  fmt.Printf(f)
  return nil
}

func main() {

  sub := sub32.SubprocessCreate()
  cmd := "C:\\WINDOWS\\System32\\cmd.exe"
  sub.ApplicationName = sptr(cmd)
  sub.CommandLine = sptr(cmd + " /c echo test")
  sub.Username = sptr("test")
  sub.Password = sptr("test321")
  sub.StdOut = &sub32.SubprocessOutputRedirect{}
  sub.StdOut.ToMemory = true
  
  sig, err := sub.Start()

  
  
  // env := sub32.GetEnvMap()
  // env["ZZZTEST"] = "VAVA"

  // r, e := sub32.CreateProcessWithLogonW("test", nil, "test321", 0, &cmd, &cmd, 0, sub32.EnvironmentMap(env), nil, nil)
  // r, e := sub32.CreateProcessW(&cmd, &cmd, sub32.EnvironmentMap(env), nil, nil)

  fmt.Printf("%s %s\n", sig, err)

  fmt.Printf("%s\n", <-sig)
}