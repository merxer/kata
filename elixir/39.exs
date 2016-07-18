prefix = fn pre -> (fn other -> pre <> " " <>  other end ) end

mrs = prefix.("Mrs")

mrs.("Smith")
|> IO.inspect

prefix.("Elixir").("Rocks")
|> IO.inspect
