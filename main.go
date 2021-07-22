package main
////////////////-------------\\\\\\\\\\\\\\
import(
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)
////////////////-------------\\\\\\\\\\\\\\
func error(file){
     return {
     "./hata/", file, ".json"
      }
}
////////////////-------------\\\\\\\\\\\\\\
func main() {
	fmt.Print("Proje Başlatıldı\n")
	e := echo.New()
	e.GET("/test", func(c echo.Context) error{
		return c.String(http.StatusOK, "Slm Dünya")
	})
	e.GET("/", func(c echo.Context) error{
		return c.File("index.json")
	})
	e.GET("/owofy", func(c echo.Context) error{
		if c.QueryParam("message") == "" {
			return c.File(error("owofy"))
		}
		owo := strings.Replace(strings.Replace(c.QueryParam("message"), "l", "w", -1), "r", "w", -1)
		return c.JSON(http.StatusOK, owo)
	})
	e.Logger.Fatal(e.Start(":3000"))
}
////////////////-------------\\\\\\\\\\\\\\
