package common

// A code for each departement of France.
// For metropolitan, the const code equal departement number.
// Else it can be different and is not dense.
type Departement byte

const (
	DepartementAin                  Departement = 1
	DepartementAisne                Departement = 2
	DepartementAllier               Departement = 3
	DepartementAlpesDeHauteProvence Departement = 4
	DepartementHautesAlpes          Departement = 5
	DepartementAlpesMaritimes       Departement = 6
	DepartementArdèche              Departement = 7
	DepartementArdennes             Departement = 8
	DepartementAriège               Departement = 9
	DepartementAube                 Departement = 10
	DepartementAude                 Departement = 11
	DepartementAveyron              Departement = 12
	DepartementBouchesDuRhône       Departement = 13
	DepartementCalvados             Departement = 14
	DepartementCantal               Departement = 15
	DepartementCharente             Departement = 16
	DepartementCharenteMaritime     Departement = 17
	DepartementCher                 Departement = 18
	DepartementCorrèze              Departement = 19
	DepartementCôteDOr              Departement = 21
	DepartementCôtesDArmor          Departement = 22
	DepartementCreuse               Departement = 23
	DepartementDordogne             Departement = 24
	DepartementDoubs                Departement = 25
	DepartementDrôme                Departement = 26
	DepartementEure                 Departement = 27
	DepartementEureEtLoir           Departement = 28
	DepartementFinistère            Departement = 29
	DepartementGard                 Departement = 30
	DepartementHauteGaronne         Departement = 31
	DepartementGers                 Departement = 32
	DepartementGironde              Departement = 33
	DepartementHérault              Departement = 34
	DepartementIlleEtVilaine        Departement = 35
	DepartementIndre                Departement = 36
	DepartementIndreEtLoire         Departement = 37
	DepartementIsère                Departement = 38
	DepartementJura                 Departement = 39
	DepartementLandes               Departement = 40
	DepartementLoirEtCher           Departement = 41
	DepartementLoire                Departement = 42
	DepartementHauteLoire           Departement = 43
	DepartementLoireAtlantique      Departement = 44
	DepartementLoiret               Departement = 45
	DepartementLot                  Departement = 46
	DepartementLotEtGaronne         Departement = 47
	DepartementLozère               Departement = 48
	DepartementMaineEtLoire         Departement = 49
	DepartementManche               Departement = 50
	DepartementMarne                Departement = 51
	DepartementHauteMarne           Departement = 52
	DepartementMayenne              Departement = 53
	DepartementMeurtheEtMoselle     Departement = 54
	DepartementMeuse                Departement = 55
	DepartementMorbihan             Departement = 56
	DepartementMoselle              Departement = 57
	DepartementNièvre               Departement = 58
	DepartementNord                 Departement = 59
	DepartementOise                 Departement = 60
	DepartementOrne                 Departement = 61
	DepartementPasDeCalais          Departement = 62
	DepartementPuyDeDôme            Departement = 63
	DepartementPyrénéesAtlantiques  Departement = 64
	DepartementHautesPyrénées       Departement = 65
	DepartementPyrénéesOrientales   Departement = 66
	DepartementBasRhin              Departement = 67
	DepartementHautRhin             Departement = 68
	DepartementRhône                Departement = 69
	DepartementHauteSaône           Departement = 70
	DepartementSaôneEtLoire         Departement = 71
	DepartementSarthe               Departement = 72
	DepartementSavoie               Departement = 73
	DepartementHauteSavoie          Departement = 74
	DepartementParis                Departement = 75
	DepartementSeineMaritime        Departement = 76
	DepartementSeineEtMarne         Departement = 77
	DepartementYvelines             Departement = 78
	DepartementDeuxSèvres           Departement = 79
	DepartementSomme                Departement = 80
	DepartementTarn                 Departement = 81
	DepartementTarnEtGaronne        Departement = 82
	DepartementVar                  Departement = 83
	DepartementVaucluse             Departement = 84
	DepartementVendée               Departement = 85
	DepartementVienne               Departement = 86
	DepartementHauteVienne          Departement = 87
	DepartementVosges               Departement = 88
	DepartementYonne                Departement = 89
	DepartementTerritoireDeBelfort  Departement = 90
	DepartementEssonne              Departement = 91
	DepartementHautsDeSeine         Departement = 92
	DepartementSeineSaintDenis      Departement = 93
	DepartementValDeMarne           Departement = 94
	DepartementValDOise             Departement = 95

	DepartementCorseDuSud Departement = 96
	DepartementHauteCorse Departement = 97

	DepartementFrançaisÉtablisHorsDeFrance Departement = 98
	DepartementSaintMartinSaintBarthélemy  Departement = 99

	DepartementGuadeloupe            Departement = 171
	DepartementMartinique            Departement = 172
	DepartementGuyane                Departement = 173
	DepartementLaRéunion             Departement = 174
	DepartementSaintPierreEtMiquelon Departement = 175
	DepartementMayotte               Departement = 176

	DepartementWallisEtFutuna     Departement = 186
	DepartementPolynésieFrançaise Departement = 187
	DepartementNouvelleCalédonie  Departement = 188
)

