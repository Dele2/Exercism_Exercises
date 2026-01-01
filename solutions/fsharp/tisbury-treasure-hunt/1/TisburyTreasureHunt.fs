module TisburyTreasureHunt

let getCoordinate (line: string * string): string =
    snd line

let convertCoordinate (coordinate: string): int * char = 
    let inline charToInt c = int c - int '0'
    let num = coordinate[0]
    let text = coordinate[1]
    num |> charToInt, text

let compareRecords (azarasData: string * string) (ruisData: string * (int * char) * string) : bool = 
    let azarasPosition = azarasData |> getCoordinate |> convertCoordinate
    let (_, ruisPosition, _) = ruisData
    azarasPosition = ruisPosition

let createRecord (azarasData: string * string) (ruisData: string * (int * char) * string) : (string * string * string * string) =
    let (location1, position1) = azarasData
    let (location2, _, quandrant) = ruisData

 
    if compareRecords (azarasData) (ruisData) then
        (position1, location2, quandrant, location1)
    else
        ("", "", "", "")
