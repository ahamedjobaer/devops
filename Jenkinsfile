pipeline {
  agent any
  tools {
    go 'go-1.19'
  }
  environment {
    GO119MODULE = 'on'
  }
  stages {
    stage('Unit test') {
      steps {
        sh 'go test -v ./...'
      }
    }
  }
}