func (d Departement) String() string {
	return departementConst2Name[d]
}

func (d Departement) Code() string {
	return departementConst2Code[d]
}

var departementConst2Name = [256]string{
	1:  "Ain",
	2:  "Aisne",
	3:  "Allier",
	4:  "Alpes-de-Haute-Provence",
	5:  "Hautes-Alpes",
	6:  "Alpes-Maritimes",
	7:  "Ardèche",
	8:  "Ardennes",
	9:  "Ariège",
	10: "Aube",
	11: "Aude",
	12: "Aveyron",
	13: "Bouches-du-Rhône",
	14: "Calvados",
	15: "Cantal",
	16: "Charente",
	17: "Charente-Maritime",
	18: "Cher",
	19: "Corrèze",
	21: "Côte-d'Or",
	22: "Côtes-d'Armor",
	23: "Creuse",
	24: "Dordogne",
	25: "Doubs",
	26: "Drôme",
	27: "Eure",
	28: "Eure-et-Loir",
	29: "Finistère",
	30: "Gard",
	31: "Haute-Garonne",
	32: "Gers",
	33: "Gironde",
	34: "Hérault",
	35: "Ille-et-Vilaine",
	36: "Indre",
	37: "Indre-et-Loire",
	38: "Isère",
	39: "Jura",
	40: "Landes",
	41: "Loir-et-Cher",
	42: "Loire",
	43: "Haute-Loire",
	44: "Loire-Atlantique",
	45: "Loiret",
	46: "Lot",
	47: "Lot-et-Garonne",
	48: "Lozère",
	49: "Maine-et-Loire",
	50: "Manche",
	51: "Marne",
	52: "Haute-Marne",
	53: "Mayenne",
	54: "Meurthe-et-Moselle",
	55: "Meuse",
	56: "Morbihan",
	57: "Moselle",
	58: "Nièvre",
	59: "Nord",
	60: "Oise",
	61: "Orne",
	62: "Pas-de-Calais",
	63: "Puy-de-Dôme",
	64: "Pyrénées-Atlantiques",
	65: "Hautes-Pyrénées",
	66: "Pyrénées-Orientales",
	67: "Bas-Rhin",
	68: "Haut-Rhin",
	69: "Rhône",
	70: "Haute-Saône",
	71: "Saône-et-Loire",
	72: "Sarthe",
	73: "Savoie",
	74: "Haute-Savoie",
	75: "Paris",
	76: "Seine-Maritime",
	77: "Seine-et-Marne",
	78: "Yvelines",
	79: "Deux-Sèvres",
	80: "Somme",
	81: "Tarn",
	82: "Tarn-et-Garonne",
	83: "Var",
	84: "Vaucluse",
	85: "Vendée",
	86: "Vienne",
	87: "Haute-Vienne",
	88: "Vosges",
	89: "Yonne",
	90: "Territoire de Belfort",
	91: "Essonne",
	92: "Hauts-de-Seine",
	93: "Seine-Saint-Denis",
	94: "Val-de-Marne",
	95: "Val-d'Oise",

	DepartementCorseDuSud: "Corse-du-Sud",
	DepartementHauteCorse: "Haute-Corse",

	DepartementFrançaisÉtablisHorsDeFrance: "Français établis hors de France",
	DepartementSaintMartinSaintBarthélemy:  "Saint-Martin/Saint-Barthélemy",

	DepartementGuadeloupe:            "Guadeloupe",
	DepartementMartinique:            "Martinique",
	DepartementGuyane:                "Guyane",
	DepartementLaRéunion:             "La Réunion",
	DepartementSaintPierreEtMiquelon: "Saint-Pierre-et-Miquelon",
	DepartementMayotte:               "Mayotte",

	DepartementWallisEtFutuna:     "Wallis et Futuna",
	DepartementPolynésieFrançaise: "Polynésie française",
	DepartementNouvelleCalédonie:  "Nouvelle-Calédonie",
}

