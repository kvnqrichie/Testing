#!/bin/sh
curl -s "https://acad.learn2earn.ng/assets/superhero/all.json"  | jq -r --arg id "$HERO_ID" '.[] | select((.id | tostring) == $id) | .connections.relatives' | awk '{printf "%s%s", (NR==1?"":"\\n"), $0} END {print ""}'
