package aggregator

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	servicemanager "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLServiceManager"

	"github.com/gorilla/mux"
)

func (agg *Aggregator) startRestServer(ctx context.Context) error {
	router := mux.NewRouter()
	router.HandleFunc("/state-root-updates", agg.handleGetStateRootUpdates).Methods("GET")

	err := http.ListenAndServe(agg.restServerIpPortAddr, router)
	if err != nil {
		agg.logger.Fatal("ListenAndServe", "err", err)
	}

	return nil
}

type GetStateRootUpdatesResponse struct {
	message     servicemanager.StateRootUpdateMessage
	aggregation MessageBlsAggregationServiceResponse
}

func (agg *Aggregator) handleGetStateRootUpdates(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	rollupId, err := strconv.ParseUint(params.Get("rollupId"), 10, 32)
	if err != nil {
		http.Error(w, "Invalid rollupId", http.StatusBadRequest)
		return
	}

	blockHeight, err := strconv.ParseUint(params.Get("blockHeight"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid blockHeight", http.StatusBadRequest)
		return
	}

	var message servicemanager.StateRootUpdateMessage
	var aggregation MessageBlsAggregationServiceResponse

	err = agg.fetchStateRootUpdate(uint32(rollupId), blockHeight, &message)
	if err != nil {
		http.Error(w, "StateRootUpdate not found", http.StatusNotFound)
		return
	}

	err = agg.fetchStateRootUpdateAggregation(uint32(rollupId), blockHeight, &aggregation)
	if err != nil {
		http.Error(w, "StateRootUpdate aggregation not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(GetStateRootUpdatesResponse{
		message:     message,
		aggregation: aggregation,
	})
}
