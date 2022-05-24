## Tasks
- Dockerize the app and optimize the image using multi-stage
- Using Jenkins to build the puipline that build that image
- Report to mail if any failure happened in the pipeline
- + Kubernetes deployment to deploy the app (tested with minikube)

## Jenkins Pipeline
- git checkout to fetch the new commits on repository (all branches)
- Build the docker image ( 50MB :) )
- if the build sucessful then Push Image To Dockerhub, else report to mail

## Diagram
