package ocpp

import (
	"time"

	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/util/request"
	"github.com/lorenzodonini/ocpp-go/ocpp"
	ocpp16 "github.com/lorenzodonini/ocpp-go/ocpp1.6"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/core"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/firmware"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/localauth"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/remotetrigger"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/reservation"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/smartcharging"
	"github.com/lorenzodonini/ocpp-go/ocppj"
)

var instance *CS

func Instance() *CS {
	if instance == nil {
		dispatcher := ocppj.NewDefaultServerDispatcher(ocppj.NewFIFOQueueMap(0))
		dispatcher.SetTimeout(request.Timeout)

		profiles := []*ocpp.Profile{core.Profile, localauth.Profile, firmware.Profile, reservation.Profile, remotetrigger.Profile, smartcharging.Profile}
		endpoint := ocppj.NewServer(nil, dispatcher, nil, profiles...)

		cs := ocpp16.NewCentralSystem(endpoint, nil)

		instance = &CS{
			log:           util.NewLogger("ocpp"),
			cps:           make(map[string]*CP),
			CentralSystem: cs,
		}

		ocppj.SetLogger(instance)

		cs.SetCoreHandler(instance)
		cs.SetNewChargePointHandler(instance.NewChargePoint)
		cs.SetChargePointDisconnectedHandler(instance.ChargePointDisconnected)
		cs.SetFirmwareManagementHandler(instance)

		go Instance().errorHandler(cs.Errors())
		go cs.Start(8887, "/{ws}")

		time.Sleep(time.Second)
	}

	return instance
}
