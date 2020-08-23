package usecases

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"

	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/goblin"
)

// A PostInteractor is an interactor for a post.
type PostInteractor struct {
	AdminRepository AdminRepository
	PostRepository  PostRepository
	JWTRepository   JWTRepository
	JSONResponse    JSONResponse
	Logger          Logger
}

// HandleIndex returns a listing of the resource.
func (pi *PostInteractor) HandleIndex(w http.ResponseWriter, r *http.Request) {
	const defaultPage = 1
	const defaultLimit = 10

	count, err := pi.PostRepository.CountAllPublish()
	if err != nil {
		pi.Logger.Error(err.Error())
		pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	paramPage := r.URL.Query().Get("page")
	var page int
	if paramPage == "" {
		page = defaultPage
	} else {
		page, err = strconv.Atoi(paramPage)
		if err != nil {
			pi.Logger.Error(err.Error())
			pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
			return
		}
	}

	paramLimit := r.URL.Query().Get("limit")
	var limit int
	if paramLimit == "" {
		limit = defaultLimit
	} else {
		limit, err = strconv.Atoi(paramLimit)
		if err != nil {
			pi.Logger.Error(err.Error())
			pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
			return
		}
	}

	var posts domain.Posts
	posts, err = pi.PostRepository.FindAllPublish(page, limit)
	if err != nil {
		pi.Logger.Error(err.Error())
		pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var pr PostResponse
	code, msg, err := pr.MakeResponseHandleIndex(posts)
	if err != nil {
		pi.Logger.Error(err.Error())
		pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	pageCount := math.Ceil(float64(count) / float64(limit))

	w.Header().Set("Pagination-Count", strconv.Itoa(count))
	w.Header().Set("Pagination-Pagecount", fmt.Sprint(pageCount))
	w.Header().Set("Pagination-Page", strconv.Itoa(page))
	w.Header().Set("Pagination-Limit", strconv.Itoa(limit))
	pi.JSONResponse.HTTPStatus(w, code, msg)
	return
}

// HandleIndexByCategory returns a listing of the resource.
func (pi *PostInteractor) HandleIndexByCategory(w http.ResponseWriter, r *http.Request) {
	const defaultPage = 1
	const defaultLimit = 10

	count, err := pi.PostRepository.CountAllPublish()
	if err != nil {
		pi.Logger.Error(err.Error())
		pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	paramPage := r.URL.Query().Get("page")
	var page int
	if paramPage == "" {
		page = defaultPage
	} else {
		page, err = strconv.Atoi(paramPage)
		if err != nil {
			pi.Logger.Error(err.Error())
			pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
			return
		}
	}

	paramLimit := r.URL.Query().Get("limit")
	var limit int
	if paramLimit == "" {
		limit = defaultLimit
	} else {
		limit, err = strconv.Atoi(paramLimit)
		if err != nil {
			pi.Logger.Error(err.Error())
			pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
			return
		}
	}

	name := goblin.GetParam(r.Context(), "name")

	var posts domain.Posts
	posts, err = pi.PostRepository.FindAllPublishByCategory(page, limit, name)
	if err != nil {
		pi.Logger.Error(err.Error())
		pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var pr PostResponse
	code, msg, err := pr.MakeResponseHandleIndex(posts)
	if err != nil {
		pi.Logger.Error(err.Error())
		pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	pageCount := math.Ceil(float64(count) / float64(limit))

	w.Header().Set("Pagination-Count", strconv.Itoa(count))
	w.Header().Set("Pagination-Pagecount", fmt.Sprint(pageCount))
	w.Header().Set("Pagination-Page", strconv.Itoa(page))
	w.Header().Set("Pagination-Limit", strconv.Itoa(limit))
	pi.JSONResponse.HTTPStatus(w, code, msg)
	return
}

// HandleIndexByTag returns a listing of the resource.
func (pi *PostInteractor) HandleIndexByTag(w http.ResponseWriter, r *http.Request) {
	const defaultPage = 1
	const defaultLimit = 10

	count, err := pi.PostRepository.CountAllPublish()
	if err != nil {
		pi.Logger.Error(err.Error())
		pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	paramPage := r.URL.Query().Get("page")
	var page int
	if paramPage == "" {
		page = defaultPage
	} else {
		page, err = strconv.Atoi(paramPage)
		if err != nil {
			pi.Logger.Error(err.Error())
			pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
			return
		}
	}

	paramLimit := r.URL.Query().Get("limit")
	var limit int
	if paramLimit == "" {
		limit = defaultLimit
	} else {
		limit, err = strconv.Atoi(paramLimit)
		if err != nil {
			pi.Logger.Error(err.Error())
			pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
			return
		}
	}

	name := goblin.GetParam(r.Context(), "name")

	var posts domain.Posts
	posts, err = pi.PostRepository.FindAllPublishByTag(page, limit, name)
	if err != nil {
		pi.Logger.Error(err.Error())
		pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var pr PostResponse
	code, msg, err := pr.MakeResponseHandleIndex(posts)
	if err != nil {
		pi.Logger.Error(err.Error())
		pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	pageCount := math.Ceil(float64(count) / float64(limit))

	w.Header().Set("Pagination-Count", strconv.Itoa(count))
	w.Header().Set("Pagination-Pagecount", fmt.Sprint(pageCount))
	w.Header().Set("Pagination-Page", strconv.Itoa(page))
	w.Header().Set("Pagination-Limit", strconv.Itoa(limit))
	pi.JSONResponse.HTTPStatus(w, code, msg)
	return
}

// HandleIndexPrivate returns a listing of the resource.
func (pi *PostInteractor) HandleIndexPrivate(w http.ResponseWriter, r *http.Request) {
	const defaultPage = 1
	const defaultLimit = 10

	count, err := pi.PostRepository.CountAll()
	if err != nil {
		pi.Logger.Error(err.Error())
		pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return

	}

	paramPage := r.URL.Query().Get("page")
	var page int
	if paramPage == "" {
		page = defaultPage
	} else {
		page, err = strconv.Atoi(paramPage)
		if err != nil {
			pi.Logger.Error(err.Error())
			pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
			return
		}
	}

	paramLimit := r.URL.Query().Get("limit")
	var limit int
	if paramLimit == "" {
		limit = defaultLimit
	} else {
		limit, err = strconv.Atoi(paramLimit)
		if err != nil {
			pi.Logger.Error(err.Error())
			pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
			return
		}
	}

	var posts domain.Posts
	posts, err = pi.PostRepository.FindAll(page, limit)
	if err != nil {
		pi.Logger.Error(err.Error())
		pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var pr PostResponse
	code, msg, err := pr.MakeResponseHandleIndexPrivate(posts)
	if err != nil {
		pi.Logger.Error(err.Error())
		pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	pageCount := math.Ceil(float64(count) / float64(limit))

	w.Header().Set("Pagination-Count", strconv.Itoa(count))
	w.Header().Set("Pagination-Pagecount", fmt.Sprint(pageCount))
	w.Header().Set("Pagination-Page", strconv.Itoa(page))
	w.Header().Set("Pagination-Limit", strconv.Itoa(limit))
	pi.JSONResponse.HTTPStatus(w, code, msg)
	return
}

// HandleShow display the specified resource.
func (pi *PostInteractor) HandleShow(w http.ResponseWriter, r *http.Request) {
	title := goblin.GetParam(r.Context(), "title")

	var post domain.Post
	post, err := pi.PostRepository.FindByTitle(title)
	if err != nil {
		pi.Logger.Error(err.Error())
		pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var pr PostResponse
	code, msg, err := pr.MakeResponseHandleShow(post)
	if err != nil {
		pi.Logger.Error(err.Error())
		pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	pi.JSONResponse.HTTPStatus(w, code, msg)
	return
}

// HandleShowPrivate display the specified resource.
func (pi *PostInteractor) HandleShowPrivate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(goblin.GetParam(r.Context(), "id"))
	if err != nil {
		pi.Logger.Error(err.Error())
		pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var post domain.Post
	post, err = pi.PostRepository.FindByID(id)
	if err != nil {
		pi.Logger.Error(err.Error())
		pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var pr PostResponse
	code, msg, err := pr.MakeResponseHandleShowPrivate(post)
	if err != nil {
		pi.Logger.Error(err.Error())
		pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	pi.JSONResponse.HTTPStatus(w, code, msg)
	return
}

// HandleStorePrivate stores a newly created resource in storage.
func (pi *PostInteractor) HandleStorePrivate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		pi.Logger.Error(err.Error())
		pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var j domain.JWT
	var accessUUID string
	accessUUID, err = j.GetAccessUUID(r.Header.Get("Authorization"))
	if err != nil {
		pi.Logger.Error(err.Error())
		pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var adminID int
	adminID, err = pi.JWTRepository.FindIDByAccessUUID(accessUUID)
	if err != nil {
		pi.Logger.Error(err.Error())
		pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	req := RequestPost{
		AdminID: adminID,
	}
	err = json.Unmarshal(body, &req)
	if err != nil {
		pi.Logger.Error(err.Error())
		pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	err = pi.PostRepository.Save(req)
	if err != nil {
		pi.Logger.Error(err.Error())
		pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	pi.JSONResponse.HTTPStatus(w, http.StatusCreated, nil)
	return
}

// HandleUpdatePrivate updates the specified resource in storage.
func (pi *PostInteractor) HandleUpdatePrivate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		pi.Logger.Error(err.Error())
		pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var j domain.JWT
	var accessUUID string
	accessUUID, err = j.GetAccessUUID(r.Header.Get("Authorization"))
	if err != nil {
		pi.Logger.Error(err.Error())
		pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var adminID int
	adminID, err = pi.JWTRepository.FindIDByAccessUUID(accessUUID)
	if err != nil {
		pi.Logger.Error(err.Error())
		pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	req := RequestPost{
		AdminID: adminID,
	}
	err = json.Unmarshal(body, &req)
	if err != nil {
		pi.Logger.Error(err.Error())
		pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	id, err := strconv.Atoi(goblin.GetParam(r.Context(), "id"))
	if err != nil {
		pi.Logger.Error(err.Error())
		pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}
	err = pi.PostRepository.SaveByID(req, id)
	if err != nil {
		pi.Logger.Error(err.Error())
		pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	pi.JSONResponse.HTTPStatus(w, http.StatusCreated, nil)
	return
}

// HandleDestroyPrivate removes the specified resource from storage.
func (pi *PostInteractor) HandleDestroyPrivate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(goblin.GetParam(r.Context(), "id"))
	if err != nil {
		pi.Logger.Error(err.Error())
		pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}
	count, err := pi.PostRepository.DeleteByID(id)
	if err != nil {
		pi.Logger.Error(err.Error())
		pi.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	if count == 0 {
		pi.JSONResponse.HTTPStatus(w, http.StatusNotFound, nil)
		return
	}

	pi.JSONResponse.HTTPStatus(w, http.StatusOK, nil)
	return
}
