pipeline {
    agent any

    stages {
        stage('Clone project from GitHub') {
            steps {
                git branch: 'main', url: 'https://github.com/0xdreamerr/CT1520-task.git'
            }
        }

	stage('Build and run'){ 
    steps{ 
        script{ 
            sshPublisher(publishers: [ 
                sshPublisherDesc(configName: "gelasimov", 
                    transfers: [ 
                        sshTransfer( 
                            sourceFiles: "**", 
                            cleanRemote: true, 
                            excludes : '*.log, *.lock, .svn/, .git/', 
                            remoteDirectory: "/home/gelasimovyi/shortener",                                                  
                            execCommand: "docker-compose -f shortener/docker-compose.shortener.yml up -d --build && docker-compose -f shortener/docker-compose.nginx.yml up -d --build"
                        ), 
                    ], 
                    verbose: true
                ) 
            ]) 
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
