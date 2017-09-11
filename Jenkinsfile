//# vi: ft=groovy
properties(
    [
        [
            $class: 'builddiscarderproperty',
            strategy: [$class: 'logrotator', numtokeepstr: '10']
        ],
        pipelinetriggers([cron('H 13-21 * * 1-5')]),
    ]
)

stage('check') {
  node('docker') {
    docker.image('golang:latest').inside {
      timeout(time: 240, unit: 'SECONDS') {
        sh "./check.sh"
      }
    }
  }
}
