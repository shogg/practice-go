#! /bin/bash

if [[ "$TRAVIS_PULL_REQUEST" == "false" ]]
then
    exit
fi

COMMITTER=$(echo ${TRAVIS_PULL_REQUEST_SLUG} | sed -r 's#(.*)/.*#\1#')
SLUG=${TRAVIS_REPO_SLUG}
COMMIT=${TRAVIS_PULL_REQUEST_SHA}

PROBLEM_DIR=$(git show --name-status $COMMIT | tail -n 1 | grep -oP '(?<=\s).*(?=/)')
if [[ -z "$PROBLEM_DIR" || "${PROBLEM_DIR:0:1}" == "." ]]
then
    exit
fi

SOLUTION_LINK=$(echo "https://github.com/$SLUG/blob/$COMMIT/$PROBLEM_DIR/$PROBLEM_DIR.go")

cd $PROBLEM_DIR

# benchmark markdown
echo $COMMITTER "[solution]($SOLUTION_LINK)"
echo '```'
go test -bench . -benchmem
echo '```'
