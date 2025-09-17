package api

import (
    "net/http"

    "github.com/yourname/starcraft-concurrency-war/internal/domain"
    "github.com/yourname/starcraft-concurrency-war/internal/logging"
    "github.com/yourname/starcraft-concurrency-war/internal/services"
)

// Handlers groups HTTP handlers.
type Handlers struct {
    // Fields to include:
    // - svc: *services.GameService
    // - logger: logging.Logger
}

// NewHandlers constructs handlers.
// Steps to implement:
// - Build GameService with dependencies in router.
func NewHandlers(d Dependencies) *Handlers {
    // Implement construction here.
    // Steps:
    // - Build a GameService via services.NewGameService.
    // - Return &Handlers with svc and logger.
    return nil
}

// Players handles POST /players to create a new player.
// Request: { "name": "..." }
// Response: 201 { "id": "...", "name": "..." }
func (h *Handlers) Players(w http.ResponseWriter, r *http.Request) {
    // Implement POST /players here.
    // Packages to use: encoding/json for decoding request and encoding response.
    // Steps:
    // - Ensure method is POST; else 405.
    // - Decode JSON body: { "name": "..." }.
    // - context.WithTimeout for a short timeout (e.g., 3s) when calling service.
    // - Call h.svc.CreatePlayer(ctx, name).
    // - On success, write 201 with JSON player.
    // - On error, write appropriate status (400 for validation, 500 otherwise).
}

// Squads handles POST /squads to create a new squad for a player.
// Request: { "player_id": "...", "name": "...", "position": {"x":0,"y":0} }
// Response: 201 { squad }
func (h *Handlers) Squads(w http.ResponseWriter, r *http.Request) {
    // Implement POST /squads here.
    // Request schema: { "player_id": "...", "name": "...", "position": {"x":0,"y":0} }
    // Steps:
    // - Validate method.
    // - Decode JSON into a struct.
    // - Call h.svc.CreateSquad(ctx, domain.ID(playerID), name, position).
    // - Respond with 201 and the created squad as JSON.
}

// Commands handles POST /commands to issue commands for a squad.
// Request: { "player_id":"...", "squad_id":"...", "type":"MOVE|ATTACK|SCOUT", "target_pos":..., "target_squad":"..." }
func (h *Handlers) Commands(w http.ResponseWriter, r *http.Request) {
    // Implement POST /commands here.
    // Request: { player_id, squad_id, type, target_pos?, target_squad? }.
    // Steps:
    // - Validate method POST.
    // - Decode request JSON to a struct.
    // - Build domain.Command (convert strings to domain.IDs) and call h.svc.IssueCommand.
    // - Respond 202 Accepted on success.
}

// World handles GET /world to return a snapshot of the current state.
func (h *Handlers) World(w http.ResponseWriter, r *http.Request) {
    // Implement GET /world here.
    // Steps:
    // - Ensure method GET.
    // - Call h.svc.GetWorld(r.Context()).
    // - Respond with JSON snapshot.
}

// SquadByID handles GET /squads/{id}
// Minimal path parsing without a router lib; consider chi for production.
func (h *Handlers) SquadByID(w http.ResponseWriter, r *http.Request) {
    // Implement GET /squads/{id} here.
    // Steps (ServeMux):
    // - Parse ID by trimming prefix "/squads/" and splitting by "/".
    // - Validate ID is present; on missing, return 400.
    // - With a short timeout context, call h.svc.GetSquad(ctx, domain.ID(id)).
    // - If storage.ErrNotFound, return 404; else 500 for other errors.
    // - Respond with JSON for the squad.
}

// Note: Use storage.ErrNotFound in error handling (import storage in this file when implementing).