var departementConst2Code = [256]string{
	1:  "1",
	2:  "2",
	3:  "3",
	4:  "4",
	5:  "5",
	6:  "6",
	7:  "7",
	8:  "8",
	9:  "9",
	10: "10",
	11: "11",
	12: "12",
	13: "13",
	14: "14",
	15: "15",
	16: "16",
	17: "17",
	18: "18",
	19: "19",
	21: "21",
	22: "22",
	23: "23",
	24: "24",
	25: "25",
	26: "26",
	27: "27",
	28: "28",
	29: "29",
	30: "30",
	31: "31",
	32: "32",
	33: "33",
	34: "34",
	35: "35",
	36: "36",
	37: "37",
	38: "38",
	39: "39",
	40: "40",
	41: "41",
	42: "42",
	43: "43",
	44: "44",
	45: "45",
	46: "46",
	47: "47",
	48: "48",
	49: "49",
	50: "50",
	51: "51",
	52: "52",
	53: "53",
	54: "54",
	55: "55",
	56: "56",
	57: "57",
	58: "58",
	59: "59",
	60: "60",
	61: "61",
	62: "62",
	63: "63",
	64: "64",
	65: "65",
	66: "66",
	67: "67",
	68: "68",
	69: "69",
	70: "70",
	71: "71",
	72: "72",
	73: "73",
	74: "74",
	75: "75",
	76: "76",
	77: "77",
	78: "78",
	79: "79",
	80: "80",
	81: "81",
	82: "82",
	83: "83",
	84: "84",
	85: "85",
	86: "86",
	87: "87",
	88: "88",
	89: "89",
	90: "90",
	91: "91",
	92: "92",
	93: "93",
	94: "94",
	95: "95",

	DepartementCorseDuSud: "2A",
	DepartementHauteCorse: "2B",

	DepartementFrançaisÉtablisHorsDeFrance: "ZZ",
	DepartementSaintMartinSaintBarthélemy:  "ZX",

	DepartementGuadeloupe:            "971",
	DepartementMartinique:            "972",
	DepartementGuyane:                "973",
	DepartementLaRéunion:             "974",
	DepartementSaintPierreEtMiquelon: "975",
	DepartementMayotte:               "976",

	DepartementWallisEtFutuna:     "986",
	DepartementPolynésieFrançaise: "987",
	DepartementNouvelleCalédonie:  "988",
}

