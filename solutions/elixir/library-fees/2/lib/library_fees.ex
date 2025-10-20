defmodule LibraryFees do
  def datetime_from_string(string) do
    NaiveDateTime.from_iso8601!(string)
  end

  def before_noon?(datetime) do
    Time.before?(NaiveDateTime.to_time(datetime), ~T[12:00:00.000])
  end

  def return_date(checkout_datetime) do
    days = if before_noon?(checkout_datetime), do: 28, else: 29
    Date.add(checkout_datetime, days)

  end

  def days_late(planned_return_date, actual_return_datetime) do
   actual_return_datetime
   |> Date.diff(planned_return_date)
   |> max(0)
  end

  def monday?(datetime) do
    Date.day_of_week(datetime) == 1
  end

   def calculate_late_fee(checkout, return, rate) do
    planned_return_date = return_date(datetime_from_string(checkout))
    actual_return_date = datetime_from_string(return)
    days_late = days_late(planned_return_date, actual_return_date)
    if monday?(actual_return_date) do
      div(days_late * rate, 2)
    else
      days_late * rate
    end
  end

end