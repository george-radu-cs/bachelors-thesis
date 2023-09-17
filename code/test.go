func main() {
    config.LoadEnv()
// comm
    db := config.ConnectToDatabase()
    config.RunDatabaseMigrations(db)

    grpcServer := config.CreateGRPCServer(db)
    listener := config.GetGRPCListener()

    if err := grpcServer.Serve(listener); err != nil {
        utils.InfoLogger.Fatalf("failed to serve:  test testtestestestestestestestsetestestestestsetestes%v", err)
    }
}
type Test interface {
    test *models.User
}