package main
 
import (
	"fmt"
	"log"
	"net/http"
)
 
const page = `<!DOCTYPE html>
<html lang="es">
<head>
<meta charset="UTF-8">
<title>Completa el for</title>
<style>
  body {
    font-family: -apple-system, Arial, sans-serif;
    background: #f2f2f2;
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    margin: 0;
  }
  .card {
    background: #fff;
    padding: 32px 40px;
    border-radius: 12px;
    box-shadow: 0 4px 12px rgba(0,0,0,0.15);
    text-align: center;
    min-width: 320px;
  }
  .code {
    font-family: 'Courier New', monospace;
    font-size: 20px;
    background: #1e1e1e;
    color: #d4d4d4;
    padding: 12px 16px;
    border-radius: 8px;
    margin-bottom: 20px;
  }
  .blank {
    color: #ffcc00;
    font-weight: bold;
  }
  .variable {
    font-size: 16px;
    color: #555;
    margin-bottom: 20px;
    min-height: 20px;
  }
  .buttons button {
    font-family: 'Courier New', monospace;
    font-size: 18px;
    padding: 8px 18px;
    margin: 0 6px;
    border: none;
    border-radius: 6px;
    background: #007acc;
    color: #fff;
    cursor: pointer;
  }
  .buttons button:hover {
    background: #005f99;
  }
</style>
</head>
<body>
  <div class="card">
    <div class="code">for i <span id="blank" class="blank">___</span> 0; i &lt; 10; i++</div>
    <div class="variable" id="variable">Elige un operador</div>
    <div class="buttons">
      <button onclick="elegir('=')">=</button>
      <button onclick="elegir('<')">&lt;</button>
      <button onclick="elegir(':=')">:=</button>
    </div>
  </div>
 
  <script>
    function elegir(valor) {
      document.getElementById('blank').textContent = valor;
      document.getElementById('variable').textContent = 'Elegiste: ' + valor;
    }
  </script>
</body>
</html>`
 
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, page)
}
 
func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}