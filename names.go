package main

import (
	"fmt"
)

func makeName(sep string) string {
	first := pickRandomString(firstNames)
	last := pickRandomString(lastNames)

	return fmt.Sprintf("%s%s%s", first, sep, last)
}

var firstNames = []string{
	"Abigail",
	"Alexandra",
	"Alison",
	"Amanda",
	"Amelia",
	"Amy",
	"Andrea",
	"Angela",
	"Anna",
	"Anne",
	"Audrey",
	"Ava",
	"Bella",
	"Bernadette",
	"Carol",
	"Caroline",
	"Carolyn",
	"Chloe",
	"Claire",
	"Deirdre",
	"Diana",
	"Diane",
	"Donna",
	"Dorothy",
	"Elizabeth",
	"Ella",
	"Emily",
	"Emma",
	"Faith",
	"Felicity",
	"Fiona",
	"Gabrielle",
	"Grace",
	"Hannah",
	"Heather",
	"Irene",
	"Jan",
	"Jane",
	"Jasmine",
	"Jennifer",
	"Jessica",
	"Joan",
	"Joanne",
	"Julia",
	"Karen",
	"Katherine",
	"Kimberly",
	"Kylie",
	"Lauren",
	"Leah",
	"Lillian",
	"Lily",
	"Lisa",
	"Madeleine",
	"Maria",
	"Mary",
	"Megan",
	"Melanie",
	"Michelle",
	"Molly",
	"Natalie",
	"Nicola",
	"Olivia",
	"Penelope",
	"Pippa",
	"Rachel",
	"Rebecca",
	"Rose",
	"Ruth",
	"Sally",
	"Samantha",
	"Sarah",
	"Sonia",
	"Sophie",
	"Stephanie",
	"Sue",
	"Theresa",
	"Tracey",
	"Una",
	"Vanessa",
	"Victoria",
	"Virginia",
	"Wanda",
	"Wendy",
	"Yvonne",
	"Zoe",
	"Adam",
	"Adrian",
	"Alan",
	"Alexander",
	"Andrew",
	"Anthony",
	"Austin",
	"Benjamin",
	"Blake",
	"Boris",
	"Brandon",
	"Brian",
	"Cameron",
	"Carl",
	"Charles",
	"Christian",
	"Christopher",
	"Colin",
	"Connor",
	"Dan",
	"David",
	"Dominic",
	"Dylan",
	"Edward",
	"Eric",
	"Evan",
	"Frank",
	"Gavin",
	"Gordon",
	"Harry",
	"Ian",
	"Isaac",
	"Jack",
	"Jacob",
	"Jake",
	"James",
	"Jason",
	"Joe",
	"John",
	"Jonathan",
	"Joseph",
	"Joshua",
	"Julian",
	"Justin",
	"Keith",
	"Kevin",
	"Leonard",
	"Liam",
	"Lucas",
	"Luke",
	"Matt",
	"Max",
	"Michael",
	"Nathan",
	"Neil",
	"Nicholas",
	"Oliver",
	"Owen",
	"Paul",
	"Peter",
	"Phil",
	"Piers",
	"Richard",
	"Robert",
	"Ryan",
	"Sam",
	"Sean",
	"Sebastian",
	"Simon",
	"Stephen",
	"Steven",
	"Stewart",
	"Thomas",
	"Tim",
	"Trevor",
	"Victor",
	"Warren",
	"William",
}

var lastNames = []string{
	"Abraham",
	"Allan",
	"Alsop",
	"Anderson",
	"Arnold",
	"Avery",
	"Bailey",
	"Baker",
	"Ball",
	"Bell",
	"Berry",
	"Black",
	"Blake",
	"Bond",
	"Bower",
	"Brown",
	"Buckland",
	"Burgess",
	"Butler",
	"Cameron",
	"Campbell",
	"Carr",
	"Chapman",
	"Churchill",
	"Clark",
	"Clarkson",
	"Coleman",
	"Cornish",
	"Davidson",
	"Davies",
	"Dickens",
	"Dowd",
	"Duncan",
	"Dyer",
	"Edmunds",
	"Ellison",
	"Ferguson",
	"Fisher",
	"Forsyth",
	"Fraser",
	"Gibson",
	"Gill",
	"Glover",
	"Graham",
	"Grant",
	"Gray",
	"Greene",
	"Hamilton",
	"Hardacre",
	"Harris",
	"Hart",
	"Hemmings",
	"Henderson",
	"Hill",
	"Hodges",
	"Howard",
	"Hudson",
	"Hughes",
	"Hunter",
	"Ince",
	"Jackson",
	"James",
	"Johnston",
	"Jones",
	"Kelly",
	"Kerr",
	"King",
	"Knox",
	"Lambert",
	"Langdon",
	"Lawrence",
	"Lee",
	"Lewis",
	"Lyman",
	"MacDonald",
	"Mackay",
	"Mackenzie",
	"MacLeod",
	"Manning",
	"Marshall",
	"Martin",
	"Mathis",
	"May",
	"McDonald",
	"McLean",
	"McGrath",
	"Metcalfe",
	"Miller",
	"Mills",
	"Mitchell",
	"Morgan",
	"Morrison",
	"Murray",
	"Nash",
	"Newman",
	"Nolan",
	"North",
	"Ogden",
	"Oliver",
	"Paige",
	"Parr",
	"Parsons",
	"Paterson",
	"Payne",
	"Peake",
	"Peters",
	"Piper",
	"Poole",
	"Powell",
	"Pullman",
	"Quinn",
	"Rampling",
	"Randall",
	"Rees",
	"Reid",
	"Roberts",
	"Robertson",
	"Ross",
	"Russell",
	"Rutherford",
	"Sanderson",
	"Scott",
	"Sharp",
	"Short",
	"Simpson",
	"Skinner",
	"Slater",
	"Smith",
	"Springer",
	"Stewart",
	"Sutherland",
	"Taylor",
	"Terry",
	"Thomson",
	"Tucker",
	"Turner",
	"Underwood",
	"Vance",
	"Vaughan",
	"Walker",
	"Wallace",
	"Walsh",
	"Watson",
	"Welch",
	"White",
	"Wilkins",
	"Wilson",
	"Wright",
	"Young",
}
