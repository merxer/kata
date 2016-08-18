ast = quote do
  if var!(meaning_to_life) == 42 do
    "it's true"
  else
    "It remains to be seen"
  end
end

ast |> IO.inspect
