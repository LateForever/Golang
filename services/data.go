package main

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Year   int     `json:"year"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Year: 1957, Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Year: 1960, Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Year: 1955, Price: 39.99},
	{ID: "4", Title: "Alone Together", Artist: "Bill Evans", Year: 1959, Price: 57.99},
	{ID: "5", Title: "The Modern Sound of Betty Carter", Artist: "Betty Carter", Year: 1962, Price: 49.99},
	{ID: "6", Title: "The Quintessence", Artist: "Joe Henderson", Year: 1963, Price: 37.99},
	{ID: "7", Title: "Something Blue", Artist: "Ella Fitzgerald", Year: 1958, Price: 23.99},
	{ID: "8", Title: "The Best of Miles Davis", Artist: "Miles Davis", Year: 1958, Price: 56.99},
	{ID: "9", Title: "The Best of Bill Evans", Artist: "Bill Evans", Year: 1961, Price: 18.99},
	{ID: "10", Title: "The Best of Charlie Parker", Artist: "Charlie Parker", Year: 1952, Price: 56.99},
	{ID: "11", Title: "The Best of Chet Baker", Artist: "Chet Baker", Year: 1955, Price: 56.99},
	{ID: "12", Title: "The Best of John Coltrane", Artist: "John Coltrane", Year: 1957, Price: 56.99},
	{ID: "13", Title: "The Best of Art Blakey", Artist: "Art Blakey", Year: 1958, Price: 56.99},
	{ID: "14", Title: "The Best of Thelonious Monk", Artist: "Thelonious Monk", Year: 1957, Price: 56.99},
	{ID: "15", Title: "The Best of Sonny Rollins", Artist: "Sonny Rollins", Year: 1957, Price: 56.99},
	{ID: "16", Title: "The Best of Coleman Hawkins", Artist: "Coleman Hawkins", Year: 1957, Price: 56.99},
	{ID: "17", Title: "The Best of Dexter Gordon", Artist: "Dexter Gordon", Year: 1957, Price: 56.99},
	{ID: "18", Title: "The Best of Wes Montgomery", Artist: "Wes Montgomery", Year: 1957, Price: 56.99},
	{ID: "19", Title: "The Best of Cannonball Adderley", Artist: "Cannonball Adderley", Year: 1957, Price: 56.99},
	{ID: "20", Title: "The Best of Miles Davis", Artist: "Miles Davis", Year: 1957, Price: 56.99},
	{ID: "21", Title: "The Best of Miles Davis", Artist: "Miles Davis", Year: 1957, Price: 56.99},
	{ID: "22", Title: "The Best of Miles Davis", Artist: "Miles Davis", Year: 1957, Price: 56.99},
	{ID: "23", Title: "The Best of Miles Davis", Artist: "Miles Davis", Year: 1957, Price: 56.99},
	{ID: "24", Title: "The Best of Miles Davis", Artist: "Miles Davis", Year: 1957, Price: 56.99},
	{ID: "25", Title: "The Best of Miles Davis", Artist: "Miles Davis", Year: 1957, Price: 56.99},
	{ID: "26", Title: "The Best of Miles Davis", Artist: "Miles Davis", Year: 1957, Price: 56.99},
	{ID: "27", Title: "The Best of Miles Davis", Artist: "Miles Davis", Year: 1957, Price: 56.99},
	{ID: "28", Title: "The Best of Miles Davis", Artist: "Miles Davis", Year: 1957, Price: 56.99},
	{ID: "29", Title: "The Best of Miles Davis", Artist: "Miles Davis", Year: 1957, Price: 56.99},
	{ID: "30", Title: "The Best of Miles Davis", Artist: "Miles Davis", Year: 1957, Price: 56.99},
	{ID: "31", Title: "The Best of Miles Davis", Artist: "Miles Davis", Year: 1957, Price: 56.99},
	{ID: "32", Title: "The Best of Miles Davis", Artist: "Miles Davis", Year: 1957, Price: 56.99},
	{ID: "33", Title: "The Best of Miles Davis", Artist: "Miles Davis", Year: 1957, Price: 56.99},
	{ID: "34", Title: "The Best of Miles Davis", Artist: "Miles Davis", Year: 1957, Price: 56.99},
	{ID: "35", Title: "The Best of Miles Davis", Artist: "Miles Davis", Year: 1957, Price: 56.99},
	{ID: "36", Title: "The Best of Miles Davis", Artist: "Miles Davis", Year: 1957, Price: 56.99},
	{ID: "37", Title: "The Best of Miles Davis", Artist: "Miles Davis", Year: 1957, Price: 56.99},
	{ID: "38", Title: "The Best of Miles Davis", Artist: "Miles Davis", Year: 1957, Price: 56.99},
	{ID: "39", Title: "The Best of Miles Davis", Artist: "Miles Davis", Year: 1957, Price: 56.99},
	{ID: "40", Title: "The Best of Miles Davis", Artist: "Miles Davis", Year: 1957, Price: 56.99},
}