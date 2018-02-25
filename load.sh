#!/bin/bash

cd `dirname $0` || exit
ulimit -c unlimited

start(){
	stop
	sleep 1
	setsid /apps/frontEnd/bin/supervise.frontEnd -u /apps/frontEnd/status/frontEnd env GOTRACEBACK=crash /apps/frontEnd/bin/frontEnd -config /apps/frontEnd/conf/frontEnd.toml
}

stop(){
    killall -9 supervise.frontEnd
	killall -9 frontEnd
}

restart(){
	stop
    sleep 1
    start
}


case C"$1" in
	Cstart)
		start
		echo "start Done!"
		;;
	Cstop)
		stop
		echo "stop Done!"
		;;
	Crestart)
		restart
		echo "restart Done!"
		;;
	C*)
		echo "Usage: $0 {start|stop|restart}"
		;;
esac
