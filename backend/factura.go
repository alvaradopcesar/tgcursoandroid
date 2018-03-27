// Cesar Alvarado Julio - 2016 @alvaradopcesar

package backend

import (
	"log"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	// "google.golang.org/appengine/log"

	// "errors"

	"github.com/gin-gonic/gin"
)

// Factura Structura de la tabla Ruc
type Factura struct {
	UID          *datastore.Key `json:"uid" datastore:"-"`
	Numero       string         `json:"numero"`
	Ruc          string         `json:"ruc"`
	Fecha        time.Time      `json:"fecha"`
	ImporteTotal float64        `json:"importetotal"`
	Estado       string         `json:"estado"` // ESTADO DEL CONTRIBUYENTE|
}

// PostFactura sdlfkj
func PostFactura(gc *gin.Context) {
	var factura Factura
	gc.Bind(&factura)
	if factura.Ruc != "" {
		// log.Println("PostFactura")
		// log.Println(factura.Numero)
		// log.Println(factura.Ruc)
		// log.Println(factura.Fecha)
		// log.Println(factura.ImporteTotal)
		// log.Println(factura.Estado)

		c := appengine.NewContext(gc.Request)
		k := datastore.NewIncompleteKey(c, "Factura", nil)
		t := &Factura{Numero: factura.Numero, Ruc: factura.Ruc, Fecha: factura.Fecha, ImporteTotal: factura.ImporteTotal, Estado: factura.Estado}
		k, err := datastore.Put(c, k, t)
		if err != nil {
			gc.JSON(422, gin.H{"error": err})
		}
		t.UID = k
		gc.JSON(200, gin.H{"status": "ok"})

	} else {
		gc.JSON(422, gin.H{"error": "fields are empty"})

	}
	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\"  }" http://localhost:8080/api/v1/users
}

// GetFacturaRuc  Obtener Rucs por codigo
func GetFacturaRuc(gc *gin.Context) {
	// var ruc Ruc
	// gc.Bind(&ruc)

	ruc := gc.Params.ByName("ruc")
	log.Println(ruc)
	var facturas []Factura

	c := appengine.NewContext(gc.Request)
	// Query
	q := datastore.NewQuery("Factura").Filter("Ruc =", ruc)

	_, err := q.GetAll(c, &facturas)
	if err != nil {
		gc.JSON(422, gin.H{"error": err})
	}
	if facturas == nil {
		gc.JSON(201, gin.H{"error": "Ruc no encontrado"})
	}

	gc.JSON(200, facturas)
}

// GetFacturaNumero sdkfjh
func GetFacturaNumero(gc *gin.Context) {
	// var ruc Ruc
	// gc.Bind(&ruc)

	numero := gc.Params.ByName("numero")
	log.Println(numero)
	var facturas []Factura

	c := appengine.NewContext(gc.Request)
	// Query
	q := datastore.NewQuery("Factura").Filter("Numero =", numero)

	_, err := q.GetAll(c, &facturas)
	if err != nil {
		gc.JSON(422, gin.H{"error": err})
	}
	if facturas == nil {
		gc.JSON(201, gin.H{"error": "Numero no encontrado"})
	}

	gc.JSON(200, facturas)
}

// PutFacturaNumero Actualizacion de Factura
func PutFacturaNumero(gc *gin.Context) {

	numeroParametro := gc.Params.ByName("numero")

	var factura Factura
	gc.Bind(&factura)
	if numeroParametro == "" {
		gc.JSON(422, gin.H{"error": "No ingreso Numero"})
	}

	log.Println(numeroParametro)
	log.Println(factura.Numero)

	var facturas []Factura

	c := appengine.NewContext(gc.Request)
	// Query
	q := datastore.NewQuery("Factura").Filter("Numero =", numeroParametro)

	key, err := q.GetAll(c, &facturas)
	if err != nil {
		gc.JSON(422, gin.H{"error": err})
	}
	if facturas == nil {
		gc.JSON(201, gin.H{"error": "Numero no encontrado"})
	}

	facturas[0].Ruc = factura.Ruc
	facturas[0].Fecha = factura.Fecha
	facturas[0].ImporteTotal = factura.ImporteTotal
	facturas[0].Estado = factura.Estado

	_, err = datastore.Put(c, key[0], &facturas[0])
	if err != nil {
		gc.JSON(422, gin.H{"error": err})
	}

	gc.JSON(200, facturas)
}
