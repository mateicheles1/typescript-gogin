package data

import "gogin/models"

var Data = models.AppData{
	Albums: []models.Album{
		{Id: 1, Title: "72 Seasons", Artist: "Metallica", Price: 39.99},
		{Id: 2, Title: "Let Them Talk", Artist: "Hugh Laurie", Price: 29.99},
		{Id: 3, Title: "B.B. King & Friends - 80", Artist: "B.B. King", Price: 22.50},
	},
}
