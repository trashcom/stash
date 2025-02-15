package api

import (
	"context"
	"strconv"
	"strings"

	"github.com/stashapp/stash/internal/api/loaders"
	"github.com/stashapp/stash/internal/api/urlbuilders"
	"github.com/stashapp/stash/pkg/gallery"
	"github.com/stashapp/stash/pkg/image"
	"github.com/stashapp/stash/pkg/models"
)

// Checksum is deprecated
func (r *performerResolver) Checksum(ctx context.Context, obj *models.Performer) (*string, error) {
	return nil, nil
}

func (r *performerResolver) Aliases(ctx context.Context, obj *models.Performer) (*string, error) {
	if !obj.Aliases.Loaded() {
		if err := r.withReadTxn(ctx, func(ctx context.Context) error {
			return obj.LoadAliases(ctx, r.repository.Performer)
		}); err != nil {
			return nil, err
		}
	}

	ret := strings.Join(obj.Aliases.List(), ", ")
	return &ret, nil
}

func (r *performerResolver) AliasList(ctx context.Context, obj *models.Performer) ([]string, error) {
	if !obj.Aliases.Loaded() {
		if err := r.withReadTxn(ctx, func(ctx context.Context) error {
			return obj.LoadAliases(ctx, r.repository.Performer)
		}); err != nil {
			return nil, err
		}
	}

	return obj.Aliases.List(), nil
}

func (r *performerResolver) Height(ctx context.Context, obj *models.Performer) (*string, error) {
	if obj.Height != nil {
		ret := strconv.Itoa(*obj.Height)
		return &ret, nil
	}
	return nil, nil
}

func (r *performerResolver) HeightCm(ctx context.Context, obj *models.Performer) (*int, error) {
	return obj.Height, nil
}

func (r *performerResolver) Birthdate(ctx context.Context, obj *models.Performer) (*string, error) {
	if obj.Birthdate != nil {
		ret := obj.Birthdate.String()
		return &ret, nil
	}
	return nil, nil
}

func (r *performerResolver) ImagePath(ctx context.Context, obj *models.Performer) (*string, error) {
	baseURL, _ := ctx.Value(BaseURLCtxKey).(string)
	imagePath := urlbuilders.NewPerformerURLBuilder(baseURL, obj).GetPerformerImageURL()
	return &imagePath, nil
}

func (r *performerResolver) Tags(ctx context.Context, obj *models.Performer) (ret []*models.Tag, err error) {
	if !obj.TagIDs.Loaded() {
		if err := r.withReadTxn(ctx, func(ctx context.Context) error {
			return obj.LoadTagIDs(ctx, r.repository.Performer)
		}); err != nil {
			return nil, err
		}
	}

	var errs []error
	ret, errs = loaders.From(ctx).TagByID.LoadAll(obj.TagIDs.List())
	return ret, firstError(errs)
}

func (r *performerResolver) SceneCount(ctx context.Context, obj *models.Performer) (ret *int, err error) {
	var res int
	if err := r.withReadTxn(ctx, func(ctx context.Context) error {
		res, err = r.repository.Scene.CountByPerformerID(ctx, obj.ID)
		return err
	}); err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *performerResolver) ImageCount(ctx context.Context, obj *models.Performer) (ret *int, err error) {
	var res int
	if err := r.withReadTxn(ctx, func(ctx context.Context) error {
		res, err = image.CountByPerformerID(ctx, r.repository.Image, obj.ID)
		return err
	}); err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *performerResolver) GalleryCount(ctx context.Context, obj *models.Performer) (ret *int, err error) {
	var res int
	if err := r.withReadTxn(ctx, func(ctx context.Context) error {
		res, err = gallery.CountByPerformerID(ctx, r.repository.Gallery, obj.ID)
		return err
	}); err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *performerResolver) Scenes(ctx context.Context, obj *models.Performer) (ret []*models.Scene, err error) {
	if err := r.withReadTxn(ctx, func(ctx context.Context) error {
		ret, err = r.repository.Scene.FindByPerformerID(ctx, obj.ID)
		return err
	}); err != nil {
		return nil, err
	}

	return ret, nil
}

func (r *performerResolver) StashIds(ctx context.Context, obj *models.Performer) ([]*models.StashID, error) {
	if err := r.withReadTxn(ctx, func(ctx context.Context) error {
		return obj.LoadStashIDs(ctx, r.repository.Performer)
	}); err != nil {
		return nil, err
	}

	return stashIDsSliceToPtrSlice(obj.StashIDs.List()), nil
}

func (r *performerResolver) Rating(ctx context.Context, obj *models.Performer) (*int, error) {
	if obj.Rating != nil {
		rating := models.Rating100To5(*obj.Rating)
		return &rating, nil
	}
	return nil, nil
}

func (r *performerResolver) Rating100(ctx context.Context, obj *models.Performer) (*int, error) {
	return obj.Rating, nil
}

func (r *performerResolver) DeathDate(ctx context.Context, obj *models.Performer) (*string, error) {
	if obj.DeathDate != nil {
		ret := obj.DeathDate.String()
		return &ret, nil
	}
	return nil, nil
}

func (r *performerResolver) Movies(ctx context.Context, obj *models.Performer) (ret []*models.Movie, err error) {
	if err := r.withReadTxn(ctx, func(ctx context.Context) error {
		ret, err = r.repository.Movie.FindByPerformerID(ctx, obj.ID)
		return err
	}); err != nil {
		return nil, err
	}

	return ret, nil
}

func (r *performerResolver) MovieCount(ctx context.Context, obj *models.Performer) (ret *int, err error) {
	var res int
	if err := r.withReadTxn(ctx, func(ctx context.Context) error {
		res, err = r.repository.Movie.CountByPerformerID(ctx, obj.ID)
		return err
	}); err != nil {
		return nil, err
	}

	return &res, nil
}
