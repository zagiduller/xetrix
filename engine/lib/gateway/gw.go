package gateway

import (
	"crypto/tls"
	"engine/lib/helper"
	"engine/lib/payments"
	"engine/lib/services/events"
	"engine/lib/structs"
	"engine/lib/websocket"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/olebedev/config"
	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"net/http"
	"os"
	"os/exec"
	r2 "runtime"
	"strings"
)

var srvc_auth structs.ServiceAuthClient

func Run(grpcConn *grpc.ClientConn, c *config.Config, bus *events.Bus, payreg *payments.Registry) {

	srvc_auth = structs.NewServiceAuthClient(grpcConn)

	serveMux := runtime.NewServeMux(runtime.WithMetadata(func(ctx context.Context, r *http.Request) metadata.MD {
		md := metadata.Pairs()
		if tok := r.Context().Value("token"); tok != nil {
			md.Set("token", tok.(string))
		}
		if pid := r.Context().Value("pid"); pid != nil {
			log.Printf("Participant id: %s", pid)
			md.Set("pid", pid.(string))
		}
		if aid := r.Context().Value("admin-id"); aid != nil {
			log.Printf("Adimn id: %s", aid)
			md.Set("admin-id", aid.(string))
		}
		return md
	}))

	ctx := context.Background()

	structs.RegisterServiceAuthHandler(ctx, serveMux, grpcConn)
	structs.RegisterServiceCurrencyHandler(ctx, serveMux, grpcConn)
	structs.RegisterServiceUserHandler(ctx, serveMux, grpcConn)

	//structs.RegisterServiceAccountHandler(ctx, serveMux, grpcConn)
	//structs.RegisterServiceAccountBalanceHandler(ctx, serveMux, grpcConn)
	structs.RegisterServiceOrderHandler(ctx, serveMux, grpcConn)
	//structs.RegisterServiceCommissionHandler(ctx, serveMux, grpcConn)
	structs.RegisterServiceTransactionHandler(ctx, serveMux, grpcConn)
	structs.RegisterServiceTransactionProcessingHandler(ctx, serveMux, grpcConn)

	mux := http.NewServeMux()

	// Register pprof handlers
	//mux.HandleFunc("/debug/pprof/", pprof.Index)
	//mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	//mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	//mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	//
	//mux.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	//mux.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	//mux.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	//mux.Handle("/debug/pprof/block", pprof.Handler("block"))

	mux.Handle("/_v1-a/", AdminAuthMiddleware(serveMux))
	mux.Handle("/_v1/", AuthMiddleware(serveMux))
	mux.Handle("/v1/", serveMux)

	//Регистрация хуков платежных систем
	if payreg.HooksMux != nil {
		mux.Handle("/payments/", http.StripPrefix("/payments", payreg.HooksMux))
	}

	// Init websocket
	websocket.InitWS(grpcConn, mux, bus)

	frontFS := http.FileServer(http.Dir(c.UString("rest_gateway.views.front")))
	viewsFS := http.FileServer(http.Dir(c.UString("rest_gateway.views.admin")))

	stripPrefix := c.UString("rest_gateway.views.stripfront")
	//stripAdmin := c.UString("rest_gateway.views.stripadmin")

	mux.Handle(stripPrefix, http.StripPrefix(stripPrefix, frontFS))
	mux.Handle("/admin/", http.StripPrefix("/admin/", viewsFS))
	mux.Handle("/", frontFS)

	tp := c.UString("rest_gateway.type")
	port := c.UString("rest_gateway.port")
	log.Printf("REST gateway Host: %s type: %s\n", port, tp)
	if tp == "https" {
		certDir := c.UString("rest_gateway.certsdir")

		domains := strings.Split(c.UString("rest_gateway.domains"), ",")

		certManager := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(domains...), //Your domain here
			Cache:      autocert.DirCache(certDir),         //Folder for storing certificates
		}

		server := &http.Server{
			Addr:    ":https",
			Handler: handlers.LoggingHandler(os.Stdout, mux),
			TLSConfig: &tls.Config{
				GetCertificate: certManager.GetCertificate,
			},
		}

		go http.ListenAndServe(":http", certManager.HTTPHandler(nil))

		log.Fatal(server.ListenAndServeTLS("", "")) //Key and cert are coming from Let's Encrypt
	} else {
		openbrowser("http://localhost:" + port)
		log.Fatalln(http.ListenAndServe(":"+port, handlers.LoggingHandler(os.Stdout, mux)))
	}

}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := helper.GetTokenStringFromReq(r)
		if len(tokenString) > 20 {
			user, err := srvc_auth.DecodeSession(context.Background(), &structs.Session{Token: tokenString})
			if err != nil {
				log.Printf("srvc_auth: %s", err)
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte("Token error "))
				return
			}
			log.Printf("Token validation ok. UserId %s", user.Id)
			ctx2 := context.WithValue(context.WithValue(r.Context(), "pid", user.Id), "token", tokenString)
			next.ServeHTTP(w, r.WithContext(ctx2))
			return

		} else {
			md, ok := metadata.FromIncomingContext(r.Context())
			fmt.Printf("\n%#v %v\n", md, ok)
			if md, ok := metadata.FromIncomingContext(r.Context()); ok {
				if len(md["pid"]) > 0 {
					log.Printf("AuthMiddleware: pid set %s ", md["pid"][0])
				}
			}
		}

		log.Println("Token extracting error")
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Token error "))
		return
	})
}

func AdminAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := helper.GetTokenStringFromReq(r)
		if len(tokenString) > 20 {
			user, err := srvc_auth.DecodeSession(context.Background(), &structs.Session{Token: tokenString})
			if err != nil {
				log.Printf("srvc_auth: %s", err)
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte("Token error "))
				return
			}

			if user.Status != structs.UserStatus_ADMINISTRATOR {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte("Operation Not Allowed"))
				return
			}

			log.Printf("AdminAuthMiddleware: Admin (%s) %s", user.Name, user.Id)

			ctx2 := context.WithValue(context.WithValue(r.Context(), "admin-id", user.Id), "token", tokenString)
			next.ServeHTTP(w, r.WithContext(ctx2))
			return

		}

		log.Println("Token extracting error")
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Token error "))
		return
	})
}

func openbrowser(url string) {
	var err error

	switch r2.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Println(err)
	}

}
