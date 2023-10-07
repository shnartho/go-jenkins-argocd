pipeline {
    agent any
    environment {
        DOCKERHUB_CREDENTIALS = credentials('dockerhub')
        GIT_REPO_URL = 'https://github.com/shnartho/go-jenkins-argocd-deployment.git'
        GIT_REPO_CREDENTIALS = credentials('git_credentials') 
    }

    stages {
        stage('Build image') {
            steps {
                sh 'docker build -t shnartho/go-jenkins-argocd:$BUILD_NUMBER .'
            }
        }

        stage('Login to dockerhub') {
            steps {
                sh 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'
            }
        }

        stage('Push image') {
            steps {
                sh 'docker push shnartho/go-jenkins-argocd:$BUILD_NUMBER'
            }
        }

        stage('Update Deployment File') {
            steps {
                //sh 'rm -rf repository'
                withCredentials([usernamePassword(credentialsId: 'git_credentials',             usernameVariable: 'GIT_USERNAME', passwordVariable: 'GIT_PASSWORD')]) {
                    sh "git clone $GIT_REPO_URL repository"
                }
                dir('repository') {
                    // Update your deployment file with the new image tag
                    sh "sed -i 's|image: shnartho/go-jenkins-argocd:[0-9]*|image: shnartho/go-jenkins-argocd:$BUILD_NUMBER|' deployment.yaml"
                    
                    // Commit and push the changes
                    sh 'git config user.email "shnartho@gmail.com"'
                    sh 'git config user.name "shnartho"'
                    sh 'git add deployment.yaml'
                    sh 'git commit -m "Update deployment.yaml"'
                    //sh 'git push --set-upstream origin main'
                    sh 'https://${GIT_USERNAME}:${GIT_PASSWORD}@github.com/shnartho/go-jenkins-argocd-deployment.git HEAD:main'
                }
            }
        }
    }
    post {
        always {
            sh 'docker logout'
        }
    }
}