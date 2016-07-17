{status, file } = File.open("15.exs")

IO.inspect status
IO.inspect file

#open non existent file
{status, file } = File.open("non-existent-file") 
IO.inspect status
IO.inspect file
