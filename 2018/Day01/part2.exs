File.read!("input.txt")
    |> String.split
    |> Enum.map(&String.to_integer/1)
    |> Stream.cycle
    |> Enum.reduce_while({0, MapSet.new([0])}, fn(input, {prev_offset, offsets}) ->
           new_offset = input + prev_offset
           if new_offset in offsets do
               {:halt, new_offset}
           else
               {:cont, {new_offset, MapSet.put(offsets, new_offset)}}
           end
       end)
    |> IO.puts
