{:ok, file} = File.open "hello", [:write]

IO.binwrite file, "world"
File.close file


IO.inspect File.read "hello"
