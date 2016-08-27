list = [1,3,5,7,9]
list |> IO.inspect

a = Enum.map list, fn elem -> elem * 2 end
a |> IO.inspect
