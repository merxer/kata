l = &length/1
l.([1,3,5,7])
|> IO.inspect


len = &Enum.count/1
len.([1,2,3,4])
|> IO.inspect


m = &Kernel.min/2
m.(99,88)
|> IO.puts
