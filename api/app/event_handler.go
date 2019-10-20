package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"

	"github.com/mergeforces/mergeforces-service/pkg/models"
	"github.com/mergeforces/mergeforces-service/pkg/repository"
)

func (app *App) HandleListEvents(w http.ResponseWriter, r *http.Request) {
	events, err := repository.ListEvents(app.db)
	if err != nil {
		app.logger.Warn().Err(err).Msg("")

		w.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(w, `{"error": "%v"}`, appErrDataAccessFailure)
		if err != nil {
			app.logger.Warn().Err(err).Msg("Failed to write error response")
		}
		return
	}

	if events == nil {
		_, err := fmt.Fprint(w, "[]")
		if err != nil {
			app.logger.Warn().Err(err).Msg("Failed to write error response")
		}
		return
	}

	dtos := events.ToDto()
	if err := json.NewEncoder(w).Encode(dtos); err != nil {
		app.logger.Warn().Err(err).Msg("")

		w.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(w, `{"error": "%v"}`, appErrJsonCreationFailure)
		if err != nil {
			app.logger.Warn().Err(err).Msg("Failed to write error response")
		}
		return
	}
}

func (app *App) HandleCreateEvent(w http.ResponseWriter, r *http.Request) {
	form := &models.EventForm{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		app.logger.Warn().Err(err).Msg("")

		w.WriteHeader(http.StatusUnprocessableEntity)
		_, err := fmt.Fprintf(w, `{"error": "%v"}`, appErrFormDecodingFailure)
		if err != nil {
			app.logger.Warn().Err(err).Msg("Failed to write error response")
		}
		return
	}

	if err := app.validator.Struct(form); err != nil {
		app.logger.Warn().Err(err).Msg("")

		w.WriteHeader(http.StatusUnprocessableEntity)
		_, err := fmt.Fprintf(w, `{"error": "%v"}`, err.Error())
		if err != nil {
			app.logger.Warn().Err(err).Msg("Failed to write error response")
		}
		return
	}

	eventModel, err := form.ToModel()
	if err != nil {
		app.logger.Warn().Err(err).Msg("")

		w.WriteHeader(http.StatusUnprocessableEntity)
		_, err := fmt.Fprintf(w, `{"error": "%v"}`, appErrFormDecodingFailure)
		if err != nil {
			app.logger.Warn().Err(err).Msg("Failed to write error response")
		}

		return
	}

	event, err := repository.CreateEvent(app.db, eventModel)
	if err != nil {
		app.logger.Warn().Err(err).Msg("")

		w.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(w, `{"error": "%v"}`, appErrDataCreationFailure)
		if err != nil {
			app.logger.Warn().Err(err).Msg("Failed to write error response")
		}
		return
	}

	app.logger.Info().Msgf("New event created: %d", event.ID)
}

func (app *App) HandleReadEvent(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 0, 64)
	if err != nil || id == 0 {
		app.logger.Info().Msgf("can not parse ID: %v", id)

		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	event, err := repository.ReadEvent(app.db, uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		app.logger.Warn().Err(err).Msg("")

		w.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(w, `{"error": "%v"}`, appErrDataAccessFailure)
		if err != nil {
			app.logger.Warn().Err(err).Msg("Failed to write error response")
		}
		return
	}

	dto := event.ToDto()
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		app.logger.Warn().Err(err).Msg("")

		w.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(w, `{"error": "%v"}`, appErrJsonCreationFailure)
		if err != nil {
			app.logger.Warn().Err(err).Msg("Failed to write error response")
		}
		return
	}
}

func (app *App) HandleUpdateEvent(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 0, 64)
	if err != nil || id == 0 {
		app.logger.Info().Msgf("can not parse ID: %v", id)

		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	form := &models.EventForm{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		app.logger.Warn().Err(err).Msg("")

		w.WriteHeader(http.StatusUnprocessableEntity)
		_, err := fmt.Fprintf(w, `{"error": "%v"}`, appErrFormDecodingFailure)
		if err != nil {
			app.logger.Warn().Err(err).Msg("Failed to write error response")
		}
		return
	}

	if err := app.validator.Struct(form); err != nil {
		app.logger.Warn().Err(err).Msg("")

		w.WriteHeader(http.StatusUnprocessableEntity)
		_, err := fmt.Fprintf(w, `{"error": "%v"}`, err.Error())
		if err != nil {
			app.logger.Warn().Err(err).Msg("Failed to write error response")
		}
		return
	}

	eventModel, err := form.ToModel()
	if err != nil {
		app.logger.Warn().Err(err).Msg("")

		w.WriteHeader(http.StatusUnprocessableEntity)
		_, err := fmt.Fprintf(w, `{"error": "%v"}`, appErrFormDecodingFailure)
		if err != nil {
			app.logger.Warn().Err(err).Msg("Failed to write error response")
		}
		return
	}

	eventModel.ID = uint(id)
	if err := repository.UpdateEvent(app.db, eventModel); err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		app.logger.Warn().Err(err).Msg("")

		w.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(w, `{"error": "%v"}`, appErrDataUpdateFailure)
		if err != nil {
			app.logger.Warn().Err(err).Msg("Failed to write error response")
		}
		return
	}

	app.logger.Info().Msgf("Event updated: %d", id)
	w.WriteHeader(http.StatusAccepted)
}

func (app *App) HandleDeleteEvent(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 0, 64)
	if err != nil || id == 0 {
		app.logger.Info().Msgf("can not parse ID: %v", id)

		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if err := repository.DeleteEvent(app.db, uint(id)); err != nil {
		app.logger.Warn().Err(err).Msg("")

		w.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(w, `{"error": "%v"}`, appErrDataAccessFailure)
		if err != nil {
			app.logger.Warn().Err(err).Msg("Failed to write error response")
		}
		return
	}

	app.logger.Info().Msgf("Event deleted: %d", id)
}
