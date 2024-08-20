package deliveries

import (
	"context"
	"encoding/json"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"go.uber.org/zap"
	"net/http"
	"sagara-msib-test/internal/entities"
	"sagara-msib-test/internal/services"
	jaegerLog "sagara-msib-test/pkg/log"
	"strconv"
)

type BajuHandler struct {
	bajuServices services.BajuServices
	tracer       opentracing.Tracer
	logger       jaegerLog.Factory
}

func NewBajuHandler(bajuServices services.BajuServices, tracer opentracing.Tracer, logger jaegerLog.Factory) (handler *BajuHandler) {
	handler = &BajuHandler{
		bajuServices: bajuServices,
		tracer:       tracer,
		logger:       logger,
	}

	return handler
}

func (bh *BajuHandler) HandleClient(w http.ResponseWriter, r *http.Request) {
	var (
		serviceResult interface{}
		ctx           context.Context
		err           error
		statusCode    int
	)

	spanCtx, _ := bh.tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	span := bh.tracer.StartSpan("HandleClient", ext.RPCServerOption(spanCtx))
	defer span.Finish()

	ctx = opentracing.ContextWithSpan(r.Context(), span)
	bh.logger.For(ctx).Info("HTTP request received", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	if err == nil {
		urlQueryLength := len(r.URL.Query())

		switch r.Method {
		case http.MethodGet:
			switch urlQueryLength {
			case 2:
				_, stokOK := r.URL.Query()["stok"]
				_, kondisiOK := r.URL.Query()["kondisi"]

				if stokOK && kondisiOK {
					stok, _ := strconv.Atoi(r.FormValue("stok"))
					kondisi := r.FormValue("kondisi")

					bh.logger.For(ctx).Info("Running service", zap.String("service", "Get Baju By Stok and Kondisi"))
					serviceResult, err = bh.bajuServices.GetBajuOrderByStok(stok, kondisi)
					if err != nil {
						statusCode = http.StatusNotFound
						break
					}
				}
			case 1:
				_, bajuIdOK := r.URL.Query()["bajuId"]
				_, stokOK := r.URL.Query()["stok"]

				if bajuIdOK {
					bajuId, err := strconv.Atoi(r.FormValue("bajuId"))
					if err != nil {
						statusCode = http.StatusBadRequest
						break
					}

					bh.logger.For(ctx).Info("Running service", zap.String("service", "Get Baju By Id"))
					serviceResult, err = bh.bajuServices.GetBajuByID(bajuId)
					if err != nil {
						statusCode = http.StatusNotFound
						break
					}
				} else if stokOK {
					stok := r.FormValue("stok")
					if stok == "empty" {
						bh.logger.For(ctx).Info("Running service", zap.String("service", "Get Baju with Empty Stok"))
						serviceResult, err = bh.bajuServices.GetBajuOrderByEmptyStok()
						if err != nil {
							statusCode = http.StatusNotFound
							break
						}
					}
				}
			default:
				bh.logger.For(ctx).Info("Running service", zap.String("service", "Get All Baju Data"))
				serviceResult, err = bh.bajuServices.GetAllBaju()
				if err != nil {
					statusCode = http.StatusNotFound
					break
				}
			}
		case http.MethodPost:
			switch urlQueryLength {
			default:
				var (
					bajuRequest entities.Baju
				)
				err = json.NewDecoder(r.Body).Decode(&bajuRequest)
				if err != nil {
					statusCode = http.StatusBadRequest
				}

				bh.logger.For(ctx).Info("Running service", zap.String("service", "Create New Baju"))
				err = bh.bajuServices.CreateBaju(bajuRequest)
				if err != nil {
					statusCode = http.StatusInternalServerError
				}
			}
		case http.MethodPut:
			switch urlQueryLength {
			default:
				var (
					bajuRequest entities.Baju
				)
				err = json.NewDecoder(r.Body).Decode(&bajuRequest)
				if err != nil {
					statusCode = http.StatusBadRequest
				}

				bh.logger.For(ctx).Info("Running service", zap.String("service", "Delete Baju by Id"))
				err = bh.bajuServices.UpdateBaju(bajuRequest)
				if err != nil {
					statusCode = http.StatusInternalServerError
					break
				}
			}
		case http.MethodDelete:
			switch urlQueryLength {
			case 1:
				_, bajuIdOK := r.URL.Query()["bajuId"]
				if bajuIdOK {
					bajuId, err := strconv.Atoi(r.FormValue("bajuId"))
					if err != nil {
						statusCode = http.StatusBadRequest
						break
					}

					bh.logger.For(ctx).Info("Running service", zap.String("service", "Delete Baju by Id"))
					err = bh.bajuServices.DeleteBaju(bajuId)
					if err != nil {
						statusCode = http.StatusInternalServerError
						break
					}
				}
			}
		}
	}

	if err != nil {
		http.Error(w, err.Error(), statusCode)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(serviceResult)
}
