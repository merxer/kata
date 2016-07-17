states = %{ "AL" => "Alabama", "WI" => "Wisconsin"}
states
|> IO.inspect

responses = %{ { :error, :enoent } => :fatal, {:error, :busy} => :retry}
responses
|> IO.inspect

colors = %{ :red => 0xff0000, :green => 0x00ff00, :blue => 0x0000ff }
colors
|> IO.inspect

