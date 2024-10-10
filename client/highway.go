package client

import (
	"fmt"

	highway2 "github.com/LagrangeDev/LagrangeGo/client/packets/highway"
	"github.com/pkg/errors"
	"github.com/sealdice/MiraiGo/binary"
)

func (c *QQClient) ensureHighwayServers() error {
	if c.highwaySession.SsoAddr == nil || c.highwaySession.SigSession == nil || c.highwaySession.SessionKey == nil {
		fmt.Println(c.highwaySession.SsoAddr)
		packet, err := highway2.BuildHighWayUrlReq(c.sig.TGT)
		if err != nil {
			return err
		}
		payload, err := c.sendAndWaitDynamic(c.uniPacket("HttpConn.0x6ff_501", packet))
		if err != nil {
			return fmt.Errorf("get highway server: %v", err)
		}
		resp, err := highway2.ParseHighWayUrlReq(payload)
		if err != nil {
			return fmt.Errorf("parse highway server: %v", err)
		}
		c.highwaySession.SigSession = resp.HttpConn.SigSession
		c.highwaySession.SessionKey = resp.HttpConn.SessionKey
		for _, info := range resp.HttpConn.ServerInfos {
			if info.ServiceType != 1 {
				continue
			}
			for _, addr := range info.ServerAddrs {
				c.debug(fmt.Sprintf("add highway server %s:%d", binary.UInt32ToIPV4Address(addr.IP), addr.Port))
				c.highwaySession.AppendAddr(addr.IP, addr.Port)
			}
		}
	}
	if c.highwaySession.SsoAddr == nil || c.highwaySession.SigSession == nil || c.highwaySession.SessionKey == nil {
		return errors.New("empty highway servers")
	}
	return nil
}
