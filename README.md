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
```