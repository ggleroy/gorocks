package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/qor5/admin/presets"
	"github.com/qor5/web"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type BlogPost struct {
	ID      uint `gorm:"primaryKey"`
	Title   string
	Content string
}

var db *gorm.DB

// Custom DataOperator implementation for QOR5
type GormDataOperator struct {
	db *gorm.DB
}

// Create a new record
func (g *GormDataOperator) Create(model interface{}, ctx *web.EventContext) error {
	return g.db.Create(model).Error
}

// Save method (handles both create & update)
func (g *GormDataOperator) Save(model interface{}, id string, ctx *web.EventContext) error {
	if id == "" {
		return g.db.Create(model).Error
	}
	return g.db.Save(model).Error
}

// ✅ Corrected Delete method
func (g *GormDataOperator) Delete(model interface{}, id string, ctx *web.EventContext) error {
	err := g.db.First(model, id).Error
	if err != nil {
		return err
	}
	return g.db.Delete(model).Error
}

// ✅ Corrected Fetch method
func (g *GormDataOperator) Fetch(model interface{}, id string, ctx *web.EventContext) (interface{}, error) {
	instance := model
	err := g.db.First(instance, id).Error
	if err != nil {
		return nil, err
	}
	return instance, nil
}

// Query all records
func (g *GormDataOperator) Query(model interface{}, ctx *web.EventContext) error {
	return g.db.Find(model).Error
}

// Search method for QOR5 DataOperator
func (g *GormDataOperator) Search(model interface{}, params *presets.SearchParams, ctx *web.EventContext) (interface{}, int, error) {
	var results []BlogPost
	query := g.db.Model(&BlogPost{})

	// Apply search filter if a keyword is provided
	if params.Keyword != "" {
		query = query.Where("title LIKE ?", "%"+params.Keyword+"%")
	}

	// Count total matching records
	var totalCount int64
	query.Count(&totalCount)

	// Fetch results
	err := query.Find(&results).Error
	if err != nil {
		return nil, 0, err
	}

	return results, int(totalCount), nil
}

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&BlogPost{})

	// Wrap GORM with a custom DataOperator
	dataOperator := &GormDataOperator{db: db}

	// Initialize QOR5 admin panel
	presetsBuilder := presets.New()
	admin := presetsBuilder.URIPrefix("/admin").DataOperator(dataOperator)
	admin.Model(&BlogPost{})

	// Setup HTTP handlers
	mux := http.NewServeMux()
	mux.Handle("/admin/", admin)
	mux.HandleFunc("/", blogHomepage)

	// Start the web server
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}

func blogHomepage(w http.ResponseWriter, r *http.Request) {
	var posts []BlogPost
	db.Find(&posts)

	tmpl := `
	<html>
	<head><title>Blog da Loira</title></head>
	<body>
		<h1>Feito por Guilherme Leroy</h1>
		{{range .}}
			<h2>{{.Title}}</h2>
			<p>{{.Content}}</p>
			<hr>
		{{end}}
	</body>
	</html>
	`

	t, _ := template.New("blog").Parse(tmpl)
	t.Execute(w, posts)
}
