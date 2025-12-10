pipeline {
    agent { label 'dev' }

    stages {
        stage('Pull SCM') {
            steps {
                git branch: 'main', url: 'https://github.com/IDNDec/financial-record-go-mysql.git'
            }
        }
        
        stage('Build') {
            steps {
                sh'''
                cd app
                go mod tidy
                '''
            }
        }
        
        stage('Code Review') {
            steps {
                sh'''
                cd app
                sonar-scanner   -Dsonar.projectKey=app-amar   -Dsonar.sources=.   -Dsonar.host.url=http://172.23.10.17:9000   -Dsonar.token=sqp_cd2c2fda78200a0e895a965ca4cf1679babc922a
                '''
            }
        }
        
        stage('Deploy') {
            steps {
                sh'''
                docker compose up --build -d
                '''
            }
        }
        
        stage('Backup') {
            steps {
                 sh 'docker compose push' 
            }
        }
    }
}