// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

/*
 * Nudm_PP
 *
 * Nudm Parameter Provision Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package parameterprovision_test

import (
	"testing"
)

// Update - provision parameters
func TestUpdate(t *testing.T) {
	/*go handler.Handle()
	go func() { // udm server
		router := gin.Default()
		Nudm_PP_Server.AddService(router)

		udmLogPath := path_util.Free5gcPath("free5gc/udmsslkey.log")
		udmPemPath := path_util.Free5gcPath("free5gc/support/TLS/udm.pem")
		udmKeyPath := path_util.Free5gcPath("free5gc/support/TLS/udm.key")
		server, err := http2_util.NewServer(":29503", udmLogPath, router)
		if err == nil && server != nil {
			logger.InitLog.Infoln(server.ListenAndServeTLS(udmPemPath, udmKeyPath))
			assert.True(t, err == nil)
		}
	}()

	go func() { // fake udr server
		router := gin.Default()

		router.PATCH("/nudr-dr/v1/:gpsi/pp-data", func(c *gin.Context) {
			gpsi := c.Param("gpsi")
			fmt.Println("==========CreateEeSubscription - Subscribe==========")
			fmt.Println("ueIdentity: ", gpsi)
			var testppData models.PpData
			if err := c.ShouldBindJSON(&testppData); err != nil {
				fmt.Println("fake udm server error: ", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{})
				return
			}
			fmt.Println("patchItems: ", testppData)
			c.JSON(http.StatusCreated, gin.H{})
		})

		udrLogPath := path_util.Free5gcPath("free5gc/udrsslkey.log")
		udrPemPath := path_util.Free5gcPath("free5gc/support/TLS/udr.pem")
		udrKeyPath := path_util.Free5gcPath("free5gc/support/TLS/udr.key")

		server, err := http2_util.NewServer(":29504", udrLogPath, router)
		if err == nil && server != nil {
			logger.InitLog.Infoln(server.ListenAndServeTLS(udrPemPath, udrKeyPath))
			assert.True(t, err == nil)
		}
	}()

	udm_context.Init()
	cfg := Nudm_PP_Client.NewConfiguration()
	cfg.SetBasePath("https://localhost:29503")
	clientAPI := Nudm_PP_Client.NewAPIClient(cfg)

	var ppData models.PpData
	ppData.SupportedFeatures = "Test001"
	gpsi := "SDM1234"
	resp, err := clientAPI.SubscriptionDataUpdateApi.Update(context.Background(), gpsi, ppData)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("resp: ", resp)
	}
	*/
}
