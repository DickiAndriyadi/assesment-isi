package nasabah

import (
	"assesment-isi/config"
	"assesment-isi/internal/constant"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (ctrl *NasabahController) RegisterNasabah(c echo.Context) (err error) {
	var req struct {
		Nama string `json:"nama"`
		NIK  string `json:"nik"`
		NoHP string `json:"no_hp"`
	}
	if err := c.Bind(&req); err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"req": req,
		})
		return c.JSON(http.StatusBadRequest, constant.RestResponse{
			Remark: err.Error(),
			Code:   http.StatusBadRequest,
			Data:   nil,
		})
	}

	result, err := ctrl.nasabah.RegisterNasabah(req.Nama, req.NIK, req.NoHP)
	if err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"req": req,
		})
		return c.JSON(http.StatusBadRequest, constant.RestResponse{
			Remark: err.Error(),
			Code:   http.StatusBadRequest,
			Data:   nil,
		})
	}

	resp := map[string]string{"no_rekening": result.NoRekening}

	return c.JSON(http.StatusOK, constant.RestResponse{
		Remark: "Nasabah berhasil dibuat!",
		Code:   http.StatusOK,
		Data:   resp,
	})
}
