pipeline {
    agent any

    stages {

        stage('Build') {
            steps {
                sh 'go build -o app'
            }
        }

        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }

        stage('Deploy') {
            steps {
                sh '''
                    cp app /opt/jenkins-bitbucket-demo/app
                    sudo systemctl restart jenkins-bitbucket-demo
                '''
            }
        }
    }
}