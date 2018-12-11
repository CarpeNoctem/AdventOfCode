File.read!("input.txt")
    |> String.split
    |> Enum.map(&String.to_integer/1)
    |> Enum.sum
    |> IO.puts
