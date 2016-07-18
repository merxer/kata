greeter = fn name ->
  (fn -> "Hello #{name}" end)
end

dave_greeter = greeter.("Dave")

dave_greeter.()
|> IO.inspect
