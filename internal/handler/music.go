package handler

import (
	"dynamodb-eg/internal/model"
	"dynamodb-eg/internal/response"
	"encoding/json"
	"fmt"
	"net/http"

	typesvc "dynamodb-eg/internal/types"
)

func (app *Application) Store() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req typesvc.StoreItemRequest

		// read the request body
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			fmt.Println("error decoding request body : ", err)
			response.BadRequest(w, "invalid request body")

			return
		}

		// check table exists
		err = app.model.Music.DescribeTable(ctx)
		if err != nil {
			err = app.model.Music.CreateTable(ctx)
			if err != nil {
				fmt.Println("error creating table : ", err)
				response.InternalServerError(w, "error while creating table")

				return
			}
		}

		musicItem := &model.Music{
			Artist:      req.Artist,
			SongTitle:   req.SongTitle,
			Description: req.Description,
			Views:       req.Views,
		}

		err = app.model.Music.Store(ctx, musicItem)
		if err != nil {
			fmt.Println("error storing new item : ", err)
			response.InternalServerError(w, "error while storing item in database")

			return
		}

		response.Created(w, "successfully stored item")
	}
}

func (app *Application) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		music, err := app.model.Music.Get(ctx)
		if err != nil {
			fmt.Println("error fetching item : ", err)
			response.InternalServerError(w, "error while fetching item from database")

			return
		}

		response.Success(w, music, "")
	}
}

func (app *Application) UpdateByName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req typesvc.StoreItemRequest

		// read the request body
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			fmt.Println("error decoding request body : ", err)
			response.BadRequest(w, "invalid request body")

			return
		}

		// prepare item to update
		item := model.Music{
			Artist:      req.Artist,
			SongTitle:   req.SongTitle,
			Description: req.Description,
			Views:       req.Views,
		}

		err = app.model.Music.UpdateByName(ctx, &item)
		if err != nil {
			fmt.Println("error updating item : ", err)
			response.InternalServerError(w, "error while updating item from database")

			return
		}

		response.Success(w, nil, "successfully updated item")
	}
}

func (app *Application) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req typesvc.StoreItemRequest

		// read the request body
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			fmt.Println("error decoding request body : ", err)
			response.BadRequest(w, "invalid request body")

			return
		}

		// prepare item to update
		item := model.Music{
			Artist:      req.Artist,
			SongTitle:   req.SongTitle,
			Description: req.Description,
			Views:       req.Views,
		}

		err = app.model.Music.Delete(ctx, &item)
		if err != nil {
			fmt.Println("error deleting item : ", err)
			response.InternalServerError(w, "error while deleting item from database")

			return
		}

		response.Success(w, nil, "successfully deleted item")
	}
}
