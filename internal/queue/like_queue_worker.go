package queue

import (
	"log"
)

func (l *likeQueueImplement) startWorkers() {
	for i := 0; i < l.workers; i++ {
		go l.worker(i)
	}
}

func (l *likeQueueImplement) worker(workerID int) {
	log.Println("Starting handle queue with workerID:", workerID)

	select {
	case likeForChan, ok := <-l.qLike:
		if !ok {
			log.Printf("Worker: %d, channel is closed.", workerID)
			return
		}

		l.handleQueue(likeForChan)

	case <-l.stop:
		log.Printf("Catch signal that channel is closed")
		return
	}
}

func (l *likeQueueImplement) handleQueue(likeForChan LikeForChan) (int, error) {
	likeID, err := l.likeService.CreateLike(likeForChan.UserID, likeForChan.PostID)
	if err != nil {
		return 0, err
	}
	return likeID, err
}
