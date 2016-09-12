defmodule MyList do
  def len([]), do: 0
  def len([_|tail]), do: 1+len(tail)
end


MyList.len([1,2,3,4,5,6])
|> IO.puts
