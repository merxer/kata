case {1, 2, 3} do
  {4, 5, 6} -> IO.puts "This clause won't match"
  {1, x, 3} when x > 0 -> IO.puts "This clause will match and bind x to 2"
  _ -> IO.puts "This clause would match any value"
end