var DepartementName2Const = map[string]Departement{
	"Ain":                     1,
	"Aisne":                   2,
	"Allier":                  3,
	"Alpes-de-Haute-Provence": 4,
	"Hautes-Alpes":            5,
	"Alpes-Maritimes":         6,
	"Ardèche":                 7,
	"Ardennes":                8,
	"Ariège":                  9,
	"Aube":                    10,
	"Aude":                    11,
	"Aveyron":                 12,
	"Bouches-du-Rhône":        13,
	"Calvados":                14,
	"Cantal":                  15,
	"Charente":                16,
	"Charente-Maritime":       17,
	"Cher":                    18,
	"Corrèze":                 19,
	"Côte-d'Or":               21,
	"Côtes-d'Armor":           22,
	"Creuse":                  23,
	"Dordogne":                24,
	"Doubs":                   25,
	"Drôme":                   26,
	"Eure":                    27,
	"Eure-et-Loir":            28,
	"Finistère":               29,
	"Gard":                    30,
	"Haute-Garonne":           31,
	"Gers":                    32,
	"Gironde":                 33,
	"Hérault":                 34,
	"Ille-et-Vilaine":         35,
	"Indre":                   36,
	"Indre-et-Loire":          37,
	"Isère":                   38,
	"Jura":                    39,
	"Landes":                  40,
	"Loir-et-Cher":            41,
	"Loire":                   42,
	"Haute-Loire":             43,
	"Loire-Atlantique":        44,
	"Loiret":                  45,
	"Lot":                     46,
	"Lot-et-Garonne":          47,
	"Lozère":                  48,
	"Maine-et-Loire":          49,
	"Manche":                  50,
	"Marne":                   51,
	"Haute-Marne":             52,
	"Mayenne":                 53,
	"Meurthe-et-Moselle":      54,
	"Meuse":                   55,
	"Morbihan":                56,
	"Moselle":                 57,
	"Nièvre":                  58,
	"Nord":                    59,
	"Oise":                    60,
	"Orne":                    61,
	"Pas-de-Calais":           62,
	"Puy-de-Dôme":             63,
	"Pyrénées-Atlantiques":    64,
	"Hautes-Pyrénées":         65,
	"Pyrénées-Orientales":     66,
	"Bas-Rhin":                67,
	"Haut-Rhin":               68,
	"Rhône":                   69,
	"Haute-Saône":             70,
	"Saône-et-Loire":          71,
	"Sarthe":                  72,
	"Savoie":                  73,
	"Haute-Savoie":            74,
	"Paris":                   75,
	"Seine-Maritime":          76,
	"Seine-et-Marne":          77,
	"Yvelines":                78,
	"Deux-Sèvres":             79,
	"Somme":                   80,
	"Tarn":                    81,
	"Tarn-et-Garonne":         82,
	"Var":                     83,
	"Vaucluse":                84,
	"Vendée":                  85,
	"Vienne":                  86,
	"Haute-Vienne":            87,
	"Vosges":                  88,
	"Yonne":                   89,
	"Territoire de Belfort":   90,
	"Essonne":                 91,
	"Hauts-de-Seine":          92,
	"Seine-Saint-Denis":       93,
	"Val-de-Marne":            94,
	"Val-d'Oise":              95,

	"Corse-du-Sud": DepartementCorseDuSud,
	"Haute-Corse":  DepartementHauteCorse,

	"Saint-Martin/Saint-Barthélemy":   DepartementSaintMartinSaintBarthélemy,
	"Français établis hors de France": DepartementFrançaisÉtablisHorsDeFrance,

	"Guadeloupe":               DepartementGuadeloupe,
	"Martinique":               DepartementMartinique,
	"Guyane":                   DepartementGuyane,
	"La Réunion":               DepartementLaRéunion,
	"Saint-Pierre-et-Miquelon": DepartementSaintPierreEtMiquelon,
	"Mayotte":                  DepartementMayotte,

	"Wallis et Futuna":    DepartementWallisEtFutuna,
	"Polynésie française": DepartementPolynésieFrançaise,
	"Nouvelle-Calédonie":  DepartementNouvelleCalédonie,
}
