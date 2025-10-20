defmodule LibraryFees do
  def datetime_from_string(string) do
    {:ok, naive_datetime} = NaiveDateTime.from_iso8601(string)
    naive_datetime
  end

  def before_noon?(datetime) do
    Time.before?(NaiveDateTime.to_time(datetime), ~T[12:00:00.000])
  end

  def return_date(checkout_datetime) do
    if before_noon?(checkout_datetime) do
      NaiveDateTime.add(checkout_datetime, 28, :day)
      |> NaiveDateTime.to_date
    else
      NaiveDateTime.add(checkout_datetime, 29, :day)
      |> NaiveDateTime.to_date
    end
  end

  def days_late(planned_return_date, actual_return_datetime) do
    date = Date.diff(NaiveDateTime.to_date(actual_return_datetime), planned_return_date)
    if date > 0, do: date, else: 0
  end

  def monday?(datetime) do
    Date.day_of_week(datetime) == 1
  end

  def calculate_late_fee(checkout, return, rate) do
    actual_return_date = LibraryFees.datetime_from_string(return)
    actual_days_late =LibraryFees.days_late(LibraryFees.return_date(LibraryFees.datetime_from_string(checkout)), actual_return_date)
    if LibraryFees.monday?(actual_return_date) do
      floor(actual_days_late * rate * 0.5)
    else
      actual_days_late * rate
    end
  end
end
