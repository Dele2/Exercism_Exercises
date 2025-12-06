defmodule BoutiqueInventory do
  def sort_by_price(inventory) do
    Enum.sort_by(inventory, &(&1.price))
  end

  def with_missing_price(inventory) do
    Enum.filter(inventory, &(&1.price == nil))
  end

  def update_names(inventory, old_word, new_word) do
    Enum.map(inventory, fn map ->
      Map.put(map, :name, String.replace(map.name, old_word, new_word)) end)
  end

  def increase_quantity(item, count) do
    %{quantity_by_size: qty} = item
    %{item | quantity_by_size: process_quantity_by_size(qty, count)}
  end
  
  defp process_quantity_by_size(inventory, count) do
    Map.new(inventory,  fn {k, v} -> {k, v + count} end)
  end

  def total_quantity(item) do
    %{quantity_by_size: qty} = item
    Enum.reduce(qty, 0, fn({_key, value}, acc) -> value + acc end)
  end
end
