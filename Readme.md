#KMS Wrapper

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

Note: When running container locally, must mount in user's `.aws` folder read only (`ro`)

```
Volume Syntax
"~/.aws:/root/.aws:ro"
```

##Usage

```
kms -h for help
```

###Adding secrets

```
kmscli --app appname --env env add secretName
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

