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

  CheckAndAddKey(app, env)

  fmt.Println("Adding secret called:", name)

  //prompt for secret value
  secret := GetInput("Enter secret value: ")

  if(secret == "") {
    panic("No Secret Added!")
  }

  secretsExist, secretPath := CheckForSecretsFile(env, true)

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
  secretsExist, secretPath := CheckForSecretsFile(env, false)

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

  secretsExist, secretPath := CheckForSecretsFile(env, false)

  if(secretsExist) {
    encryptedSecrets := ReadFile(secretPath)
    result := Decrypt(GetKMSSession(), app, env, encryptedSecrets)
    fmt.Println(string(result))
  } else {
    fmt.Println("No secrets found for app / env")
    os.Exit(1)
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

func CheckAndAddKey(app, env string) {
  session := GetKMSSession()
  aliases := ListAliases(session)
  aliasExists := AliasExists(GetAliasName(app, env), aliases)

  if(!aliasExists) {
    createKey := BoolQuestion("Master Key doesn't exist, create?")

    if(createKey) {
      CreateKeyWithAlias(session, app, env)
      fmt.Println("Key Created")
    } else {
      fmt.Println("Could not save secret without key")
      os.Exit(1)
    }
  }
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
        fmt.Println("File written")
        return nil
      },
    },
    {
      Name: "remove",
      Aliases: []string{"r"},
      Usage: "Remove from encrypted file for Environment",
      ArgsUsage: "[name]",
      Action: func(c *cli.Context) error {
        fmt.Println("Remove secret")
        RemoveSecrets(c.Args().Get(0), appName, env)
        fmt.Println("Secret Removed")
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
