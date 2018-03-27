// Cesar Alvarado Julio - 2016 @alvaradopcesar

package backend

import (
	// "log"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"

	// "errors"

	"github.com/gin-gonic/gin"
)

// Cliente Structura de la tabla Ruc
type Cliente struct {
	UID               *datastore.Key `json:"uid" datastore:"-"`
	Ruc               string         `json:"ruc"`
	NombreRazonSocial string         `json:"nombrerazonsocial"`
	Estado            string         `json:"estado"` // ESTADO DEL CONTRIBUYENTE|
}

// PostCliente sdlfkj
func PostCliente(gc *gin.Context) {
	var cliente Cliente
	gc.Bind(&cliente)
	if cliente.Ruc != "" {
		// log.Println("PostRuc")
		// log.Println(ruc.Ruc)

		c := appengine.NewContext(gc.Request)
		k := datastore.NewIncompleteKey(c, "Cliente", nil)
		t := &Cliente{Ruc: cliente.Ruc, NombreRazonSocial: cliente.NombreRazonSocial, Estado: cliente.Estado}
		k, err := datastore.Put(c, k, t)
		if err != nil {
			gc.JSON(422, gin.H{"error": err})
			// log.Warningf(gc, format, args)
		}
		t.UID = k
		gc.JSON(200, gin.H{"status": "ok"})

	} else {
		gc.JSON(422, gin.H{"error": "fields are empty"})

	}
}

// GetClienteRuc Obtener Rucs por codigo
func GetClienteRuc(gc *gin.Context) {
	// var ruc Ruc
	// gc.Bind(&ruc)

	ruc := gc.Params.ByName("ruc")
	// log.Println(ruc)
	var clientes []Cliente

	c := appengine.NewContext(gc.Request)
	// Query
	q := datastore.NewQuery("Cliente").Filter("Ruc =", ruc)

	_, err := q.GetAll(c, &clientes)
	if err != nil {
		gc.JSON(422, gin.H{"error": err})
	}
	if clientes == nil {
		gc.JSON(201, gin.H{"error": "Ruc no encontrado"})
	}

	gc.JSON(200, clientes)
}

// GetClientesRuc ksdjfh
func GetClientesRuc(gc *gin.Context) {

	var clientes []Cliente

	c := appengine.NewContext(gc.Request)
	// Query
	q := datastore.NewQuery("Cliente")

	_, err := q.GetAll(c, &clientes)
	if err != nil {
		gc.JSON(422, gin.H{"error": err})
	}
	if clientes == nil {
		gc.JSON(201, gin.H{"error": "Ruc no encontrado"})
	}

	gc.JSON(200, clientes)
}

// PutClienteRuc Actualizacion de Ruc
func PutClienteRuc(gc *gin.Context) {

	rucParametro := gc.Params.ByName("ruc")

	var cliente Cliente
	gc.Bind(&cliente)
	if rucParametro == "" {
		gc.JSON(422, gin.H{"error": "No ingreso Ruc"})
	}

	// log.Println(rucParametro)
	// log.Println(cliente.NombreRazonSocial)
	var clientes []Cliente

	c := appengine.NewContext(gc.Request)
	// Query
	q := datastore.NewQuery("Cliente").Filter("Ruc =", rucParametro)

	key, err := q.GetAll(c, &clientes)
	if err != nil {
		gc.JSON(422, gin.H{"error": err})
	}
	if clientes == nil {
		gc.JSON(201, gin.H{"error": "Ruc no encontrado"})
	}

	clientes[0].NombreRazonSocial = cliente.NombreRazonSocial
	_, err = datastore.Put(c, key[0], &clientes[0])
	if err != nil {
		gc.JSON(422, gin.H{"error": err})
	}

	gc.JSON(200, cliente)
}
