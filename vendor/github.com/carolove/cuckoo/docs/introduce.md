

### UML diagrams

You can also render sequence diagrams like this:

```sequence
Condidate->Follower:RequestVote
Follower->Condidate:Ack
Leader->Follower:AppendEntries
Follower->Leader:Ack
Leader->Follower:InstallSnapshot 
Follower->Leader:Ack
```

And flow charts like this:

```flow
st=>start: Start
e=>end
op=>operation: My Operation
cond=>condition: Yes or No?

st->op->cond
cond(yes)->e
cond(no)->op
```
