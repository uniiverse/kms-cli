package main

import (
  "fmt"
  "testing"
)

func TestGetKMSSession(t *testing.T) {
  got := GetKMSSession()

  fmt.Println(got)
}

func TestListAliases(t *testing.T) {
  session := GetKMSSession()

  got := ListAliases(session)

  fmt.Println(got)
}

func TestEncrypt(t *testing.T) {
  svc := GetKMSSession()

  //keyId := "1b4a9160-0e9a-4c4f-ae96-ff4f656ba8e2"
  keyId := "alias/web-staging"
  payload := []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
  app := "web"
  env := "staging"

  result := Encrypt(svc,keyId,app,env,payload)

  fmt.Println(result)
}

func TestDecrypt(t *testing.T) {

  svc := GetKMSSession()

  //keyId := "1b4a9160-0e9a-4c4f-ae96-ff4f656ba8e2"
  keyId := "alias/web-staging"
  payload := []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
  app := "web"
  env := "staging"

  encryptResult := Encrypt(svc,keyId,app,env,payload)

  decryptResult := Decrypt(svc, app, env, encryptResult.CiphertextBlob)

  fmt.Println(string(decryptResult.Plaintext))
}
