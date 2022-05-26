# Instabug 2022 Infrastructure Task

## Functionality
1. Dockerize a violin GOlang app 
    1. Multi-step build process to minimize the final image output (52.1MB)
    2. PORT environment variable to mimic all code features 
    3. Non root user for added security
2. Structure a CI/CD Jenkins pipeline
    1. Build an image with the commit number & latest tags
    2. Push mentioned image to docker hub
    4. Re-deploy the app to AKS to automatically update the deployment image to the latest
    3. Send email upon build failure as a reporting service
3. Create 3 kubernetes resources for deployment
    1. Deployment
    2. ClusterIP
    3. Ingress
4. App is deployed on an AKS cluster [20.121.167.232](20.121.167.232)

## How to run
- Docker file
    - Build
        ``` shell
            docker build -t aayman1/violin-app:imageTag .
        ```
    - Run with default port
        ``` shell
            docker run -p <PORT_NUMBER>:8080 aayman1/violin-app:imageTag
        ```
    - Run with a specified port 
        ```shell
            docker run -p <PORT_NUMBER>:<CONTAINER_PORT_NUMBER> --env PORT=CONTAINER_PORT_NUMBER aayman1/violin-app:imageTag
        ```
- k8s minikube
    1. Enable ingress on minikube 
    2. Run
        ``` shell
            kubectl apply -f k8s 
        ```
- AKS 
    1. Enable ingress controller on AKS
    2. Upload k8s folder 
    3. Run
        ``` shell
            kubectl apply -f k8s 
        ```
