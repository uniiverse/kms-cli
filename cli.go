package main

import (
  "os"
  "github.com/urfave/cli"
  "fmt"
  "encoding/json"
)

func CheckApp(app string) {
  if app == "" {
    //TODO: Try to get app name from containing folder
    panic("No App Provided")
  }
}

func CheckEnv(env string) {
  if env == "" {
    panic("No Env Provided")
  }
}

func CheckName(name string) {
  if(name == "") {
    panic("No name provided for secret!")
  }
}

func addSecret(app, env, name string) {
  CheckApp(app)
  CheckEnv(env)
  CheckName(name)

  fmt.Println("Adding secret called:", name)

  //prompt for secret value
  secret := GetInput("Enter secret value: ")

  if(secret == "") {
    panic("No Secret Added!")
  }

  secretsExist, secretPath := CheckForSecretsFile(env)

  if(!secretsExist) {
    //Create
    secrets := map[string]string{}
    secrets[name] = secret
    ParseEncryptWrite(secrets, app, env, secretPath)
  } else {
    secrets := ReadDecryptParse(secretPath, app, env)
    secrets[name] = secret
    ParseEncryptWrite(secrets, app, env, secretPath)
  }
}

func RemoveSecrets(name, app, env string) {
  CheckName(name)
  CheckEnv(env)
  CheckApp(app)
  fmt.Println("env", env)
  //Check for secrets file 
  secretsExist, secretPath := CheckForSecretsFile(env)

  if(secretsExist) {
    //Decrypt and parse secrets
    secrets := ReadDecryptParse(secretPath, app, env)
    delete(secrets, name)
    ParseEncryptWrite(secrets, app, env, secretPath)
  } else {
    fmt.Println("No secrets file for env")
  }
}

func ListSecrets(app, env string) {
  CheckEnv(env)
  CheckApp(app)

  secretsExist, secretPath := CheckForSecretsFile(env)

  if(secretsExist) {
    encryptedSecrets := ReadFile(secretPath)
    result := Decrypt(GetKMSSession(), app, env, encryptedSecrets)
    fmt.Println(string(result))
  }
}

func ReadDecryptParse(path, app, env string) map[string]interface{} {
  session := GetKMSSession()
  encryptedSecrets := ReadFile(path)
  decryptedSecrets := Decrypt(session, app, env, encryptedSecrets)
  secrets := UnmarshalSecrets(decryptedSecrets)
  return secrets
}

func ParseEncryptWrite(input interface{}, app, env, path string) {
  session := GetKMSSession()
  newJson := MarshalSecrets(input)
  encryptedPayload := Encrypt(session, app, env, newJson)
  WriteFile(path,encryptedPayload)

}

func UnmarshalSecrets(input []byte) map[string]interface{} {
  var dat map[string]interface{}

  err := json.Unmarshal(input, &dat)

  if(err != nil) {
    panic(err)
  }
  return dat
}

func MarshalSecrets(input interface{}) []byte {
  data, err := json.Marshal(input)

  if(err != nil) {
    panic(err)
  }
  return data
}

func main() {
  var env string
  var appName string

  app := cli.NewApp()

  app.Name = "AWS KMS Secrets Wrapper"
  app.Usage = "Add or remove application secrets"

  app.Authors = []cli.Author{
    cli.Author{
      Name: "kyle.white",
      Email: "kyle.white@universe.com",
    },
  }
  app.Flags = []cli.Flag{
    cli.StringFlag{
      Name: "env",
      Usage: "The application environment to target",
      Destination: &env,
    },
    cli.StringFlag{
      Name: "app",
      Usage: "The Application to target",
      Destination: &appName,
    },
  }

  app.Commands = []cli.Command{
    {
      Name: "add",
      Aliases: []string{"a"},
      Usage: "Add to encrypted file for environment",
      ArgsUsage: "[name]",
      Action: func(c *cli.Context) error {
        fmt.Println("Add to encrypted file")
        fmt.Println("Env", env)
        addSecret(appName, env, c.Args().Get(0))
        return nil
      },
    },
    {
      Name: "remove",
      Aliases: []string{"r"},
      Usage: "Remove from encrypted file for Environment",
      ArgsUsage: "[name]",
      Action: func(c *cli.Context) error {
        fmt.Println("Remove from encrypted file")
        RemoveSecrets(c.Args().Get(0), appName, env)
        return nil
      },
    },
    {
      Name: "list",
      Aliases: []string{"l"},
      Usage: "List secrets for a given environment",
      Action: func(c *cli.Context) error {
        ListSecrets(appName, env)
        return nil
      },
    },
  }

  app.Run(os.Args)
}
