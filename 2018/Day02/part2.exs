defmodule Day2 do

    defp find_closest_match(_last_item, [], {high_score, id1, id2}) do
        {high_score, id1, id2}
    end

    defp find_closest_match(id, [next_id | ids], {high_score, id1, id2}) do
        pair_distance = String.jaro_distance(id, next_id)
        if pair_distance > high_score do
            find_closest_match(id, ids, {pair_distance, id, next_id})
        else
            find_closest_match(id, ids, {high_score, id1, id2})
        end
    end

    defp find_closest_match([id | ids]) do
        find_closest_match(id, ids, {0, 0, 0})
    end

    defp find_best_pair([_last_item], {overall_best, id1, id2}) do
        {overall_best, id1, id2}
    end

    defp find_best_pair([_id | ids], {overall_best, id1, id2}) do
        {local_best, lid1, lid2} = find_closest_match(ids)
        if local_best > overall_best do
            find_best_pair(ids, {local_best, lid1, lid2})
        else
            find_best_pair(ids, {overall_best, id1, id2})
        end
    end

    def find_best_pair(ids) do
        find_best_pair(ids, find_closest_match(ids))
    end

end

{_score, id1, id2} = File.read!("input.txt") |> String.split |> Day2.find_best_pair
String.myers_difference(id1, id2) |> Keyword.get_values(:eq) |> Enum.join |> IO.puts

