// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package subscriberdatamanagement

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/free5gc/http_wrapper"
	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"
	"github.com/free5gc/udm/logger"
	"github.com/free5gc/udm/producer"
)

// GetSmData - retrieve a UE's Session Management Subscription Data
func HTTPGetSmData(c *gin.Context) {
	req := http_wrapper.NewRequest(c.Request, nil)
	req.Params["supi"] = c.Param("supi")
	req.Query.Set("plmn-id", c.Query("plmn-id"))
	req.Query.Set("dnn", c.Query("dnn"))
	req.Query.Set("single-nssai", c.Query("single-nssai"))
	req.Query.Set("supported-features", c.Query("supported-features"))

	rsp := producer.HandleGetSmDataRequest(req)

	responseBody, err := openapi.Serialize(rsp.Body, "application/json")
	if err != nil {
		logger.SdmLog.Errorln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(rsp.Status, "application/json", responseBody)
	}
}
