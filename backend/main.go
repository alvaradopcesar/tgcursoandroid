// Cesar Alvarado Julio - 2016 @alvaradopcesar

package backend

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// This function's name is a must. App Engine uses it to drive the requests properly.
func init() {
	// Starts a new Gin instance with no middle-ware
	r := gin.New()

	// Define your handlers
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.POST("/cliente/", PostCliente)
	r.GET("/cliente/:ruc", GetClienteRuc)
	r.GET("/clientes/", GetClientesRuc)
	r.PUT("/cliente/:ruc", PutClienteRuc)

	r.POST("/factura/", PostFactura)
	r.GET("/facturaruc/:ruc", GetFacturaRuc)
	r.GET("/factura/:numero", GetFacturaNumero)
	r.PUT("/factura/:numero", PutFacturaNumero)
	// Handle all requests using net/http
	http.Handle("/", r)
}
