handle_open = fn
  {:ok, file} -> "Read data: #{IO.read(file, :line)}"
  {_, error} -> "Error: #{:file.format_error(error)}"
end

handle_open.(File.open("./passwd"))
|> IO.puts

handle_open.(File.open("nonexistent"))
|> IO.puts
