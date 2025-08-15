package ue2024

import (
	"time"

	"github.com/HuguesGuilleus/site-vote/common"
)

const (
	voteID   = "2024_06_09_UE"
	voteName = "2024-06-09 Européenne"
)

var voteDate = time.Date(2024, time.June, 9, 0, 0, 0, 0, time.UTC)

const url = "https://static.data.gouv.fr/resources/resultats-des-elections-europeennes-du-9-juin-2024/20240613-154804/resultats-definitifs-par-bureau-de-vote.csv"

const header = `Code localisation;Libellé localisation;Code département;Libellé département;Code commune;Libellé commune;Code BV;Inscrits;Votants;% Votants;Abstentions;% Abstentions;Exprimés;% Exprimés/inscrits;% Exprimés/votants;Blancs;% Blancs/inscrits;% Blancs/votants;Nuls;% Nuls/inscrits;% Nuls/votants;Numéro de panneau 1;Nuance liste 1;Libellé abrégé de liste 1;Libellé de liste 1;Voix 1;% Voix/inscrits 1;% Voix/exprimés 1;Sièges 1;Numéro de panneau 2;Nuance liste 2;Libellé abrégé de liste 2;Libellé de liste 2;Voix 2;% Voix/inscrits 2;% Voix/exprimés 2;Sièges 2;Numéro de panneau 3;Nuance liste 3;Libellé abrégé de liste 3;Libellé de liste 3;Voix 3;% Voix/inscrits 3;% Voix/exprimés 3;Sièges 3;Numéro de panneau 4;Nuance liste 4;Libellé abrégé de liste 4;Libellé de liste 4;Voix 4;% Voix/inscrits 4;% Voix/exprimés 4;Sièges 4;Numéro de panneau 5;Nuance liste 5;Libellé abrégé de liste 5;Libellé de liste 5;Voix 5;% Voix/inscrits 5;% Voix/exprimés 5;Sièges 5;Numéro de panneau 6;Nuance liste 6;Libellé abrégé de liste 6;Libellé de liste 6;Voix 6;% Voix/inscrits 6;% Voix/exprimés 6;Sièges 6;Numéro de panneau 7;Nuance liste 7;Libellé abrégé de liste 7;Libellé de liste 7;Voix 7;% Voix/inscrits 7;% Voix/exprimés 7;Sièges 7;Numéro de panneau 8;Nuance liste 8;Libellé abrégé de liste 8;Libellé de liste 8;Voix 8;% Voix/inscrits 8;% Voix/exprimés 8;Sièges 8;Numéro de panneau 9;Nuance liste 9;Libellé abrégé de liste 9;Libellé de liste 9;Voix 9;% Voix/inscrits 9;% Voix/exprimés 9;Sièges 9;Numéro de panneau 10;Nuance liste 10;Libellé abrégé de liste 10;Libellé de liste 10;Voix 10;% Voix/inscrits 10;% Voix/exprimés 10;Sièges 10;Numéro de panneau 11;Nuance liste 11;Libellé abrégé de liste 11;Libellé de liste 11;Voix 11;% Voix/inscrits 11;% Voix/exprimés 11;Sièges 11;Numéro de panneau 12;Nuance liste 12;Libellé abrégé de liste 12;Libellé de liste 12;Voix 12;% Voix/inscrits 12;% Voix/exprimés 12;Sièges 12;Numéro de panneau 13;Nuance liste 13;Libellé abrégé de liste 13;Libellé de liste 13;Voix 13;% Voix/inscrits 13;% Voix/exprimés 13;Sièges 13;Numéro de panneau 14;Nuance liste 14;Libellé abrégé de liste 14;Libellé de liste 14;Voix 14;% Voix/inscrits 14;% Voix/exprimés 14;Sièges 14;Numéro de panneau 15;Nuance liste 15;Libellé abrégé de liste 15;Libellé de liste 15;Voix 15;% Voix/inscrits 15;% Voix/exprimés 15;Sièges 15;Numéro de panneau 16;Nuance liste 16;Libellé abrégé de liste 16;Libellé de liste 16;Voix 16;% Voix/inscrits 16;% Voix/exprimés 16;Sièges 16;Numéro de panneau 17;Nuance liste 17;Libellé abrégé de liste 17;Libellé de liste 17;Voix 17;% Voix/inscrits 17;% Voix/exprimés 17;Sièges 17;Numéro de panneau 18;Nuance liste 18;Libellé abrégé de liste 18;Libellé de liste 18;Voix 18;% Voix/inscrits 18;% Voix/exprimés 18;Sièges 18;Numéro de panneau 19;Nuance liste 19;Libellé abrégé de liste 19;Libellé de liste 19;Voix 19;% Voix/inscrits 19;% Voix/exprimés 19;Sièges 19;Numéro de panneau 20;Nuance liste 20;Libellé abrégé de liste 20;Libellé de liste 20;Voix 20;% Voix/inscrits 20;% Voix/exprimés 20;Sièges 20;Numéro de panneau 21;Nuance liste 21;Libellé abrégé de liste 21;Libellé de liste 21;Voix 21;% Voix/inscrits 21;% Voix/exprimés 21;Sièges 21;Numéro de panneau 22;Nuance liste 22;Libellé abrégé de liste 22;Libellé de liste 22;Voix 22;% Voix/inscrits 22;% Voix/exprimés 22;Sièges 22;Numéro de panneau 23;Nuance liste 23;Libellé abrégé de liste 23;Libellé de liste 23;Voix 23;% Voix/inscrits 23;% Voix/exprimés 23;Sièges 23;Numéro de panneau 24;Nuance liste 24;Libellé abrégé de liste 24;Libellé de liste 24;Voix 24;% Voix/inscrits 24;% Voix/exprimés 24;Sièges 24;Numéro de panneau 25;Nuance liste 25;Libellé abrégé de liste 25;Libellé de liste 25;Voix 25;% Voix/inscrits 25;% Voix/exprimés 25;Sièges 25;Numéro de panneau 26;Nuance liste 26;Libellé abrégé de liste 26;Libellé de liste 26;Voix 26;% Voix/inscrits 26;% Voix/exprimés 26;Sièges 26;Numéro de panneau 27;Nuance liste 27;Libellé abrégé de liste 27;Libellé de liste 27;Voix 27;% Voix/inscrits 27;% Voix/exprimés 27;Sièges 27;Numéro de panneau 28;Nuance liste 28;Libellé abrégé de liste 28;Libellé de liste 28;Voix 28;% Voix/inscrits 28;% Voix/exprimés 28;Sièges 28;Numéro de panneau 29;Nuance liste 29;Libellé abrégé de liste 29;Libellé de liste 29;Voix 29;% Voix/inscrits 29;% Voix/exprimés 29;Sièges 29;Numéro de panneau 30;Nuance liste 30;Libellé abrégé de liste 30;Libellé de liste 30;Voix 30;% Voix/inscrits 30;% Voix/exprimés 30;Sièges 30;Numéro de panneau 31;Nuance liste 31;Libellé abrégé de liste 31;Libellé de liste 31;Voix 31;% Voix/inscrits 31;% Voix/exprimés 31;Sièges 31;Numéro de panneau 32;Nuance liste 32;Libellé abrégé de liste 32;Libellé de liste 32;Voix 32;% Voix/inscrits 32;% Voix/exprimés 32;Sièges 32;Numéro de panneau 33;Nuance liste 33;Libellé abrégé de liste 33;Libellé de liste 33;Voix 33;% Voix/inscrits 33;% Voix/exprimés 33;Sièges 33;Numéro de panneau 34;Nuance liste 34;Libellé abrégé de liste 34;Libellé de liste 34;Voix 34;% Voix/inscrits 34;% Voix/exprimés 34;Sièges 34;Numéro de panneau 35;Nuance liste 35;Libellé abrégé de liste 35;Libellé de liste 35;Voix 35;% Voix/inscrits 35;% Voix/exprimés 35;Sièges 35;Numéro de panneau 36;Nuance liste 36;Libellé abrégé de liste 36;Libellé de liste 36;Voix 36;% Voix/inscrits 36;% Voix/exprimés 36;Sièges 36;Numéro de panneau 37;Nuance liste 37;Libellé abrégé de liste 37;Libellé de liste 37;Voix 37;% Voix/inscrits 37;% Voix/exprimés 37;Sièges 37;Numéro de panneau 38;Nuance liste 38;Libellé abrégé de liste 38;Libellé de liste 38;Voix 38;% Voix/inscrits 38;% Voix/exprimés 38;Sièges 38`

