#KMS CLI

##About
KMS uses AWS Key Management Service (KMS) to encrypt JSON containing application secrets. 
IAM roles can be used to restrict access to Customer Master Keys (CMK)

IAM role is determined via:

- AWS cli tools (`~.aws/credentials`) (local dev machine)
- Introspected via the AWS EC2 instance
- Service Role (ECS)

##Installation
**GO**

```
go get github.com/uniiverse/kmscli
go install github.com/uniiverse/kmscli
```
**Docker**

Note: The following assumes alpine linux, running as root

```
Dockerfile
ADD https://github.com/uniiverse/kmscli/releases/download/v0.1/kmscli_linux_amd64.gz /tmp

RUN gzip -d /tmp/kmscli.gz && mv /tmp/kmscli /usr/bin && chmod +x /usr/bin/kmscli
```

Note: When running container locally (ie. docker compose), must mount in user's `.aws` folder read only (`ro`)

```
Volume Syntax
"~/.aws:/root/.aws:ro"
```

##Usage

```
kmscli -h for help
```

###Adding secrets

```
kmscli --app appname --env env add secretName
```

###Listing secrets

```
kmscli --app appname --env env list
```

###Removing secrets

```
kmscli --app appname --env env remove secretName
```

##Ruby Gem

TODO

##Node Module

TODO

##Building
To build for different environments, install GO with cross compilation

```
brew install go --with-cc-common # Installs go with cross compilation support
```

To build for linux x64

From the project directory

```
GOOS=linux GOARCH=amd64 go build -o kmscli
gzip kmscli
```

##Releasing
- Upload `kmscli.gz` created in the build step using GitHub's release UI

