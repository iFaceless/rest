package rest

import (
	"net/http"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"context"
	"time"
)

var banner = `
               _
 _ __ ___  ___| |_
| '__/ _ \/ __| __|
| | |  __/\__ \ |_
|_|  \___||___/\__|
`

func Run(hd http.Handler, port int)  {
	server := http.Server{
		Addr:fmt.Sprintf(":%d", port),
		Handler:hd,
	}

	go func() {
		fmt.Println(banner)
		fmt.Printf("REST API Server started, listening at '%s'\n", server.Addr)
		if err := server.ListenAndServe(); err != nil {
			fmt.Printf("Exiting with error: %s\n", err)
		}
	}()

	kill := make(chan os.Signal)
	signal.Notify(kill, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	sig := <-kill
	fmt.Printf("Received signal '%s', shutdown gracefully\n", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 10 *time.Second)
	defer cancel()

	server.Shutdown(ctx)
	time.Sleep(100*time.Millisecond)
}