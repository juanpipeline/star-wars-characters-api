package crud

import (
	"log"

	"github.com/juanpipeline/star-wars-characters-api/db"
	"github.com/juanpipeline/star-wars-characters-api/models"
)

func CreateCharacter(characterCreate models.CharacterCreate) models.Character {

	character := models.Character{CharacterBase: models.CharacterBase{Name: characterCreate.Name, Height: characterCreate.Height, Mass: characterCreate.Mass, HairColor: characterCreate.HairColor, SkinColor: characterCreate.SkinColor, Masters: characterCreate.Masters, Species: characterCreate.Species}}
	conn := db.GetConnection()
	sqlconn, err := conn.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlconn.Close()

	log.Println("Starting with character creation")

	result := conn.Create(&character)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println("This is the record inserted: ", character)

	return character

}

func SearchCharacters(name string) []models.Character {
	var characters []models.Character
	conditions := make(map[string]interface{})
	conn := db.GetConnection()
	sqlconn, err := conn.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlconn.Close()

	log.Println("Starting with character Searching")

	log.Println("The names is: ", name)

	if name != "" {
		log.Println("Filetering by name ", name)
		conditions["name"] = name
	}
	// Here you can write other conditions

	log.Println("Conditions: ", conditions)

	result := conn.Where(conditions).Find(&characters)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println("Records found: ", characters)

	return characters

}

func GetCharacter(id int64) (models.Character, error) {
	var character models.Character
	conn := db.GetConnection()
	sqlconn, err := conn.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlconn.Close()

	log.Println("Starting with character get by id")

	result := conn.First(&character, id)

	if result.Error != nil {
		log.Println("DB error ", result.Error)
		return character, result.Error
	}

	log.Println("Found record: ", character)

	return character, result.Error

}

func DeleteCharacter(id int64) (models.Character, error) {
	var character models.Character
	conn := db.GetConnection()
	sqlconn, err := conn.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlconn.Close()

	log.Println("Starting with character delete")

	result := conn.Delete(&character, id)

	if result.Error != nil {
		log.Println("DB error ", result.Error)
		return character, result.Error
	}

	log.Println("Item deleted: ", character)

	return character, result.Error

}

func UpdateCharacter(id int64, characterUpdate map[string]interface{}) (models.Character, error) {
	var character models.Character
	conn := db.GetConnection()
	sqlconn, err := conn.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlconn.Close()

	log.Println("Starting with character update")

	//result := conn.Delete(&character, id)

	if err := conn.First(&character, id).Error; err != nil {
		log.Println("DB error ", err)
		return character, err

	}

	// Compare input with character:
	if name := characterUpdate["Name"].(string); name == "" {
		characterUpdate["Name"] = character.Name
	}
	if height := characterUpdate["Height"].(float64); height == 0 {
		characterUpdate["Height"] = character.Height
	}
	if mass := characterUpdate["Mass"].(float64); mass == 0 {
		characterUpdate["Mass"] = character.Mass
	}
	if hair_color := characterUpdate["HairColor"].(string); hair_color == "" {
		characterUpdate["HairColor"] = character.HairColor
	}
	if skin_color := characterUpdate["SkinColor"].(string); skin_color == "" {
		characterUpdate["SkinColor"] = character.SkinColor
	}
	if masters := characterUpdate["Masters"].(string); masters == "" {
		characterUpdate["Masters"] = character.Masters
	}
	if species := characterUpdate["Species"].(string); species == "" {
		characterUpdate["Species"] = character.Species
	}

	result := conn.Model(&character).Updates(characterUpdate)
	//.Updates(characterUpdate)

	if result.Error != nil {
		log.Println("DB error ", result.Error)
		return character, result.Error
	}

	return character, result.Error

}
