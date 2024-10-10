package client

import (
	oidb2 "github.com/LagrangeDev/LagrangeGo/client/packets/oidb"
	"github.com/LagrangeDev/LagrangeGo/client/packets/pb/service/oidb"
)

func (c *QQClient) GetPrivateImageUrl(node *oidb.IndexNode) (string, error) {
	pkt, err := oidb2.BuildPrivateImageDownloadReq(c.Uid, node)
	if err != nil {
		return "", err
	}
	resp, err := c.sendAndWaitDynamic(c.uniPacket(pkt.Cmd, pkt.Data))
	if err != nil {
		return "", err
	}
	return oidb2.ParsePrivateImageDownloadResp(resp)
}

// GetGroupImageUrl 获取群聊图片下载url
func (c *QQClient) GetGroupImageUrl(groupUin uint32, node *oidb.IndexNode) (string, error) {
	pkt, err := oidb2.BuildGroupImageDownloadReq(groupUin, node)
	if err != nil {
		return "", err
	}
	resp, err := c.sendAndWaitDynamic(c.uniPacket(pkt.Cmd, pkt.Data))
	if err != nil {
		return "", err
	}
	return oidb2.ParseGroupImageDownloadResp(resp)
}

func (c *QQClient) GetPrivateRecordUrl(node *oidb.IndexNode) (string, error) {
	pkt, err := oidb2.BuildPrivateRecordDownloadReq(c.Uid, node)
	if err != nil {
		return "", err
	}
	resp, err := c.sendAndWaitDynamic(c.uniPacket(pkt.Cmd, pkt.Data))
	if err != nil {
		return "", err
	}
	return oidb2.ParsePrivateRecordDownloadResp(resp)
}

// GetGroupRecordUrl 获取群聊语音下载url
func (c *QQClient) GetGroupRecordUrl(groupUin uint32, node *oidb.IndexNode) (string, error) {
	pkt, err := oidb2.BuildGroupRecordDownloadReq(groupUin, node)
	if err != nil {
		return "", err
	}
	resp, err := c.sendAndWaitDynamic(c.uniPacket(pkt.Cmd, pkt.Data))
	if err != nil {
		return "", err
	}
	return oidb2.ParseGroupRecordDownloadResp(resp)
}
