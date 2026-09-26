export INTERVIEW_NUM=$(head -n 179 ./streets/Buckingham_Place | tail -n 1 | tr -dc '0-9')
echo $INTERVIEW_NUM
cat ./interviews/interview-$INTERVIEW_NUM
echo $MAIN_SUSPECT