- Jenkins (Whole Pipeline)
    1. Create a pipeline in Jenkins
    2. Install plugins
        - [Kubernetes CLI](https://plugins.jenkins.io/kubernetes-cli/)
    3. Create a secret text credential that has token of the cluster with id **instabug-secret**
    5. Create a UserName password credential with id **DockerHub**
    4. Paste the Jenkins file in the pipeline project.

## Output
```
Started by user Ahmed Ayman
[Pipeline] Start of Pipeline
[Pipeline] node
Running on Jenkins in /var/jenkins_home/workspace/Instabug-Task
[Pipeline] {
[Pipeline] withCredentials
Masking supported pattern matches of $DOCKERHUB_CREDENTIALS or $DOCKERHUB_CREDENTIALS_USR or $DOCKERHUB_CREDENTIALS_PSW
[Pipeline] {
[Pipeline] sh
+ git rev-parse --short=10 HEAD
[Pipeline] withEnv
[Pipeline] {
[Pipeline] stage
[Pipeline] { (git checkout)
[Pipeline] git
The recommended git tool is: NONE
No credentials specified
 > git rev-parse --resolve-git-dir /var/jenkins_home/workspace/Instabug-Task/.git # timeout=10
Fetching changes from the remote Git repository
 > git config remote.origin.url https://github.com/aayymann/Instabug-Infrastructure-Task.git # timeout=10
Fetching upstream changes from https://github.com/aayymann/Instabug-Infrastructure-Task.git
 > git --version # timeout=10
 > git --version # 'git version 2.30.2'
 > git fetch --tags --force --progress -- https://github.com/aayymann/Instabug-Infrastructure-Task.git +refs/heads/*:refs/remotes/origin/* # timeout=10
 > git rev-parse refs/remotes/origin/master^{commit} # timeout=10
Checking out Revision 7a17b338e34979a80d26fe533bf70f2159fa1453 (refs/remotes/origin/master)
 > git config core.sparsecheckout # timeout=10
 > git checkout -f 7a17b338e34979a80d26fe533bf70f2159fa1453 # timeout=10
 > git branch -a -v --no-abbrev # timeout=10
 > git branch -D master # timeout=10
 > git checkout -b master 7a17b338e34979a80d26fe533bf70f2159fa1453 # timeout=10
Commit message: "Merge pull request #1 from aayymann/dev"
First time build. Skipping changelog.
[Pipeline] }
[Pipeline] // stage
[Pipeline] stage
[Pipeline] { (build image)
[Pipeline] sh
Warning: A secret was passed to "sh" using Groovy String interpolation, which is insecure.
		 Affected argument(s) used the following variable(s): [DOCKERHUB_CREDENTIALS_USR]
		 See https://jenkins.io/redirect/groovy-string-interpolation for details.
+ docker build -t ****/violin-app:34f21e4510 -t ****/violin-app:latest .
#2 [internal] load .dockerignore
#2 sha256:2872a516a3c6f0114897c3929b6d35ab4983f97e7551caeccf673c1cd8f38137
#2 transferring context: 32B
#2 transferring context: 34B done
#2 ...

#1 [internal] load build definition from Dockerfile
#1 sha256:1cd00cfbee1216985a0cf1509d5623934007db7b11e38db0ed7024de9da15d26
#1 transferring dockerfile: 32B done
#1 DONE 0.4s

#2 [internal] load .dockerignore
#2 sha256:2872a516a3c6f0114897c3929b6d35ab4983f97e7551caeccf673c1cd8f38137
#2 DONE 0.5s

#3 [internal] load metadata for docker.io/library/alpine:3.15.4
#3 sha256:06edcfaac690ebf32125388172af4c413e1aa88e77652c8af884b4ece9954b22
#3 DONE 0.0s

#4 [internal] load metadata for docker.io/library/golang:1.18-alpine3.15
#4 sha256:c36df16c50dbaa42b1337a580f490220e4230b77c6008333aad7fd9004ad3a1d
#4 DONE 0.0s

#5 [runner 1/8] FROM docker.io/library/alpine:3.15.4
#5 sha256:83101af27eafd4fdb4bd132610a44d51fdcbd295a6fb7049eb9abb63c7f96ca1
#5 DONE 0.0s

#13 [builder 1/4] FROM docker.io/library/golang:1.18-alpine3.15
#13 sha256:6d8e305ffe99871e47149edc0bed60ff7db3d415f36ac789f96750719df12a5c
#13 DONE 0.0s

#8 [internal] load build context
#8 sha256:06374b52dd82f537948ea6ec685789b7e81ce92cb9d2046382a4a6957889be8d
#8 transferring context: 12.17kB 0.0s done
#8 DONE 0.3s

#14 [builder 2/4] WORKDIR /violin-app
#14 sha256:d139558a66f6a228b539e7c9d3e8acc45023c2c22c02cee6ccebbe159a478b77
#14 CACHED

#6 [runner 2/8] RUN addgroup -S appgroup && adduser -S appuser -G appgroup
#6 sha256:34692d62d91f01e67b9a53122891c4bde96036cc3ea2912f0e422383e25b453d
#6 CACHED

#16 [builder 4/4] RUN go mod init     && go build -o violin-app
#16 sha256:8dbfc9188f9dc4d2680755b7c00a184b597ad6a8004aa6be93772a41a3906aca
#16 CACHED

#10 [runner 5/8] COPY ./css ./css
#10 sha256:9b586a44249dcbf350862b32701bec6b38bb067812cbfcd630e7d647a09a7f90
#10 CACHED

#9 [runner 4/8] COPY ./templates ./templates
#9 sha256:b6de6d37c56722b7780f0b4e36d6df24ffcaf4fc5bc23b7267607d81a2023ae4
#9 CACHED

#7 [runner 3/8] WORKDIR /violin-app/
#7 sha256:42daf7020605c00f038978511240fcfce2094071f337c2a6817716c509450ddb
#7 CACHED

#11 [runner 6/8] COPY ./img ./img
#11 sha256:5daa5a554319726a17e8be41b97b4743d53caa571dab92550377279ed322338b
#11 CACHED

#12 [runner 7/8] COPY ./mp3 ./mp3
#12 sha256:952a9c813ac0a08f38e5515d644639be240a3fbfd0de450c2b7512a471c6bd55
#12 CACHED

#15 [builder 3/4] COPY . .
#15 sha256:26423a474c9904b2ce8a1ff9c477888c01e933e90d26c3e7a12c99d3f37770f4
#15 CACHED

#17 [runner 8/8] COPY --from=builder /violin-app/violin-app ./
#17 sha256:957def31b302fd72640741b99558cc5ad8617daf24bc21f223de457ed11659ee
#17 CACHED

#18 exporting to image
#18 sha256:e8c613e07b0b7ff33893b694f7759a10d42e180f2b4dc349fb57dc6b71dcab00
#18 exporting layers done
#18 writing image sha256:9d46cc7e99011660c7af9b46ad1e3a5a7700e630f2cd998923a270fece8f6589 0.0s done
#18 naming to docker.io/****/violin-app:34f21e4510
#18 naming to docker.io/****/violin-app:34f21e4510 0.0s done
#18 naming to docker.io/****/violin-app:latest done
#18 DONE 0.3s
[Pipeline] }
[Pipeline] // stage
[Pipeline] stage
[Pipeline] { (dockerhub login)
[Pipeline] sh
Warning: A secret was passed to "sh" using Groovy String interpolation, which is insecure.
		 Affected argument(s) used the following variable(s): [DOCKERHUB_CREDENTIALS_USR, DOCKERHUB_CREDENTIALS_PSW]
		 See https://jenkins.io/redirect/groovy-string-interpolation for details.
+ echo ****
+ docker login -u **** --password-stdin
WARNING! Your password will be stored unencrypted in /root/.docker/config.json.
Configure a credential helper to remove this warning. See
https://docs.docker.com/engine/reference/commandline/login/#credentials-store

Login Succeeded
[Pipeline] }
[Pipeline] // stage
[Pipeline] stage
[Pipeline] { (push image)
[Pipeline] sh
Warning: A secret was passed to "sh" using Groovy String interpolation, which is insecure.
		 Affected argument(s) used the following variable(s): [DOCKERHUB_CREDENTIALS_USR]
		 See https://jenkins.io/redirect/groovy-string-interpolation for details.
+ docker push ****/violin-app:34f21e4510
The push refers to repository [docker.io/****/violin-app]
8349236976af: Preparing
cac08c745ab4: Preparing
fc0ccdacc285: Preparing
60a405261b8b: Preparing
7d2b71b17f4e: Preparing
bbc2d80f9e6b: Preparing
810a7079f773: Preparing
4fc242d58285: Preparing
bbc2d80f9e6b: Waiting
4fc242d58285: Waiting
810a7079f773: Waiting
fc0ccdacc285: Layer already exists
60a405261b8b: Layer already exists
7d2b71b17f4e: Layer already exists
cac08c745ab4: Layer already exists
bbc2d80f9e6b: Layer already exists
810a7079f773: Layer already exists
4fc242d58285: Layer already exists
8349236976af: Layer already exists
34f21e4510: digest: sha256:309720e9b51ca56f7fa4bd2462266ea52375a4cf3367bba89c70c074e8c42602 size: 1993
[Pipeline] sh
+ docker push ****/violin-app:latest
The push refers to repository [docker.io/****/violin-app]
8349236976af: Preparing
cac08c745ab4: Preparing
fc0ccdacc285: Preparing
60a405261b8b: Preparing
7d2b71b17f4e: Preparing
bbc2d80f9e6b: Preparing
810a7079f773: Preparing
4fc242d58285: Preparing
bbc2d80f9e6b: Waiting
810a7079f773: Waiting
4fc242d58285: Waiting
60a405261b8b: Layer already exists
7d2b71b17f4e: Layer already exists
8349236976af: Layer already exists
fc0ccdacc285: Layer already exists
cac08c745ab4: Layer already exists
bbc2d80f9e6b: Layer already exists
4fc242d58285: Layer already exists
810a7079f773: Layer already exists
latest: digest: sha256:309720e9b51ca56f7fa4bd2462266ea52375a4cf3367bba89c70c074e8c42602 size: 1993
[Pipeline] }
[Pipeline] // stage
[Pipeline] stage
[Pipeline] { (deploy AKS)
[Pipeline] withKubeConfig
[Pipeline] {
[Pipeline] sh
+ kubectl apply -f k8s
ingress.networking.k8s.io/violin-ingress unchanged
service/violin-cluster-ip unchanged
deployment.apps/violin-deployment configured
[Pipeline] sh
Warning: A secret was passed to "sh" using Groovy String interpolation, which is insecure.
		 Affected argument(s) used the following variable(s): [DOCKERHUB_CREDENTIALS_USR]
		 See https://jenkins.io/redirect/groovy-string-interpolation for details.
+ kubectl set image deployments/violin-deployment violin-container=****/violin-app:34f21e4510
deployment.apps/violin-deployment image updated
[Pipeline] }
[kubernetes-cli] kubectl configuration cleaned up
[Pipeline] // withKubeConfig
[Pipeline] }
[Pipeline] // stage
[Pipeline] stage
[Pipeline] { (Declarative: Post Actions)
[Pipeline] sh
+ docker logout
Removing login credentials for https://index.docker.io/v1/
[Pipeline] }
[Pipeline] // stage
[Pipeline] }
[Pipeline] // withEnv
[Pipeline] }
[Pipeline] // withCredentials
[Pipeline] }
[Pipeline] // node
[Pipeline] End of Pipeline
Finished: SUCCESS

```

