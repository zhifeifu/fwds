#!/bin/bash

SERVER="$1"
COMMAND="$2"
BASE_DIR=$PWD
INTERVAL=2

# 命令行参数，需要手动指定
ARGS=""

function start() {
  if [ "$(pgrep "$SERVER" -u $UID)" != "" ]; then
    echo "$SERVER already running"
    exit 1
  fi

  nohup "$BASE_DIR"/"$SERVER" "$ARGS" server >>"$SERVER"".nohup" 2>&1 &

  echo "sleeping..." && sleep $INTERVAL

  # check status
  if [ "$(pgrep $SERVER -u $UID)" == "" ]; then
    echo "$SERVER start failed"
    exit 1
  fi
  echo "$SERVER start success"
}

function status() {
  if [ "$(pgrep $SERVER -u $UID)" != "" ]; then
    echo $SERVER is running
  else
    echo $SERVER is not running
  fi
}

function stop() {
  if [ "$(pgrep $SERVER -u $UID)" != "" ]; then
    kill -9 $(pgrep $SERVER -u $UID)
  fi

  echo "sleeping..." && sleep $INTERVAL

  if [ "$(pgrep $SERVER -u $UID)" != "" ]; then
    echo "$SERVER stop failed"
    exit 1
  fi
}

function main() {
  if [ ! -f "$SERVER" ]; then
    echo "文件不存在"
    exit 1
  fi

  case "$COMMAND" in
  'start')
    start
    ;;
  'stop')
    stop
    ;;
  'status')
    status
    ;;
  'restart')
    stop && start
    ;;
  *)
    echo "参数错误:启动方式必传 usage: $0 {file} {start|stop|restart|status}"
    exit 1
    ;;
  esac
}

main
