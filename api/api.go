package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/juanpipeline/star-wars-characters-api/crud"
	"github.com/juanpipeline/star-wars-characters-api/models"
)

// SearchCharacters godoc
// @Summary Search Characters
// @Description Search all characteres in the DataBase
// @Accept  json
// @Produce  json
// @Param name query string false "Character name"
// @Success 200 {object} []models.CharacterResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /star-wars/characters [get]
func SearchCharacters(c *gin.Context) {
	name := c.Param("name")
	c.IndentedJSON(http.StatusOK, crud.SearchCharacters(name))
	//c.IndentedJSON(http.StatusOK, )
}

// CreateCharacter godoc
// @Summary Create Characters
// @Description create an Character
// @Accept  json
// @Produce  json
// @Param character body models.CharacterCreate true "Create account"
// @Success 200 {object} models.CharacterResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /star-wars/characters [post]
func CreateCharacter(c *gin.Context) {
	var characterCreate models.CharacterCreate

	// TODO: Json validation
	err := c.BindJSON(&characterCreate)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting with character creation: ", characterCreate)

	c.IndentedJSON(http.StatusOK, crud.CreateCharacter(characterCreate))

}

// GetCharacter godoc
// @Summary Get Character by id
// @Description Get an character by id
// @Accept  json
// @Produce  json
// @Param id path int true "Character ID"
// @Success 200 {object} []models.CharacterResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /star-wars/characters/{id} [get]
func GetCharacter(c *gin.Context) {
	_id := c.Param("id")

	id, err := strconv.ParseInt(_id, 0, 8)

	if err != nil {
		//log.Fatal(err)
		c.IndentedJSON(http.StatusBadRequest, models.APIError{ErrorCode: 401, ErrorMessage: "Bad request"})
	}

	character, apierr := crud.GetCharacter(id)

	if apierr != nil {
		//log.Fatal(err)
		c.IndentedJSON(http.StatusNotFound, models.APIError{ErrorCode: 404, ErrorMessage: "Not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, character)
	//c.IndentedJSON(http.StatusOK, )
}

// DeleteCharacter godoc
// @Summary Delete Character by id
// @Description Delete an character by id
// @Accept  json
// @Produce  json
// @Param id path int true "Character ID"
// @Success 204 {object} models.CharacterResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /star-wars/characters/{id} [delete]
func DeleteCharacter(c *gin.Context) {
	_id := c.Param("id")

	id, err := strconv.ParseInt(_id, 0, 8)

	if err != nil {
		//log.Fatal(err)
		c.IndentedJSON(http.StatusBadRequest, models.APIError{ErrorCode: 401, ErrorMessage: "Bad request"})
	}

	character, apierr := crud.DeleteCharacter(id)

	if apierr != nil {
		//log.Fatal(err)
		c.IndentedJSON(http.StatusNotFound, models.APIError{ErrorCode: 404, ErrorMessage: "Not found"})
		return
	}

	c.IndentedJSON(http.StatusNoContent, character)
	//c.IndentedJSON(http.StatusOK, )
}

// UpdateCharacter godoc
// @Summary Update Character by id
// @Description Update an character by id
// @Accept  json
// @Produce  json
// @Param id path int true "Character ID"
// @Param character body models.CharacterCreate true "Update account"
// @Success 204 {object} models.CharacterResponse
// @Failure 404 {object} models.APIError "Can not find objects"
// @Router /star-wars/characters/{id} [patch]
func UpdateCharacter(c *gin.Context) {
	var characterupdate models.CharacterBase
	_id := c.Param("id")

	id, err := strconv.ParseInt(_id, 0, 8)

	if err != nil {
		//log.Fatal(err)
		c.IndentedJSON(http.StatusBadRequest, models.APIError{ErrorCode: 401, ErrorMessage: "Bad request"})
	}

	// TODO: Json validation
	if err_bind := c.BindJSON(&characterupdate); err_bind != nil {
		//log.Fatal(err_bind)
		c.IndentedJSON(http.StatusBadRequest, models.APIError{ErrorCode: 401, ErrorMessage: err_bind.Error()})
		return
	}

	log.Println("Character update bound data: ", characterupdate)

	character, apierr := crud.UpdateCharacter(id, structs.Map(characterupdate))

	if apierr != nil {
		//log.Fatal(err)
		c.IndentedJSON(http.StatusNotFound, models.APIError{ErrorCode: 404, ErrorMessage: "Not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, character)
	//c.IndentedJSON(http.StatusOK, )
}
