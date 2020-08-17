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

Csv file with statistics will be created at `path/to/exported/dir/stats.csv`.

For each chat conversation is split into "sessions" based on **Arg 2** value.
File includes the following columns:

* **Contact**: name of the chat
* **I initiated %**: percent of sessions initiated by me (the rest is initiated by a contact)
* **My msgs/session**: average messages per session from me
* **Contact's msgs/session**: average messages per session from a contact
* **My chars/message**: average length of my messages
* **Contact's chars/message**: average length of contact's messages
* **Avg messages/session**: average messages per session (total my + contact's)
* **My avg reaction (sec)**: average amount of seconds after which I reply
* **Contact's avg reaction (sec)**: average amount of seconds after which a contact replies

Have fun sorting file by different columns to compare your's and contact's behavior! 