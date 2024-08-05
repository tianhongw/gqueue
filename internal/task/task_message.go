package task

import (
	"errors"
	"fmt"

	pb "github.com/tianhongwu/gqueue/internal/proto"
	"google.golang.org/protobuf/proto"
)

type TaskMessage struct {
	Type string

	Payload []byte

	ID string

	Queue string

	Retry int

	Retried int

	ErrorMsg string

	LastFailedAt int64

	Timeout int64

	Deadline int64

	UniqueKey string

	GroupKey string

	Retention int64

	CompletedAt int64
}

func EncodeMessage(msg *TaskMessage) ([]byte, error) {
	if msg == nil {
		return nil, errors.New("cannot encode nil message")
	}

	return proto.Marshal(&pb.TaksMessage{
		Type:         msg.Type,
		Payload:      msg.Payload,
		Id:           msg.ID,
		Queue:        msg.Queue,
		Retry:        int32(msg.Retry),
		Retried:      int32(msg.Retried),
		ErrorMsg:     msg.ErrorMsg,
		LastFailedAt: msg.LastFailedAt,
		Timeout:      msg.Timeout,
		Deadline:     msg.Deadline,
		UniqueKey:    msg.UniqueKey,
		GroupKey:     msg.GroupKey,
		Retention:    msg.Retention,
	})
}

func DecodeMessage(data []byte) (*TaskMessage, error) {
	var pbmsg pb.TaksMessage
	if err := proto.Unmarshal(data, &pbmsg); err != nil {
		return nil, fmt.Errorf("decode message failed: %v", err)
	}

	return &TaskMessage{
		Type:         pbmsg.GetType(),
		Payload:      pbmsg.GetPayload(),
		ID:           pbmsg.GetId(),
		Queue:        pbmsg.GetQueue(),
		Retry:        int(pbmsg.GetRetry()),
		Retried:      int(pbmsg.GetRetried()),
		ErrorMsg:     pbmsg.GetErrorMsg(),
		LastFailedAt: pbmsg.GetLastFailedAt(),
		Timeout:      pbmsg.GetTimeout(),
		Deadline:     pbmsg.GetDeadline(),
		UniqueKey:    pbmsg.GetUniqueKey(),
		GroupKey:     pbmsg.GetGroupKey(),
		Retention:    pbmsg.GetRetention(),
	}, nil
}
