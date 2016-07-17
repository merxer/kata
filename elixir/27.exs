with [a|_] <- [1,2,3], do: a
|> IO.inspect

with [a|_]<- nil, do: a
|> IO.inspect
