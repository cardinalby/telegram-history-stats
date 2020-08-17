## Telegram chats statistics

It's only a learning project!

The applications analyzes your chats (excluding groups) and 
calculates some interesting statistics.

Use Telegram Desktop to export your chats history (you may not include media data).
Then run:

```
tlgstats "path/to/exported/dir" 10
```

Where:
* **Arg 1** is path to exported directory and the destination for csv file at the same time
* **Arg 2** is an interval in hours used to split chat history into "sessions"

Csv file with statistics will be created inside "path/to/exported/dir/".



It includes the following stats for chats: