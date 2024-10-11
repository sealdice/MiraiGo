package client

import (
	"encoding/hex"

	"github.com/LagrangeDev/LagrangeGo/client/packets/oidb"
	"github.com/LagrangeDev/LagrangeGo/client/packets/pb/service/highway"
	oidb2 "github.com/LagrangeDev/LagrangeGo/client/packets/pb/service/oidb"
	message2 "github.com/LagrangeDev/LagrangeGo/message"
	"github.com/LagrangeDev/LagrangeGo/utils"
	"github.com/pkg/errors"
	"github.com/sealdice/MiraiGo/binary"
	highway2 "github.com/sealdice/MiraiGo/client/internal/highway"
	"github.com/sealdice/MiraiGo/internal/proto"
	"github.com/sealdice/MiraiGo/message"
)

const BlockSize = 1024 * 1024

func oidbIPv4ToNTHighwayIPv4(ipv4s []*oidb2.IPv4) []*highway.NTHighwayIPv4 {
	hwipv4s := make([]*highway.NTHighwayIPv4, len(ipv4s))
	for i, ip := range ipv4s {
		hwipv4s[i] = &highway.NTHighwayIPv4{
			Domain: &highway.NTHighwayDomain{
				IsEnable: true,
				IP:       binary.UInt32ToIPV4Address(ip.OutIP),
			},
			Port: ip.OutPort,
		}
	}
	return hwipv4s
}

func (c *QQClient) RecordUploadPrivate(targetUid string, recordRaw *message.VoiceElement) (*message.VoiceElement, error) {
	if recordRaw == nil || recordRaw.Stream == nil {
		return nil, errors.New("element type is not friend record")
	}
	record := &message2.VoiceElement{
		Md5:      recordRaw.Md5,
		Sha1:     recordRaw.Sha1,
		Size:     uint32(recordRaw.Size),
		Duration: recordRaw.Duration,
		Summary:  recordRaw.Summary,
		Stream:   recordRaw.Stream,
	}
	defer utils.CloseIO(record.Stream)
	req, err := oidb.BuildPrivateRecordUploadReq(targetUid, record)
	if err != nil {
		return nil, err
	}
	resp, err := c.sendAndWaitDynamic(c.uniPacket(req.Cmd, req.Data))
	if err != nil {
		return nil, err
	}
	uploadResp, err := oidb.ParsePrivateRecordUploadResp(resp)
	if err != nil {
		return nil, err
	}
	ukey := uploadResp.Upload.UKey.Unwrap()
	c.debug("private record upload ukey:", ukey)
	if ukey != "" {
		index := uploadResp.Upload.MsgInfo.MsgInfoBody[0].Index
		sha1, err := hex.DecodeString(index.Info.FileSha1)
		if err != nil {
			return nil, err
		}
		extend := &highway.NTV2RichMediaHighwayExt{
			FileUuid: index.FileUuid,
			UKey:     ukey,
			Network: &highway.NTHighwayNetwork{
				IPv4S: oidbIPv4ToNTHighwayIPv4(uploadResp.Upload.IPv4S),
			},
			MsgInfoBody: uploadResp.Upload.MsgInfo.MsgInfoBody,
			BlockSize:   uint32(BlockSize),
			Hash: &highway.NTHighwayHash{
				FileSha1: [][]byte{sha1},
			},
		}
		extStream, err := proto.Marshal(extend)
		if err != nil {
			return nil, err
		}
		md5, err := hex.DecodeString(index.Info.FileHash)
		if err != nil {
			return nil, err
		}
		err = c.ensureHighwayServers()
		if err != nil {
			return nil, err
		}
		input := highway2.Transaction{
			CommandID: 1007,
			Body:      record.Stream,
			Size:      int64(record.Size),
			Sum:       md5,
			Ticket:    c.highwaySession.SigSession,
			Ext:       extStream,
		}
		_, err = c.highwaySession.Upload(input)
		if err != nil {
			return nil, err
		}
	}
	recordRaw.MsgInfo = uploadResp.Upload.MsgInfo
	recordRaw.Compat = uploadResp.Upload.CompatQMsg
	return recordRaw, nil
}

func (c *QQClient) RecordUploadGroup(groupUin uint32, recordRaw *message.VoiceElement) (*message.VoiceElement, error) {
	if recordRaw == nil || recordRaw.Stream == nil {
		return nil, errors.New("element type is not voice record")
	}
	record := &message2.VoiceElement{
		Md5:      recordRaw.Md5,
		Sha1:     recordRaw.Sha1,
		Size:     uint32(recordRaw.Size),
		Duration: recordRaw.Duration,
		Summary:  recordRaw.Summary,
		Stream:   recordRaw.Stream,
	}
	defer utils.CloseIO(record.Stream)
	req, err := oidb.BuildGroupRecordUploadReq(groupUin, record)
	if err != nil {
		return nil, err
	}
	resp, err := c.sendAndWaitDynamic(c.uniPacket(req.Cmd, req.Data))
	if err != nil {
		return nil, err
	}
	uploadResp, err := oidb.ParseGroupRecordUploadResp(resp)
	if err != nil {
		return nil, err
	}
	ukey := uploadResp.Upload.UKey.Unwrap()
	c.debug("group record upload ukey:", ukey)
	if ukey != "" {
		index := uploadResp.Upload.MsgInfo.MsgInfoBody[0].Index
		sha1, err := hex.DecodeString(index.Info.FileSha1)
		if err != nil {
			return nil, err
		}
		extend := &highway.NTV2RichMediaHighwayExt{
			FileUuid: index.FileUuid,
			UKey:     ukey,
			Network: &highway.NTHighwayNetwork{
				IPv4S: oidbIPv4ToNTHighwayIPv4(uploadResp.Upload.IPv4S),
			},
			MsgInfoBody: uploadResp.Upload.MsgInfo.MsgInfoBody,
			BlockSize:   uint32(highway2.BlockSize),
			Hash: &highway.NTHighwayHash{
				FileSha1: [][]byte{sha1},
			},
		}
		extStream, err := proto.Marshal(extend)
		if err != nil {
			return nil, err
		}
		hash, err := hex.DecodeString(index.Info.FileHash)
		if err != nil {
			return nil, err
		}
		err = c.ensureHighwayServers()
		if err != nil {
			return nil, err
		}
		input := highway2.Transaction{
			CommandID: 1008,
			Body:      record.Stream,
			Size:      int64(record.Size),
			Sum:       hash,
			Ticket:    c.highwaySession.SigSession,
			Ext:       extStream,
		}
		_, err = c.highwaySession.Upload(input)
		if err != nil {
			return nil, err
		}
	}
	recordRaw.MsgInfo = uploadResp.Upload.MsgInfo
	recordRaw.Compat = uploadResp.Upload.CompatQMsg
	return recordRaw, nil
}

func (c *QQClient) UploadRecord(target message.Source, voice *message.VoiceElement) (*message.VoiceElement, error) {
	switch target.SourceType {
	case message.SourceGroup:
		return c.RecordUploadGroup(uint32(target.PrimaryID), voice)
	case message.SourcePrivate:
		return c.RecordUploadPrivate(c.GetUid(uint32(target.PrimaryID)), voice)
	}
	return nil, errors.New("unknown target type")
}
