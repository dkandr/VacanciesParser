package api

import (
	"context"
	"gitlab.com/dvkgroup/vacancies-parser/app/handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Run(c *handler.VacancyController) {
	port := ":8080"

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/public/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))).ServeHTTP(w, r)
	})

	r.Mount("/swagger", swaggerRouter())
	r.Mount("/vacancies", vacancyRouter(c))

	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	// Запуск веб-сервера в отдельном горутине
	go func() {
		log.Printf("server started on port %s\n", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Ожидание сигнала для начала завершения работы
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	// Установка тайм-аута для завершения работы
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
