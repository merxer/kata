list = [1, 3, 5, 7, 9]

double = Enum.map list, fn elem -> elem * 2 end
double
|> IO.inspect

expo = Enum.map list, fn elem -> elem * elem end
expo
|> IO.inspect

greater = Enum.map list, fn elem -> elem > 6 end
greater
|> IO.inspect
