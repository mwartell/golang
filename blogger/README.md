# stupid markdown blogpost generation for testing

Running

    task posts

will add ollama generated posts on a series of topics in `topics.txt`

* `blogger.modelfile` — contains an ollama "model" which specifies an
llm and a system prompt
* `blogger.sh` - reads topics from a file and creates `posts/{topic}.md`
for each topic in a textfile
* `Taskfile.yml` checking out https://taskfile.dev/ for make-less makes
