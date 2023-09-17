// LoadTLSCredentials returns tls credentials for server
// in case of error log fatal to trigger panic and stop the api since we want it to run with tls
func LoadTLSCredentials() (tlsCreds credentials.TransportCredentials) {
	serverCert, err := tls.LoadX509KeyPair(os.Getenv("SERVER_CERT_FILE"), os.Getenv("SERVER_PRIVATE_KEY_FILE"))
	if err != nil {
		utils.ErrorLogger.Fatalf("Error loading tls x509 key pair for server: %w", err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(tlsConfig)
}

func CreateGRPCServer(db *gorm.DB) *grpc.Server {
	// ...
	grpcServerOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(authInterceptor.Authorize()),
	}

	if os.Getenv("IN_PRODUCTION") == "true" {
		grpcServerOptions = append(grpcServerOptions, grpc.Creds(LoadTLSCredentials()))
	}

	grpcServer := grpc.NewServer(grpcServerOptions...)
	// ..
}