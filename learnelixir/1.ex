defmodule Math do
  def add42(a,b) do
    add(a,b) + 42
  end

  def add(a,b) do
    a + b
  end
end


Math.add42(1,2)
|> IO.puts
