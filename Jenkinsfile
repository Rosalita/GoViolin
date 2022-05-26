pipeline {
    agent any
    environment{
        DOCKERHUB_CREDENTIALS = credentials("DockerHub")
        GIT_COMMIT_NUMBER = sh(returnStdout: true, script: "git rev-parse --short=10 HEAD").trim()
    }
    stages {
        stage ("git checkout"){
            steps {
                git 'https://github.com/aayymann/Instabug-Infrastructure-Task.git'
            }
        }
        stage("build image"){
            steps {
                sh "docker build -t aayman1/violin-app:${env.GIT_COMMIT_NUMBER} -t aayman1/violin-app:latest ."
            }
        }
        stage("dockerhub login"){
            steps{
                sh " echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin"
            }
        }
        stage("push image"){
            steps{
                sh "docker push aayman1/violin-app:${env.GIT_COMMIT_NUMBER}"
                sh "docker push aayman1/violin-app:latest"
            }
        }
        stage("deploy AKS"){
            steps{
                withKubeConfig([credentialsId: 'instabug-secret', serverUrl: 'https://instabug-task-dns-b7529c4a.hcp.eastus.azmk8s.io:443']) {
                    sh "kubectl apply -f k8s"
                    sh "kubectl set image deployments/violin-deployment violin-container=aayman1/violin-app:${env.GIT_COMMIT_NUMBER}"
                }
            }
        }
    }
    post {
        always{
            sh "docker logout"
        }
		failure{
            mail to : "temp.on.holdd@gmail.com",
            subject : "Build ${env.BUILD_ID} Failed ",
            body: "Build Failed"
        }
	}
}