package handler

// PingPongCommand is a view on CLI.
// It shows results of `PingPongModel`.
type PingPongCommand interface {
	// RunSyncService lets Nurupoga wait updating messages to send suited reply
	RunSyncService()
}