var constOptions = common.ConstOptionsList(
	"O	LDIV	HUMANITE SOUVERAINE (POUR UNE HUMANITE SOUVERAINE)",
	"O	LDIV	POUR UNE DEMOCRATIE REELLE : DECIDONS NOUS-MEMES!",
	"XR	LREC	LA FRANCE FIERE, MENEE PAR MARION MARECHAL ET SOUTENUE PAR ÉRIC ZEMMOUR",
	"L	LFI	LA FRANCE INSOUMISE - UNION POPULAIRE",
	"XR	LRN	La FRANCE REVIENT (LA FRANCE REVIENT! AVEC JORDAN BARDELLA ET MARINE LE PEN)",
	"L	LVEC	EUROPE ÉCOLOGIE",
	"O	LDIV	FREE PALESTINE",
	"O	LDIV	PARTI ANIMALISTE - LES ANIMAUX COMPTENT, VOTRE VOIX AUSSI",
	"XL	LEXG	PARTI REVOLUTIONNAIRE COMMUNISTES",
	"L	LDIV	PARTI PIRATE",
	"R	LENS	BESOIN D'EUROPE",
	"O	LDIV	PACE - PARTI DES CITOYENS EUROPÉENS, POUR L'ARMÉE EUROPÉENNE, POUR L'EUROPE SOCIALE , POUR LA PLANÈTE !",
	"O	LECO	ÉQUINOXE : ÉCOLOGIE PRATIQUE ET RENOUVEAU DÉMOCRATIQUE",
	"R	LECO	ECOLOGIE POSITIVE ET TERRITOIRES",
	"XR	LDIV	LISTE ASSELINEAU-FREXIT, POUR LE POUVOIR D'ACHAT ET POUR LA PAIX",
	"XL	LEXG	PAIX ET DECROISSANCE",
	"O	LDIV	POUR UNE AUTRE EUROPE",
	"R	LLR	LA DROITE POUR FAIRE ENTENDRE LA VOIX DE LA FRANCE EN EUROPE",
	"XL	LEXG	LUTTE OUVRIERE - LE CAMP DES TRAVAILLEURS",
	"L	LDVG	CHANGER L'EUROPE",
	"XR	LDIV	NLP (NOUS LE PEUPLE)",
	"XR	LEXG	URGENCE REVOLUTION! (POUR UN MONDE SANS FRONTIERES NI PATRONS, URGENCE REVOLUTION!)",
	"XL	LEXG	POUR LE PAIN, LA PAIX, LA LIBERTÉ! (PRÉSENTÉE PAR LE PARTI DES TRAVAILLEURS)",
	"XR	LEXD	L'EUROPE CA SUFFIT!",
	"R	LDIV	NON! PRENONS-NOUS EN MAINS",
	"XR	LEXD	FORTERESSE EUROPE - LISTE D'UNITE NATIONALISTE #nazis",
	"C	LUG	RÉVEILLER L'EUROPE (CIA)",
	"XL	LEXG	NON À L'UE ET À L'OTAN, COMMUNISTES POUR LA PAIX ET LE PROGRÈS SOCIAL",
	"R	LDVD	ALLIANCE RURALE",
	"R	LDIV	FRANCE LIBRE",
	"L	LDVG	EUROPE TERRITOIRES ÉCOLOGIE",
	"O	LDIV	LA RUCHE CITOYENNE",
	"L	LCOM	GAUCHE UNIE POUR LE MONDE DU TRAVAIL SOUTENUE PAR FABIEN ROUSSEL",
	"O	LDIV	DEFENDRE LES ENFANTS",
	"C	LECO	EAC (ÉCOLOGIE AU CENTRE)",
	"O	LDIV	DEMOCRATIE REPRESENTATIVE",
	"O	LDIV	ESPERANTO LANGUE COMMUNE",
	"O	LDIV	LIBERTÉ DÉMOCRATIQUE FRANÇAISE",
)
