defmodule Math do
  defmodule Division do
    def divide(a, b) do
      a /b
    end
  end
end


Math.Division.divide(15,5)
|> IO.puts
