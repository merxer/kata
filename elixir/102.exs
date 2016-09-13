defmodule MyList do
  def square([]), do: []
  def square([head | tail]), do: [ head * head |square(tail)]
end

MyList.square([1,2,3,4,5]) |> IO.inspect
