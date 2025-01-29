pipeline {
    agent any

    stages {
        stage('Clone project from GitHub') {
            steps {
                git 'https://github.com/0xdreamerr/CT1520-task.git'
            }
        }

        stage('Build And Run Shortener Container') {
            steps {
                script {
                    sh "docker-compose -f docker-compose.shortener.yml up -d --build"
                }
            }
        }

        stage('Build And Run Nginx Container') {
            steps {
                script {
                    sh "docker-compose -f docker-compose.nginx.yml up -d --build"
                }
            }
        }

        
    }
    
    post {
        always {
            echo 'Pipeline finished!'
        }
    }
}
