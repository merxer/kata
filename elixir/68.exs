defmodule Times do
  def triple(n), do: n * n * n

  def double(n), do: n * n

  def quadruple(n), do: double(n) * double(n)
end


Times.double(2)
|> IO.puts

Times.triple(2)
|> IO.puts

Times.quadruple(2)
|> IO.puts
