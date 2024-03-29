#!/bin/bash

clear

ERROR_STATE=0
MODULE_NAME=test
ROOT_PATH=$GOPATH/src
# MODULE_BIN_NAME='bin/ldc_search.bin'
#PATH_TO_LIB="../lib"
# PATH_TO_WUI="../wui"
LOG_FILE="logs/run.log"
CNTR_FILE="logs/build_counter.bld"
CUT_LINE_BEGIN="--------------------------->8---------------------------------"
CUT_LINE_END="---------------------------8<---------------------------------"
read BUILD_CNT < $CNTR_FILE
# if [ $? -ne 0 ] ; then
#     echo 'First run...'
#     BUILD_CNT=0
#     echo $BUILD_CNT > $CNTR_FILE
# fi
SERVER=as-nginx1
RUNFILE="quadrube"
RESULTFILE="result.csv"
RUN="./$RUNFILE"
CODE="."
SPY=$(find . -name '*.go')
# SPY="$SPY DCpp-global.json DCpp-local.json"


function remoteRun {
    scp $RUNFILE $SERVER: && \
    ssh $SERVER $RUN && \
    scp $SERVER:$RESULTFILE ./ \
    ssh $SERVER rm -f $RUNFILE $RESULTFILE
}



echo spy: $SPY
killall $RUN
cat '' > $LOG_FILE
while true
do
    ATIME=`stat -c %Z $SPY 2>/dev/null`
    if [[ "$ATIME" != "$LTIME" ]]
    then
    clear
    echo -e "$(tput bold)autobuild: [qonetime].[\E[32m'$MODULE_NAME'\E[39m], please wait...$(tput sgr0)"
    echo -e "\E[36m"
    echo "Last build number: "$BUILD_CNT
        pgrep -f -x $RUN > /dev/null 2>&1
        if [ $? -eq 0 ] ; then
            pkill -f -x $RUN
            killall $RUN
        fi

        if [ $ERROR_STATE -ne 0 ];
        then
            echo $CUT_LINE_BEGIN

    else
            echo $CUT_LINE_BEGIN
            echo $(date +%c)' '$MODULE_NAME' build number: '$BUILD_CNT
            echo $(date +%c)' '$MODULE_NAME' build number: '$BUILD_CNT >> logs/build.bld
            echo $BUILD_CNT > $CNTR_FILE
    fi

        time go build $CODE

        if [ $? -eq 0 ]; then
            ERROR_STATE=0
            BUILD_CNT=$((BUILD_CNT+1))
            echo $CUT_LINE_END
            echo -e "\E[0;39m"
            if [[ $1 = "-r" ]]; then
                remoteRun
            else
                $RUN &
            fi
        else
            ERROR_STATE=1
        fi
        LTIME=$ATIME
    fi
    # echo -e "\E[32m"
    # echo 'TICK: '$ATIME' ? '$LTIME
    sleep 1
done

