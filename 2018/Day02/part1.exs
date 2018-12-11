defmodule Day2 do

    defp get_checksum([], twos, threes) do
        twos * threes
    end

    defp get_checksum([id | ids], twos, threes) do
        stats = parse_id(String.split(id, ""), %{}) |> Map.values
        twos = 2 in stats && twos + 1 || twos
        threes = 3 in stats && threes + 1 || threes
        get_checksum(ids, twos, threes)
    end

    def get_checksum(ids) do
        get_checksum(ids, 0, 0)
    end

    defp parse_id([], stats) do
        stats
    end

    defp parse_id([char | chars], stats) do
        parse_id(chars, Map.update(stats, char, 1, &(&1 + 1)))
    end

end

File.read!("input.txt") |> String.split |> Day2.get_checksum |> IO.puts
