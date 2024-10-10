package client

import (
	"github.com/LagrangeDev/LagrangeGo/client/entity"
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

// FetchGroupMembers 获取对应群的所有群成员信息，使用token可以获取下一页的群成员信息
func (c *QQClient) FetchGroupMembers(groupUin uint32, token string) ([]*entity.GroupMember, string, error) {
	pkt, err := oidb2.BuildFetchMembersReq(groupUin, token)
	if err != nil {
		return nil, "", err
	}
	resp, err := c.sendAndWaitDynamic(c.uniPacket(pkt.Cmd, pkt.Data))
	if err != nil {
		return nil, "", err
	}
	members, newToken, err := oidb2.ParseFetchMembersResp(resp)
	if err != nil {
		return nil, "", err
	}
	return members, newToken, nil
}

// FetchRkey 获取Rkey
func (c *QQClient) FetchRkey() (entity.RKeyMap, error) {
	pkt, err := oidb2.BuildFetchRKeyReq()
	if err != nil {
		return nil, err
	}
	resp, err := c.sendAndWaitDynamic(c.uniPacket(pkt.Cmd, pkt.Data))
	if err != nil {
		return nil, err
	}
	return oidb2.ParseFetchRKeyResp(resp)
}

// FetchFriends 获取好友列表信息，使用token可以获取下一页的群成员信息
func (c *QQClient) FetchFriends(token uint32) ([]*entity.Friend, uint32, error) {
	pkt, err := oidb2.BuildFetchFriendsReq(token)
	if err != nil {
		return nil, 0, err
	}
	resp, err := c.sendAndWaitDynamic(c.uniPacket(pkt.Cmd, pkt.Data))
	if err != nil {
		return nil, 0, err
	}
	friends, token, err := oidb2.ParseFetchFriendsResp(resp)
	if err != nil {
		return nil, 0, err
	}
	return friends, token, nil
}

// FetchGroups 获取所有已加入的群的信息
func (c *QQClient) FetchGroups() ([]*entity.Group, error) {
	pkt, err := oidb2.BuildFetchGroupsReq()
	if err != nil {
		return nil, err
	}
	resp, err := c.sendAndWaitDynamic(c.uniPacket(pkt.Cmd, pkt.Data))
	if err != nil {
		return nil, err
	}
	groups, err := oidb2.ParseFetchGroupsResp(resp)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (c *QQClient) FetchGroupMember(groupUin, memberUin uint32) (*entity.GroupMember, error) {
	pkt, err := oidb2.BuildFetchMemberReq(groupUin, c.GetUid(memberUin, groupUin))
	if err != nil {
		return nil, err
	}
	resp, err := c.sendAndWaitDynamic(c.uniPacket(pkt.Cmd, pkt.Data))
	if err != nil {
		return nil, err
	}
	members, err := oidb2.ParseFetchMemberResp(resp)
	if err != nil {
		return nil, err
	}
	return members, nil
}
