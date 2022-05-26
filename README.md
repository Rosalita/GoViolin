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
![diagram](https://user-images.githubusercontent.com/38042656/170201677-86103ceb-8a5e-4e14-82e6-2efc12f44bcd.jpeg)

## Mail Report

![report](https://user-images.githubusercontent.com/38042656/170383078-48ae9ea2-2ad4-4a5a-afc3-94e134e64218.PNG)
