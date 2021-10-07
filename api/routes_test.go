package api

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"

	"github.com/golang/mock/gomock"
)

type getCtrl struct{}

func (_ getCtrl) HandleGet(_ *gin.Context)         {}
func (_ getCtrl) UnmatchedFunction(_ *gin.Context) {}

type allCtrl struct{}

func (_ allCtrl) HandleGet(_ *gin.Context)    {}
func (_ allCtrl) HandlePost(_ *gin.Context)   {}
func (_ allCtrl) HandlePut(_ *gin.Context)    {}
func (_ allCtrl) HandleDelete(_ *gin.Context) {}

func Test_mapRoutesToHandlers(t *testing.T) {
	t.Run("happy path - single handlers", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		apiPath := "foo/api"
		r := NewMockmethodRegistrar(ctrl)
		r.EXPECT().GET(apiPath, gomock.Any())

		mapRoutesToHandlers(r, gin.H{apiPath: getCtrl{}})
	})

	t.Run("happy path - all handlers", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		apiPath := "foo/api"
		r := NewMockmethodRegistrar(ctrl)

		r.EXPECT().GET(apiPath, gomock.Any())
		r.EXPECT().POST(apiPath, gomock.Any())
		r.EXPECT().PUT(apiPath, gomock.Any())
		r.EXPECT().DELETE(apiPath, gomock.Any())

		mapRoutesToHandlers(r, gin.H{apiPath: allCtrl{}})
	})

	t.Run("happy path - multiple handlers", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		r := NewMockmethodRegistrar(ctrl)
		fooPath := "foo/api"
		r.EXPECT().GET(fooPath, gomock.Any())
		r.EXPECT().POST(fooPath, gomock.Any())
		r.EXPECT().PUT(fooPath, gomock.Any())
		r.EXPECT().DELETE(fooPath, gomock.Any())

		barPath := "bar/api"
		r.EXPECT().GET(barPath, gomock.Any())

		mapRoutesToHandlers(r, gin.H{
			fooPath: allCtrl{},
			barPath: getCtrl{},
		})
	})

	t.Run("sad path - no handler registered", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		apiPath := "foo/api"
		r := NewMockmethodRegistrar(ctrl)

		assert.Panics(t, func() {
			mapRoutesToHandlers(r, gin.H{apiPath: struct{}{}})
		})
	})
}
