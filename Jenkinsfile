pipeline {
  agent any
  stages {
    stage('git checkout') {
      steps {
        git(url: 'https://github.com/s403o/GoViolin.git', branch: 'master', changelog: true, poll: true, credentialsId: 'myGit')
      }
    }

    stage('Build Docker Image') {
      steps {
        withCredentials(bindings: [[$class: 'UsernamePasswordMultiBinding', credentialsId: 'myDocker', usernameVariable: 'DOCKER_USERNAME', passwordVariable: 'DOCKER_PASSWORD']]) {
          sh '''
		          docker build -t s403o/goapp .
	        '''
        }

      }
    }

    stage('Push Image To Dockerhub') {
      steps {
        withCredentials(bindings: [[$class: 'UsernamePasswordMultiBinding', credentialsId: 'myDocker', usernameVariable: 'DOCKER_USERNAME', passwordVariable: 'DOCKER_PASSWORD']]) {
          sh '''
              docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD
              docker push s403o/goapp
	        '''
        }

      }
    }

  }
}