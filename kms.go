package main

import (
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/kms"
)

func GetKMSSession() *kms.KMS {
  config := &aws.Config{Region: aws.String("us-east-1")}
  svc := kms.New(session.New(), config)

  return svc
}

func ListAliases(svc *kms.KMS) ([]*kms.AliasListEntry) {
  resp, err := svc.ListAliases(nil)
  if(err != nil) {
    panic(err)

    /*
    * TODO: Handle errors more gracefully
     if awsErr, ok := err.(awserr.Error); ok {
      fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
      if reqErr, ok := err.(awserr.RequestFailure); ok {
        fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
      }
    } else {
      fmt.Println(err.Error())
    }*/
  }

  return resp.Aliases
}

func GetAliasName(app, env string) string {
 return "alias/" + app + "-" + env
}

func Encrypt(svc *kms.KMS, app, env string, payload []byte) []byte { //*kms.EncryptOutput

  params := &kms.EncryptInput{
    KeyId:     aws.String(GetAliasName(app, env)), // Required
    Plaintext: payload,           // Required
    EncryptionContext: map[string]*string{
      "App": aws.String(app), // Required
      "Env": aws.String(env), // Required
    },
    //GrantTokens: []*string{
    //  aws.String("GrantTokenType"),
    //},
  }

  resp, err := svc.Encrypt(params)

  if err != nil {
    panic(err)
  }
  return resp.CiphertextBlob
}

func Decrypt(svc *kms.KMS, app, env string, payload []byte) []byte { //*kms.DecryptOutput
  params := &kms.DecryptInput{
    CiphertextBlob: payload,
    EncryptionContext: map[string]*string{
        "App": aws.String(app), // Required
        "Env": aws.String(env),
    },
  }

  resp, err := svc.Decrypt(params)

  if(err != nil) {
    panic(err)
  }

  return resp.Plaintext
}
