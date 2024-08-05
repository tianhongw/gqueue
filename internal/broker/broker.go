package broker

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/tianhongwu/gqueue/internal/task"
)

type Broker interface {
	Ping() error
	Close() error
	Enqueue(ctx context.Context, msg *task.TaskMessage) error
	EnqueueUnique(ctx context.Context, msg *task.TaskMessage, ttl time.Duration)
	Dequeue(qnames ...string) (*task.TaskMessage, time.Time, error)
	Done(ctx context.Context, msg *task.TaskMessage) error
	MarkAsComlete(ctx context.Context, msg *task.TaskMessage) error
	Requeue(ctx context.Context, msg *task.TaskMessage) error
	Schedule(ctx context.Context, msg *task.TaskMessage, processAt time.Time) error
	ScheduleUnique(ctx context.Context, msg *task.TaskMessage, processAt time.Time, ttl time.Duration) error
	Retry(ctx context.Context, msg *task.TaskMessage, processAt time.Time, errMsg string, isFailure bool) error
	Archive(ctx context.Context, msg *task.TaskMessage, errMsg string) error
	ForwardIfReady(qnames ...string) error

	AddToGroup(ctx context.Context, msg *task.TaskMessage, gname string) error
	AddToGroupUnique(ctx context.Context, msg *task.TaskMessage, groupKey string, ttl time.Duration) error
	ListGroups(qname string) ([]string, error)
	AggregationCheck(qname, gname string, t time.Time, gracePeriod, maxDely time.Duration, maxSize int) (aggregationSetID string, err error)
	ReadAggregationSet(ctx context.Context, qname, gname, aggretationSetID string) error
	ReclaimStaleAggregationSet(qname string) error

	DeleteExpiredCompletedTasks(qname string, batchSize int) error

	ListLeaseExpired(cutoff time.Time, qnames ...string) (*[]task.TaskMessage, error)
	ExtendLease(qname string, ids ...string) (time.Time, error)

	CancelationPubSub() (*redis.PubSub, error)
	PublishCancelation(id string) error

	WriteResult(qname, id string, data []byte) (n int, err error)
}
