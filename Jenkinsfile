//# vi: ft=groovy
properties(
    [
        [
            $class: 'BuildDiscarderProperty',
            strategy: [$class: 'LogRotator', numToKeepStr: '10']
        ],
        pipelineTriggers([cron('H 13-21 * * 1-5')]),
    ]
)


stage('check') {
  node('docker') {
    docker.image('golang:latest').inside {
      checkout scm
      timeout(time: 240, unit: 'SECONDS') {
        sh "./check.sh"
      }
    }
  }
}

