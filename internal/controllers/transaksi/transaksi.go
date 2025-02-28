package transaksi

import (
	"assesment-isi/config"
	"assesment-isi/internal/constant"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (ctrl *TransaksiController) Tabung(c echo.Context) (err error) {
	var req struct {
		NoRekening string `json:"no_rekening"`
		Nominal    int64  `json:"nominal"`
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

	result, err := ctrl.transaksi.Tabung(req.NoRekening, req.Nominal)
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

	resp := map[string]int64{"saldo": result.Saldo}

	return c.JSON(http.StatusOK, constant.RestResponse{
		Remark: "Saldo berhasil ditambahkan!",
		Code:   http.StatusOK,
		Data:   resp,
	})
}

func (ctrl *TransaksiController) Tarik(c echo.Context) (err error) {
	var req struct {
		NoRekening string `json:"no_rekening"`
		Nominal    int64  `json:"nominal"`
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

	result, err := ctrl.transaksi.Tarik(req.NoRekening, req.Nominal)
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

	resp := map[string]int64{"saldo": result.Saldo}

	return c.JSON(http.StatusOK, constant.RestResponse{
		Remark: "Saldo berhasil ditarik!",
		Code:   http.StatusOK,
		Data:   resp,
	})
}

func (ctrl *TransaksiController) GetSaldo(c echo.Context) (err error) {
	noRekening := c.Param("no_rekening")

	saldo, err := ctrl.transaksi.GetSaldo(noRekening)
	if err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"no_rekening": noRekening,
		})
		return c.JSON(http.StatusBadRequest, constant.RestResponse{
			Remark: err.Error(),
			Code:   http.StatusBadRequest,
			Data:   nil,
		})
	}

	resp := map[string]int64{"saldo": saldo}

	return c.JSON(http.StatusOK, constant.RestResponse{
		Remark: "Saldo berhasil ditampilkan!",
		Code:   http.StatusOK,
		Data:   resp,
	})
}
