#!/bin/bash

if [ -z $1]
then
	echo "please enter a tag"
	exit 0
fi

docker build -t afeeblechild/meal-planner:$1 .
docker push afeeblechild/meal-planner:$1
