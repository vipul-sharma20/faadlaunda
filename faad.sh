#!/bin/bash
case $1 in
    start)
        echo "Starting faadlaunda."
        bin/faadlaunda &
        ;;
    stop)
        echo "Stopping faadlaunda."
        sudo kill $(sudo lsof -t -i:8100)
        ;;
    *)
        echo "Faadlaunda."
        echo $"Usage $0 {start|stop}"
        exit 1
esac
exit 0
