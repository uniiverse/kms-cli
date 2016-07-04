package main

import (
  "os"
  "path/filepath"
  "io/ioutil"
)

//Check if file path exists
func Exists(path string) (bool, error) {
  _, err := os.Stat(path)
  if err == nil { return true, nil }
  if os.IsNotExist(err) { return false, nil }
  return true, err
}

//Returns true or false if secrets file exists
func CheckForSecretsFile(env string, mkdir bool) (bool, string) {

  dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
  if err != nil {
    panic(err)
  }
  dir += "/secrets"

  if(mkdir) {
    os.MkdirAll(dir, os.ModePerm) //Ensure dir exists
  }

  dir += "/" + env

  secretsExist, err := Exists(dir)
  return secretsExist, dir
}

//TODO: Filemode
func WriteFile(path string, contents []byte) {

  err := ioutil.WriteFile(path, contents, 0644)

  if(err != nil) {
    panic(err)
  }
}

func ReadFile(path string) []byte {
  result, err := ioutil.ReadFile(path)
  if(err != nil) {
    panic(err)
  }
  return result
}
