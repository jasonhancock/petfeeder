package main

import (
	"log"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/jasonlvhit/gocron"
	"github.com/matryer/temple"
	"github.com/tylerb/graceful"
)

type feedMsg struct {
	duration time.Duration
}

// Server is the web server.
type Server struct {
	config      *Config
	router      *mux.Router
	temple      *temple.Temple
	assetsDir   string
	assetsPath  string
	contextPath string
	feedChannel chan *feedMsg
	fdr         feeder
}

// NewServer creates a new Server.
func NewServer(c *Config, fdr feeder) (*Server, error) {
	tpl, err := temple.New(c.Templates)
	if err != nil {
		return nil, err
	}

	contextPath := "/"
	var assetsPath string

	if contextPath == "/" {
		assetsPath = "/assets"
	} else {
		assetsPath = path.Join(contextPath, "assets")
	}

	s := &Server{
		config:      c,
		router:      mux.NewRouter(),
		temple:      tpl,
		assetsDir:   c.Assets,
		assetsPath:  assetsPath,
		contextPath: strings.TrimRight(contextPath, "/"),
		fdr:         fdr,
	}

	s.buildRoutes()
	s.feedChannel = make(chan *feedMsg)
	return s, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) Run(addr string) {

	go s.handleFeed()
	go s.scheduleFeeds()

	http.Handle("/", s)
	srv := &graceful.Server{
		Timeout: 2 * time.Second,
		Server:  &http.Server{Addr: addr, Handler: s},
	}

	log.Println("Listening for connections on ", addr)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Server) renderTemplate(w http.ResponseWriter, r *http.Request, name string) {
	tpl, ok := s.temple.GetOK(name)
	if !ok {
		http.NotFound(w, r)
		return
	}
	tpl.Execute(w, s.PageInfo(r))
}

// TemplateHandler gets an http.Handler that will render the specified
// template.
func (s *Server) TemplateHandler(name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.renderTemplate(w, r, name)
	})
}

// PageInfo is the object that contains data that drives
// the templates.
type PageInfo struct {
	AssetsPath  string
	ContextPath string
}

// PageInfo gets page information.
func (s *Server) PageInfo(r *http.Request) PageInfo {
	return PageInfo{
		AssetsPath:  s.assetsPath,
		ContextPath: s.contextPath,
	}
}

func (s *Server) path(route string) string {
	if s.contextPath == "" {
		return route
	}

	return path.Join(s.contextPath, route)
}

func (s *Server) feedHandler(w http.ResponseWriter, r *http.Request) {
	s.performFeed(s.config.DefaultDuration)
	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleFeed() {
	for msg := range s.feedChannel {
		s.fdr.feed(msg.duration)
	}
}

func (s *Server) performFeed(duration time.Duration) {
	s.feedChannel <- &feedMsg{
		duration: duration,
	}
}

func (s *Server) scheduleFeeds() {
	for _, v := range s.config.Schedule {
		log.Printf("Scheduling a job at %s for %s\n", v.Time, v.Duration)
		gocron.Every(1).Day().At(v.Time).Do(s.performFeed, v.Duration)
	}

	<-gocron.Start()
}
