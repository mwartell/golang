#!/bin/sh


# generate blog posts using ollama for some topics

if [ $# -eq 0 ]; then
    echo "usage: $0 topic1 topic2 topic3"
    echo "   or: $0 topics.txt where file contains a list of topics"
    exit 1
fi

if [ $# -eq 1 -a -f "$1" ]; then
    topics=$(cat "$1")
else
    topics="$@"
fi

if [ ! -d "posts" ]; then
    mkdir posts
fi

for topic in $topics; do
    echo "generating posts/$topic.md"
    ollama run blogger "write a blog post about $topic in mathematics" > posts/"$topic".md
done
