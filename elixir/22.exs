states = %{ "AL" => "Alabama", "WI" => "Wisconsin"}

states["AL"]
|> IO.puts

states["TX"]
|> IO.inspect


response_types = %{ {:error, :enoent} => :fatal,
                    {:error, :busy}   => :retry}

response_types[{:error, :busy}]
|> IO.inspect
