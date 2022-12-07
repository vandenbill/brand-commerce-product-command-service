pipeline {
    agent any
    stages {
        stage ('get source code product command service') {
            steps {
                checkout([$class: 'GitSCM', branches: [[name: '*/main']], extensions: [], userRemoteConfigs: [[url: 'https://github.com/vandenbill/brand-commerce-product-command-service']]])
            }
        }
        stage ('build source code product command service') {
            steps {
                script {
                    sh 'docker build -t vandenbill/brand-commerce-product-command-service .'
                    sh 'rm -rf *'
                }
            }
        }
        stage ('get source code product query service') {
            steps {
                checkout([$class: 'GitSCM', branches: [[name: '*/main']], extensions: [], userRemoteConfigs: [[url: 'https://github.com/vandenbill/brand-commerce-product-query-service']]])
            }
        }
        stage ('build source code product query service') {
            steps {
                script {
                    sh 'docker build -t vandenbill/brand-commerce-product-command-service .'
                }
            }
        }
        stage ('push image to docker hub') {
            steps {
                script {
                    withCredentials([string(credentialsId: 'dockerhubpw', variable: 'dockerhubpw')]) {
                        sh 'docker login -u vandenbill -p ${dockerhubpw}'
                    }
                    sh 'docker image push vandenbill/brand-commerce-product-command-service'
                    sh 'docker image push vandenbill/brand-commerce-product-query-service'
                }
            }
        }
    }
}