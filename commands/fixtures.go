package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	fmt.Println("Dropping all records and repopulating...")

	client, ctx := initDb()
	defer client.Disconnect(ctx)
	database := client.Database("myeatopia")

	userCollection := database.Collection("user")
	userCollection.Drop(ctx)

	userCollection.InsertOne(ctx, bson.D{
		{Key: "userID", Value: 1},
		{Key: "name", Value: "Niels"},
		{Key: "isAdmin", Value: 1},
		{Key: "email", Value: "nielszwart@hotmail.com"},
		{Key: "password", Value: "niels"},
		{Key: "deleted", Value: 0},
		{Key: "created", Value: "2015-05-01 00:00:00"},
	})

	recipeCollection := database.Collection("recipe")
	recipeCollection.Drop(ctx)

	recipeCollection.InsertOne(ctx, bson.D{
		{Key: "RecipeID", Value: 1},
		{Key: "RecipeParentID", Value: nil},
		{Key: "UserID", Value: 1},
		{Key: "Title", Value: "Dahl"},
		{Key: "Description", Value: "Dit recept komt van Gale dat weer geïnspireerd is door de moeder van Cryson #ogrecipes"},
		{Key: "Instruction", Value: "<p>Was de linzen tot het water niet meer troebel is. Gooi in het kookwater een eetlepel kurkuma, een theelepel masallah en snuf zout. Kook de linzen op een rustig vuurtje. Schep het schuim wat ontstaat uit de pan. Als er te weinig water in de pan zit kun je er gewoon water bij doen.&nbsp;Hoe verser je doorkookt hoe verder de linzen uit elkaar vallen. Doen wat je fijn vind.&nbsp;</p><p>Ondertussen de tomaten, peterselie, knoflook, ui, gember en rode peper snijden.</p><p>Als de linzen op de helft van hun volledige kooktraject zitten vul je een koekepan met zonnebloemolie zodat de hele bodem vol ligt. Gooi in de koude olie mosterdzaad. Wacht tot het zaad gaat poffen. Voeg vervolgens de ui, knoflook, peterselie, rode peper, gember en komijnzaad&nbsp;toe aan de pan. Alles even goed heet laten worden. Daarna kun je het hele pannetje legen in de pan met linzen. Goed doorroeren.&nbsp;</p><p>Laat de dahl nog eventjes koken. Daarna tomaatjes erbij.&nbsp;</p><p>Hierna is het proeven en toevoegen wat je denkt dat nodig is: Zout/peper/massallah.</p>"},
		{Key: "Ingredients", Value: "<ul><li>Linzen</li><li>Tomaten</li><li>Rode peper</li><li>Peterselie</li><li>Uien</li><li>Knoflook</li><li>Gember</li><li>Kurkumapoeder</li><li>Komijnzaad</li><li>Mosterdzaad</li><li>Massalah</li><li>Zonnebloemolie</li></ul>"},
		{Key: "Persons", Value: 2},
		{Key: "Image", Value: "/files/recepten/151/dahl.jpeg"},
		{Key: "CookingTime", Value: "30 mins"},
		{Key: "Alias", Value: "dahl"},
		{Key: "IsInspiration", Value: 0},
		{Key: "Deleted", Value: 0},
		{Key: "Edited", Value: "2020-01-14 10:33:05"},
		{Key: "Created", Value: "2020-01-14 10:33:05"},
	})
	recipeCollection.InsertOne(ctx, bson.D{
		{Key: "RecipeID", Value: 2},
		{Key: "RecipeParentID", Value: nil},
		{Key: "UserID", Value: 1},
		{Key: "Title", Value: "Bolognese saus"},
		{Key: "Description", Value: "Deze saus gaat natuurlijk heel erg goed samen met zelfgemaakte pasta!"},
		{Key: "Instruction", Value: "<p>Was de linzen tot het water niet meer troebel is. Gooi in het kookwater een eetlepel kurkuma, een theelepel masallah en snuf zout. Kook de linzen op een rustig vuurtje. Schep het schuim wat ontstaat uit de pan. Als er te weinig water in de pan zit kun je er gewoon water bij doen.&nbsp;Hoe verser je doorkookt hoe verder de linzen uit elkaar vallen. Doen wat je fijn vind.&nbsp;</p><p>Ondertussen de tomaten, peterselie, knoflook, ui, gember en rode peper snijden.</p><p>Als de linzen op de helft van hun volledige kooktraject zitten vul je een koekepan met zonnebloemolie zodat de hele bodem vol ligt. Gooi in de koude olie mosterdzaad. Wacht tot het zaad gaat poffen. Voeg vervolgens de ui, knoflook, peterselie, rode peper, gember en komijnzaad&nbsp;toe aan de pan. Alles even goed heet laten worden. Daarna kun je het hele pannetje legen in de pan met linzen. Goed doorroeren.&nbsp;</p><p>Laat de dahl nog eventjes koken. Daarna tomaatjes erbij.&nbsp;</p><p>Hierna is het proeven en toevoegen wat je denkt dat nodig is: Zout/peper/massallah.</p>"},
		{Key: "Ingredients", Value: "<ul><li>Linzen</li><li>Tomaten</li><li>Rode peper</li><li>Peterselie</li><li>Uien</li><li>Knoflook</li><li>Gember</li><li>Kurkumapoeder</li><li>Komijnzaad</li><li>Mosterdzaad</li><li>Massalah</li><li>Zonnebloemolie</li></ul>"},
		{Key: "Persons", Value: 2},
		{Key: "Image", Value: "/files/recepten/151/dahl.jpeg"},
		{Key: "CookingTime", Value: "30 mins"},
		{Key: "Alias", Value: "bolognese-saus"},
		{Key: "IsInspiration", Value: 0},
		{Key: "Deleted", Value: 0},
		{Key: "Edited", Value: "2020-01-14 10:33:05"},
		{Key: "Created", Value: "2020-01-14 10:33:05"},
	})
	recipeCollection.InsertOne(ctx, bson.D{
		{Key: "RecipeID", Value: 3},
		{Key: "RecipeParentID", Value: nil},
		{Key: "UserID", Value: 1},
		{Key: "Title", Value: "Verse pesto"},
		{Key: "Description", Value: "Zelf pesto maken, als je het nog nooit gedaan hebt, denk je misschien dat het een hele klus is. Uh… nee hoor, pesto maken is ongelofelijk simpel, helemaal met dit pesto..."},
		{Key: "Instruction", Value: "<p>Was de linzen tot het water niet meer troebel is. Gooi in het kookwater een eetlepel kurkuma, een theelepel masallah en snuf zout. Kook de linzen op een rustig vuurtje. Schep het schuim wat ontstaat uit de pan. Als er te weinig water in de pan zit kun je er gewoon water bij doen.&nbsp;Hoe verser je doorkookt hoe verder de linzen uit elkaar vallen. Doen wat je fijn vind.&nbsp;</p><p>Ondertussen de tomaten, peterselie, knoflook, ui, gember en rode peper snijden.</p><p>Als de linzen op de helft van hun volledige kooktraject zitten vul je een koekepan met zonnebloemolie zodat de hele bodem vol ligt. Gooi in de koude olie mosterdzaad. Wacht tot het zaad gaat poffen. Voeg vervolgens de ui, knoflook, peterselie, rode peper, gember en komijnzaad&nbsp;toe aan de pan. Alles even goed heet laten worden. Daarna kun je het hele pannetje legen in de pan met linzen. Goed doorroeren.&nbsp;</p><p>Laat de dahl nog eventjes koken. Daarna tomaatjes erbij.&nbsp;</p><p>Hierna is het proeven en toevoegen wat je denkt dat nodig is: Zout/peper/massallah.</p>"},
		{Key: "Ingredients", Value: "<ul><li>Linzen</li><li>Tomaten</li><li>Rode peper</li><li>Peterselie</li><li>Uien</li><li>Knoflook</li><li>Gember</li><li>Kurkumapoeder</li><li>Komijnzaad</li><li>Mosterdzaad</li><li>Massalah</li><li>Zonnebloemolie</li></ul>"},
		{Key: "Persons", Value: 2},
		{Key: "Image", Value: "/files/recepten/151/dahl.jpeg"},
		{Key: "CookingTime", Value: "30 mins"},
		{Key: "Alias", Value: "verse-pesto"},
		{Key: "IsInspiration", Value: 0},
		{Key: "Deleted", Value: 0},
		{Key: "Edited", Value: "2020-01-14 10:33:05"},
		{Key: "Created", Value: "2020-01-14 10:33:05"},
	})

	fmt.Println("Done!")
}

func initDb() (*mongo.Client, context.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return client, ctx
}
