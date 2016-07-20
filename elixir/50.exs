defmodule Times do

  def double(n), do: n * n

  def triple(n), do: n * n * n

  def quadruple(n), do: double(n) * double(n)

end


Times.triple(3)
|> IO.inspect


Times.quadruple(4)
|> IO.inspect
